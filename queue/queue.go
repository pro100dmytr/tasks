package queue

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zeroValue T
	if len(q.elements) == 0 {
		return zeroValue, false
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Front() (T, bool) {
	var zeroValue T
	if len(q.elements) == 0 {
		return zeroValue, false
	}
	return q.elements[0], true
}
