package add

import (
	"testing"
)

func TestAddTable(t *testing.T) {
	tests := []struct {
		sum int

		expected int
	}{
		{

			sum:      Add(10, 7),
			expected: 17,
		},
		{
			sum:      Add(10, 7),
			expected: 17,
		},
		{

			sum: Add(10, 7),

			expected: 17,
		},
	}

	for _, test := range tests {
		result := test.sum
		// assert.Equal(t, test.expected, result)
		if result != test.expected {
			t.Errorf(" your result is %v but the rights answer is %v", test.expected, result)

		}
	}
}

// test

func TestAdd(t *testing.T) {
	request := Add(20, 9)
	if request != 19 {
		t.FailNow()

	}

}
