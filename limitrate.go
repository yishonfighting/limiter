package limiter

import (
	"sync"
	"time"
)

// LimitRate 计数器
type LimitRate struct {
	rate  int           // 阀值
	begin time.Time     // 开始时间
	cycle time.Duration // 时间周期

	reqCount int // 请求数
	locker   sync.Mutex
}

func NewLimitRate(rate int, cycle time.Duration) *LimitRate {
	return &LimitRate{
		rate:  rate,
		cycle: cycle,
		begin: time.Now(),
	}
}

func (s *LimitRate) IsLimit() (limited bool) {
	s.locker.Lock()
	defer s.locker.Unlock()

	if time.Now().After(s.begin.Add(s.cycle)) {
		s.reset()
	}

	if s.reqCount >= s.rate {
		return true
	}

	s.reqCount++
	return
}

func (s *LimitRate) reset() {
	s.begin = time.Now()
	s.reqCount = 0
}
