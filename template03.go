/* Alta3 Research | RZFeeser
   Template - Using a Map to build a template  */

package main

import (
        "html/template"
        "io"
        "log"
        "os"
        "strings" // we need to this trim the extension off our template
)

func parse(path string) {

        // template.ParseFiles() will READ a template file
        t, err := template.ParseFiles(path)

        // check to see if our template was read correctly
        if err != nil {
                log.Print(err)
                return
        }

        // get rid of ".template" on our path name
        path = strings.TrimRight(path, ".template")

        // create a file "just" named example.css (not example.css.template)
        f, err := os.Create(path)
        if err != nil {
                log.Println("create file: ", err)
                return
        }

        // A sample config
        config := map[string]string{
                "textColor":      "#abcdef",
                "linkColorHover": "#ffaacc",
        }

        // create our file using the fill in the blank info we have within our map
        err = t.Execute(f, config)
        if err != nil {
                log.Print("execute: ", err)
                return
        }
        f.Close()
}

func main() {

        // call our function, "parse()" and tell it to use the template we just made
        parse("example.css.template")

        // try opening the finished file we just created (example.css)
        f, _ := os.Open("example.css")
        
        // display the content to the screen
        io.Copy(os.Stdout, f)
        
        // close the open file (best practice)
        f.Close()
}

