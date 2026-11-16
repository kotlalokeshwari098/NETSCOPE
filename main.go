package main
import (
"fmt"
"os"
"net"
)

func main(){
  //fmt.Println(os.Args)
  fmt.Println("NETSCOPE")
  fmt.Println("----------------------------")
  if len(os.Args)<2{
    fmt.Println("Please provide a hostname")
    return
  }

  hostname :=os.Args[1]

  ips, err := net.LookupHost(hostname)

  if err != nil {
    fmt.Println("Error: ",err)
     fmt.Println("Target: ", hostname)
    fmt.Println("DNS Resolution: FAIL")
    fmt.Println("Finding: Hostname could not be resolved")
     
    fmt.Println("TCP Connection: NOT EXECUTED")
    fmt.Println("TLS Handshake: NOT EXECUTED")
    fmt.Println("HTTP Request: NOT EXECUTED")
    return
  }
  fmt.Println("Target: ",hostname) 
  fmt.Println("DNS Resolution: PASS") 
  fmt.Println("Resolved IPs: ") 
  fmt.Println( " - ",ips)
}
