package queue

// A queue consists of elements to be processed in a particular order
//or based on priority.
//A priority-based queue of orders is shown in the following code,
//structured as a heap. Operations such as enqueue, dequeue,
//and peek can be performed on queue. A queue is a linear data structure
//and a sequential collection.
//
//Elements are added to the end and are removed from the start of the collection.
//Queues are commonly used for storing tasks that need to be done,
//or incoming HTTP requests that need to be processed by a server.
//In real life, handling interruptions in real-time systems, call handling,
//and CPU task scheduling are good examples for using queues.
//
//The following code shows the queue of Orders and how the Queue type is defined:

// Queue—Array of Orders Type
type Queue []*Order

// Order class
type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}
