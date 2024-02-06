package todo

type queue[T any] struct {
	items []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		items: make([]T, 0),
	}
}

func (q *queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) Pop() T {
	// if len(q.items) == 0 {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
