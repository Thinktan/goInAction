package main

import (
    "log"
    "os"
    "fmt"
    _ "./matchers"
)

func init() {
    fmt.Println("pkg main init~")
    log.SetOutput(os.Stdout)
}

func main() {
    fmt.Println("Run main~")
}
