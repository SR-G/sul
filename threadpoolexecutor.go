package sul

import (
	"strconv"
	"sync"
)

type Runnable struct {
	Execute func()
}

type ThreadPoolExecutor[T any] struct {
	NbThreads      int                // Number of worker thread
	QueueSize      int                // Max queue size for both channels
	ChannelResults chan T             // Channel containing the results (added by the Runnable)
	DebugCallback  func(event string) // Not mandatory, allow to receive events and to log them, etc.

	channelJobs chan Runnable  // Never used directly (use tpe.EnqueueJob())
	wg          sync.WaitGroup // Internal semaphore
}

func (tpe *ThreadPoolExecutor[T]) StartThreads() {
	tpe.channelJobs = make(chan Runnable, tpe.QueueSize)
	tpe.ChannelResults = make(chan T, tpe.QueueSize)
	tpe.wg = sync.WaitGroup{}

	if tpe.DebugCallback != nil {
		tpe.DebugCallback("Starting thread pool executor with [" + strconv.Itoa(tpe.NbThreads) + "] number of threads, queue size [" + strconv.Itoa(tpe.QueueSize) + "]")
	}

	for i := 0; i < tpe.NbThreads; i++ {
		if tpe.DebugCallback != nil {
			tpe.DebugCallback("Starting thread seed [" + strconv.Itoa(i+1) + "]")
		}
		tpe.wg.Add(1)
		go func() {
			defer tpe.wg.Done()
			for ch := range tpe.channelJobs {
				ch.Execute()
			}
		}()
	}
}

func (tpe *ThreadPoolExecutor[T]) StopThreads() {
	// close job channel to signal workers to finish, wait for them, then close results
	close(tpe.channelJobs)
	if tpe.DebugCallback != nil {
		tpe.DebugCallback("Awaiting thread jobs termination ...")
	}
	tpe.wg.Wait()
	close(tpe.ChannelResults)
	if tpe.DebugCallback != nil {
		tpe.DebugCallback("All jobs finished & channels closed")
	}
}

func (tpe *ThreadPoolExecutor[T]) EnqueueJob(executionWorker Runnable) {
	if tpe.DebugCallback != nil {
		tpe.DebugCallback("Enqueing new job ...")
	}
	tpe.channelJobs <- executionWorker
}

func (tpe *ThreadPoolExecutor[T]) EnqueueResult(result T) {
	if tpe.DebugCallback != nil {
		tpe.DebugCallback("Enqueing result ...")
	}
	tpe.ChannelResults <- result
}
