package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	f func(x int) error
}

func NewTask(funcArg func(x int) error) *Task {
	return &Task{
		f: funcArg,
	}
}

type Pool struct {
	RunningWorkers int64
	Capacity       int64
	JobCh          chan *Task
	sync.Mutex
}

func NewPool(capacity int64, taskNum int) *Pool {
	return &Pool{
		Capacity: capacity,
		JobCh:    make(chan *Task, taskNum),
	}
}

func (p *Pool) GetCap() int64 {
	return p.Capacity
}

func (p *Pool) incRunning() {
	atomic.AddInt64(&p.RunningWorkers, 1)
}

func (p *Pool) decRunning() {
	atomic.AddInt64(&p.RunningWorkers, -1)
}

func (p *Pool) GetRunningWorkers() int64 {
	return atomic.LoadInt64(&p.RunningWorkers)
}

func (p *Pool) run() {
	p.incRunning()
	go func() {
		defer func() {
			p.decRunning()
		}()

		for task := range p.JobCh {
			i := 1
			err := task.f(i)
			if err != nil {
				return
			}
		}
	}()
}

func (p *Pool) AddTask(task *Task) {
	p.Lock()
	defer p.Unlock()

	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	p.JobCh <- task
}

func main() {
	pool := NewPool(3, 10)

	for i := 0; i < 20; i++ {

		pool.AddTask(NewTask(func(i int) error {
			fmt.Printf("I am running a task %d\n", i)

			return nil
		}))
	}

	time.Sleep(10 * time.Second)
}
