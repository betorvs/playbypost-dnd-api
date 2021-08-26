package utils

import (
	"testing"

	"github.com/betorvs/playbypost-dnd/domain/rule"
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

func TestPrintSliceString(t *testing.T) {
	test1 := []string{"a", "b"}
	res1 := PrintSliceString(test1)
	assert.Contains(t, res1, "*a*")
}

func TestPrintMapStringString(t *testing.T) {
	test1 := map[string]string{"a": "b"}
	res1 := PrintMapStringString(test1)
	assert.Contains(t, res1, "*a:*")
}

func TestFormatMessage(t *testing.T) {
	test1 := "a"
	res1 := FormatMessage(test1)
	expected1 := new(rule.ReturnMessage)
	expected1.Message = "a"
	assert.Equal(t, res1, expected1)
}

func TestExtractWholeInt(t *testing.T) {
	test1 := "d20"
	res1 := ExtractWholeInt(test1)
	assert.Equal(t, res1, 20)
	test2 := "d#$%"
	res2 := ExtractWholeInt(test2)
	assert.NotNil(t, res2)
}

func TestExtractInSlice(t *testing.T) {
	test1 := []string{"a", "b"}
	res1 := ExtractInSlice("b", test1)
	assert.Contains(t, res1, "b")
	res2 := ExtractInSlice("c", test1)
	assert.Contains(t, res2, "")
}

func TestEven(t *testing.T) {
	test1 := 10
	res1 := Even(test1)
	assert.True(t, res1)
	test2 := 3
	res2 := Even(test2)
	assert.False(t, res2)
}

func TestRemoveItemSlice(t *testing.T) {
	test1 := []string{"a", "b", "c"}
	res1 := RemoveItemSlice(test1, "b")
	expected1 := []string{"a", "c"}
	assert.Equal(t, res1, expected1)
	test2 := []string{"a", "b", "c"}
	res2 := RemoveItemSlice(test2, "d")
	expected2 := []string{"a", "b", "c"}
	assert.Equal(t, res2, expected2)
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
