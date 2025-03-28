package main

import (
	"fmt"

	"github.com/pborman/uuid"
	"github.com/waditya/cryptit/decrypt"
	"github.com/waditya/cryptit/encrypt"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	encryptedString := encrypt.Nimbus("KodeKloud")
	fmt.Println("Encrypted String is : ", encryptedString)
	fmt.Println("Decrypted String is : ", decrypt.Nimbus(encryptedString))

}
