package main

import (
  "os"
  "fmt"
  "net"
  "bufio"
  "io/ioutil"

  "../imageHide"
)

const (
  HOST = "0.0.0.0"
  PORT = ":3000"
  TYPE = "tcp"
)

// to connect to this server: telnet <SERVER IP> 3000

func hendleConnection(c net.Conn) {
  buf := make([]byte, 1024)

  for {
    n, err := c.Read(buf)
    if err != nil {
      c.Close()
      break
    }

    err := ioutil.WriteFile("imgRead.bmp", message, 0644)
    if err != nil {
      return err
    }

    message, err := imageHide.readMessage()
    if err != nil {
      fmt.Println("Error while reading th image file.")
      c.Close()
      os.Exit(1)
      return
    }

    messageBytes, err := ioutil.ReadFile("messageOut.txt")
    if err != nil {
      fmt.Println("Error while reading the message file.")
      c.Close()
      os.Exit(1)
      return
    }

    messageReaded := string(messageBytes)

    fmt.Println(">>>\t", messageReaded, "\n")

    var messageWriten string
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, err := reader.ReadString('\n')
    if err != nil {
      fmt.Println("Error while reading the message input file.")
      c.Close()
      os.Exit(1)
      return
    }

    imageHide.WriteMessage([]byte(text))

    messageBytes, err := ioutil.ReadFile("imgOut.bmp")
    if err != nil {
      fmt.Println("Error while reading the image response file.")
      c.Close()
      os.Exit(1)
      return
    }

    n, err = c.Write(messageBytes)
    if err != nil {
      c.Close()
      os.Exit(1)
      break
    }
  }
}

func main() {
  if len(os.Args) < 3 {

  } else {
    // is a server

    ln, err := net.Listen(TYPE, HOST + PORT)
    if err != nil {
      os.Exit(1)
      return
    }

    for {
      conn, err := ln.Accept()
      if err != nil {
        continue
      }

      go hendleConnection(conn)
    }
  }
}
