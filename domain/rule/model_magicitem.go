package rule

// MagicItem struct
type MagicItem struct {
	Name                  string        `json:"name"`
	Title                 string        `json:"title"`
	Content               string        `json:"content"`
	Category              string        `json:"category"`
	Rarity                string        `json:"rarity"`
	HoardTable            []string      `json:"hoard_table"`
	AttunementRestriction []string      `json:"attunement_restriction,omitempty"`
	RequiredAttunement    bool          `json:"required_attunement"`
	RolePlay              bool          `json:"roleplay"`
	Forbidden             bool          `json:"forbidden"`
	Shape                 string        `json:"shape,omitempty"`
	Feature               *CoreFeatures `json:"magic_feature,omitempty"`
	Power                 *CorePowers   `json:"power,omitempty"`
	Scroll                *Scroll       `json:"scroll,omitempty"`
}

// Scroll struct
type Scroll struct {
	Content               string `json:"content,omitempty"`
	DifficultClass        int    `json:"difficult_class,omitempty"`
	SavingsDifficultClass int    `json:"savings_difficult_class,omitempty"`
	Attack                int    `json:"attack,omitempty"`
	WizardDifficult       int    `json:"wizard_difficult,omitempty"`
}
