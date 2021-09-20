package rule

//SimpleList struct
type SimpleList struct {
	List []string `json:"list"`
}

//FullDescription struct
type FullDescription struct {
	Description map[string]string `json:"description"`
}

//ReturnMessage struct
type ReturnMessage struct {
	Message string `json:"message"`
}

//ReturnACMessage struct
type ReturnACMessage struct {
	ArmorClassStealthDisvantage  bool     `json:"armor_class_stealth_disvantage"`
	ArmorClassDisvantages        []string `json:"armor_class_disvantages"`
	ArmorClassAutomaticallyFails []string `json:"armor_class_auto_fails"`
	ArmorClassSpeedReduced       bool     `json:"armor_class_speed_reduced"`
	ArmorClassMaxDexterity       int      `json:"armor_class_max_dexterity"`
	ArmorClass                   int      `json:"armor_class"`
}

// ReturnCalcMessage struct
type ReturnCalcMessage struct {
	Message        string          `json:"message"`
	RolledMessage  string          `json:"rolled_message,omitempty"`
	SuccessMessage string          `json:"success_message,omitempty"`
	Success        bool            `json:"success,omitempty"`
	RolledValue    int             `json:"rolled_value,omitempty"`
	DamageValue    int             `json:"damage_value,omitempty"`
	DamageType     string          `json:"damage_type,omitempty"`
	HealingValue   int             `json:"healing_value,omitempty"`
	MonstersResult []MonsterResult `json:"monsters_result,omitempty"`
	// TwoWeaponFighting       bool            `json:"two_weapon_fighting"`
}

//MonsterResult struct
type MonsterResult struct {
	Name           string `json:"name"`
	Message        string `json:"message"`
	Success        bool   `json:"success"`
	DamageMessage  string `json:"damage_message,omitempty"`
	SuccessMessage string `json:"success_message,omitempty"`
	RolledMessage  string `json:"rolled_message,omitempty"`
	RolledValue    int    `json:"rolled_value,omitempty"`
	DamageValue    int    `json:"damage_value,omitempty"`
	DamageType     string `json:"damage_type,omitempty"`
}

// ReturnCondition struct
type ReturnCondition struct {
	Name        string   `json:"name"`
	Description []string `json:"description"`
	Disvantages []string `json:"disvantages"`
	AutoFail    []string `json:"auto_fail"`
}

//MonsterRoll struct
// used by a master when need test monster against players
type MonsterRoll struct {
	Name                       string   `json:"name"`
	Check                      string   `json:"check"`
	Weapon                     string   `json:"weapon"`
	DifficultClass             int      `json:"difficult_class"`
	Disvantages                []string `json:"disvantages"`
	Advantages                 []string `json:"advantages"`
	AutoFail                   []string `json:"auto_fail"`
	EnemyRage                  bool     `json:"enemy_rage"`
	EnemyDamageVulnerabilities []string `json:"enemy_damage_vulnerabilities"`
	EnemyDamageImmunities      []string `json:"enemy_damage_immunities"`
	EnemyConditionImmunities   []string `json:"enemy_condition_immunities"`
	EnemyDamageResistances     []string `json:"enemy_damage_resistances"`
	Externald20                int      `json:"external_d20"`
}

//MonsterTurn struct
type MonsterTurn struct {
	Level         int            `json:"level"`
	ClassFeatures []string       `json:"class_features"`
	Ability       map[string]int `json:"ability"`
	MonsterList   []string       `json:"monster_list"`
}

//NewCharacter struct
type NewCharacter struct {
	Level                     int            `json:"level"`
	Class                     string         `json:"class"`
	Race                      string         `json:"race"`
	Subrace                   string         `json:"subrace,omitempty"`
	Background                string         `json:"background"`
	Ability                   map[string]int `json:"ability"`
	ChosenLanguages           []string       `json:"chosen_languages,omitempty"`
	ChosenSkills              []string       `json:"chosen_skills,omitempty"`
	ChosenAbility             []string       `json:"chosen_ability,omitempty"`
	ChosenAbilityByLevel      []string       `json:"chosen_ability_level,omitempty"`
	ChosenClassFeatures       []string       `json:"chosen_class_features,omitempty"`
	ChosenSkillsByFeatures    []string       `json:"chosen_skills_features,omitempty"`
	ChosenLanguagesByFeatures []string       `json:"chosen_languages_features,omitempty"`
}

