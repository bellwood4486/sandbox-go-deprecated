package main

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func newSomeTargeter(id uint64) vegeta.Targeter {
	type entity struct {
		ID   uint64 `json:"entityId"`
		Name string `json:"entityName"`
	}
	return func(t *vegeta.Target) (err error) {
		t.Method = "POST"
		t.URL = "http://localhost:8080/users"

		t.Body, err = json.Marshal(&entity{
			Name: "burger",
			ID:   atomic.AddUint64(&id, 1),
		})

		return err
	}
}

func main() {
	rate := vegeta.Rate{Freq: 10, Per: time.Second}
	duration := 1 * time.Second
	targeter := newSomeTargeter(0)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("%+v  \n", metrics)
	//fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
