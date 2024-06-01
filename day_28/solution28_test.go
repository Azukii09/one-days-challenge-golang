package day_28

import "testing"

var tests = []struct {
	description string
	input       int
	expected    int
	err         string
}{
	{
		description: "first prime",
		input:       1,
		expected:    2,
		err:         "",
	},
	{
		description: "second prime",
		input:       2,
		expected:    3,
		err:         "",
	},
	{
		description: "sixth prime",
		input:       6,
		expected:    13,
		err:         "",
	},
	{
		description: "big prime",
		input:       10001,
		expected:    104743,
		err:         "",
	},
	{
		description: "there is no zeroth prime",
		input:       0,
		expected:    0,
		err:         "there is no zeroth prime",
	},
}

func TestNth(t *testing.T) {
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := Nth(tc.input)
			switch {
			case tc.err != "":
				if err == nil {
					t.Fatalf("Nth(%d) expected error: %q, got: %d", tc.input, tc.err, actual)
				}
			case err != nil:
				t.Fatalf("Nth(%d) returned error: %v, want: %d", tc.input, err, tc.expected)
			case actual != tc.expected:
				t.Fatalf("Nth(%d) = %d, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkNth(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Nth(10001)
	}
}
