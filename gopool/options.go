package gopool

import (
	"time"
)

// options provides all optional parameters
type options struct {
	queueCap     int32
	minWorkers   int32
	maxWorkers   int32
	tasksBelowN  int32 // in shrinkPreiod  表示在“shrinkPeriod”期间，当任务数量低于这个值时，可能会触发某些操作（例如减少工作单元数量）
	shrinkPeriod time.Duration
	panicHandler func(any)
}

// Option function
type Option func(*options)

func setOptions(optL ...Option) options {
	opts := options{
		queueCap:     128,
		minWorkers:   8,
		maxWorkers:   256,
		tasksBelowN:  1024,
		shrinkPeriod: time.Minute,
	}

	for _, opt := range optL {
		opt(&opts) //调用当前遍历optL中的 Option类型函数，并将上面构建的opts 的地址作为参数传递给它。这样，Option 函数就可以修改 opts 的字段值。
	}
	return opts
}

//这下面就没什么好看的了，都是饿setter()函数 ===================================

// QueueCap set the capacity of the pool's queue
func QueueCap(v int32) Option {
	if v < 1 {
		panic("gopool:QueueCap: param is illegal")
	}
	return func(o *options) {
		o.queueCap = v
	}
}

// MinWorkers set min workers
func MinWorkers(v int32) Option {
	if v < 1 {
		panic("gopool:MinWorkers: param is illegal")
	}
	return func(o *options) {
		o.minWorkers = v
	}
}

// MaxWorkers set max workers
func MaxWorkers(v int32) Option {
	return func(o *options) {
		if v < 1 {
			panic("gopool:MinWorkers: param is illegal")
		}
		o.maxWorkers = v
	}
}

// ShrinkPeriod set shrink cycle
func ShrinkPeriod(v time.Duration) Option {
	if v < 1 {
		panic("gopool:ShrinkPeriod: param is illegal")
	}
	return func(o *options) {
		o.shrinkPeriod = v
	}
}

// TasksBelowNToShrink set shrink condition
func TasksBelowNToShrink(v int32) Option {
	if v < 1 {
		panic("gopool:TasksBelowNToShrink: param is illegal")
	}
	return func(o *options) {
		o.tasksBelowN = v
	}
}

// PanicHandler set panic handler
func PanicHandler(fn func(any)) Option {
	if fn == nil {
		panic("gopool:PanicHandler: param is illegal")
	}
	return func(o *options) {
		o.panicHandler = fn
	}
}
