package queues

type Queues struct {
	items []int
}

// enqueues
func (q *Queues) Enqueues(i int) {
	q.items = append(q.items, i)
}

// deenqueues
func (q *Queues) Dequeues() {
	// toRemove := q.items[0]
	q.items = q.items[1:]
	// return toRemove
}
