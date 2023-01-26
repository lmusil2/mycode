/* Alta3 Research | rzfeeser@alta3.com

   To check for racing conditions use the "-race" flag
   
   Example:
   go run -race goroutines03.go
   
*/

package main

import (
    "fmt"
    "sync"
)

var (
    counter int32          // counter is a variable incremented by all goroutines.
    wg      sync.WaitGroup // wg is used to wait for the program to finish.
)

func main() {
    wg.Add(4) // Add a count of two, one for each goroutine.

    go SuperPower("controlling elements", 2)
    go SuperPower("energy projection", 5)
    go SuperPower("flight", 6)
    go SuperPower("healing", 2)

    wg.Wait() // Wait for the goroutines to finish.
    fmt.Println("Total Mutant Powers Used:", counter)

}

func SuperPower(power string, num int) {
    defer wg.Done() // Schedule the call to Done to tell main we are done.

    for i := 0; i < num; i++ {
        {
            fmt.Println(power)
            counter++
        }
    }
}

