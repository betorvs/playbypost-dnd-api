package player

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Players struct
type Players struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CampaignID                string             `bson:"campaign_id" json:"campaign_id"`
	CampaignTitle             string             `bson:"campaign_title" json:"campaign_title"`
	Name                      string             `bson:"name" json:"name"`
	Username                  string             `bson:"username" json:"username"`
	SlackID                   string             `bson:"slack_id" json:"slack_id"`
	SlackUsername             string             `bson:"slack_username" json:"slack_username"`
	SlackChannelID            string             `bson:"slack_channel_id" json:"slack_channel_id"`
	Description               Description        `bson:"description" json:"description"`
	Alignment                 string             `bson:"alignment" json:"alignment"`
	Level                     int                `bson:"level" json:"level"`
	Class                     string             `bson:"class" json:"class"`
	ClassFeatures             []string           `bson:"class_features" json:"class_features"`
	Race                      string             `bson:"race" json:"race"`
	Subrace                   string             `bson:"subrace" json:"subrace"`
	RaceFeatures              []string           `bson:"race_features" json:"race_features"`
	Background                string             `bson:"background" json:"background"`
	Ability                   map[string]int     `bson:"ability" json:"ability"`
	AbilityWithoutMagic       map[string]int     `bson:"ability_without_magic" json:"ability_without_magic"`
	AbilityModifier           map[string]int     `bson:"ability_modifier" json:"ability_modifier"`
	ChosenLanguages           []string           `bson:"chosen_languages" json:"chosen_languages,omitempty"`
	ChosenSkills              []string           `bson:"chosen_skills" json:"chosen_skills,omitempty"`
	ChosenAbility             []string           `bson:"chosen_ability" json:"chosen_ability,omitempty"`
	ChosenAbilityByLevel      []string           `bson:"chosen_ability_level" json:"chosen_ability_level,omitempty"`
	ChosenClassFeatures       []string           `bson:"chosen_class_features" json:"chosen_class_features,omitempty"`
	ChosenSkillsByFeatures    []string           `bson:"chosen_skills_features" json:"chosen_skills_features,omitempty"`
	ChosenLanguagesByFeatures []string           `bson:"chosen_languages_features" json:"chosen_languages_features,omitempty"`
	Proficiency               int                `bson:"proficiency" json:"proficiency"`
	HitDice                   string             `bson:"hit_dice" json:"hit_dice"`
	Size                      string             `bson:"size" json:"size"`
	Speed                     int                `bson:"speed" json:"speed"`
	SpeedMeasure              string             `bson:"speed_measure" json:"speed_measure"`
	Skills                    []string           `bson:"skills" json:"skills"`
	Language                  []string           `bson:"language" json:"language"`
	Savings                   []string           `bson:"savings" json:"savings"`
	ArmorProficiency          []string           `bson:"armor_proficiency" json:"armor_proficiency"`
	ArmorClass                int                `bson:"armor_class" json:"armor_class"`
	ArmorClassBonus           int                `bson:"armor_class_bonus" json:"armor_class_bonus"`
	ArmorName                 string             `bson:"armor_name" json:"armor_name"`
	ShieldName                string             `bson:"shield_name" json:"shield_name"`
	WeaponName                string             `bson:"weapon_name" json:"weapon_name"`
	WeaponBonus               int                `bson:"weapon_bonus" json:"weapon_bonus"`
	HPMax                     int                `bson:"hp_max" json:"hp_max"`
	HPTemp                    int                `bson:"hp_temp" json:"hp_temp"`
	SpellKnown                int                `bson:"spell_known" json:"spell_known,omitempty"`
	SpellList                 []string           `bson:"spell_list" json:"spell_list,omitempty"`
	SpellListLevel            map[string][]int   `bson:"spell_list_level" json:"spell_list_level,omitempty"`
	SpellMaxLevel             int                `bson:"spell_max_level" json:"spell_max_level,omitempty"`
	CantripsKnown             int                `bson:"cantrips_known" json:"cantrips_known,omitempty"`
	SorceryPoints             int                `json:"sorcery_points,omitempty"`
	BarbarianRage             int                `bson:"rage" json:"rage,omitempty"`
	BarbarianDamage           int                `bson:"rage_damage" json:"rage_damage,omitempty"`
	MonkMartial               string             `bson:"martial" json:"martial,omitempty"`
	MonkKi                    int                `bson:"ki" json:"ki,omitempty"`
	MonkMovement              string             `bson:"movement" json:"movement,omitempty"`
	RogueSneak                string             `bson:"sneack_attack" json:"sneack_attack,omitempty"`
	WarlockSpellSlots         int                `bson:"warlock_spell_slots" json:"warlock_spell_slots,omitempty"`
	WarlockSlotLevel          string             `bson:"warlock_slot_level" json:"warlock_slot_level,omitempty"`
	WarlockInvocationsKnown   int                `bson:"warlock_invocation_known" json:"warlock_invocation_known,omitempty"`
	XP                        int                `bson:"xp" json:"xp"`
	XPNextLevel               int                `bson:"xp_next_level" json:"xp_next_level"`
	SpellsUsedByLevel         SpellsUsedByLevel  `bson:"spells_used" json:"spells_used,omitempty"`
	Disvantages               []string           `bson:"disvantages" json:"disvantages"`
	Advantages                []string           `bson:"advantages" json:"advantages"`
	AutoFail                  []string           `bson:"auto_fail" json:"auto_fail"`
	DamageVulnerabilities     []string           `bson:"damage_vulnerabilities" json:"damage_vulnerabilities"`
	DamageImmunities          []string           `bson:"damage_immunities" json:"damage_immunities"`
	ConditionImmunities       []string           `bson:"condition_immunities" json:"condition_immunities"`
	DamageResistence          []string           `bson:"damage_resistance" json:"damage_resistance"`
	MagicalEffect             []string           `bson:"magical_effect" json:"magical_effect,omitempty"`
}

