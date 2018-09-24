package main

import (
	"fmt"
	"testing"
)

func TestAnnotate(t *testing.T) {
	resp, err := getTextAnnotation()
	if err != nil {
		t.Fail()
	}
	fmt.Println(resp)
}
