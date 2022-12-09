package strutils_test

import (
	"apisrv/pkg/utils/strutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlicesXOR(t *testing.T) {

	// check empty
	assert.Equal(
		t,
		[]string{},
		strutils.StringSlicesXOR([]string{}, []string{}),
	)

	// check duplicate elements
	assert.Equal(
		t,
		[]string{"1"},
		strutils.StringSlicesXOR([]string{"1", "2"}, []string{"2", "2", "2"}),
	)

	// check both slice has different elements
	assert.Equal(
		t,
		[]string{"1", "3"},
		strutils.StringSlicesXOR([]string{"1", "2", "2", "2", "2"}, []string{"2", "3", "2"}),
	)

	// check different order but same elements
	assert.Equal(
		t,
		[]string{},
		strutils.StringSlicesXOR([]string{"1", "2"}, []string{"2", "1"}),
	)
}

func TestTrimAllIndent(t *testing.T) {

	ansStr := `{"query":{"term":{"user.id":{"value":"kimchy","boost":1.0}}}}`
	ansStr = strutils.TrimAllIndent(ansStr)

	testStr := `
        {
            "query":{
                "term":{
                    "user.id":{
                        "value":"kimchy",
                        "boost":1.0
                    }
                }
            }
        }
    `
	resStr := strutils.TrimAllIndent(testStr)

	assert.Equal(t, ansStr, resStr)
}
