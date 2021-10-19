package random

import (
	rrand "crypto/rand"
	"encoding/hex"
	"io"
	"math"

	"github.com/fufuok/random/utils"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 31 / letterIdxBits   // # of letter indices fitting in 31 bits
)

// Rand goroutine-safe, use Rand.xxx instead of rand.xxx
var Rand = NewRand()

// RandInt (>=)min - (<)max
func RandInt(min, max int) int {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return FastIntn(max-min) + min
}

// RandUint32 (>=)min - (<)max
func RandUint32(min, max uint32) uint32 {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return FastRandn(max-min) + min
}

// FastIntn this is similar to rand.Intn, but faster.
// A non-negative pseudo-random number in the half-open interval [0,n).
// Return 0 if n <= 0.
func FastIntn(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= math.MaxUint32 {
		return int(FastRandn(uint32(n)))
	}
	return int(Rand.Int63n(int64(n)))
}

// FastRandn this is similar to fastrand() % n, but faster.
// See https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func FastRandn(n uint32) uint32 {
	return uint32(uint64(FastRand()) * uint64(n) >> 32)
}

// RandString a random string, which may contain uppercase letters, lowercase letters and numbers.
// Ref: stackoverflow.icza
func RandString(n int) string {
	if n < 1 {
		return ""
	}
	b := make([]byte, n)
	for i, cache, remain := n-1, FastRand(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = FastRand(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return utils.B2S(b)
}

// RandHex a random string containing only the following characters: 0123456789abcdef
func RandHex(nHalf int) string {
	return hex.EncodeToString(FastRandBytes(nHalf))
}

// RandBytes random bytes
func RandBytes(n int) []byte {
	if n < 1 {
		return nil
	}

	b := make([]byte, n)
	if _, err := io.ReadFull(rrand.Reader, b); err != nil {
		return nil
	}

	return b
}

// FastRandBytes random bytes, but faster.
func FastRandBytes(n int) []byte {
	return utils.S2B(RandString(n))
}
