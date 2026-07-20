package sul

import (
	"strconv"
	"sync"
	"time"
)

type Runnable[T any] struct {
	Execute func() T
}

type ThreadPoolExecutor[T any] struct {
	NbThreads      int                // Number of worker thread
	QueueSize      int                // Max queue size for both channels
	ChannelResults chan T             // Channel containing the results (added by the Runnable)
	DebugCallback  func(event string) // Not mandatory, allow to receive events and to log them, etc.

	channelJobs chan Runnable[T] // Never used directly (use tpe.EnqueueJob())
	wg          sync.WaitGroup   // Internal semaphore

	startTimestamp time.Time     // timestamp at which the workers are started
	elapsed        time.Duration // full duration of the TPE
}

func (tpe *ThreadPoolExecutor[T]) Event(s string) {
	if tpe.DebugCallback != nil {
		tpe.DebugCallback(s)
	}
}

func (tpe *ThreadPoolExecutor[T]) EventWorker(seed int, s string) {
	tpe.Event("[#" + strconv.Itoa(seed) + "] " + s)
}

// T is the result type
func (tpe *ThreadPoolExecutor[T]) StartThreads() {
	tpe.channelJobs = make(chan Runnable[T], tpe.QueueSize)
	tpe.ChannelResults = make(chan T, tpe.QueueSize)
	tpe.wg = sync.WaitGroup{}
	tpe.startTimestamp = time.Now()

	tpe.Event("Starting thread pool executor with [" + strconv.Itoa(tpe.NbThreads) + "] number of threads, queue size [" + strconv.Itoa(tpe.QueueSize) + "]")

	for i := 0; i < tpe.NbThreads; i++ {
		tpe.wg.Add(1)
		go func() {
			defer tpe.wg.Done()
			tpe.EventWorker(i, "Starting")
			for ch := range tpe.channelJobs {
				tpe.EventWorker(i, "Processing")
				tpe.ChannelResults <- ch.Execute()
			}
		}()
	}
}

func (tpe *ThreadPoolExecutor[T]) StopThreads() {
	// close job channel to signal workers to finish, wait for them, then close results
	close(tpe.channelJobs)
	tpe.Event("Awaiting thread jobs termination ...")
	tpe.wg.Wait()
	close(tpe.ChannelResults)
	tpe.elapsed = time.Since(tpe.startTimestamp)
	tpe.Event("All jobs finished & channels closed, thread executed during " + HumanizeDuration(tpe.elapsed))
}

func (tpe *ThreadPoolExecutor[T]) ElapsedTime() time.Duration {
	return tpe.elapsed
}

func (tpe *ThreadPoolExecutor[T]) EnqueueJob(executionWorker Runnable[T]) {
	tpe.Event("Enqueing new job ...")
	tpe.channelJobs <- executionWorker
}
