/* RZFeeser | Alta3 Research
   CHALLEGE 01 - Create a struct named Astro */

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

}

