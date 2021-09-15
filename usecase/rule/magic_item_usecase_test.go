package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestGetMagicItem(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetMagicItem(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"winged-boots"}
	res1 := GetMagicItem(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["title"] = []string{"Winged Boots"}
	res2 := GetMagicItem(value2)
	assert.NotEmpty(t, res2)
}

func TestGetMagicItemByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := GetMagicItemByName("winged-boots")
	assert.NotEmpty(t, res)
}

func TestGetMagicItemByHoardTable(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := getMagicItemByHoardTable("F")
	assert.NotEmpty(t, res)
}

func TestRandomMagicItemTable(t *testing.T) {}
