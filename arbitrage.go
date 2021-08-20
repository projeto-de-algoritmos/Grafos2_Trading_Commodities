package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./rates.txt")
	check(err)
	fmt.Print(dat)

	f, err := os.Open("./rates.txt")
	check(err)

	b1 := make([]byte, 1)
	V, err := f.Read(b1)
	check(err)
	fmt.Printf("\n%s\n", string(b1[:V]))
	v, err := strconv.Atoi(string(b1[:V]))
	//fmt.Printf("v: %d", v)

	name := make([]string, v)

}
