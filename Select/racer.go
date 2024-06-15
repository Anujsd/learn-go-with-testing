package selectPack

import (
	"fmt"
	"net/http"
	"time"
)

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// func Racer(slowUrl, fastUrl string) string {
// 	aDuration := measureResponseTime(slowUrl)
// 	bDuration := measureResponseTime(fastUrl)

// 	fmt.Printf("Slow URL: %v \n", aDuration)
// 	fmt.Printf("Fast URL: %v \n", bDuration)
// 	if aDuration < bDuration {
// 		return slowUrl
// 	}

//		return fastUrl
//	}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
