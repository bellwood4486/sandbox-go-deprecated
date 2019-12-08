package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func in(r io.Reader) {
	for {
		var s string
		_, err := fmt.Fscanln(r, &s)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("readed: %s\n", s)
	}
}

func out(w io.Writer) {
	signal.Ignore(syscall.SIGPIPE)

	for i := 0; i < 5; i++ {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				break
			}
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	flag := os.Args[1]
	switch flag {
	case "out":
		out(os.Stdout)
	case "in":
		in(os.Stdin)
	default:
		log.Fatal("unknown flag")
	}
}
