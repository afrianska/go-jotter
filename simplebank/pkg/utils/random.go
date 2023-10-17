// ./random.go
package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const consonant = "bcdfghjklmnpqrstvwxyz"
const vocal = "aiueo"

func init()  {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt to generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString to generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	c := len(consonant)
	v := len(vocal)
	a := len(alphabet)
	
	for i := 0; i < n; i++ {
		var r byte

		if x := rand.Intn(3); x == 1 {
			r = consonant[rand.Intn(c)]
		} else if x := rand.Intn(3); x == 2 {
			r = vocal[rand.Intn(v)]
		} else {
			r = alphabet[rand.Intn(a)]
		}
		sb.WriteByte(r)
	}

	return sb.String()
}

// RandomOwner to generate a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney to generate a random amount of balance
func RandomMoney() int64 {
	return RandomInt(0, 1500)
}

// RandomCurrency to generate a random of currency code
func RandomCurrency() string {
	currency := []string{"USD", "EUR", "CAD", "GBP"}
	n := len(currency)
	return currency[rand.Intn(n)]
}