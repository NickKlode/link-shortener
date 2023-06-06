package generator

import (
	"math/rand"
	"time"
)

const alphabet = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateToken() string {
	rand.Seed(time.Now().UnixNano())
	token := ""
	for i := 0; i < 10; i++ {
		ch := rand.Intn(len(alphabet))
		token += string(alphabet[ch])
	}
	return token
}
