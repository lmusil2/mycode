/* RZFeeser | Alta3 Research
   CHALLEGE 03 - Building complex data struct(ures) */

package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}

func main() {

    ast1 := Astro{"Megan McArthur", 35, "ISS", false}
    ast2 := Astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := Astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)
    
    //   name   slice of Astro
    p := []Astro{ast1, ast2, ast3}
    
    fmt.Println(p)
    
    // select data from the struct
    fmt.Println(p[1].mission)
    
    // initialize a nasaMission struct, "s"
    s := nasaMission{p, 3, "success"}
    
    // display "s" without fields
    fmt.Println(s)
    
    // display "s" with fields
    fmt.Printf("%+v", s)

}

