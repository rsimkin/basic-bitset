package basic_bitset

import "testing"

func TestOnEvenNumbers(t *testing.T) {
	set := fillSet()

	for i := 0; i < 2000000; i++ {
		if (i % 2 == 0) != set.Has(i) {
			t.Error("Expected ", (i % 2 == 0), " got ", set.Has(i),  " on ", i)
		}
	}
}

func TestIteration(t *testing.T) {
	set := fillSet()

	set.Iterate(func(value int) {
		if value % 2 > 0 {
			t.Error("Odd number unexpected ", value)
		}
	})
}

func fillSet() *BasicBitset {
	set := NewBasicBitset()

	for i := 0; i < 2000000; i += 2 {
		set.Add(i)
	}

	return set
}