//Character struct
type Character struct {
	Level                     int               `json:"level"`
	Class                     string            `json:"class"`
	ClassFeatures             []string          `json:"class_features"`
	Race                      string            `json:"race"`
	Subrace                   string            `json:"subrace,omitempty"`
	RaceFeatures              []string          `json:"race_features"`
	Background                string            `json:"background"`
	Proficiency               int               `json:"proficiency"`
	HitDice                   string            `json:"hit_dice"`
	Size                      string            `json:"size"`
	SpellKnown                int               `json:"spell_known,omitempty"`
	SpellListLevel            map[string][]int  `json:"spell_list_level,omitempty"`
	SpellMaxLevel             int               `json:"spell_max_level,omitempty"`
	CantripsKnown             int               `json:"cantrips_known,omitempty"`
	XPNextLevel               int               `json:"xp_next_level"`
	BarbarianRage             int               `json:"rage,omitempty"`
	BarbarianDamage           int               `json:"rage_damage,omitempty"`
	MonkMartial               string            `json:"martial,omitempty"`
	MonkKi                    int               `json:"ki,omitempty"`
	MonkMovement              string            `json:"movement,omitempty"`
	RogueSneak                string            `json:"sneack_attack,omitempty"`
	SorceryPoints             int               `json:"sorcery_points,omitempty"`
	WarlockSpellSlots         int               `json:"warlock_spell_slots,omitempty"`
	WarlockSlotLevel          string            `json:"warlock_slot_level,omitempty"`
	WarlockInvocationsKnown   int               `json:"warlock_invocation_known,omitempty"`
	Speed                     int               `json:"speed"`
	SpeedMeasure              string            `json:"speed_measure"`
	Ability                   map[string]int    `json:"ability"`
	AbilityModifier           map[string]int    `json:"ability_modifier"`
	ChosenLanguages           []string          `json:"chosen_languages,omitempty"`
	ChosenSkills              []string          `json:"chosen_skills,omitempty"`
	ChosenAbility             []string          `json:"chosen_ability,omitempty"`
	ChosenAbilityByLevel      []string          `json:"chosen_ability_level,omitempty"`
	ChosenClassFeatures       []string          `json:"chosen_class_features,omitempty"`
	ChosenSkillsByFeatures    []string          `json:"chosen_skills_features,omitempty"`
	ChosenLanguagesByFeatures []string          `json:"chosen_languages_features,omitempty"`
	HPMax                     int               `json:"hp_max"`
	HPTemp                    int               `json:"hp_temp"`
	Language                  []string          `json:"language"`
	Savings                   []string          `json:"savings"`
	ArmorProficiency          []string          `json:"armor_proficiency"`
	Skills                    []string          `json:"skills,omitempty"`
	Disvantages               []string          `json:"disvantages,omitempty"`
	Advantages                []string          `json:"advantages,omitempty"`
	AutoFail                  []string          `json:"auto_fail,omitempty"`
	DamageResistence          []string          `json:"damage_resistance"`
	DamageVulnerabilities     []string          `json:"damage_vulnerabilities"`
	DamageImmunities          []string          `json:"damage_immunities"`
	ConditionImmunities       []string          `json:"condition_immunities"`
	MagicalEffect             []string          `json:"magical_effect"`
	IncompleteOptions         map[string]string `json:"incomplete_options"`
}

//Attack struct
// used by player to calculate one attack
type Attack struct {
	Level            int            `json:"level"`
	Race             string         `json:"race,omitempty"`
	Subrace          string         `json:"subrace,omitempty"`
	Ability          map[string]int `json:"ability"`
	ClassFeatures    []string       `json:"class_features"`
	ArmorProficiency []string       `json:"armor_proficiency"`
	Weapon           []string       `json:"weapon"`
	SecundaryWeapon  string         `json:"secundary_weapon"`
	TwoHands         bool           `json:"two_hands"`
	Rage             bool           `json:"rage"`
	TemporaryBonus   int            `json:"temporary_bonus,omitempty"`
	MagicBonus       int            `json:"magic_bonus"`
	UsingFeature     string         `json:"using_feature"`
	UsingFeatureType string         `json:"using_feature_type"`
	UsingFeatureSlot int            `json:"using_feature_slot"`
	DifficultClass   int            `json:"difficult_class"`
	Monster          string         `json:"monster"`
	Disvantages      []string       `json:"disvantages"`
	Advantages       []string       `json:"advantages"`
	AutoFail         []string       `json:"auto_fail"`
	Externald20      int            `json:"external_d20"`
}

