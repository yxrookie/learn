package rollingwindow

import (
	"sync"
	"time"

	"github.com/zeromicro/go-zero/core/timex"
)

type rollingwindow struct {
	lock          sync.RWMutex
	size          int
	win           *window
	interval      time.Duration
	offset        int
	lastTime      time.Duration 
}

func Newrollingwindow(size int, interval time.Duration) *rollingwindow {
	if size < 1 {
		panic("size must be greater than 0")
	}
	return &rollingwindow {
		// 多余
		//lock: sync.RWMutex{},
		size: size,
		win: Newwindow(size),
		interval: interval,
		lastTime: timex.Now(),
	}
}

func (r *rollingwindow) Span()int {
	
	offest := int(timex.Since(r.lastTime) / r.interval)
	if 0 <= offest && offest < r.size {
		return offest
	}
	return r.size
}



func (r *rollingwindow) Update() {
	span := r.Span()
	if span <= 0 {
		return
	}
	// 将窗口过期的桶数据清零
	for i := r.offset; i < r.offset+span; i++ {
		r.win.Reset((i+1) % r.size)
	}
	r.offset = (r.offset+span) % r.size
	// my fault
	//r.win.buckets[(r.offset+span) % r.size ].Add(v)
	// my second fault 
	// when span == 0 it return, so can't put it here
	// because span == 0, need to add data
	//r.win.Add(v, r.offset)
	now := timex.Now()
	r.lastTime = now - (now-r.lastTime)%r.interval
}

 func (r *rollingwindow) Add(v float64) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// 原始算法做的不好，没有对旧窗口中的数据没有清理
	r.Update()
	r.win.Add(v, r.offset)
} 

func (r *rollingwindow) Reduce (fn func(b *Bucket)) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	var diff int
	span := r.Span()
	diff = r.size-span
	if diff > 0 {
		// 注意传入函数的 offset 的开始位置，并不是当前滑动窗口偏移量的位置
		offset := (r.offset + span + 1) % r.size
		r.win.Reduce(offset, diff, fn)
	}
}