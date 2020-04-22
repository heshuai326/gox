package gox

import "fmt"

func CountBitsSize(i int64) uint {
	var bitsSize uint = 0
	for bitsSize < 64 {
		i = i >> 1
		if i == 0 {
			break
		}
		bitsSize++
	}
	return bitsSize
}

func KeepRightBits(i int64, bitSize uint) int64 {
	return ((i >> bitSize) << bitSize) ^ i
}

func LeftMultiRight(n int64) int64 {
	bitsSize := CountBitsSize(n)
	halfSize := bitsSize / 2
	left := n >> halfSize
	right := KeepRightBits(n, halfSize)
	if right == 0 {
		right = 1
	}
	return left * right
}

type ByteUnit int64

const (
	_           = iota
	KB ByteUnit = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
)

func (b ByteUnit) HumanReadable() string {
	return ByteCountToDisplaySize(int64(b))
}

func ByteCountToDisplaySize(count int64) string {
	n := ByteUnit(count)
	if n < KB {
		return fmt.Sprintf("%d B", count)
	} else if n < MB {
		return fmt.Sprintf("%.2f KB", float64(count)/float64(KB))
	} else if n < GB {
		return fmt.Sprintf("%.2f MB", float64(count)/float64(MB))
	} else if n < TB {
		return fmt.Sprintf("%.2f GB", float64(count)/float64(GB))
	} else if n < PB {
		return fmt.Sprintf("%.2f TB", float64(count)/float64(TB))
	} else {
		return fmt.Sprintf("%.2f PB", float64(count)/float64(PB))
	}
}
