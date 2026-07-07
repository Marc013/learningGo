# Concurrency Flow for src/demo

This branch demonstrates a small Go concurrency pipeline with goroutines and channels.

[Runtime flow in HTML][1]

## Source

- Main program: [src/demo/main.go](../src/demo/main.go)
- Types and status formatting: [src/demo/order.go](../src/demo/order.go)

## Generated Visuals

- Editable diagram (draw.io): [docs/media/goroutine-channel-flow.drawio](./media/goroutine-channel-flow.drawio)
- Diagram export used by the HTML page: [docs/media/goroutine-channel-flow.drawio.svg](./media/goroutine-channel-flow.drawio.svg)
- Interactive explanation page: [docs/media/goroutine-channel-flow.html](./media/goroutine-channel-flow.html)

## Current Program Flow (Branch-Accurate)

1. main creates three channels:
   - receivedOrdersCh (order)
   - validOrderCh (order)
   - invalidOrderCh (invalidOrder)
2. main starts two worker goroutines:
   - receiveOrders(receivedOrdersCh)
   - validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)
3. receiveOrders loops through rawOrders, unmarshals JSON, and sends each parsed order into receivedOrdersCh.
4. validateOrders reads exactly one order from receivedOrdersCh and routes it:
   - Quantity <= 0 goes to invalidOrderCh
   - Quantity > 0 goes to validOrderCh
5. Two printer goroutines wait on validOrderCh and invalidOrderCh. Only one receives a value in this version.
6. main waits with a WaitGroup and exits after the routed order is printed.

## Notes About This Branch

- validateOrders processes only the first received order because it performs one channel read and does not loop.
- After the first routed value, later values from receiveOrders are not validated in this version because validateOrders does not continue reading.
- With the current rawOrders list, the first order is valid and the program prints:

  Valid order received: ProductCode: 1111, Quantity: 5, Status: new

If you change the first order quantity from 5 to -5, the invalid path is taken instead.

[1]: https://rawcdn.githack.com/Marc013/learningGo/494c70c1cf861bb02c086cd104e53ae24f1ba7ca/docs/media/goroutine-channel-flow.html
