package main

func main() {
	set := NewBasicBitset()

	for i := 0; i < 13000000; i += 2 {
		set.Add(i)
	}

	PrintMemUsage()
}