package sim

import (
	"container/heap"
	"container/list"
	"sync"
//	"fmt"
//	"log"
)

// EventQueue are a queue of event ordered by the time of events
type EventQueue interface {
	Push(evt Event) bool
	Pop() Event
	Len() int
	Peek() Event
}

// EventQueueImpl provides a thread safe event queue
type EventQueueImpl struct {
	sync.Mutex
	events eventHeap
}

// NewEventQueue creates and returns a newly created EventQueue
func NewEventQueue() *EventQueueImpl {
	q := new(EventQueueImpl)
	q.events = make([]Event, 0)
	heap.Init(&q.events)
	return q
}

// Push adds an event to the event queue
func (q *EventQueueImpl) Push(evt Event) bool {
	q.Lock()
	heap.Push(&q.events, evt)
	q.Unlock()
    return true
}

// Pop returns the next earliest event
func (q *EventQueueImpl) Pop() Event {
	q.Lock()
	e := heap.Pop(&q.events).(Event)
	q.Unlock()
	return e
}

// Len returns the number of event in the queue
func (q *EventQueueImpl) Len() int {
	q.Lock()
	l := q.events.Len()
	q.Unlock()
	return l
}

// Peek returns the event in front of the queue without removing it from the
// queue
func (q *EventQueueImpl) Peek() Event {
	q.Lock()
	evt := q.events[0]
	q.Unlock()
	return evt
}

type eventHeap []Event

// Len returns the length of the event queue
func (h eventHeap) Len() int {
	return len(h)
}

// Less determines the order between two events. Less returns true if the i-th
// event happens before the j-th event.
func (h eventHeap) Less(i, j int) bool {
	return h[i].Time() < h[j].Time()
}

// Swap changes the position of two events in the event queue
func (h eventHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an event into the event queue
func (h *eventHeap) Push(x interface{}) {
	event := x.(Event)
	*h = append(*h, event)
}

// Pop removes and returns the next event to happen
func (h *eventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	event := old[n-1]
	*h = old[0 : n-1]
	return event
}

// InsertionQueue is a queue that is based on insertion sort
type InsertionQueue struct {
	lock sync.RWMutex
	l    *list.List
    lastevt Event
}

// NewInsertionQueue returns a new InsertionQueue
func NewInsertionQueue() *InsertionQueue {
	q := new(InsertionQueue)
	q.l = list.New()
    q.lastevt = nil
	return q
}

// Push add an event to the event queue
func (q *InsertionQueue) Push(evt Event) bool{
	if q.lastevt == nil {
	    q.lastevt = evt
		q.l.PushBack(evt)
    } else {
        //log.Printf("size %d %.2f %.2f\n",q.Len(),q.lastevt.Time()*1e9,evt.Time()*1e9)
        if q.lastevt.Time() <= evt.Time() {
		    q.l.PushBack(evt)
            q.lastevt=evt
        } else {
            return false
        }
	}
    return true
}

// Pop returns the event with the smallest time, and removes it from the queue
func (q *InsertionQueue) Pop() Event {
	evt := q.l.Remove(q.l.Front())
    if q.l.Len() == 0 {
        q.lastevt = nil
    }
	return evt.(Event)
}

// Len return the number of events in the queue
func (q *InsertionQueue) Len() int {
//	q.lock.RLock()
	l := q.l.Len()
//	q.lock.RUnlock()
	return l
}

// Peek returns the event at the front of the queue without removing it from
// the queue.
func (q *InsertionQueue) Peek() Event {
//	q.lock.RLock()
	evt := q.l.Front().Value.(Event)
//	q.lock.RUnlock()
	return evt
}
