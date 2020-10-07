package kmp

import (
	"fmt"
	"testing"
)

func TestKmp(t *testing.T) {
	kmpCase, _ := Init("xxxx.;12114AGFSzzzz")
	listCase := kmpCase.Match("aaaaxxxx.;12114AGFSzzzzddddxxxx.;12114AGFSzzzz")
	fmt.Println("kmpCase: ", listCase)

	kmpNocase, _ := Init("xxXx.;12114AGfSZzZz", true)
	listNocase := kmpNocase.Match("aaaaXXXX.;12114aGFSzzzzddddxxxx.;12114AGFSzzzz")
	fmt.Println("kmpNocase: ", listNocase)

	return
}
