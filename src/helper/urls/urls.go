package urls

import (
	"math/rand"
	"time"
)

// Generate creates count number of random URLs
// of length length and return its collection
func Generate(count, length int) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	url := make([]string, count)
	avail := []rune("abcdefghijklmnopqrstuvwxyz")
	for ctr := 0; ctr < count; ctr++ {
		temp_url := make([]rune, length)
		for i := range temp_url {
			temp_url[i] = avail[rand.Intn(len(avail))]
		}
		url[ctr] = "www." + string(temp_url) + ".com"
	}
	return url
}
