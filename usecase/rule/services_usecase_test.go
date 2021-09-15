package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllServices(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllServices(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"poor"}
	res1 := GetAllServices(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["source"] = []string{"lifestyle"}
	res2 := GetAllServices(value2)
	assert.NotEmpty(t, res2)
}

func TestServiceByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ServiceByName("poor")
	assert.NotEmpty(t, res)
}

func TestServicesNameList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ServicesNameList()
	assert.NotEmpty(t, res)
}
