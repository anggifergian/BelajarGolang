package udemygolangunittest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Before...")
	m.Run()
	fmt.Println("After...")
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Anggi)",
			request:  "Anggi",
			expected: "Hello Anggi",
		},
		{
			name:     "HelloWorld(Mona)",
			request:  "Mona",
			expected: "Hello Mona",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("First", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, " is Equal.")
		fmt.Println("Test FIRST is done.")
	})

	t.Run("Second", func(t *testing.T) {
		result := HelloWorld("Anggi")
		assert.Equal(t, "Hello Anggi", result, " is Equal.")
		fmt.Println("Test SECOND is done.")
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, " is Equal.")
	fmt.Println("Test HelloWorldAssertion is done.")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		t.FailNow()
		// t.Fail()
	}

	fmt.Println("TestHelloWord")
}
