package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {

	note := "This is a recursion call function"
	fmt.Println(len(note))

	for i := 0; i < len(note); i++ {
		fmt.Printf("%x", note[i])

	}
	decoded, err := hex.DecodeString("54686973206973206120726563757273696f6e2063616c6c2066756e6374696f6e")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", decoded)

}
