package basic_bitset

import "testing"

func TestOnEvenNumbers(t *testing.T) {
	set := NewBasicBitset()

	for i := 0; i < 2000000; i += 2 {
		set.Add(i)
	}

	for i := 0; i < 2000000; i++ {
		if (i % 2 == 0) != set.Has(i) {
			t.Error("Expected ", (i % 2 == 0), " got ", set.Has(i),  " on ", i)
		}
	}
}