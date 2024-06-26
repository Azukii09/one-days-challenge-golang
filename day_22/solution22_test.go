package day_22

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	description string
	limit       int
	expected    []int
}{
	{
		description: "no primes under two",
		limit:       1,
		expected:    []int{},
	},
	{
		description: "find first prime",
		limit:       2,
		expected:    []int{2},
	},
	{
		description: "find primes up to 10",
		limit:       10,
		expected:    []int{2, 3, 5, 7},
	},
	{
		description: "limit is prime",
		limit:       13,
		expected:    []int{2, 3, 5, 7, 11, 13},
	},
	{
		description: "find primes up to 1000",
		limit:       1000,
		expected:    []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919, 929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997},
	},
}

func TestSieve(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Sieve(tc.limit)
			// use len() to allow either nil or empty list, because they are not equal by DeepEqual
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Sieve(%d)\n got:%#v\nwant:%#v", tc.limit, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSieve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sieve(tc.limit)
		}
	}
}
