package main

import (
	"errors"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/bellwood4486/sandbox-go/vegeta-lib/sampler"

	"github.com/mroth/weightedrand"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	sampler, ok := samplers()["service1"]
	if !ok {
		log.Fatal("sampler not found")
	}

	ts, err := sampler.Templates()
	if err != nil {
		log.Fatal("sampler not found")
	}

	rand.Seed(time.Now().UnixNano())
	var choices []weightedrand.Choice
	for _, t := range ts {
		choices = append(choices, weightedrand.Choice{Item: t, Weight: t.Weight})
	}
	chooser := weightedrand.NewChooser(choices...)

	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 2 * time.Second
	targeter := newWeightedRandomTargeter(chooser)
	attacker := vegeta.NewAttacker()

	enc := vegeta.NewEncoder(os.Stdout)
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		if err := sampler.HandleResponse(res.Method, res.URL, res.Code, res.Body); err != nil {
			log.Fatal(err)
		}
		if err := enc.Encode(res); err != nil {
			log.Fatal(err)
		}
	}
}

func samplers() map[string]sampler.Sampler {
	s := make(map[string]sampler.Sampler)

	s["service1"] = &sampler.Service1Sampler{}
	s["service2"] = &sampler.Service1Sampler{}
	s["service3"] = &sampler.Service1Sampler{}
	s["service4"] = &sampler.Service1Sampler{}
	s["service5"] = &sampler.Service1Sampler{}

	return s
}

func newWeightedRandomTargeter(chooser weightedrand.Chooser) vegeta.Targeter {
	var mu sync.Mutex
	return func(tgt *vegeta.Target) error {
		mu.Lock()
		defer mu.Unlock()

		tmpl, ok := chooser.Pick().(sampler.RequestTemplate)
		if !ok {
			return errors.New("not tmpl")
		}
		tgt.Method = tmpl.Method
		tgt.URL = tmpl.URL
		tgt.Header = tmpl.Header
		tgt.Body = []byte(tmpl.Body)

		return nil
	}
}
