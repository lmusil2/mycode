/* Alta3 Research | RZFeeser
   Template - HTML template  */

package main


import (
    "html/template"
    "os"
)

// This represents a single task that appears within our slice of tasks
type Todo struct {
    Title string
    Done  bool
}

// this represents the page data we need to render a finished TODO list
type TodoPageData struct {
    PageTitle string
    Todos     []Todo     // slice of the type Todo (we created this)
}


func main() {

    // check our template for potential errors with Must
    tmpl := template.Must(template.ParseFiles("index.html"))
    
        // this is the data we will use to render the page
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Laundry", Done: false},
                {Title: "Pull weeds in the garden", Done: true},
                {Title: "Sweep the stairs", Done: true},
            },
        }
        tmpl.Execute(os.Stdout, data)
}

