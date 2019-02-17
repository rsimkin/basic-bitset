package main

import "basic_bitset"

func main() {
	set := basic_bitset.NewBasicBitset()

	for i := 0; i < 13000000; i += 2 {
		set.Add(i)
	}

	PrintMemUsage()
}