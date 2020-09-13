package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func doSomething(i, j int) error {
	if i == 5 && j == 3 {
		return fmt.Errorf("error: goroutine(%d)", i)
	}
	time.Sleep(100 * time.Millisecond) // do something
	return nil
}

// エラーが起こってもすべてのゴルーチンの完了を待つ。
func before() {
	eg := errgroup.Group{}

	for i := 0; i < 20; i++ {
		i := i
		eg.Go(func() error {
			begin := time.Now()
			for j := 0; j < i; j++ {
				if err := doSomething(i, j); err != nil {
					fmt.Printf("goroutine(%d) -> error\n", i)
					return err
				}
			}
			fmt.Printf("goroutine(%d) -> done %v\n", i, time.Since(begin))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
}

// エラーが発生したら、他の未完了のゴルーチンは処理を中断する
func after() {
	eg, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 20; i++ {
		i := i
		eg.Go(func() error {
			begin := time.Now()
			for j := 0; j < i; j++ {
				select {
				case <-ctx.Done():
					fmt.Printf("goroutine(%d) -> canceled\n", i)
					return nil
				default:
					if err := doSomething(i, j); err != nil {
						fmt.Printf("goroutine(%d) -> error\n", i)
						return err
					}
				}
			}
			fmt.Printf("goroutine(%d) -> done %v\n", i, time.Since(begin))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
}

func main() {
	start := time.Now()
	before()
	fmt.Printf("[before] total time: %v\n", time.Since(start))

	fmt.Println("----")

	start = time.Now()
	after()
	fmt.Printf("[after] total time: %v\n", time.Since(start))
}
