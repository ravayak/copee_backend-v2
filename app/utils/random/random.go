package randomutils

import (
	"math/rand"
	"time"
)

func GenerateEstimateNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(max-min) + min
	return v
}