//SpellcastAbility struct
//  used by a player to calculate a spell
type SpellcastAbility struct {
	Level           int            `json:"level"`
	Class           string         `json:"class"`
	Race            string         `json:"race,omitempty"`
	Subrace         string         `json:"subrace,omitempty"`
	RacialTrait     bool           `json:"racial_trait,omitempty"`
	Ability         map[string]int `json:"ability"`
	ClassFeatures   []string       `json:"class_features"`
	MagicBonus      int            `json:"magic_bonus"`
	DifficultClass  int            `json:"difficult_class"`
	SpellName       string         `json:"spell_name"`
	SpellDamage     string         `json:"spell_damage"`
	SpellDamageType string         `json:"spell_damage_type"`
	SpellLevel      int            `json:"spell_level"`
	SpellSaving     string         `json:"spell_saving"`
	Rage            bool           `json:"rage"`
	Monster         []string       `json:"monster"`
	Disvantages     []string       `json:"disvantages"`
	Advantages      []string       `json:"advantages"`
	AutoFail        []string       `json:"auto_fail"`
	Externald20     int            `json:"external_d20"`
}

//SkillOrAbilityCheck struct
// used by a player to calcular a skill or ability check
type SkillOrAbilityCheck struct {
	Level             int            `json:"level"`
	Race              string         `json:"race,omitempty"`
	Subrace           string         `json:"subrace,omitempty"`
	Ability           map[string]int `json:"ability"`
	ClassFeatures     []string       `json:"class_features"`
	Check             string         `json:"check"`
	Skills            []string       `json:"skills"`
	TemporaryBonus    int            `json:"temporary_bonus,omitempty"`
	MagicBonus        int            `json:"magic_bonus"`
	DoubleProficiency bool           `json:"double_proficiency"`
	Rage              bool           `json:"rage"`
	DifficultClass    int            `json:"difficult_class"`
	Disvantages       []string       `json:"disvantages"`
	Advantages        []string       `json:"advantages"`
	AutoFail          []string       `json:"auto_fail"`
	Externald20       int            `json:"external_d20"`
}

//SavingsCheck struct
// used by a player to calculate a saving
type SavingsCheck struct {
	Level          int            `json:"level"`
	Race           string         `json:"race,omitempty"`
	Subrace        string         `json:"subrace,omitempty"`
	Ability        map[string]int `json:"ability"`
	ClassFeatures  []string       `json:"class_features"`
	Check          string         `json:"check"`
	Saving         string         `json:"saving"`
	Savings        []string       `json:"savings"`
	MagicBonus     int            `json:"magic_bonus"`
	TemporaryBonus int            `json:"temporary_bonus,omitempty"`
	Rage           bool           `json:"rage"`
	DifficultClass int            `json:"difficult_class"`
	Disvantages    []string       `json:"disvantages"`
	Advantages     []string       `json:"advantages"`
	AutoFail       []string       `json:"auto_fail"`
	Externald20    int            `json:"external_d20"`
}

//ArmorClass struct
// used to calculate armor class for a player
type ArmorClass struct {
	Ability          map[string]int `json:"ability"`
	ArmorProficiency []string       `json:"armor_proficiency"`
	Armor            string         `json:"armor"`
	ClassFeatures    []string       `json:"class_features"`
	Shield           string         `json:"shield"`
	ArmorMagicBonus  int            `json:"armor_magic_bonus"`
	ShieldMagicBonus int            `json:"shield_magic_bonus"`
}

//SpecialRaceFeature struct
//  used by a player to calculate a spell
type SpecialRaceFeature struct {
	Name           string         `json:"name"`
	Level          int            `json:"level"`
	Class          string         `json:"class"`
	Race           string         `json:"race,omitempty"`
	Subrace        string         `json:"subrace,omitempty"`
	Ability        map[string]int `json:"ability"`
	ClassFeatures  []string       `json:"class_features"`
	MagicBonus     int            `json:"magic_bonus"`
	DifficultClass int            `json:"difficult_class"`
	Damage         string         `json:"damage"`
	DamageType     string         `json:"damage_type"`
	Saving         string         `json:"saving"`
	Rage           bool           `json:"rage"`
	Monster        []string       `json:"monster"`
	Disvantages    []string       `json:"disvantages"`
	Advantages     []string       `json:"advantages"`
	AutoFail       []string       `json:"auto_fail"`
	Externald20    int            `json:"external_d20"`
}

