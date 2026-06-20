package v09

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.stack1 = append(q.stack1, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) > 0 {
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
			q.stack1 = q.stack1[:len(q.stack1)-1]
		}
	}
	if len(q.stack2) == 0 {
		return -1
	}
	value := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return value
}
