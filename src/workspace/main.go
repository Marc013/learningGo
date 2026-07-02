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

	for _, salesRegion := range salesData {
		// Invoke to calculate Total for Each Region
		go calculateRegionSales(salesRegion, salesCh)
	}

    // Allow goroutines to complete
    // Receive total sales from each goroutine asynchronously
    for i := 0; i < len(salesData); i++ {
      totalSales += <- salesCh
    }

	fmt.Printf("Total %d , Time taken to calculate %s \n", totalSales, time.Since(startTime))
}

// Function to Calculate Total Sales per Region
func calculateRegionSales(salesRegion []int, salesCh chan <- int) {
	regionTotal := 0
	for _, storeSales := range salesRegion {
		// Calculate region total
		regionTotal += storeSales
		time.Sleep(100 * time.Millisecond)
	}

	// Post Calculation of regionTotal
	salesCh <- regionTotal
}


/* ---- Channel Direction ----
You can specify the direction of a channel in its type signature to restrict
its usage to sending or receiving operations. This helps enforce
communication protocols and prevent misuse of channels. Here's how you
specify channel direction:

func sendData(ch chan<- int) {
  // Send data into channel
}

func receiveData(ch <-chan int) {
  // Receive data from channel
}
*/

/* ---- A simple example of using channel direction in Go ----

func sendData(ch chan<- int) {
 // Send data into the channel
 ch <- 10
 ch <- 20
 ch <- 30
 close(ch) // Close the channel after sending all values
}

func main() {
 // Create an unbuffered channel of type int
 ch := make(chan int)

 // Start a goroutine to send data into the channel
 go sendData(ch)

 // Receive data from the channel
 for {
   // Attempt to receive a value from the channel
   value, ok := <-ch
   if !ok {
   // Channel closed, exit the loop
   break
  }
  // Print the received value
  fmt.Println("Received:", value)
  }
}
*/

/* ---- Output ----
Received: 10
Received: 20
Received: 30
*/
