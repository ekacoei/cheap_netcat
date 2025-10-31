/* Cheap Netcat */

package main

import "fmt"
import "net"
import "os"
import "io"

func help() {
  fmt.Println( "Usage:\n   cheap_netcat <PORT>")
  os.Exit(1)
}

func main() {
  if len( os.Args) != 2 {
    help()
  }
  ln, err := net.Listen("tcp", ":" + os.Args[1])
  defer ln.Close()
  if err != nil {
//  	fmt.Fprintln (os.Stderr, "Fehler beim Socket öffnen")
  	fmt.Fprintln (os.Stderr, "Error opening Socket")
    os.Exit(1)
  }
  for {
  	conn, err := ln.Accept()
  	if err != nil {
//  		fmt.Fprintln (os.Stderr, "Fehler beim Client via TCP annehmen, setze mit nächstem Client fort...")
  		fmt.Fprintln (os.Stderr, "Failed to accept a TCP client, but will continue as soon as the next TCP connection comes in...")
  	}
    go func( c net.Conn) {
      for {
        io.Copy( os.Stdout, c)
      }
    }( conn)
  }
}