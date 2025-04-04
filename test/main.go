package main

import (
	"fmt"
	"os"

	"github.com/chanmaoganda/anydrop/common"
)

func main() {
	arg := os.Args[1]
	file, err := os.Open(arg)
	if err != nil {
		fmt.Println(err)
	}

	sha, err := common.CheckSumSha256(file)
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", sha)
}
