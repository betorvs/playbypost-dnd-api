package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/domain/rule"
)

// UsePower func
func UsePower(power *rule.CorePowers) error {
	// receives power and rolls all contents
	if power == nil {
		return fmt.Errorf("power not found")
	}
	// check all purposes and create a rule for each

	return nil
}
