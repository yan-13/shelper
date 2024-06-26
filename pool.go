package shelper

import (
	"log"
	"runtime"
)

// Task 任务接口
type Task interface {
	Execute()
}

// Pool 协程池
type Pool struct {
	TaskChannel chan Task // 任务队列
}

var poolInstance *Pool

func GetPoolInstance() *Pool {
	if nil == poolInstance {
		poolInstance = NewPool()
		log.Println("任务池已启动")
	}
	return poolInstance
}

// NewPool 创建一个协程池
func NewPool() *Pool {
	//cpu num
	cpuNum := runtime.NumCPU()
	log.Println("cpu num", cpuNum)

	p := &Pool{TaskChannel: make(chan Task, cpuNum)}

	// 创建指定数量 worker 从任务队列取出任务执行
	for i := 0; i < cpuNum; i++ {
		go func() {
			for {
				task := <-p.TaskChannel
				task.Execute()
			}
		}()
	}
	return p
}

// Submit 提交任务
func (p *Pool) Submit(t Task) {
	p.TaskChannel <- t
}
