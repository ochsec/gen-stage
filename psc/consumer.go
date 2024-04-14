package psc

func Consumer[T any](in <-chan T, processed chan<- T, processor Processor[T], ready chan<- struct{}, batchSize int) {
    count := 0
    for item := range in {
        processed <- processor(item)
        count++
        if count == batchSize {
            count = 0
            select {
            case ready <- struct{}{}:  // Signal that it's ready for more items
            default:
            }
        }
    }
    close(processed)
    if count > 0 {  // Ensure to signal if there were remaining items less than a batch size
        close(ready)
    }
}
