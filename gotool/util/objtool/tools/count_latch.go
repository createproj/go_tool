package tools

/*
 * 文件名：count_latch.go
 * 原子计数工具类
 *　@Author: Jimmy
 */
import (
	"sync"
	"time"
)

// 实体
type CountLatch struct {
	count int
	mu    sync.Mutex
	wg    sync.WaitGroup
}

// 初始化数字
func NewCountLatch(count int) *CountLatch {
	return &CountLatch{count: count}
}

// 数字递减
func (l *CountLatch) CountDown() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.count > 0 {
		l.count--
		time.Sleep(time.Second) // 模拟耗时操作
		l.wg.Done()
	}
}

// 数字递增
func (l *CountLatch) CountUp() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.count > 0 {
		l.count++
		time.Sleep(time.Second) // 模拟耗时操作
		l.wg.Done()
	}
}

// 阻塞
func (l *CountLatch) Wait() {
	l.wg.Wait()
}

// 返回数字
func (l *CountLatch) getNum() int {
	return l.count
}
