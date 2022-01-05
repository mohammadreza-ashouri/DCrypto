package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	_, err := ethclient.Dial("https://kovan.infura.io/v3/0595376196d94dbeb70356e748b82e89")

	if err != nil {

		log.Fatalf("Error!", err)

	} else {
		fmt.Println("We have the connection!")
	}
}
