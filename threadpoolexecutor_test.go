package sul

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTPEWithQueueSizeSmallerThanExpectedJobs(t *testing.T) {
	maxThreads := 5
	maxJobs := 20
	queueSize := 2
	tpe := ThreadPoolExecutor[bool]{NbThreads: maxThreads, QueueSize: queueSize, DebugCallback: func(event string) {
		fmt.Println(" > " + event)
	}}
	tpe.StartThreads()

	// As queue size < jobs, we have to start consumers *before* injecting jobs
	go func() {
		for i := range maxJobs {
			<-tpe.ChannelResults
			fmt.Println("job : " + strconv.Itoa(i) + " finished")
		}
	}()

	for range maxJobs {
		tpe.EnqueueJob(Runnable[bool]{
			Execute: func() bool {
				return true
			},
		})
	}

	tpe.StopThreads()
}

func TestTPEWithUnbufferedChannels(t *testing.T) {
	maxThreads := 5
	maxJobs := 20
	queueSize := 0 // will trigger unbuffered channels at GO level
	tpe := ThreadPoolExecutor[bool]{NbThreads: maxThreads, QueueSize: queueSize, DebugCallback: func(event string) {
		fmt.Println(" > " + event)
	}}
	tpe.StartThreads()

	// As we have unbuffered channels (due to queue size == 0), we have to launch consumers *before* injecting jobs
	go func() {
		for i := range maxJobs {
			<-tpe.ChannelResults
			fmt.Println("job : " + strconv.Itoa(i) + " finished")
		}
	}()

	for range maxJobs {
		tpe.EnqueueJob(Runnable[bool]{
			Execute: func() bool {
				return true
			},
		})
	}

	tpe.StopThreads()
}

func TestTPEWithoutQueueSizeAndByStartingProcessingResultsAtTheEnd(t *testing.T) {
	maxThreads := 5
	maxJobs := 20
	queueSize := 20
	tpe := ThreadPoolExecutor[bool]{NbThreads: maxThreads, QueueSize: queueSize, DebugCallback: func(event string) {
		fmt.Println(" > " + event)
	}}
	tpe.StartThreads()

	for range maxJobs {
		tpe.EnqueueJob(Runnable[bool]{
			Execute: func() bool {
				return true
			},
		})
	}

	// As we have queue size >= number of jobs, we can process result channels *after* having injected all jobs
	for i := range maxJobs {
		<-tpe.ChannelResults
		fmt.Println("job : " + strconv.Itoa(i) + " finished")
	}

	tpe.StopThreads()
}
