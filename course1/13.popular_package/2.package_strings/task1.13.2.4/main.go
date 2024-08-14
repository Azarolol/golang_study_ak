package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	rand.Seed(time.Now().UnixNano())
	output := make([]string, 4)
	for i := 0; i < 4; i++ {
		buf := strings.Builder{}
		for i := 0; i < 4; i++ {
			r := rand.Int31n(43) + 48
			for r > 57 && r < 65 {
				r = rand.Int31n(43) + 48
			}
			buf.WriteRune(r)
		}
		output[i] = buf.String()
	}
	return strings.Join(output, "-")
}

func main() {
	activationKey := generateActivationKey()
	fmt.Println(activationKey) // UQNI-NYSI-ZVYB-ZEFQ
}
