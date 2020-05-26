package queue

// Queue xxx
type Queue []int

// Push xxx
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pop xxx
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// IsEmpty xxx
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
