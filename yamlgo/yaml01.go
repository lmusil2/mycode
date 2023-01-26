/* RZFeeser | Alta3 Research
   Unmarshal the contents of a YAML file into a Go map.
   
   This example uses a map with "empty interfaces" (aka interface{} or any)
   for the key and value. This technique is *not* recommended, as it ignores
   all types. Typing is a good thing, in that it helps promote consistency.
   It becomes difficult to process data if the type of data is an unknown. */


package main

import (
     "fmt"
     "io"
     "log"
     "os"

     "gopkg.in/yaml.v3"
)

func main() {

     // open the file, "startrek.yaml"
     fs, err := os.Open("startrek.yaml")

     // read the contents of, "startrek.yaml"
     yfile, err := io.ReadAll(fs)

     // ensure everything worked
     if err != nil {
          log.Fatal(err)     // provides standard datetime formatting and then calls os.exit(1)
     }

     // A map in which we read the data is defined         
     // The "interface{}" is also known as the "empty interface"
     // Recently, go created the keyword, "any", which is synonymous with "interface{}"
     data := make(map[interface{}]interface{})

     // Unmarshal the data into the map
     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {

          log.Fatal(err2)    // provides standard datetime formatting and then calls os.exit(1)
     }

     // k - Key from the map (left side)
     // v - Value from the map (right side)
     for k, v := range data {
          fmt.Printf("%s -> %d\n", k, v)
     }
}

