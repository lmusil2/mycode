/* Alta3 Research | RZFeeser
   Map - associative data types   */

package main

import "fmt"

func main() {

    // create an empty map using "make"
    // make(map[key-type]val-type)
    m := make(map[string]int)             // m is a map

    m["key1"] = 400                 // name[key] = val
    m["key2"] = 1000                // name[key] = val

    // this will show all of the key/value pairs
    fmt.Println("map:", m)

    // return the value associated with "key", key1
    v1 := m["key1"]
    fmt.Println("v1: ", v1)

    // determine number of key/value associations
    fmt.Println("len:", len(m))

    // get rid of the "key", key2
    delete(m, "key2")
    
    // no more "key", key2
    fmt.Println("map:", m)
    
    // if you try to retireve a key
    // that does not exist
    // it will simply provide the "zero-value" of the "value" type (the zero-value of int is 0)
    fmt.Println(m["skeleton"]) // there is no key named "skeleton" within our map, so it simply writes out 0
    
    
    // that result doesn't tell us if "skeleton" didn't exist, OR, if the value stored with "skeleton" is actually 0
    

    // this operation actually returns TWO values: value, bool (where bool is if the key exists)
    // The following operations says, "save only the second return value"
    // We can use this second value can be used to determine if a key exists
    _, present := m["cookiemonster"]          // _ is the "blank identifier", it says, "don't bother assigning memory to this"
    fmt.Println("present:", present)               // present will be "false" as the key "cookiemonster" does not exist within the map

    // declare and initialize a map all one line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

