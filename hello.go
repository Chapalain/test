package main

import (
	"fmt"
        "flag"
)

// Deal with input's request
var msg string

func init() {
        const (
                default_message = "Hello World!"
        )
        flag.StringVar(&msg, "msg", default_message, "message to be displaying")

}

func main() {
        flag.Parse()

	fmt.Println(msg)
}

