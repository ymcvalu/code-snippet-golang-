package retry

import (
	"github.com/ymcvalu/code-snippet-golang/backoff"
	"fmt"
	"time"
)

var strategy = backoff.New(time.Second * 5)

func MaxTimes(times int, fn func() error) (err error) {
	for i := 0; i < times; i++ {
		func() {
			defer func() {
				if ee := recover(); ee != nil {
					err = fmt.Errorf("[panic] %v", ee)
				}
			}()
			err = fn()
		}()
		if err == nil || i == times-1 {
			break
		}

		delay := strategy.Backoff(i)
		time.Sleep(delay)
	}

	return
}

func DelayMaxTimes(delay, times int, fn func() error) (err error) {
	for i := 0; i < times; i++ {
		func() {
			defer func() {
				if ee := recover(); ee != nil {
					err = fmt.Errorf("[panic] %v", ee)
				}
			}()
			err = fn()
		}()
		if err == nil || i == times-1 {
			break
		}
		time.Sleep(time.Duration(delay) * time.Second)
	}

	return
}
