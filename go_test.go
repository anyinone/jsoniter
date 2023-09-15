package jsoniter

import (
	"fmt"
	"testing"
)

type AAA struct {
	A int `json:"a"`
}

func TestXXX(t *testing.T) {
	a := AAA{}
	UnmarshalFromString("{\"a\":\"-10000\"}", &a)
	fmt.Print(a)
}
