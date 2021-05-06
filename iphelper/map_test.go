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
			iphelper.KeyCity:   "beijing",
			iphelper.KeyRegion: "beijing",
		},
		"124.123.123.123": {
			iphelper.KeyCity:   "shanghai",
			iphelper.KeyRegion: "shanghai",
		},
	}

	mapKV := repository.NewMapKV(make(map[string]interface{}))

	ipHelper := iphelper.NewIPHelperWithMap(ipDB, mapKV)

	include, err := ipHelper.CheckIPAddressExistsInRegions([]string{"beijing"}, "123.123.123.123", "include")
	assert.Equal(t, true, include)
	assert.Nil(t, err)

	exclude, err := ipHelper.CheckIPAddressExistsInRegions([]string{"beijing"}, "123.123.123.123", "exclude")
	assert.Equal(t, false, exclude)
	assert.Nil(t, err)
}
