/* Alta3 Research | RZFeeser@alta3.com
   SOLUTION 01 */
   
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// seed some entropy into the program
func init(){
    // rand.Seed(1)  // randomness has this "seed" value by default
    rand.Seed(time.Now().Unix()) 
}


func main(){

    // create a slice of strings
    instructorSlice := []string{"Jane", "Zach", "Rey", "Mike", "Alice"}

    // take a random selection from our instructorSlice
    // x := instructorSlice[rand.Intn(len(instructorSlice))]

    // x is local in scope to the if-else-if-else block
    // when we declare and initialize in the same line as the "if" keyword
    if x := instructorSlice[rand.Intn(len(instructorSlice))]; x == "Zach" {
        fmt.Println("Zach works out of the NorthEast")
    } else if x == "Jane" {
        fmt.Println("Jane lives in the Pacific NorthWest")
    } else {
        fmt.Println("Hmm, I don't know much about the instructor,", x)
    }

    // this should error, x is local to the if-else-if-else block
    //fmt.Println(x)

}

