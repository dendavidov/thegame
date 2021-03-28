package main

import (
	"flag"
	"fmt"
)

const SecretLengthQ1 = 69
const SecretLengthQ2 = 7

var (
	secret *string
)

func init() {
	secret = flag.String("s", "", "the key")
}

func main() {
	flag.Parse()

	switch len(*secret) {
	case SecretLengthQ1:
		fmt.Println("The game has started.")

		df := Decrypt(q1, *secret)

		fmt.Println(string(df))
	case SecretLengthQ2:
		df := Decrypt(q2, *secret)

		fmt.Println(string(df))

	default:
		fmt.Println("Game over.")
	}
}
