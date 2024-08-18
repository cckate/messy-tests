package main

import (
	"flag"
	"fmt"

	"github.com/knights-analytics/tokenizers"
)

var flagString = flag.String("input", "", "input string that need to be encode into tokens")

func main() {
	flag.Parse()

	tk, err := tokenizers.FromFile("./qwen2/tokenizer.json")
	if err != nil {
		panic(err)
	}
	// release native resources
	defer tk.Close()

	fmt.Println("Vocab size:", tk.VocabSize())
	println("input: ", *flagString)

	retInt, retString := tk.Encode(*flagString, true)
	println("length: ", len(retString))

	for _, word := range retInt {
		println(tk.Decode([]uint32{word}, true))
	}
	println("decode:", tk.Decode(retInt, true))
}
