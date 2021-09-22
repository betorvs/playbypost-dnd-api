package rule

// Class struct
type Class struct {
	Name                    string           `json:"name"`
	Title                   string           `json:"title"`
	Description             string           `json:"description"`
	HitDice                 int              `json:"hit_dice"`
	Savings                 []string         `json:"savings"`
	ArmorProficiency        []string         `json:"armor_proficiency"`
	SkillNumber             int              `json:"skill_number"`
	SkillList               []string         `json:"skill_list"`
	AbilityForSpell         string           `json:"ability_for_spell,omitempty"`
	SpellSlotsMultiClass    int              `json:"spell_slots_multiclass,omitempty"`
	Features                map[int][]string `json:"features"`
	SpellKnown              map[int]int      `json:"spell_known,omitempty"`
	CantripsKnown           map[int]int      `json:"cantrips,omitempty"`
	SpellsPerLevel          map[int][]int    `json:"spell_per_level,omitempty"`
	BarbarianRage           map[int]int      `json:"barbarian_rage,omitempty"`
	BarbarianRageDamage     map[int]int      `json:"barbarian_rage_damage,omitempty"`
	MonkMartial             map[int]string   `json:"monk_martial,omitempty"`
	MonkKi                  map[int]int      `json:"monk_ki,omitempty"`
	MonkMovement            map[int]string   `json:"monk_movement,omitempty"`
	RogueSneakAttack        map[int]string   `json:"rogue_sneak_attack,omitempty"`
	WarlockSpellSlots       map[int]int      `json:"warlock_spells_slot,omitempty"`
	WarlockSlotLevel        map[int]string   `json:"warlock_slot_level,omitempty"`
	WarlockInvocationsKnown map[int]int      `json:"warlock_invocations_known,omitempty"`
}
