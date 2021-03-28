package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecrypt(t *testing.T) {
	p2 := Decrypt([]byte{161, 192, 50, 205, 249, 209, 73, 41, 244, 215, 130, 154, 112, 94, 2, 134, 3, 34, 67, 103, 135, 211, 131, 83, 215, 95, 62, 40, 108, 245, 214, 25, 176, 165, 236, 80, 87, 105, 80, 38, 141, 188}, "passphrase")

	assert.Equal(t, "Awesome phrase", string(p2), "they should be equal")
}
