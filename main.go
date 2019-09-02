package main

import (
	"github.com/rsimkin/basic-bitset/bitset"
	"fmt"
)

func main() {
	set := bitset.NewBasicBitset()

	for i := 0; i < 13000000; i += 2 {
		set.Add(i)
	}

	set.Iterate(func(value int) {
		fmt.Println(value)
	})

	PrintMemUsage()
}
