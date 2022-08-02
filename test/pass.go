package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	sum256 := sha256.Sum256([]byte("1hjkllllllllllllllllllllllllllllllllyghghfffffffftyeddrtsdrtdtfyfytgu"))
	s := hex.EncodeToString(sum256[:])
	fmt.Println(s)
}
