/* RZFeeser | Alta3 Research
   Creating text files
   See documentation @ https://pkg.go.dev/bufio  */       
     
package main
 
import (
    "bufio"
    "log"
    "os"
)
 
func main() {

    // create a slice of strings
    sampledata := []string{"West of House.",
        "You are standing in an open field west of a white house,",
        "with a boarded front door.",
        "There is a small mailbox here.",
    }

    // os.OpenFile() - You can specify any flag when opening a file based on your requirements.
    // If you call the OpenFile() method with the O_CREATE flag if the file does not exist, 
    // it is created with the mode permission before umask.
    file, err := os.OpenFile("ZorkOpening.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
 
    datawriter := bufio.NewWriter(file)
 
    for _, data := range sampledata {
        _, _ = datawriter.WriteString(data + "\n")
    }
 
    datawriter.Flush()
    file.Close()
}

