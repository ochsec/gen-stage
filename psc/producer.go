package psc

func Producer[T any](data []T, out chan<- T, ready <-chan struct{}, batchSize int) {
    index := 0
    for index < len(data) {
        <-ready  // Wait for the signal to send more items
        end := index + batchSize
        if end > len(data) {
            end = len(data)
        }
        for ; index < end; index++ {
            out <- data[index]
        }
    }
    close(out)
}
