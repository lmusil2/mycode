/* RZFeeser | Alta3 Research
   HTTP GET and parameters */

package main

import (
        "fmt"
        "io"
        "log"
        "net/http"
        "net/url"
)

func main() {

        // define two string variables within the main function block
        var name string = "John Doe"
        var occupation string = "gardener"

        // url.QueryEscape makes your string "URI safe" (i.e. replaces white space with %20 or +)
        params := "name=" + url.QueryEscape(name) + "&" +
                "occupation=" + url.QueryEscape(occupation)
                
        path := fmt.Sprintf("https://httpbin.org/get?%s", params)

        resp, err := http.Get(path)

        if err != nil {
                // on an error, the error will be wrapped in a logging timestamp and send to std out
                // followed by a call to os.Exit(1)
                log.Fatal(err)
        }

        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)

        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(string(body))
}

