package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

//StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//PrintSliceString func
func PrintSliceString(slice []string) (result string) {
	for _, v := range slice {
		result += fmt.Sprintf("*%s* \n", v)
	}
	return result
}

//PrintMapStringString func
func PrintMapStringString(localMap map[string]string) (result string) {
	for k, v := range localMap {
		if v != "" {
			result += fmt.Sprintf("*%s:* %s \n", k, v)
		}
	}
	return result
}

//FormatMessage func
func FormatMessage(value string) *rule.ReturnMessage {
	res := new(rule.ReturnMessage)
	res.Message = value
	return res
}

//ExtractWholeInt func
func ExtractWholeInt(value string) int {
	var re = regexp.MustCompile(`[^0-9]+`)
	temp := re.ReplaceAllString(value, "")
	numbers, err := strconv.Atoi(strings.TrimSpace(temp))
	if err != nil {
		fmt.Printf("error extractWholeInt convert string to int %v", err)
	}
	return numbers
}

//ExtractInSlice returns a string if a slice contains it
func ExtractInSlice(a string, list []string) string {
	for _, b := range list {
		if b == a {
			return b
		}
	}
	return ""
}

//Even func test
func Even(number int) bool {
	return number%2 == 0
}

//RemoveItemSlice func
func RemoveItemSlice(items []string, item string) []string {
	newitems := []string{}

	for _, i := range items {
		if i != item {
			newitems = append(newitems, i)
		}
	}

	return newitems
}

// CleanTrimSlice returns a []string after TrimLeft and TrimRight
func CleanTrimSlice(list []string) (clean []string) {
	for _, v := range list {
		left := strings.TrimLeft(v, " ")
		right := strings.TrimRight(left, " ")
		clean = append(clean, right)
	}
	return clean
}

// CleanTrimString returns a string after TrimLeft and TrimRight
func CleanTrimString(list string) (clean string) {
	left := strings.TrimLeft(list, " ")
	clean = strings.TrimRight(left, " ")
	return clean
}

// GetD20ToRoll func receive 2 boolean and answer a string
func GetD20ToRoll(adv, dis bool) string {
	if adv && dis {
		// advantage and disvantage at the same time doesn't stack
		// return a simple 1d20
		return "1d20"
	}
	if dis {
		// returning a disvantage d20 roll: roll 2 d20 and keep higher value
		return "2d20kh1"
	}
	if adv {
		// returning a advantage d20 roll: roll 2 d20 and discard higher value
		return "2d20dh1"
	}
	// simple d20
	return "1d20"
}
