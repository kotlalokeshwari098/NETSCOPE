package main
import (
"fmt"
"os"
)

func main(){
  fmt.Println(os.Args)

  if len(os.Args)<2{
    fmt.Println("Please provide a hostname")
    return
  }

  hostname :=os.Args[1]
  fmt.Println("Target :",hostname)
}
