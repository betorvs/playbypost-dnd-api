package rule

// Race struct
type Race struct {
	Name                string         `json:"name"`
	Subrace             string         `json:"subrace"`
	Description         string         `json:"description"`
	Size                string         `json:"size"`
	Speedmeasure        string         `json:"speed_measure"`
	Speed               int            `json:"speed"`
	Ability             map[string]int `json:"ability"`
	Special             []string       `json:"special"`
	ArmorProficiency    []string       `json:"armor_proficiency"`
	Language            []string       `json:"language"`
	AdditionalLanguages int            `json:"additional_languages"`
	Skills              []string       `json:"skills"`
	Resistance          []string       `json:"resistance"`
	Advantages          []string       `json:"advantages"`
	Conditions          []string       `json:"conditions"`
	Disvantages         []string       `json:"disvantages"`
}
