package rule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListConditions(t *testing.T) {
	test1 := "blinded"
	assert.Contains(t, ListConditions(), test1)
}

func TestListOfDamageTypes(t *testing.T) {
	test1 := "acid"
	assert.Contains(t, ListOfDamageTypes(), test1)
}

func TestListOfWeaponDamageType(t *testing.T) {
	test1 := "bludgeoning"
	assert.Contains(t, listOfWeaponDamageType(), test1)
}

func TestConditionsCheck(t *testing.T) {
	test1 := []string{
		"blinded", "charmed", "deafened", "frightened", "grappled", "incapacitated", "invisible", "paralyzed", "petrified", "poisoned", "prone", "pestrained", "stunned", "unconscious",
	}
	for _, v := range test1 {
		desc, _, _ := conditionsCheck(v, 1)
		assert.NotEmpty(t, desc)
	}
	// "exhaustion"
	levels := []int{0, 1, 2, 3, 4, 5, 6}
	for _, v := range levels {
		desc, _, _ := conditionsCheck("exhaustion", v)
		assert.NotEmpty(t, desc)
	}
}

func TestConditionsMap(t *testing.T) {
	res1 := conditionsMap("blinded")
	assert.NotEmpty(t, res1)
}

func TestDamageTypes(t *testing.T) {
	test1 := []string{"acid", "bludgeoning", "cold", "fire", "force", "lightning", "necrotic", "piercing", "poison", "psychic", "radiant", "slashing", "thunder"}
	for _, v := range test1 {
		res := damageTypes(v)
		assert.NotEmpty(t, res)
	}
	res2 := damageTypes("")
	assert.Empty(t, res2)
}

func TestDamageTypeMap(t *testing.T) {
	res1 := damageTypeMap("thunder")
	assert.NotEmpty(t, res1)
}

func TestGetConditions(t *testing.T) {
	test := "blinded"
	res := GetConditions(test, 1)
	assert.NotEmpty(t, res)
}
