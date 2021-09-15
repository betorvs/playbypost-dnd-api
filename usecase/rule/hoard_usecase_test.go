package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllHoard(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllHoard(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"zurite"}
	res1 := GetAllHoard(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["kind"] = []string{"gemstone"}
	res2 := GetAllHoard(value2)
	assert.NotEmpty(t, res2)
}

func TestHoardByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := HoardByName("zurite")
	assert.NotEmpty(t, res)
}

func TestHoardNameList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := HoardNameList("gemstone", 10)
	assert.NotEmpty(t, res)
}
