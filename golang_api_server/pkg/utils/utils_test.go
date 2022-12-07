package utils_test

import (
	"api-server/pkg/utils"
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
		utils.StringSlicesXOR([]string{}, []string{}),
	)

	// check duplicate elements
	assert.Equal(
		t,
		[]string{"1"},
		utils.StringSlicesXOR([]string{"1", "2"}, []string{"2", "2", "2"}),
	)

	// check both slice has different elements
	assert.Equal(
		t,
		[]string{"1", "3"},
		utils.StringSlicesXOR([]string{"1", "2", "2", "2", "2"}, []string{"2", "3", "2"}),
	)

	// check different order but same elements
	assert.Equal(
		t,
		[]string{},
		utils.StringSlicesXOR([]string{"1", "2"}, []string{"2", "1"}),
	)
}

func TestStringTrimAllIndent(t *testing.T) {

	fmt.Println("Test utils/utils.go")
	fmt.Println("> StringTrimAllIndent(inStr string) (string, error)")

	testStr := `
    {
      "query": {
        "term": {
          "user.id": {
            "value": "kimchy",
            "boost": 1.0
          }
        }
      }
    }
    `
	ansStr := `{"query":{"term":{"user.id":{"value":"kimchy","boost":1.0}}}}`

	resStr, err := utils.StringTrimAllIndent(testStr)
	assert.NoError(t, err)
	assert.Equal(t, ansStr, resStr)
}
