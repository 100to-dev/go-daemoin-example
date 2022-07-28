package core

import (
	"math"
	"math/rand"
)

type CoinSide int

const (
	Head CoinSide = iota
	Tails
	Middle
)

type CoinFlipMethod func() CoinSide

func DefaultCoinFlip() CoinSide {
	n := 10

	middle := int(math.Ceil(float64(n) / 2))

	if p := rand.Intn(n); p < middle {
		return Head
	} else if p >= middle && p < n-1 {
		return Tails
	} else {
		return Middle
	}
}