//Description struct
type Description struct {
	Age    string `bson:"age" json:"age"`
	Height string `bson:"height" json:"height"`
	Weight string `bson:"weight" json:"weight"`
	Eyes   string `bson:"eyes" json:"eyes"`
	Skin   string `bson:"skin" json:"skin"`
	Hair   string `bson:"hair" json:"hair"`
}

// Inventory struct
type Inventory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	PlayerID     string             `bson:"player_id" json:"player_id"`
	Armory       Armory             `bson:"armory" json:"armory"`
	MagicItems   []string           `bson:"magic_items" json:"magic_items"`               // without attunement require
	AttunedItems []string           `bson:"attuned_items" json:"attuned_items,omitempty"` // with attunement requirement; max 3
	Treasure     Treasure           `bson:"treasure" json:"treasure"`
	Items        []string           `bson:"others_items" json:"others_items,omitempty"`
}

//Armory struct
type Armory struct {
	ArmorMagicBonus  int    `bson:"armor_magic_bonus" json:"armor_magic_bonus"`   // magical bonus
	Armor            string `bson:"armor" json:"armor"`                           // armor name from armory database
	ArmorName        string `bson:"armor_name" json:"armor_name"`                 // used by magic items only
	Shield           string `bson:"shield" json:"shield"`                         // shield name if exists
	ShieldName       string `bson:"shield_name" json:"shield_name"`               // used by magic items only
	ShieldMagicBonus int    `bson:"shield_magic_bonus" json:"shield_magic_bonus"` // magic bonus
	Weapon           string `bson:"weapon" json:"weapon"`                         // weapon name from weapons database
	WeaponName       string `bson:"weapon_name" json:"weapon_name"`               // used by magic items only
	WeaponMagicBonus int    `bson:"weapon_magic_bonus" json:"weapon_magic_bonus"` // magic bonus to attack and damage
}

//AddCampaign struct
type AddCampaign struct {
	CampaignID     string `json:"campaign_id"`
	CampaignTitle  string `json:"campaign_title"`
	SlackChannelID string `json:"slack_channel_id"`
}

// Treasure struct
type Treasure struct {
	Copper   int `bson:"copper" json:"copper,omitempty"`
	Silver   int `bson:"silver" json:"silver,omitempty"`
	Electrum int `bson:"electrum" json:"electrum,omitempty"`
	Gold     int `bson:"gold" json:"gold,omitempty"`
	Platinum int `bson:"platinum" json:"platinum,omitempty"`
}

// SpellsUsedByLevel struct
type SpellsUsedByLevel struct {
	Level0 int `bson:"level0" json:"level0"`
	Level1 int `bson:"level1" json:"level1"`
	Level2 int `bson:"level2" json:"level2"`
	Level3 int `bson:"level3" json:"level3"`
	Level4 int `bson:"level4" json:"level4"`
	Level5 int `bson:"level5" json:"level5"`
	Level6 int `bson:"level6" json:"level6"`
	Level7 int `bson:"level7" json:"level7"`
	Level8 int `bson:"level8" json:"level8"`
	Level9 int `bson:"level9" json:"level9"`
}

// NPC struct
type NPC struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EncounterID string             `bson:"encounter_id,omitempty" json:"encounter_id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Monster     string             `bson:"monster" json:"monster"`
	HitPoints   int                `bson:"hp" json:"hp"`
	ArmorClass  int                `bson:"armor_class" json:"armor_class"`
	XP          int                `bson:"xp" json:"xp"`
	Conditions  []string           `bson:"conditions" json:"conditions"`
	Disvantages []string           `bson:"disvantages" json:"disvantages"`
	Advantages  []string           `bson:"advantages" json:"advantages"`
	AutoFail    []string           `bson:"auto_fail" json:"auto_fail"`
}

// Condition struct
type Condition struct {
	Conditions  []string `bson:"conditions" json:"conditions"`
	Disvantages []string `bson:"disvantages" json:"disvantages"`
	Advantages  []string `bson:"advantages" json:"advantages"`
	AutoFail    []string `bson:"auto_fail" json:"auto_fail"`
}

//NPCDamage struct
type NPCDamage struct {
	Damage int `bson:"hp" json:"hp"`
}

//NPCCondition struct
type NPCCondition struct {
	Condition   []string `bson:"condition" json:"condition"`
	Disvantages []string `bson:"disvantages" json:"disvantages"`
	Advantages  []string `bson:"advantages" json:"advantages"`
	AutoFail    []string `bson:"auto_fail" json:"auto_fail"`
}
