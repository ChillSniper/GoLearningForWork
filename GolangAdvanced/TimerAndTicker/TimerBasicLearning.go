package TimerAndTicker

import (
	"fmt"
	"time"
)

func BuildTimer() {
	timer := time.NewTimer(time.Second)
	<-timer.C
	fmt.Println("Time out !")
}

func CloseTimer() {
	timer := time.NewTimer(2 * time.Second)
	//<-timer.C
	res := timer.Stop()
	fmt.Println(res)
}

func ResetTimer() {
	timer := time.NewTimer(time.Second * 3)

	<-timer.C
	fmt.Println("time outA")

	resA := timer.Stop()
	fmt.Printf("resA is %t\n", resA)

	timer.Reset(time.Second * 2)
	resB := timer.Stop()
	fmt.Printf("resB is %t\n", resB)
}

func AfterTime() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- "test"
	}()

	select {
	case val := <-ch:
		fmt.Printf("val is %s\n", val)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}

func Watch() chan struct{} {
	ticker := time.NewTicker(time.Second)

	ch := make(chan struct{})

	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("watch!")
			case <-ch:
				fmt.Println("close !")
				return
			}
		}
	}(ticker)
	return ch
}

func TickerLearning() {
	ch := Watch()
	time.Sleep(time.Second * 5)
	ch <- struct{}{}
	close(ch)
}
