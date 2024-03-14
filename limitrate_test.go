package limiter

import (
	"fmt"
	"testing"
	"time"
)

func Test_LimitRateIsLimit(t *testing.T) {
	lr := NewLimitRate(10, time.Minute)

	for i := 0; i < 100; i++ {
		go func() {
			if lr.IsLimit() {
				fmt.Println("limited")
			} else {
				fmt.Println("ok")
			}
		}()

	}
}
