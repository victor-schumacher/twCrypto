package main

import (
	"cryptoTwitter/server/crypto"
	"fmt"
)

func main() {

	a := crypto.GetCryptoData("ETC", "BRL")

	fmt.Println(a)
}
