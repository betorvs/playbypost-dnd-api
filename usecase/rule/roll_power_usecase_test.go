package rule

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/stretchr/testify/assert"
)

func TestUsePower(t *testing.T) {
	test1 := UsePower(nil)
	assert.Error(t, test1)
	test2 := new(rule.CorePowers)
	res2 := UsePower(test2)
	assert.NoError(t, res2)
}
