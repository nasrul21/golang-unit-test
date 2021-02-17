package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nasrul")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Nasrul",
			request:  "Nasrul",
			expected: "Hello Nasrul",
		},
		{
			name:     "Faizin",
			request:  "Faizin",
			expected: "Hello Faizin",
		},
		{
			name:     "Muslim",
			request:  "Muslim",
			expected: "Hello Muslim",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubtest(t *testing.T) {
	t.Run("Nasrul", func(t *testing.T) {
		result := HelloWorld("Nasrul")
		require.Equal(t, "Hello Nasrul", result, "Result must be 'Hello Nasrul'")
	})

	t.Run("Faizin", func(t *testing.T) {
		result := HelloWorld("Faizin")
		require.Equal(t, "Hello Faizin", result, "Result must be 'Hello Faizin'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Cannot run on Linux")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}
