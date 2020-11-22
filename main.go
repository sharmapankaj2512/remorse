package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/naoina/toml"
	morse "github.com/sharmapankaj2512/remorse/morse"
	"os"	
)

var usage = `Remorse - morse code translator.
Usage:
  remorse decode <morse_code>
  remorse encode <text>  
Options:
  -h --help     Show this screen.
  --version     Show version.`

func main() {
	codes := makeMorseCodes("morse_code.toml")
	morseTree, _ := morse.Make(morse.MorseCodes{codes.Preorder, codes.Inorder})		
	arguments, _ := docopt.ParseArgs(usage, nil, "Remorse 1.0")
	if arguments["decode"] == true {						
		fmt.Println(morseTree.Decode(arguments["<morse_code>"].(string)))
	}
	if arguments["encode"] == true {				
		fmt.Println(morseTree.Encode(arguments["<text>"].(string)))
	}
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
