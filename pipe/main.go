package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGPIPE)
	go func() {
		<-c // ignore
	}()

	for {
		_, err := fmt.Println("hello")
		if err != nil {
			if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
				break
			} else {
				panic(err)
			}
		}
	}
}

//func main() {
//	signal.Ignore(syscall.SIGPIPE)
//	for {
//		_, err := fmt.Println("hello")
//		if err != nil {
//			if errors.Is(err, syscall.EPIPE) {
//				break
//			} else {
//				panic(err)
//			}
//		}
//	}
//}

//func main() {
//	signal.Ignore(syscall.SIGPIPE)
//	for {
//		_, err := fmt.Println("hello")
//		if err != nil {
//			if e, ok := err.(*os.PathError); ok && e.Err == syscall.EPIPE {
//				break
//			} else {
//				panic(err)
//			}
//		}
//	}
//}

//func main() {
//	for {
//		fmt.Println("hello")
//	}
//}
