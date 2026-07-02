package main

import (
	"fmt"
	"time"
)

// Calculate Total Sales
func main() {
	startTime := time.Now() // Record the current time
	var totalSales int
    // Creates an unbuffered channel of type int
    salesCh := make(chan int)
    // Create a channel for error communication
    errorCh := make(chan error)

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSales(salesRegion, salesCh, errorCh)
	}

    // Allow goroutines to complete
    // Receive total sales from each goroutine asynchronously
    for i := 0; i < len(salesData); i++ {
      select {
       case total := <-salesCh:
        totalSales += total
       case err := <-errorCh:
        fmt.Printf("Error: %v\n", err)
       case <-time.After(3000 * time.Millisecond):
        fmt.Println("Timeout occurred while waiting for sales data")
        return
       }
    }

	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, salesCh chan<- int, errorCh chan<- error) {
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
        if storeSales < 0 {
          errorCh <- fmt.Errorf("SKIPPED REGION! Found store with sale value [%d] less than 0. Please recheck data for region", storeSales)
          return
        }
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}

	// Post Calculation of regionTotal
	salesCh <- regionTotal
}


/* ---- example of error propagation using channels ----
// Create a channel for error communication
errCh := make(chan error)

func doTask(resultCh chan int, errCh chan error) {
 resultCh <- 42              // Simulate a calculation
 errCh <- errors.New("something went wrong") // Simulate an error
}

func main() {
 resultCh := make(chan int)  // Channel for result
 errCh := make(chan error)   // Channel for error

go doTask(resultCh, errCh)  // Start the task goroutine

 select {
  case result := <-resultCh:
   fmt.Println("Result:", result)
  case err := <-errCh:
   fmt.Println("Error:", err)
  }
}
*/


/*
In this example:

The `doTask` function sends the result, `42`, and an error, `("something went wrong")`, directly to their respective channels.
Channels for both the result and the error communication are created.
The `main` function starts the task goroutine using `go doTask(resultCh, errCh)`.
The select statement handles either receiving the result or the error from their respective channels and prints them accordingly.
*/
