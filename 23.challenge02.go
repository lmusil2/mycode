/* RZFeeser | Alta3 Research
   CHALLEGE 02 - Create a slice of Astros */

package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

func main() {

    ast1 := Astro{"Megan McArthur", 35, "ISS", false}
    ast2 := Astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := Astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
    
    // slice named people made up of Astro
    p := []Astro{ast1, ast2, ast3}

    // display the slice
    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[2].mission)

}

