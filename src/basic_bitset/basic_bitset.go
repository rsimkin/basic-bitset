package basic_bitset

type BasicBitset struct {
    data map[uint8]uint64
}

func (b *BasicBitset) Add(value int) {
	var base uint8 = uint8(value / 64)
	var rest uint8 = uint8(value % 64)
	var bitValue uint64 = 1 << rest

	b.data[base] |= bitValue
}

func (b *BasicBitset) Has(value int) bool {
	var base uint8 = uint8(value / 64)
	var rest uint8 = uint8(value % 64)
	var bitValue uint64 = 1 << rest

	storage, exists := b.data[base]

	if exists && storage & bitValue > 0 {
		return true
	}

	return false
}

func NewBasicBitset() *BasicBitset {
	var b BasicBitset
	b.data = make(map[uint8]uint64)
	return &b
}