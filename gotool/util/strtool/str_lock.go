// Package strtool, 文件名： str_lock.go， 字符串锁
package strtool

import (
	"sync"
)

// StringLocks 用于管理字符串锁的结构体
type StringLocks struct {
	locks map[string]*sync.Mutex
	mutex sync.Mutex
}

var StringLocksGlobal = NewStringLocks()

// NewStringLocks 创建一个新的 StringLocks 实例
func NewStringLocks() *StringLocks {
	return &StringLocks{
		locks: make(map[string]*sync.Mutex),
	}
}

// Lock 锁定一个字符串标识的资源
func (sl *StringLocks) Lock(key string) {
	sl.mutex.Lock()
	if _, exists := sl.locks[key]; !exists {
		sl.locks[key] = &sync.Mutex{}
	}
	sl.mutex.Unlock()

	sl.locks[key].Lock()
}

// TryLock 锁定一个字符串标识的资源。如果没获取到锁，则返回失败
func (sl *StringLocks) TryLock(key string) bool {
	sl.mutex.Lock()
	if _, exists := sl.locks[key]; !exists {
		sl.locks[key] = &sync.Mutex{}
	}
	sl.mutex.Unlock()

	return sl.locks[key].TryLock()
}

// Unlock 解锁一个字符串标识的资源
func (sl *StringLocks) Unlock(key string) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	sl.locks[key].Unlock()
}
