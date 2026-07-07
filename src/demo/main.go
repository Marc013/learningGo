package main

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
	go func() {
		order := <-validOrderCh
		fmt.Printf("Valid order received: %v\n", order)
		wg.Done()
	}()
	go func() {
		order := <-invalidOrderCh
		fmt.Printf("Invalid order received: %v, Issue: %v\n", order.order, order.err)
		wg.Done()
	}()
	wg.Wait()
}

func validateOrders(in, out chan order, errCh chan invalidOrder) {
	order := <-in
	if order.Quantity <= 0 {
		errCh <- invalidOrder{order: order, err:
			errors.New("quantity must be greater than zero")}
	} else {
		out <- order
	}
}

func receiveOrders(out chan order) {
	for _, rawOrder := range rawOrders {
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			log.Print(err)
			continue
		}
		out <- newOrder
	}
}

var rawOrders = []string{
	`{"ProductCode": 1111, "Quantity": 5, "Status": 1}`, // to make this fail, change "5" to "-5"
	`{"ProductCode": 2222, "Quantity": 42.3, "Status": 1}`,
	`{"ProductCode": 3333, "Quantity": 19, "Status": 1}`,
	`{"ProductCode": 4444, "Quantity": 8, "Status": 1}`,
}
