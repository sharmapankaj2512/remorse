package main

import (
	"os"	
	"fmt"
	"io/ioutil"
	"github.com/docopt/docopt-go"
	"github.com/naoina/toml"
	"github.com/sharmapankaj2512/remorse/morse"	
)

var usage = `Remorse - morse code translator.
Usage:
  remorse decode --file <file>
  remorse encode --file <file>
Options:
  -h --help     Show this screen.
  --version     Show version.`

func main() {
	const start = "$"
	codes := makeMorseCodes("morse_code.toml")
	morseTree, _ := morse.Make(morse.MorseCodes{codes.Preorder, codes.Inorder})
	arguments, _ := docopt.ParseArgs(usage, nil, "Remorse 1.0")		
	if arguments["decode"].(bool) {
		decode(morseTree, start, arguments)		
	}
	if arguments["encode"].(bool) {
		encode(morseTree, start, arguments)		
	}
}

func decode(morseTree *morse.MorseTree, start string, args map[string]interface{}) {
	if args["--file"].(bool) {
		file := args["<file>"].(string)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic("Cannot read file")
		}
		fmt.Println(morseTree.Decode(start, string(data)))
	}
}

func encode(morseTree *morse.MorseTree, start string, args map[string]interface{}) {
	if args["--file"].(bool) {
		file := args["<file>"].(string)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic("Cannot read file")
		}
		fmt.Println(morseTree.Encode(start, string(data)))
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
