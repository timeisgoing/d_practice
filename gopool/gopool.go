package gopool

import (
	"sync/atomic"
	"time"
)

// GoPool is a minimalistic goroutine pool that provides a pure Go implementation
type GoPool struct {
	noCopy

	queueLen atomic.Int32
	doTaskN  atomic.Int32
	workerN  atomic.Int32 //提交到需要执行的任务数
	options  options

	workerSem chan struct{} //控制协程数量的信号量
	queue     chan func()   //存储任务的通道 queue
}

// NewGoPool provite fixed number of goroutines, reusable. M:N model
//
// M: the number of reusable goroutines,
// N: the capacity for asynchronous task queue.
func NewGoPool(opts ...Option) *GoPool {
	opt := setOptions(opts...)
	if opt.minWorkers <= 0 {
		panic("GoPool: min workers <= 0")
	}
	if opt.minWorkers > opt.maxWorkers {
		panic("GoPool: min workers > max workers")
	}
	p := &GoPool{
		options:   opt,
		workerSem: make(chan struct{}, opt.maxWorkers),
		queue:     make(chan func(), opt.queueCap),
	}
	for i := int32(0); i < p.options.minWorkers; i++ { // pre spawn
		p.workerSem <- struct{}{} //往信号量channel中装对象
		go p.worker(func() {})    //加入任务方法到任务的chan
	}
	go p.shrink()
	return p
}

// QueueFree returns (capacity of task-queue - length of task-queue)
func (p *GoPool) QueueFree() int {
	return int(p.options.queueCap - p.queueLen.Load())
}

// Workers returns current the number of workers
func (p *GoPool) Workers() int {
	return int(p.workerN.Load())
}

// Go submits a task to this pool.
func (p *GoPool) Go(task func()) {
	if task == nil {
		panic("GoPool: Go task is nil")
	}
	select {
	case p.queue <- task: //放置任务
		p.queueLen.Add(1)
	case p.workerSem <- struct{}{}: //同时存一个信号量对象
		go p.worker(task) //开启一个协程去执行这个任务
	}
}

func (p *GoPool) worker(task func()) {
	p.workerN.Add(1) //添加一个任务，workerN+1
	defer func() {
		<-p.workerSem     //workerSem`（可能是一个信号量）中接收一个值
		p.workerN.Add(-1) //当工作线程或协程结束时，`workerN` 的值减1
		if e := recover(); e != nil {
			if p.options.panicHandler != nil {
				p.options.panicHandler(e)
			}
		}
	}()

	for {
		task()           //执行当前任务
		task = <-p.queue //从队列中取出下一个任务并赋值给 `task`
		//  在 Go(task func())中task已经放入 p.queue 了，为何还要go p.worker(task) 当作参数传入呢？) worker(task func())方法中task() 和task = <-p.queue会不会同一个task执行了两次呢？

		if task == nil {
			break //一旦遇到nil，这个循环就不执行了
		}
		p.doTaskN.Add(1)   //当一个任务被取出并执行时doTaskN+1
		p.queueLen.Add(-1) //当一个任务被取出时，队列的长度减1
	}
}

// 这个方法的主要目的是定期检查任务的数量，并在任务数量低于某个阈值时，减少工作线程或协程的数量
func (p *GoPool) shrink() {
	ticker := time.NewTicker(p.options.shrinkPeriod)
	defer ticker.Stop()

	for { //无限循环，用于持续监听定时器的事件。
		select {
		case <-ticker.C:
			doTaskN := p.doTaskN.Load() //从 doTaskN 原子变量中加载当前【正在执行的任务数量】。
			p.doTaskN.Store(0)          //将 doTaskN 原子变量的值重置为 0。这意味着我们不关心当前正在执行的任务，只关心任务队列中的任务。
			if doTaskN < p.options.tasksBelowN {
				closeN := p.workerN.Load() - p.options.minWorkers //要关闭的工作线=前工作线程数量减去最小工作线程数量
				for closeN > 0 {                                  //当 closeN 大于 0 时，不断
					p.queue <- nil //向任务队列发送一个 nil 值。这通常表示队列的结束，并且工作线程或协程在接收到 nil 值后会退出
					closeN--
				}
			}
		}
	}
}

// Detecting illegal struct copies using `go vet`
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
