package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/stretchr/testify/assert"
)

func TestUsePotion(t *testing.T) {
	test1 := new(rule.Potion)
	test1.Name = "amazing-pills"
	res1, err1 := UsePotion(test1)
	assert.NotEmpty(t, res1)
	assert.Error(t, err1)
	test1.Name = "potion-of-climbing"
	res2, err2 := UsePotion(test1)
	assert.NotEmpty(t, res2)
	assert.NoError(t, err2)
	// test1 := UsePotion("amazing-pills")
}
