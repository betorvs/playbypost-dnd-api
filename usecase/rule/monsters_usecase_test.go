package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestMonsterForNPC(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := MonsterForNPC(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"kobold"}
	res1 := MonsterForNPC(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["xp"] = []string{"25"}
	res2 := MonsterForNPC(value2)
	assert.NotEmpty(t, res2)
	value3 := make(url.Values)
	value3["type"] = []string{"humanoid"}
	res3 := MonsterForNPC(value3)
	assert.NotEmpty(t, res3)
}

func TestMosterByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := MosterByName("kobold")
	assert.NotEmpty(t, res)
}

func TestReturnMosterByType(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := returnMosterByType("humanoid")
	assert.NotEmpty(t, res)
}

func TestReturnMonsterPaladinEnemy(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res1 := returnMonsterPaladinEnemy("dretch")
	assert.True(t, res1)
	res2 := returnMonsterPaladinEnemy("skeleton")
	assert.True(t, res2)
	res3 := returnMonsterPaladinEnemy("kobold")
	assert.False(t, res3)
}

func TestMonsterTypeSingular(t *testing.T) {
	for _, v := range monsterTypeList() {
		res := monsterTypeSingular(v)
		assert.NotEmpty(t, res)
	}
	test := monsterTypeSingular("unknown")
	assert.Empty(t, test)
}
