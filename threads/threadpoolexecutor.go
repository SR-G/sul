package threads

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	sulstrings "github.com/SR-G/sul/strings"
)

var (
	ErrExecutorNotStarted = errors.New("thread pool executor is not started")
	ErrExecutorStopped    = errors.New("thread pool executor is stopped")
)

type Runnable[T any] struct {
	Execute func() T
}

// ThreadPoolExecutor runs jobs on NbThreads workers.
// Results are pushed on ChannelResults, whose capacity is QueueSize: consumers must read it
// concurrently (or QueueSize must be >= the number of jobs), otherwise workers block.
type ThreadPoolExecutor[T any] struct {
	NbThreads      int                // Number of worker threads (defaults to 1 if <= 0)
	QueueSize      int                // Max queue size for both channels
	ChannelResults chan T             // Channel containing the results (added by the Runnable)
	DebugCallback  func(event string) // Not mandatory, allow to receive events and to log them, etc.
	ErrorCallback  func(err error)    // Not mandatory, called when a job panics (a zero value result is still emitted)

	channelJobs chan Runnable[T] // Never used directly (use tpe.EnqueueJob())
	wg          sync.WaitGroup   // Internal semaphore

	mu      sync.RWMutex // guards started/stopped, elapsed and the closing of channelJobs
	started bool
	stopped bool

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

// StartThreads starts the workers. Calling it on an already started executor is a no-op.
func (tpe *ThreadPoolExecutor[T]) StartThreads() {
	tpe.mu.Lock()
	defer tpe.mu.Unlock()
	if tpe.started {
		return
	}
	if tpe.NbThreads <= 0 {
		tpe.NbThreads = 1
	}
	tpe.channelJobs = make(chan Runnable[T], tpe.QueueSize)
	tpe.ChannelResults = make(chan T, tpe.QueueSize)
	tpe.startTimestamp = time.Now()
	tpe.started = true

	if tpe.QueueSize <= 0 {
		tpe.Event("Starting thread pool executor with [" + strconv.Itoa(tpe.NbThreads) + "] number of threads, unbuffered channels")
	} else {
		tpe.Event("Starting thread pool executor with [" + strconv.Itoa(tpe.NbThreads) + "] number of threads, channels queue size [" + strconv.Itoa(tpe.QueueSize) + "]")
	}

	for i := 0; i < tpe.NbThreads; i++ {
		tpe.wg.Add(1)
		go tpe.worker(i)
	}
}

func (tpe *ThreadPoolExecutor[T]) worker(id int) {
	defer tpe.wg.Done()
	tpe.EventWorker(id, "Starting")
	for job := range tpe.channelJobs {
		tpe.EventWorker(id, "Processing")
		tpe.ChannelResults <- tpe.run(id, job)
	}
}

// run executes a job, turning a panic into a zero value result so consumers counting results don't hang.
func (tpe *ThreadPoolExecutor[T]) run(id int, job Runnable[T]) (result T) {
	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("job panicked in worker [#%d]: %v", id, r)
			tpe.Event(err.Error())
			if tpe.ErrorCallback != nil {
				tpe.ErrorCallback(err)
			}
			var zero T
			result = zero
		}
	}()
	return job.Execute()
}

// StopThreads waits for all enqueued jobs to complete, then closes the results channel. It is idempotent.
func (tpe *ThreadPoolExecutor[T]) StopThreads() {
	tpe.mu.Lock()
	if !tpe.started || tpe.stopped {
		tpe.mu.Unlock()
		return
	}
	tpe.stopped = true
	close(tpe.channelJobs)
	tpe.mu.Unlock()

	tpe.Event("Awaiting thread jobs termination ...")
	tpe.wg.Wait()
	close(tpe.ChannelResults)
	tpe.mu.Lock()
	tpe.elapsed = time.Since(tpe.startTimestamp)
	elapsed := tpe.elapsed
	tpe.mu.Unlock()
	tpe.Event("All jobs finished & channels closed, thread executed during " + sulstrings.HumanizeDuration(elapsed))
}

func (tpe *ThreadPoolExecutor[T]) ElapsedTime() time.Duration {
	tpe.mu.RLock()
	defer tpe.mu.RUnlock()
	return tpe.elapsed
}

// EnqueueJob blocks while the jobs queue is full.
func (tpe *ThreadPoolExecutor[T]) EnqueueJob(executionWorker Runnable[T]) error {
	return tpe.EnqueueJobContext(context.Background(), executionWorker)
}

// EnqueueJobContext is like EnqueueJob but gives up when ctx is done.
func (tpe *ThreadPoolExecutor[T]) EnqueueJobContext(ctx context.Context, executionWorker Runnable[T]) error {
	tpe.mu.RLock()
	defer tpe.mu.RUnlock()
	if !tpe.started {
		return ErrExecutorNotStarted
	}
	if tpe.stopped {
		return ErrExecutorStopped
	}
	tpe.Event("Enqueing new job ...")
	select {
	case tpe.channelJobs <- executionWorker:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
