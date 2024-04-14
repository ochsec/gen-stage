package psc

func ProducerConsumer[T any](in <-chan T, processor Processor[T], ready chan<- struct{}, batchSize int) <-chan T {
	processed := make(chan T)
	go Consumer(in, processed, processor, ready, batchSize)
	return processed
}
