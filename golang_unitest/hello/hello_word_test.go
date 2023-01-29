package hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// test Main
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNITEST")
	m.Run()
	// after
	fmt.Println("AFTER UNITEST")

}

// assertion

func TestHelloAssertion(t *testing.T) {
	actual := Hello("Anna")
	assert.Equal(t, "Hai Anna", actual, "Actual is Hai Anna")
	fmt.Println("extecuted")

}

// require
func TestHelloRequire(t *testing.T) {
	actual := Hello("Anna")
	require.Equal(t, "Hai Anna", actual, "Actual is Hai Anna")
	fmt.Println("extecuted")

}

func TestHello(t *testing.T) {
	actual := Hello("Jono")
	if actual != "Hai Jono" {
		// error
		t.Error("Print is not Hai Jono")

	}
	fmt.Println("TestHello Done")

}
func TestHelloType(t *testing.T) {
	result := Hello("Anna")
	if result != "Hai Anna" {
		// tanpa menyelesaikan fungsi
		t.FailNow()

	}
	fmt.Println("TestHelloType Done")

}
func TestSubTestHello(t *testing.T) {
	t.Run("Aku", func(t *testing.T) {
		actual := Hello("Aku")
		require.Equal(t, "Hai Aku", actual, "Actual is Hai Aku")
		fmt.Println("extecuted")
	})
	t.Run("Anna", func(t *testing.T) {
		actual := Hello("Anna")
		require.Equal(t, "Hai Anna", actual, "Actual is Hai Anna")
		fmt.Println("extecuted")
	})

}

func TestHelloTable(t *testing.T) {
	tests := []struct {
		name     string
		actual   string
		expected string
	}{
		{
			name:   "Hello(Ani)",
			actual: "Ani",

			expected: "Hai Ani",
		},
		{
			name:   "Hello(Uni)",
			actual: "Uni",

			expected: "Hai Uni",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Hello(test.actual)
			assert.Equal(t, test.expected, result)
		})

	}

}
// benchmarck test

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("AnJ")

	}

}

func BenchmarkHelloMonyet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("Monyet")

	}

}


// bechmark table test
