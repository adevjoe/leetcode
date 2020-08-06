package common

type Queue struct {
	values []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return len(q.values)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Put(i interface{}) {
	q.values = append(q.values, i)
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}

func (q *Queue) Touch() interface{} {
	if q.Empty() {
		return nil
	}
	return q.values[0]
}

func (q *Queue) GetIndex(i int) interface{} {
	if q.Len() <= i {
		return nil
	}
	v := q.values[i]
	q.values = append(q.values[:i], q.values[i+1:]...)
	return v
}
