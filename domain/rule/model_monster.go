package rule

//MonsterNPC struct
type MonsterNPC struct {
	Name                   string                 `json:"name"`
	Size                   string                 `json:"size"`
	Type                   string                 `json:"type"`
	Aligment               string                 `json:"aligment"`
	Senses                 string                 `json:"senses"`
	Darkvision             string                 `json:"darkvision"`
	Blindsight             string                 `json:"blindsight"`
	Tremorsense            string                 `json:"tremorsense"`
	Truesight              string                 `json:"truesight"`
	Languages              []string               `json:"languages"`
	Challenge              float64                `json:"challenge"`
	ArmorClass             int                    `json:"armor_class"`
	HitPoints              int                    `json:"hit_points"`
	XP                     int                    `json:"xp"`
	Actions                []string               `json:"actions"`
	WeaponAttack           []WeaponAttack         `json:"weapon_attack"`
	SpecialAttack          []SpecialAttack        `json:"special_attack"`
	SpellCastAbility       SpellCastAbility       `json:"spellcast_abilty,omitempty"`
	InnateSpellCastAbility InnateSpellCastAbility `json:"innate_spellcast_abilty,omitempty"`
	Ability                map[string]int         `json:"ability"`
	Savings                map[string]int         `json:"savings"`
	Skills                 map[string]int         `json:"skills"`
	DamageVulnerabilities  []string               `json:"damage_vulnerabilities"`
	DamageImmunities       []string               `json:"damage_immunities"`
	ConditionImmunities    []string               `json:"condition_immunities"`
	DamageResistances      []string               `json:"damage_resistances"`
	PassivePerception      int                    `json:"passive_perception"`
	Traits                 []string               `json:"traits"`
	LegendaryActions       []string               `json:"legendary_actions"`
	ImgURL                 string                 `json:"img_url"`
}

// WeaponAttack struct
type WeaponAttack struct {
	Name            string `json:"name"`
	Attack          int    `json:"attack"`
	AverageDamage   int    `json:"average_damage"`
	Damage          string `json:"damage"`
	DamageType      string `json:"damage_type"`
	SavingThrows    string `json:"saving_throws,omitempty"`
	DifficultClass  int    `json:"difficult_class,omitempty"`
	ExtraDamage     string `json:"extra_damage,omitempty"`
	ExtraDamageType string `json:"extra_damage_type,omitempty"`
}

// SpecialAttack struct
type SpecialAttack struct {
	Name           string `json:"name"`
	SavingThrows   string `json:"saving_throws"`
	DifficultClass int    `json:"difficult_class"`
	AverageDamage  int    `json:"average_damage,omitempty"`
	Damage         string `json:"damage,omitempty"`
	DamageType     string `json:"damage_type,omitempty"`
	Content        string `json:"content,omitempty"`
}

//SpellCastAbility struct
type SpellCastAbility struct {
	Level          int                 `json:"level,omitempty"`
	DifficultClass int                 `json:"difficult_class,omitempty"`
	Attack         int                 `json:"attack,omitempty"`
	Ability        string              `json:"ability,omitempty"`
	List           string              `json:"list,omitempty"`
	CantripsList   []string            `json:"cantrips_list,omitempty"`
	SlotsPerLevel  map[string]int      `json:"slots_per_level,omitempty"`
	ListPerLevel   map[string][]string `json:"list_per_level,omitempty"`
}

//InnateSpellCastAbility struct
type InnateSpellCastAbility struct {
	DifficultClass  int      `json:"difficult_class,omitempty"`
	Ability         string   `json:"ability,omitempty"`
	List            string   `json:"list,omitempty"`
	AtWillList      []string `json:"at_will_list,omitempty"`
	OnePerDayList   []string `json:"one_per_day_list,omitempty"`
	TwoPerDayList   []string `json:"two_per_day_list,omitempty"`
	ThreePerDayList []string `json:"three_per_day_list,omitempty"`
}
