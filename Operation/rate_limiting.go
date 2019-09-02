package main

import "time"
import "fmt"

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// request 1 2019-09-02 10:06:01.620587 +0300 +03 m=+0.200715809
// request 2 2019-09-02 10:06:01.821263 +0300 +03 m=+0.401388656
// request 3 2019-09-02 10:06:02.021113 +0300 +03 m=+0.601234779
// request 4 2019-09-02 10:06:02.220659 +0300 +03 m=+0.800776716
// request 5 2019-09-02 10:06:02.420668 +0300 +03 m=+1.000782269
// request 1 2019-09-02 10:06:02.420849 +0300 +03 m=+1.000962806
// request 2 2019-09-02 10:06:02.420918 +0300 +03 m=+1.001031476
// request 3 2019-09-02 10:06:02.420931 +0300 +03 m=+1.001044886
// request 4 2019-09-02 10:06:02.625116 +0300 +03 m=+1.205226153
// request 5 2019-09-02 10:06:02.821209 +0300 +03 m=+1.401315713
