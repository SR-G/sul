package sul

import (
	"fmt"
	"strconv"
	"testing"
)

func TestTPEWithQueueSize(t *testing.T) {
	maxThreads := 5
	maxJobs := 20
	queueSize := 2
	tpe := ThreadPoolExecutor[bool]{NbThreads: maxThreads, QueueSize: queueSize, DebugCallback: func(event string) {
		fmt.Println(" > " + event)
	}}
	tpe.StartThreads()

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

func TestTPEWithoutQueueSize(t *testing.T) {
	maxThreads := 5
	maxJobs := 20
	queueSize := 0
	tpe := ThreadPoolExecutor[bool]{NbThreads: maxThreads, QueueSize: queueSize, DebugCallback: func(event string) {
		fmt.Println(" > " + event)
	}}
	tpe.StartThreads()

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

	for i := range maxJobs {
		<-tpe.ChannelResults
		fmt.Println("job : " + strconv.Itoa(i) + " finished")
	}

	tpe.StopThreads()
}
