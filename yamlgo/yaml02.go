/* RZFeeser | Alta3 Research
   Unmarshal YAML data into a Go Map  */


package main

import (
     "fmt"
     "io"
     "os"
     "log"

     "gopkg.in/yaml.v3"
)

// User struct represents one user record in the file
type User struct {
     Name       string
     Occupation string
}

func main() {

     // open the file, "villains.yaml"
     fs, err := os.Open("villains.yaml")

     // read the contents of, "villains.yaml"
     yfile, err := io.ReadAll(fs)

     // error checking
     if err != nil {
          log.Fatal(err)   // provides standard datetime formatting and then calls os.exit(1)
     }

     data := make(map[string]User)

     // We deserialize the data into the map or users
     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {

          log.Fatal(err2)     // provides standard datetime formatting and then calls os.exit(1)
     }

     for k, v := range data {

          fmt.Printf("%s: %s\n", k, v)
     }
}

