package queues

import (
	"sync"
)

type Queue struct {
	lock   *sync.Mutex
	Values []string
	Count  int
}

func Init() Queue {
	return Queue{&sync.Mutex{}, make([]string, 0), 0}
}

func (q *Queue) Enqueue(x string) {
	for {
		q.lock.Lock()
		q.Values = append(q.Values, x)
		q.lock.Unlock()
		q.Count++
		return
	}
}

func (q *Queue) Dequeue() *string {
	for {
		if len(q.Values) > 0 {
			q.lock.Lock()
			x := q.Values[0]
			q.Values = q.Values[1:]
			q.lock.Unlock()
			q.Count--
			return &x
		}
		return nil
	}
	return nil
}
