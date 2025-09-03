package bit_flags

type BitFlags map[uint8]uint64

func NewBitFlags() BitFlags {
	return make(BitFlags)
}

func SetFlags(bitFlags BitFlags, flags ...uint8) map[uint8]uint64 {
	bitFlags = validateBitFlags(bitFlags)

	var flvl uint8
	for _, flag := range flags {
		flvl = fLvl(flag)

		var bit uint64
		if b, ok := bitFlags[flvl]; ok {
			bit = b
		}

		bitFlags[flvl] = 1<<fbp(flag) | bit
	}

	return bitFlags
}

func RemoveFlags(bitFlags BitFlags, flags ...uint8) map[uint8]uint64 {
	bitFlags = validateBitFlags(bitFlags)

	var flvl uint8
	for _, flag := range flags {
		flvl = fLvl(flag)
		var bit uint64
		if b, ok := bitFlags[flvl]; ok {
			bit = b
		}

		bitFlags[flvl] = bit &^ (1 << fbp(flag))
	}

	return bitFlags
}

func HasFlag(bitFlags BitFlags, flag uint8) bool {

	if bit, ok := bitFlags[fLvl(flag)]; ok {
		return bit&(1<<fbp(flag)) != 0
	}

	return false
}

// fLvl return flag level
func fLvl(p uint8) uint8 {
	return (p + 62) / 63
}

// fbp flag bit position
func fbp(f uint8) uint8 {
	f = f % 63

	if f == 0 {
		f = 63
	}

	return f
}

func validateBitFlags(m BitFlags) map[uint8]uint64 {
	if m == nil {
		return make(map[uint8]uint64)
	}
	return m
}
