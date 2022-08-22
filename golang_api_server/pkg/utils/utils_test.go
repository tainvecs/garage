package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlicesXOR(t *testing.T) {

	fmt.Println("Test utils/utils.go")
	fmt.Println("> StringSlicesXOR(ss1, ss2 []string) []string")

	// check empty
	assert.Equal(
		t,
		[]string{},
		StringSlicesXOR([]string{}, []string{}),
	)

	// check duplicate elements
	assert.Equal(
		t,
		[]string{"1"},
		StringSlicesXOR([]string{"1", "2"}, []string{"2", "2", "2"}),
	)

	// check both slice has different elements
	assert.Equal(
		t,
		[]string{"1", "3"},
		StringSlicesXOR([]string{"1", "2", "2", "2", "2"}, []string{"2", "3", "2"}),
	)

	// check different order but same elements
	assert.Equal(
		t,
		[]string{},
		StringSlicesXOR([]string{"1", "2"}, []string{"2", "1"}),
	)
}
