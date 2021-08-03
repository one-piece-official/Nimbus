package iphelper_test

import (
	"testing"

	"github.com/one-piece-official/Nimbus/repository"
	"github.com/stretchr/testify/assert"

	"github.com/one-piece-official/Nimbus/iphelper"
)

func TestIPHelperChekc(t *testing.T) {
	t.Parallel()

	ipDB := map[string]map[string]string{
		"123.123.123.123": {
			iphelper.KeyCity:    "beijing",
			iphelper.KeyRegion:  "beijing",
			iphelper.KeyCountry: "中国",
		},
		"124.123.123.123": {
			iphelper.KeyCity:    "shanghai",
			iphelper.KeyRegion:  "shanghai",
			iphelper.KeyCountry: "中国",
		},
	}

	mapKV := repository.NewMapKV(make(map[string]interface{}))

	ipHelper := iphelper.NewIPHelperWithMap(ipDB, mapKV)
	testIP := "123.123.123.123"

	value, err := ipHelper.CheckIPAddressExistsInRegions([]string{"beijing"}, testIP, "include")
	assert.Equal(t, true, value)
	assert.Nil(t, err)

	value, err = ipHelper.CheckIPAddressExistsInRegions([]string{"beijing"}, testIP, "exclude")
	assert.Equal(t, false, value)
	assert.Nil(t, err)

	ipDB[testIP][iphelper.KeyCountry] = "123"
	value, err = ipHelper.CheckIPAddressExistsInRegions([]string{"beijing"}, testIP, "include")
	assert.Equal(t, false, value)
	assert.NotNil(t, err)
}
