package rule

//SpellDescription struct
type SpellDescription struct {
	Name              string   `json:"name"`
	Level             int      `json:"level"`
	Title             string   `json:"title"`
	Subtitle          string   `json:"subtitle"`
	CastingTime       string   `json:"casting_time"`
	Range             string   `json:"range"`
	Components        string   `json:"components"`
	Duration          string   `json:"duration"`
	Description       string   `json:"description"`
	AtHigherLevels    string   `json:"at_higher_levels,omitempty"`
	DamageIncrease    string   `json:"damage_increase,omitempty"`
	DamageDice        string   `json:"damage_dice,omitempty"`
	DamageType        string   `json:"damage_type,omitempty"`
	SavingThrow       string   `json:"saving_throw,omitempty"`
	HealDice          string   `json:"heal_dice,omitempty"`
	HealingIncreases  string   `json:"healing_increase,omitempty"`
	ExtraDice         string   `json:"extra_dice,omitempty"`
	ExtraDiceUsage    []string `json:"extra_dice_usage,omitempty"`
	BonusArmorClass   int      `json:"bonus_armor_class,omitempty"`
	BaseArmorClass    int      `json:"base_armor_class,omitempty"`
	MinimumArmorClass int      `json:"minimum_armor_class,omitempty"`
	AttackRolls       bool     `json:"attack_rolls,omitempty"`
}

//SpellListByClass struct
type SpellListByClass struct {
	Bard struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Bard"`
	Cleric struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Cleric"`
	Druid struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Druid"`
	Paladin struct {
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
	} `json:"Paladin"`
	Ranger struct {
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
	} `json:"Ranger"`
	Sorcerer struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Sorcerer"`
	Warlock struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Warlock"`
	Wizard struct {
		Level0 []string `json:"level0"`
		Level1 []string `json:"level1"`
		Level2 []string `json:"level2"`
		Level3 []string `json:"level3"`
		Level4 []string `json:"level4"`
		Level5 []string `json:"level5"`
		Level6 []string `json:"level6"`
		Level7 []string `json:"level7"`
		Level8 []string `json:"level8"`
		Level9 []string `json:"level9"`
	} `json:"Wizard"`
}
