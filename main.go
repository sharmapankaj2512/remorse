package main

import (
	"fmt"
	"github.com/naoina/toml"
	"github.com/sharmapankaj2512/remorse/morse"
	"os"
)

func main() {
	codes := makeMorseCodes("morse_code.toml")
	fmt.Println(codes.Preorder)
	fmt.Println(codes.Inorder)
}

func makeMorseCodes(tomlFile string) morse.MorseCodes {
	f, err := os.Open(tomlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var codes morse.MorseCodes
	if err := toml.NewDecoder(f).Decode(&codes); err != nil {
		panic(err)
	}
	return codes
}
