# gen-stage

An attempt at an implementation of Elixir's [GenStage](https://hexdocs.pm/gen_stage/GenStage.html) in Go

Example:

Functions:
- ateFn: fruit -> "ate: {fruit}":
- titleFn: fruit -> Fruit

Agents:
- Producer: ["apple", "banana", "cherry", "date"]
- ProducerConsumer: (fruit, titleFn, ready, batchSize) -> ["Apple", "Banana", "Cherry", "Date"]
- Consumer: (Fruit, ateFn, ready, batchSize) -> ["ate: Apple", "ate: Banana", "ate: Chery", "ate: Date"]

Consumer signals producer that it's ready for more items with a boolean channel. ProducerConsumer passes this channel to the Consumer.
