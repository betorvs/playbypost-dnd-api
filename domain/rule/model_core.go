package rule

// CoreFeatures struct
// this kind of magic is always activated when in use
type CoreFeatures struct {
	AttackBonus           int            `json:"attack_bonus,omitempty"`           // in attack only
	DamageBonus           int            `json:"damage_bonus,omitempty"`           // in Damage only
	CombateBonus          int            `json:"combate_bonus,omitempty"`          // on attack and damage
	DamageBonusCritical   int            `json:"damage_bonus_critical,omitempty"`  // add bonus to damage if was a critical
	DamageDiceCritical    string         `json:"damage_dice_critical,omitempty"`   // add dice bonus to damage if was a critical
	DamageType            string         `json:"damage_type,omitempty"`            // type of damage from allDamageType
	SpellBonus            int            `json:"spell_bonus,omitempty"`            // on spell attack rolls and difficult class
	SpellAttackBonus      int            `json:"spell_attack_bonus,omitempty"`     // on spell attack rolls
	ArmorClassBonus       int            `json:"armor_class_bonus,omitempty"`      // AC
	SavingBonus           int            `json:"saving_bonus,omitempty"`           // on saving roll
	ExtraHitPointsLevel   int            `json:"extra_hit_points_level,omitempty"` // extra hit points per level
	AbilityBonus          int            `json:"ability_bonus,omitempty"`          // on ability checks roll
	NewAbility            map[string]int `json:"new_ability,omitempty"`            // Change Ability int
	IncreaseAbility       map[string]int `json:"increase_ability,omitempty"`       // Increase Ability int by
	SkillBonus            map[string]int `json:"skill_bonus,omitempty"`            // map skill and add bonus int in check
	WithProficiency       []string       `json:"with_proficiency,omitempty"`       // gives to user that proficiency
	Disvantages           []string       `json:"disvantages,omitempty"`            // list of Disvantages to add
	Advantages            []string       `json:"advantages,omitempty"`             // list of Disvantages to add
	AutoFail              []string       `json:"auto_fail,omitempty"`              // autoFail in something, like ability, skill, attack, anything, add to your list
	DamageResistance      []string       `json:"damage_resistance,omitempty"`      // damageResistance to add from allDamageType list one or more
	DamageVulnerabilities []string       `json:"damage_vulnerabilities,omitempty"` // DamageVulnerabilities to add from allDamageType list one or more
	DamageImmunities      []string       `json:"damage_immunities,omitempty"`      // DamageImmunities to add from allDamageType list one or more
	ConditionImmunities   []string       `json:"condition_immunities,omitempty"`   // ConditionImmunities to add from allDamageType list one or more
	CancelDisvantage      []string       `json:"cancel_disvantage,omitempty"`      // Removes one kind of Disvantages from your list
	CancelCondition       []string       `json:"cancel_condition,omitempty"`       // Removes one kind of Condition from your list
	EnemyDisvantages      []string       `json:"enemy_disvantages,omitempty"`      // list of Disvantages to your enemies
	Curse                 bool           `json:"curse,omitempty"`                  // add disvantage when activated
	OverrideDamageType    string         `json:"override_damage_type,omitempty"`   // used in weapons to cause only one type of damage instead weapons one
	Regeneration          int            `json:"regeneration,omitempty"`           // recovery HP per turn, after his own action
	ProficiencyBonus      int            `json:"proficiency_bonus,omitempty"`      // Increase Character proficiency bonus in all tests
	SpellImmunity         []string       `json:"spell_immunity,omitempty"`         // just a Spell Immunity list
}

// CorePowers struct
// this kind of magic should be trigger or used
type CorePowers struct {
	Purpose                   string   `json:"purpose,omitempty"`                     // in CoreDnDSystem which rule should be used
	Dice                      string   `json:"dice,omitempty"`                        // Dice to roll
	DamageType                string   `json:"damage_type,omitempty"`                 // type of damage from allDamageType
	DamageDice                string   `json:"damage_dice,omitempty"`                 // if have damage, use this dices
	DifficultClass            int      `json:"difficult_class,omitempty"`             // to enemy use
	SavingThrow               string   `json:"saving_throw,omitempty"`                // which ability to use to check
	SpellName                 string   `json:"spell_name,omitempty"`                  // spell name for use
	SpellList                 []string `json:"spell_list,omitempty"`                  // spell list name to be use as spell name
	SpellFreeList             []string `json:"spell_free_list,omitempty"`             // spell free list name to be use as spell name without charge cost
	SpellLevel                int      `json:"spell_level,omitempty"`                 // spell level to be used if apply for that magic
	Duration                  int      `json:"duration,omitempty"`                    // duration time
	AttackRoll                int      `json:"attack_roll,omitempty"`                 // if need to attack, use this value
	AttackNumber              int      `json:"attack_number,omitempty"`               // number of attacks possible
	ConditionMultiple         int      `json:"condition_multiple,omitempty"`          // Condition to trigger any SavingThrow
	ConditionHitPoints        int      `json:"condition_hit_points,omitempty"`        // Hit Points Condition to trigger any SavingThrow
	ArmorClassBonus           int      `json:"armor_class_bonus,omitempty"`           // AC
	EnemyAttackType           string   `json:"enemy_attack_type,omitempty"`           // Enemy Attack Type used to trigger any defense type
	ReduceDamageDice          string   `json:"reduce_damage_dice,omitempty"`          // If enemy attack type matches, reduce damage received
	ReduceDamageAbilitiy      string   `json:"reduce_damage_abality,omitempty"`       // which ability to use with reduce damage dice
	DamageResistance          []string `json:"damage_resistance,omitempty"`           // damageResistance to add from allDamageType list one or more
	Advantages                []string `json:"advantages,omitempty"`                  // add Advantages if enemy list match
	Disvantages               []string `json:"disvantages,omitempty"`                 // add Disvantages in enemy list temporary
	Condition                 []string `json:"condition,omitempty"`                   // add Condition in enemy list temporary
	CancelCondition           []string `json:"cancel_condition,omitempty"`            // Removes one kind of Condition from your list
	CombateMastery            int      `json:"combate_mastery,omitempty"`             // You can choose where to use your bonus, attack/damage or armor class
	Curse                     bool     `json:"curse,omitempty"`                       // add disvantage when activated
	ChargeType                bool     `json:"charge_type,omitempty"`                 // Any kind of item who needs to have charges, like: staff, wands
	Charges                   int      `json:"charges,omitempty"`                     // number of charges
	DiceCharges               string   `json:"dice_charges,omitempty"`                // random way to generate number of charges, used by hoard generator
	RecoveryDiceCharges       string   `json:"recovery_dice_charges,omitempty"`       // random way to recover charges / per time
	ZeroChargesRoll           bool     `json:"zero_charges_roll,omitempty"`           // if reaches 0 charges, roll a d20, on 1, destroyed
	WeaponPropertyRestriction string   `json:"weapon_property_restriction,omitempty"` // if used weapon doenst have that property, will not work
	DamageBonus               int      `json:"damage_bonus,omitempty"`                // add bonus to damage if power in use
	RequireEnemyType          []string `json:"require_enemy_type,omitempty"`          // used to activate a extra damage if enemy type matches
	CombateBonus              int      `json:"combate_bonus,omitempty"`               // on attack and damage
	DamageBonusCritical       int      `json:"damage_bonus_critical,omitempty"`       // add bonus to damage if was a critical
	ExtraHitPoints            int      `json:"extra_hit_points,omitempty"`            // extra Hit Points
}