//Feature struct
type Feature struct {
	Level            int            `json:"level"`
	Ability          map[string]int `json:"ability"`
	Name             string         `json:"name"`
	ClassFeatures    []string       `json:"class_features"`
	UsingFeatureSlot int            `json:"using_feature_slot"`
	GenericList      []string       `json:"generic_list,omitempty"`
	MonsterList      []string       `json:"monster_list"`
}

//PreparedSpellsList struct
type PreparedSpellsList struct {
	Level          int            `json:"level"`
	Ability        map[string]int `json:"ability"`
	Class          string         `json:"class"`
	Verified       bool           `json:"verified"`
	PreparedSpells []string       `json:"prepared_spells"`
}

//KnownSpellsList struct
type KnownSpellsList struct {
	Level         int      `json:"level"`
	Class         string   `json:"class"`
	Verified      bool     `json:"verified"`
	SpellMaxLevel int      `json:"spell_max_level,omitempty"`
	KnownSpells   int      `json:"known_spells"`
	SpellList     []string `json:"spell_list"`
}

//KnownCantripList struct
type KnownCantripList struct {
	Class         string   `json:"class"`
	ClassFeatures []string `json:"class_features"`
	Verified      bool     `json:"verified"`
	CantripsKnown int      `json:"cantrips_known,omitempty"`
	CantripsList  []string `json:"cantrips_list"`
}

//Potion struct
//  used by a player to calculate potion effect
type Potion struct {
	Name                  string         `json:"name"`
	Ability               map[string]int `json:"ability"`
	Message               string         `json:"message"`
	HealingValue          int            `json:"healing_value,omitempty"`
	DamageValue           int            `json:"damage_value,omitempty"`
	DamageType            string         `json:"damage_type,omitempty"`
	Disvantages           []string       `json:"disvantages"`
	Advantages            []string       `json:"advantages"`
	AutoFail              []string       `json:"auto_fail"`
	AttackMagicBonus      int            `json:"attack_magic_bonus"`
	DamageResistence      []string       `json:"damage_resistance"`
	DamageVulnerabilities []string       `json:"damage_vulnerabilities"`
	DamageImmunities      []string       `json:"damage_immunities"`
	ConditionImmunities   []string       `json:"condition_immunities"`
	MagicalEffect         []string       `json:"magical_effect"`
	DifficultClass        int            `json:"difficult_class"`
	Conditions            []string       `json:"conditions"`
	SavingThrow           string         `json:"saving_throw,omitempty"`
}

// ShoppingCart struct
type ShoppingCart struct {
	Items            []string `json:"items"`
	Cost             Treasure `json:"costs"`
	ServicesItems    []string `json:"services_items"`
	UnavailableItems []string `json:"unavailable_items"`
}

// Treasure struct
type Treasure struct {
	Copper   int `json:"copper,omitempty"`
	Silver   int `json:"silver,omitempty"`
	Electrum int `json:"electrum,omitempty"`
	Gold     int `json:"gold,omitempty"`
	Platinum int `json:"platinum,omitempty"`
}

//RandomTreasure struct
type RandomTreasure struct {
	Challenge int      `json:"challenge"`
	Treasure  Treasure `json:"treasure"`
	Message   string   `json:"message"`
}

//RandomTreasureHoard struct
type RandomTreasureHoard struct {
	Challenge   int               `json:"challenge"`
	Treasure    Treasure          `json:"treasure"`
	MagicItems  []string          `json:"magic_items"`
	ArmoryItems map[string]string `json:"armory_items"`
	ChargeItems map[string]int    `json:"charge_items"`
	Gemstones   []string          `json:"gemstones"`
	ArtObjects  []string          `json:"art-objects"`
	Message     string            `json:"message"`
}

// CoinTable struct
type CoinTable struct {
	Type         string
	Dice         string
	AverageValue int
	Multiple     int
}
