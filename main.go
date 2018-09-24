package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Test")
	vocab, err := loadvocab()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%d", vocab["apple"])
}

func loadvocab() (map[string]int, error) {
	vocabFile, err := os.Open("vocab.json")
	if err != nil {
		return nil, err
	}
	defer vocabFile.Close()
	vocabdata, err := ioutil.ReadAll(vocabFile)
	if err != nil {
		return nil, err
	}
	var vocab map[string]int
	if err = json.Unmarshal(vocabdata, &vocab); err != nil {
		return nil, err
	}
	fmt.Println("", vocab)
	return vocab, nil
}
