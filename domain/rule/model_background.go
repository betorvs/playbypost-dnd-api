package rule

// Background struct
type Background struct {
	Name                string   `json:"name"`
	Title               string   `json:"title"`
	Language            []string `json:"language,omitempty"`
	Skills              []string `json:"skills,omitempty"`
	AdditionalLanguages int      `json:"additional_languages,omitempty"`
	Extra               string   `json:"extra,omitempty"`
	Description         string   `json:"description,omitempty"`
}
