package queues

type IQueue interface {
	Enqueue(x string)
	Dequeue() *string
}
