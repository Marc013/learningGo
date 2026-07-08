package main

// import "fmt"

// func main() {
// 	ch := make(chan string, 1)

// 	ch <- "message"

// 	fmt.Println(<-ch)
// }

// Select statement example

// func main() {
// 	ch1 := make(chan int, 1)
// 	ch2 := make(chan string, 1)

// 	ch1 <- 999
// 	ch2 <- "message"

// 	select {
// 	case msg := <-ch1:
// 		fmt.Println("Received from ch1:", msg)
// 	case msg := <-ch2:
// 		fmt.Println("Received from ch2:", msg)
// 	}
// }

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var receivedOrdersCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)

	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func(validOrderCh <-chan order, invalidOrder <-chan invalidOrder) {
	loop:
		for {  // this is an infinite for loop
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid order received: %v\n", order)
				} else {
					break loop
				}
			case order, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid order received: %v, Issue: %v\n", order.order, order.err)
				} else {
					break loop
				}
			}
		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) { // receive-only channel (in), send-only channel (out), send-only channel (errCh)
	// order := <-in
	for order := range in {
		if order.Quantity <= 0 {
			errCh <- invalidOrder{order: order, err: errors.New("quantity must be greater than zero")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}

func receiveOrders(out chan<- order) { // send-only channel
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
	close(out)
}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`, // to make this fail, change "5" to "-5"
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}
