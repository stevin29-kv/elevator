package main

func enqueue(queue []int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []int) (int, []int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []int{}
		return element, tmp
	}

	return element, queue[1:] // Slice off the element once it is dequeued.
}
