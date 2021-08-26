package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringInSlice(t *testing.T) {
	testSlice := []string{"foo", "bar", "test"}
	testString := "test"
	testResult := StringInSlice(testString, testSlice)
	assert.True(t, testResult)

	testSlice1 := []string{"foo", "bar", "testa"}
	testString1 := "testb"
	testResult1 := StringInSlice(testString1, testSlice1)
	assert.False(t, testResult1)

}

func TestCleanTrimSlice(t *testing.T) {
	testSlice := []string{" left", "right ", " both ", " two words "}
	expectedSlice := []string{"left", "right", "both", "two words"}
	res := CleanTrimSlice(testSlice)
	for k, v := range res {
		assert.Equal(t, expectedSlice[k], v)
	}
}

func TestCleanTrimString(t *testing.T) {
	left := " left"
	right := "right "
	both := " both "
	words := " two words "
	expected1 := "left"
	expected2 := "right"
	expected3 := "both"
	expected4 := "two words"
	res1 := CleanTrimString(left)
	assert.Equal(t, expected1, res1)
	res2 := CleanTrimString(right)
	assert.Equal(t, expected2, res2)
	res3 := CleanTrimString(both)
	assert.Equal(t, expected3, res3)
	res4 := CleanTrimString(words)
	assert.Equal(t, expected4, res4)
}

func TestGetD20ToRoll(t *testing.T) {
	adv := false
	dis := false
	res1 := GetD20ToRoll(adv, dis)
	assert.Equal(t, "1d20", res1)
	adv = true
	res2 := GetD20ToRoll(adv, dis)
	assert.Equal(t, "2d20dh1", res2)
	adv = false
	dis = true
	res3 := GetD20ToRoll(adv, dis)
	assert.Equal(t, "2d20kh1", res3)
	adv = true
	dis = true
	res4 := GetD20ToRoll(adv, dis)
	assert.Equal(t, "1d20", res4)
}
