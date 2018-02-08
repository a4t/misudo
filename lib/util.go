package misudo

import (
	"math/rand"
	"time"
)

var strings = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = strings[rand.Intn(len(strings))]
	}
	return string(b)
}
