package psc

type Processor[T any] func(T) T
