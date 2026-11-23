package objtool

import (
	"container/list"
	"sync"
)

type ThreadSafeList struct {
	List *list.List
	Lock sync.RWMutex
}

func (l *ThreadSafeList) PushBack(v interface{}) {
	l.Lock.Lock()
	l.List.PushBack(v)
	l.Lock.Unlock()
}

func (l *ThreadSafeList) PushFront(v interface{}) {
	l.Lock.Lock()
	l.List.PushFront(v)
	l.Lock.Unlock()
}

func (l *ThreadSafeList) Len() int {
	l.Lock.RLock()
	defer l.Lock.RUnlock()
	return l.List.Len()
}

func (l *ThreadSafeList) Front() *list.Element {
	l.Lock.RLock()
	defer l.Lock.RUnlock()
	return l.List.Front()
}

func (l *ThreadSafeList) Back() *list.Element {
	l.Lock.RLock()
	defer l.Lock.RUnlock()
	return l.List.Back()
}

func (l *ThreadSafeList) ToList() list.List {
	return *l.List
}
func (l *ThreadSafeList) PushBackList(list *list.List) {
	l.List.PushBackList(list)
}
func (l *ThreadSafeList) PushBackSafeList(list *ThreadSafeList) {
	l.List.PushBackList(list.List)
}

func (l *ThreadSafeList) PushBackArr(arr []interface{}) {
	lis := l.List
	if arr == nil || len(arr) == 0 {
		return
	}
	for i, size := 0, len(arr); i < size; i++ {
		lis.PushBack(arr[i])
	}
}

func (l *ThreadSafeList) ToArray() []any {
	lis := l.ToList()
	arr := make([]any, lis.Len())
	if l.Len() == 0 {
		return arr
	}
	arrayIdx := 0
	for el := lis.Front(); el != nil; el = el.Next() {
		arr[arrayIdx] = el.Value
		arrayIdx++
	}
	return arr
}

func NewThreadSafeList() *ThreadSafeList {
	return &ThreadSafeList{List: list.New()}
}
