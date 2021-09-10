package test

import (
	"encoding/json"
	"fmt"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
)

var (
	MockMonsterDatabase = []byte(`[{
        "name": "kobold",
        "size": "small",
        "type": "humanoid",
        "aligment": "lawful-evil",
        "senses": "Darkvision 60 ft.,  Passive Perception 8",
        "darkvision": "60 ft.",
        "blindsight": "",
        "tremorsense": "",
        "truesight": "",
        "languages": [
            "common",
            "draconic"
        ],
        "challenge": 0.12,
        "armor_class": 12,
        "hit_points": 5,
        "xp": 25,
        "actions": [
            "Dagger.   Melee Weapon Attack:  +4 to hit, reach 5 ft., one target.  Hit:  4 (1d4 + 2) piercing damage.  ",
            "Sling.   Ranged Weapon Attack:  +4 to hit, range 30/120 ft., one target.  Hit:  4 (1d4 + 2) bludgeoning damage. "
        ],
        "weapon_attack": [
            {
                "name": "dagger",
                "attack": 4,
                "average_damage": 4,
                "damage": "1d4+2",
                "damage_type": "piercing"
            },
            {
                "name": "sling",
                "attack": 4,
                "average_damage": 4,
                "damage": "1d4+2",
                "damage_type": "bludgeoning"
            }
        ],
        "special_attack": [],
        "spellcast_abilty": {},
        "innate_spellcast_abilty": {},
        "ability": {
            "charisma": 8,
            "constitution": 9,
            "dexterity": 15,
            "intelligence": 8,
            "strength": 7,
            "wisdom": 7
        },
        "savings": {
            "charisma": -1,
            "constitution": -1,
            "dexterity": 2,
            "intelligence": -1,
            "strength": -2,
            "wisdom": -2
        },
        "skills": {},
        "damage_vulnerabilities": [],
        "damage_immunities": [],
        "condition_immunities": [],
        "damage_resistances": [],
        "passive_perception": 8,
        "traits": [
            "Sunlight Sensitivity.  While in sunlight, the kobold has disadvantage on attack rolls, as well as on Wisdom (Perception) checks that rely on sight.  ",
            "Pack Tactics.  The kobold has advantage on an attack roll against a creature if at least one of the kobold's allies is within 5 feet of the creature and the ally isn't incapacitated. "
        ],
        "legendary_actions": [],
        "img_url": "https://media-waterdeep.cursecdn.com/avatars/thumbnails/0/379/1000/1000/636252780450300625.jpeg"
    },
    {
        "name": "dretch",
        "size": "small",
        "type": "fiend",
        "aligment": "chaotic-evil",
        "senses": "Darkvision 60 ft.,  Passive Perception 9",
        "darkvision": "60 ft.",
        "blindsight": "",
        "tremorsense": "",
        "truesight": "",
        "languages": [
            "abyssal",
            "telepathy 60 ft."
        ],
        "challenge": 0.25,
        "armor_class": 11,
        "hit_points": 18,
        "xp": 50,
        "actions": [
            "Multiattack.  The dretch makes two attacks: one with its bite and one with its claws.  ",
            "Bite.   Melee Weapon Attack:  +2 to hit, reach 5 ft., one target.  Hit:  3 (1d6) piercing damage.  ",
            "Claws.   Melee Weapon Attack:  +2 to hit, reach 5 ft., one target.  Hit:  5 (2d4) slashing damage.  ",
            "Fetid Cloud (1/Day).  A 10­-foot radius of disgusting green gas extends out from the dretch. The gas spreads around corners, and its area is lightly obscured. It lasts for 1 minute or until a strong wind disperses it. Any creature that starts its turn in that area must succeed on a DC 11 Constitution saving throw or be poisoned until the start of its next turn. While poisoned in this way, the target can take either an action or a bonus action on its turn, not both, and can't take reactions. "
        ],
        "weapon_attack": [
            {
                "name": "bite",
                "attack": 2,
                "average_damage": 3,
                "damage": "1d6",
                "damage_type": "piercing"
            },
            {
                "name": "claws",
                "attack": 2,
                "average_damage": 5,
                "damage": "2d4",
                "damage_type": "slashing"
            }
        ],
        "special_attack": [
            {
                "name": "fetid-cloud",
                "saving_throws": "constitution",
                "difficult_class": 11,
                "damage_type": "poisoned"
            }
        ],
        "spellcast_abilty": {},
        "innate_spellcast_abilty": {},
        "ability": {
            "charisma": 3,
            "constitution": 12,
            "dexterity": 11,
            "intelligence": 5,
            "strength": 11,
            "wisdom": 8
        },
        "savings": {
            "charisma": -4,
            "constitution": 1,
            "dexterity": 0,
            "intelligence": -3,
            "strength": 0,
            "wisdom": -1
        },
        "skills": {},
        "damage_vulnerabilities": [],
        "damage_immunities": [
            "poison"
        ],
        "condition_immunities": [
            "poisoned"
        ],
        "damage_resistances": [
            "cold",
            "fire",
            "lightning"
        ],
        "passive_perception": 9,
        "traits": null,
        "legendary_actions": [],
        "img_url": "https://media-waterdeep.cursecdn.com/avatars/thumbnails/0/293/1000/1000/636252771253285096.jpeg"
    },
    {
        "name": "skeleton",
        "size": "medium",
        "type": "undead",
        "aligment": "lawful-evil",
        "senses": "Darkvision 60 ft.,  Passive Perception 9",
        "darkvision": "60 ft.",
        "blindsight": "",
        "tremorsense": "",
        "truesight": "",
        "languages": [
            "understands all languages it knew in life but can't speak"
        ],
        "challenge": 0.25,
        "armor_class": 13,
        "hit_points": 13,
        "xp": 50,
        "actions": [
            "Shortsword.   Melee Weapon Attack:  +4 to hit, reach 5 ft., one target.  Hit:  5 (1d6 + 2) piercing damage.  ",
            "Shortbow.   Ranged Weapon Attack:  +4 to hit, range 80/320 ft., one target.  Hit:  5 (1d6 + 2) piercing damage. "
        ],
        "weapon_attack": [
            {
                "name": "shortsword",
                "attack": 4,
                "average_damage": 5,
                "damage": "1d6+2",
                "damage_type": "piercing"
            },
            {
                "name": "shortbow",
                "attack": 4,
                "average_damage": 5,
                "damage": "1d6+2",
                "damage_type": "piercing"
            }
        ],
        "special_attack": [],
        "spellcast_abilty": {},
        "innate_spellcast_abilty": {},
        "ability": {
            "charisma": 5,
            "constitution": 15,
            "dexterity": 14,
            "intelligence": 6,
            "strength": 10,
            "wisdom": 8
        },
        "savings": {
            "charisma": -3,
            "constitution": 2,
            "dexterity": 2,
            "intelligence": -2,
            "strength": 0,
            "wisdom": -1
        },
        "skills": {},
        "damage_vulnerabilities": [
            "bludgeoning"
        ],
        "damage_immunities": [
            "poison"
        ],
        "condition_immunities": [
            "exhaustion",
            "poisoned"
        ],
        "damage_resistances": [],
        "passive_perception": 9,
        "traits": null,
        "legendary_actions": [],
        "img_url": "https://media-waterdeep.cursecdn.com/avatars/thumbnails/16/472/315/315/636376294573239565.jpeg"
    }]`)
	MockSpellDescriptionDatabase = []byte(`[{
        "name": "magic-missile",
        "level": 1,
        "title": "Magic Missile",
        "subtitle": "1st-level evocation",
        "casting_time": "1 action",
        "range": "120 feet",
        "components": "V, S",
        "duration": "Instantaneous",
        "description": " You create three glowing darts of magical force. Each dart hits a creature of your choice that you can see within range. A dart deals 1d4 + 1 force damage to its target. The darts all strike simultaneously, and you can direct them to hit one creature or several. *At Higher Levels.* When you cast this spell using a spell slot of 2nd level or higher, the spell creates one more dart for each slot level above 1st.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 2nd level or higher, the spell creates one more dart for each slot level above 1st."
    },
    {
        "name": "spiritual-weapon",
        "level": 2,
        "title": "Spiritual Weapon",
        "subtitle": "2nd-level evocation",
        "casting_time": "1 bonus action",
        "range": "60 feet",
        "components": "V, S",
        "duration": "1 minute",
        "description": " You create a floating, spectral weapon within range that lasts for the duration or until you cast this spell again. When you cast the spell, you can make a melee spell attack against a creature within 5 feet of the weapon. On a hit, the target takes force damage equal to 1d8 + your spellcasting ability modifier. As a bonus action on your turn, you can move the weapon up to 20 feet and repeat the attack against a creature within 5 feet of it. The weapon can take whatever form you choose. Clerics of deities who are associated with a particular weapon (as St. Cuthbert is known for his mace and Thor for his hammer) make this spell’s effect resemble that weapon. *At Higher Levels.* When you cast this spell using a spell slot of 3rd level or higher, the damage increases by 1d8 for every two slot levels above 2nd.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 3rd level or higher, the damage increases by 1d8 for every two slot levels above 2nd.",
        "attack_rolls": true
    },
    {
        "name": "fireball",
        "level": 3,
        "title": "Fireball",
        "subtitle": "3rd-level evocation",
        "casting_time": "1 action",
        "range": "150 feet",
        "components": "V, S, M (a tiny ball of bat guano and sulfur)",
        "duration": "Instantaneous",
        "description": " A bright streak flashes from your pointing finger to a point you choose within range and then blossoms with a low roar into an explosion of flame. Each creature in a 20-foot-radius sphere centered on that point must make a Dexterity saving throw. A target takes 8d6 fire damage on a failed save, or half as much damage on a successful one. The fire spreads around corners. It ignites flammable objects in the area that aren’t being worn or carried. *At Higher Levels.* When you cast this spell using a spell slot of 4th level or higher, the damage increases by 1d6 for each slot level above 3rd.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 4th level or higher, the damage increases by 1d6 for each slot level above 3rd.",
        "damage_dice": "8d6",
        "damage_type": "fire",
        "saving_throw": "dexterity"
    },
    {
        "name": "produce-flame",
        "level": 0,
        "title": "Produce Flame",
        "subtitle": "conjuration cantrip",
        "casting_time": "1 action",
        "range": "Self",
        "components": "V, S",
        "duration": "10 minutes",
        "description": " A flickering flame appears in your hand. The flame remains there for the duration and harms neither you nor your equipment. The flame sheds bright light in a 10-foot radius and dim light for an additional 10 feet. The spell ends if you dismiss it as an action or if you cast it again. You can also attack with the flame, although doing so ends the spell. When you cast this spell, or as an action on a later turn, you can hurl the flame at a creature within 30 feet of you. Make a ranged spell attack. On a hit, the target takes 1d8 fire damage. This spell’s damage increases by 1d8 when you reach 5th level (2d8), 11th level (3d8), and 17th level (4d8).",
        "damage_increase": "This spell’s damage increases by 1d8 when you reach 5th level (2d8), 11th level (3d8), and 17th level (4d8).",
        "damage_dice": "1d8",
        "damage_type": "fire",
        "attack_rolls": true
    },
    {
        "name": "acid-splash",
        "level": 0,
        "title": "Acid Splash",
        "subtitle": "conjuration cantrip",
        "casting_time": "1 action",
        "range": "60 feet",
        "components": "V, S",
        "duration": "Instantaneous",
        "description": " You hurl a bubble of acid. Choose one creature within range, or choose two creatures within range that are within 5 feet of each other. A target must succeed on a Dexterity saving throw or take 1d6 acid damage. This spell’s damage increases by 1d6 when you reach 5th level (2d6), 11th level (3d6), and 17th level (4d6).",
        "damage_increase": "This spell’s damage increases by 1d6 when you reach 5th level (2d6), 11th level (3d6), and 17th level (4d6).",
        "saving_throw": "dexterity"
    },
    {
        "name": "chill-touch",
        "level": 0,
        "title": "Chill Touch",
        "subtitle": "necromancy cantrip",
        "casting_time": "1 action",
        "range": "120 feet",
        "components": "V, S",
        "duration": "1 round",
        "description": " You create a ghostly, skeletal hand in the space of a creature within range. Make a ranged spell attack against the creature to assail it with the chill of the grave. On a hit, the target takes 1d8 necrotic damage, and it can’t regain hit points until the start of your next turn. Until then, the hand clings to the target. If you hit an undead target, it also has disadvantage on attack rolls against you until the end of your next turn. This spell’s damage increases by 1d8 when you reach 5th level (2d8), 11th level (3d8), and 17th level (4d8).",
        "damage_increase": "This spell’s damage increases by 1d8 when you reach 5th level (2d8), 11th level (3d8), and 17th level (4d8).",
        "damage_dice": "1d8",
        "damage_type": "necrotic",
        "attack_rolls": true
    },
    {
        "name": "alter-self",
        "level": 2,
        "title": "Alter Self",
        "subtitle": "2nd-level transmutation",
        "casting_time": "1 action",
        "range": "Self",
        "components": "V, S",
        "duration": "Concentration, up to 1 hour",
        "description": " You assume a different form. When you cast the spell, choose one of the following options, the effects of which last for the duration of the spell. While the spell lasts, you can end one option as an action to gain the benefits of a different one. *Aquatic Adaptation.* You adapt your body to an aquatic environment, sprouting gills and growing webbing between your fingers. You can breathe underwater and gain a swimming speed equal to your walking speed. *Change Appearance.* You transform your appearance. You decide what you look like, including your height, weight, facial features, sound of your voice, hair length, coloration, and distinguishing characteristics, if any. You can make yourself appear as a member of another race, though none of your statistics change. You also can’t appear as a creature of a different size than you, and your basic shape stays the same; if you’re bipedal, you can’t use this spell to become quadrupedal, for instance. At any time for the duration of the spell, you can use your action to change your appearance in this way again. *Natural Weapons.* You grow claws, fangs, spines, horns, or a different natural weapon of your choice. Your unarmed strikes deal 1d6 bludgeoning, piercing, or slashing damage, as appropriate to the natural weapon you chose, and you are proficient with your unarmed strikes. Finally, the natural weapon is magic and you have a +1 bonus to the attack and damage rolls you make using it."
    },
    {
        "name": "burning-hands",
        "level": 1,
        "title": "Burning Hands",
        "subtitle": "1st-level evocation",
        "casting_time": "1 action",
        "range": "Self (15-foot cone)",
        "components": "V, S",
        "duration": "Instantaneous",
        "description": " As you hold your hands with thumbs touching and fingers spread, a thin sheet of flames shoots forth from your outstretched fingertips. Each creature in a 15-foot cone must make a Dexterity saving throw. A creature takes 3d6 fire damage on a failed save, or half as much damage on a successful one. The fire ignites any flammable objects in the area that aren’t being worn or carried. *At Higher Levels.* When you cast this spell using a spell slot of 2nd level or higher, the damage increases by 1d6 for each slot level above 1st.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 2nd level or higher, the damage increases by 1d6 for each slot level above 1st.",
        "damage_dice": "3d6",
        "damage_type": "fire",
        "saving_throw": "dexterity"
    },
    {
        "name": "cure-wounds",
        "level": 1,
        "title": "Cure Wounds",
        "subtitle": "1st-level evocation",
        "casting_time": "1 action",
        "range": "Touch",
        "components": "V, S",
        "duration": "Instantaneous",
        "description": " A creature you touch regains a number of hit points equal to 1d8 + your spellcasting ability modifier. This spell has no effect on undead or constructs. *At Higher Levels.* When you cast this spell using a spell slot of 2nd level or higher, the healing increases by 1d8 for each slot level above 1st.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 2nd level or higher, the healing increases by 1d8 for each slot level above 1st.",
        "heal_dice": "1d8",
        "healing_increase": "healing increases by 1d8 for each slot level above 1st."
    },
    {
        "name": "hellish-rebuke",
        "level": 1,
        "title": "Hellish Rebuke",
        "subtitle": "1st-level evocation",
        "casting_time": "1 reaction, which you take in response to being damaged by a creature within 60 feet of you that you can see",
        "range": "60 feet",
        "components": "V, S",
        "duration": "Instantaneous",
        "description": " You point your finger, and the creature that damaged you is momentarily surrounded by hellish flames. The creature must make a Dexterity saving throw. It takes 2d10 fire damage on a failed save, or half as much damage on a successful one. *At Higher Levels.* When you cast this spell using a spell slot of 2nd level or higher, the damage increases by 1d10 for each slot level above 1st.",
        "at_higher_levels": "At Higher Levels. When you cast this spell using a spell slot of 2nd level or higher, the damage increases by 1d10 for each slot level above 1st.",
        "damage_dice": "2d10",
        "damage_type": "fire",
        "saving_throw": "dexterity"
    }]`)
	MockSpellListByClass = []byte(`{
        "Bard": {
          "level0": [
            "Dancing Lights",
            "Light",
            "Mage Hand",
            "Mending",
            "Message",
            "Minor Illusion",
            "Prestidigitation",
            "True Strike"
          ]
        },
        "Cleric": {
          "level0": [
            "Guidance",
            "Light",
            "Mending",
            "Resistance",
            "Sacred Flame",
            "Thaumaturgy"
          ],
          "level1": ["Cure Wounds"]
        },
        "Druid": {
          "level0": [
            "Guidance",
            "Mending",
            "Produce Flame",
            "Resistance",
            "Shillelagh"
          ]
        },
        "Ranger": {
          "level1": [
            "Alarm",
            "Cure Wounds",
            "Detect Magic",
            "Detect Poison and Disease",
            "Fog Cloud",
            "Jump",
            "Longstrider",
            "Speak with Animals"
          ]
        },
        "Paladin": {
          "level1": [
            "Bless",
            "Command",
            "Cure Wounds",
            "Detect Evil and Good",
            "Detect Magic",
            "Detect Poison and Disease",
            "Divine Favor",
            "Heroism",
            "Protection from Evil and Good",
            "Purify Food and Drink",
            "Shield of Faith"
          ]
        },
        "Sorcerer": {
          "level0": [
            "Acid Splash",
            "Chill Touch",
            "Dancing Lights",
            "Light",
            "Mage Hand",
            "Mending",
            "Message",
            "Minor Illusion",
            "Prestidigitation",
            "Ray of Frost",
            "Shocking Grasp",
            "True Strike"
          ],
          "level1": [ "Burning Hands" ],
          "level2": [ "Alter Self" ]
        },
        "Warlock": {
          "level0": [
            "Chill Touch",
            "Mage Hand",
            "Minor Illusion",
            "Prestidigitation",
            "True Strike"
          ]
        },
        "Wizard": {
          "level0": [
            "Acid Splash",
            "Chill Touch",
            "Dancing Lights",
            "Light",
            "Mage Hand",
            "Mending",
            "Message",
            "Minor Illusion",
            "Prestidigitation",
            "Ray of Frost",
            "Shocking Grasp",
            "True Strike"
          ],
          "level1": ["Burning Hands"],
          "level2": ["Acid Arrow"],
          "level3": ["Fireball"]
        }
      }
      `)
	MockMagicItemDatabase = []byte(`[{
        "name": "winged-boots",
        "title": "Winged Boots",
        "content": "While you wear these boots, you have a flying speed equal to your walking speed. You can use the boots to fly for up to 4 hours, all at once or in several shorter flights, each one using a minimum of 1 minute from the duration. If you are flying when the duration expires, you descend at a rate of 30 feet per round until you land.The boots regain 2 hours of flying capability for every 12 hours they aren’t in use.",
        "category": "woundrous items",
        "rarity": "uncommon",
        "hoard_table": [
            "F"
        ],
        "required_attunement": true,
        "roleplay": true,
        "forbidden": false
    },
    {
        "name": "potion-of-climbing",
        "title": "Potion of Climbing",
        "content": "When you drink this potion, you gain a climbing speed equal to your walking speed for 1 hour. During this time, you have advantage on Strength (Athletics) checks you make to climb. The potion is separated into brown, silver, and gray layers resembling bands of stone. Shaking the bottle fails to mix the colors.",
        "category": "potions",
        "rarity": "common",
        "hoard_table": [
            "A"
        ],
        "required_attunement": false,
        "roleplay": false,
        "forbidden": false,
        "magic_feature": {
            "advantages": [
                "athletics"
            ]
        }
    }]`)
	MockArmorDatabase = []byte(`[{
        "name": "padded",
        "title": "Padded",
        "kind": "light-armor",
        "cost": 5,
        "coin_type": "gold",
        "armor_class": 11,
        "dexterity_modifier": 99,
        "stealth": true,
        "strength": 0,
        "weight": 8,
        "measure": "lb"
    },
    {
        "name": "plate",
        "title": "Plate",
        "kind": "heavy-armor",
        "cost": 1500,
        "coin_type": "gold",
        "armor_class": 18,
        "dexterity_modifier": 0,
        "stealth": true,
        "strength": 15,
        "weight": 65,
        "measure": "lb"
    }]`)
	MockWeaponDatabase = []byte(`[{
        "name": "longsword",
        "title": "Longsword",
        "kind": "martial-weapon",
        "cost": 15,
        "coin_type": "gold",
        "damage": "1d8",
        "damage_two_hands": "1d10",
        "damage_type": "slashing",
        "weight": 3,
        "measure": "lb",
        "properties": "Versatile (1d10)"
    }]`)
	MockGearDatabase = []byte(`[{
        "name": "bottle glass",
        "title": "Bottle glass",
        "kind": "other",
        "cost": 2,
        "coin_type": "gold",
        "weight": 2,
        "measure": "lb",
        "number": 0
    },
    {
        "name": "bucket",
        "title": "Bucket",
        "kind": "other",
        "cost": 5,
        "coin_type": "copper",
        "weight": 2,
        "measure": "lb",
        "number": 0
    },
    {
        "name": "fake-electrum",
        "title": "Fake",
        "kind": "other",
        "cost": 5,
        "coin_type": "electrum",
        "weight": 2,
        "measure": "lb",
        "number": 0
    },
    {
        "name": "fake-platinum",
        "title": "Platinum",
        "kind": "other",
        "cost": 5,
        "coin_type": "platinum",
        "weight": 2,
        "measure": "lb",
        "number": 0
    }]`)
	MockPacksDatabase = []byte(`[{
        "name": "explorers pack",
        "title": "Explorers Pack ",
        "description": "Includes a backpack, a bedroll, a mess kit, a tinderbox, 10 torches, 10 days of rations, and a waterskin",
        "cost": 10,
        "coin_type": "gold"
    }]`)
	MockToolsDatabase = []byte(`[{
        "name": "cooks utensils",
        "title": "Cooks utensils",
        "kind": "artisans tools",
        "cost": 1,
        "coin_type": "gold",
        "weight": 8,
        "measure": "lb",
        "description": "These special tools include the items needed to pursue a craft or trade. The table shows examples of the most common types of tools, each providing items related to a single craft. Proficiency with a set of artisan’s tools lets you add your proficiency bonus to any ability checks you make using the tools in your craft. Each type of artisan’s tools requires a separate proficiency."
    }]`)
	MockMountsDatabase = []byte(`[{
        "name": "donkey or mule",
        "title": "Donkey or mule",
        "cost": 8,
        "coin_type": "gold",
        "carrying_capacity": 420,
        "carrying_capacity_measure": "lb",
        "speed": 40,
        "speed_measure": "ft"
    }]`)
	MockHoardDatabase = []byte(`[{
        "name": "zurite",
        "kind": "gemstone",
        "description": "opaque mottled deep blue",
        "value": 10,
        "coin_type": "gold"
    },
    {
        "name": "andedagate",
        "kind": "gemstone",
        "description": "translucent striped brown, blue, white, or red",
        "value": 10,
        "coin_type": "gold"
    },
    {
        "name": "loodstone",
        "kind": "gemstone",
        "description": "opaque dark gray with red flecks",
        "value": 50,
        "coin_type": "gold"
    },
    {
        "name": "Ewer",
        "kind": "art-object",
        "description": "Silver ewer",
        "value": 25,
        "coin_type": "gold"
    }]`)
	MockServicesDatabase = []byte(`[{
        "name": "lifestyle poor",
        "title": "Lifestyle Poor",
        "cost": 2,
        "coin_type": "silver",
        "unit": "per day",
        "source": "lifestyle"
    }]`)
)

//MockDatabase struct
type MockDatabase struct {
	MonsterDatabase          []byte
	SpellDescriptionDatabase []byte
	SpellListByClass         []byte
	MagicItemDatabase        []byte
	ArmorDatabase            []byte
	WeaponDatabase           []byte
	GearDatabase             []byte
	PacksDatabase            []byte
	ToolsDatabase            []byte
	MountsDatabase           []byte
	HoardDatabase            []byte
	ServicesDatabase         []byte
}

// DatabaseRepository interface
type DatabaseRepository interface {
	appcontext.Component
	//GetMonsterDatabase returns monster json database
	GetMonsterDatabase() []rule.MonsterNPC
	// GetSpellListByClass return list of spell by class
	GetSpellListByClass() rule.SpellListByClass
	//GetSpellDescriptionDatabase returns spell descriptions json database
	GetSpellDescriptionDatabase() []rule.SpellDescription
	// GetMagicItemDatabase returns magic items json database
	GetMagicItemDatabase() []rule.MagicItem
	// GetArmorDatabase returns armors json database
	GetArmorDatabase() []rule.Armor
	// GetWeaponDatabase returns weapons json database
	GetWeaponDatabase() []rule.Weapon
	// GetGearDatabase returns advantures gear to shop
	GetGearDatabase() []rule.Gear
	// GetPacksDatabase returns adventures gear packs to shop
	GetPacksDatabase() []rule.Packs
	// GetToolsDatabase returns Tools database to shop
	GetToolsDatabase() []rule.Tools
	// GetMountsDatabase returns Mounts database to shop
	GetMountsDatabase() []rule.Mounts
	// GetHoardDatabase returns TreasureHoard database to create random treasure hoards
	GetHoardDatabase() []rule.TreasureHoard
	// GetServicesDatabase returns services database
	GetServicesDatabase() []rule.Services
}

func (db MockDatabase) GetMonsterDatabase() []rule.MonsterNPC {
	var monsters []rule.MonsterNPC

	err := json.Unmarshal(MockMonsterDatabase, &monsters)
	if err != nil {
		fmt.Println(err)
	}
	return monsters
}

func (db MockDatabase) GetSpellListByClass() rule.SpellListByClass {
	var spellList rule.SpellListByClass

	err := json.Unmarshal(MockSpellListByClass, &spellList)
	if err != nil {
		fmt.Println("3", err)
	}
	return spellList
}

func (db MockDatabase) GetSpellDescriptionDatabase() []rule.SpellDescription {
	var spellList []rule.SpellDescription

	err := json.Unmarshal(MockSpellDescriptionDatabase, &spellList)
	if err != nil {
		fmt.Println(err)
	}

	return spellList
}

func (db MockDatabase) GetMagicItemDatabase() []rule.MagicItem {
	var magicItems []rule.MagicItem

	err := json.Unmarshal(MockMagicItemDatabase, &magicItems)
	if err != nil {
		fmt.Println(err)
	}

	return magicItems
}

func (db MockDatabase) GetArmorDatabase() []rule.Armor {
	var armors []rule.Armor

	err := json.Unmarshal(MockArmorDatabase, &armors)
	if err != nil {
		fmt.Println(err)
	}

	return armors
}

func (db MockDatabase) GetWeaponDatabase() []rule.Weapon {
	var weapons []rule.Weapon

	err := json.Unmarshal(MockWeaponDatabase, &weapons)
	if err != nil {
		fmt.Println(err)
	}
	unarmed := rule.Weapon{
		Name:       "unarmed",
		Kind:       "",
		Cost:       0,
		CoinType:   "",
		Damage:     "1",
		DamageType: "bludgeoning",
		Weight:     0,
		Measure:    "",
		Properties: "",
	}
	weapons = append(weapons, unarmed)

	return weapons
}

func (db MockDatabase) GetGearDatabase() []rule.Gear {
	var gears []rule.Gear

	err := json.Unmarshal(MockGearDatabase, &gears)
	if err != nil {
		fmt.Println(err)
	}

	return gears
}

func (db MockDatabase) GetPacksDatabase() []rule.Packs {
	var packs []rule.Packs

	err := json.Unmarshal(MockPacksDatabase, &packs)
	if err != nil {
		fmt.Println(err)
	}

	return packs
}

func (db MockDatabase) GetToolsDatabase() []rule.Tools {
	var tools []rule.Tools

	err := json.Unmarshal(MockToolsDatabase, &tools)
	if err != nil {
		fmt.Println(err)
	}

	return tools
}

func (db MockDatabase) GetMountsDatabase() []rule.Mounts {
	var mounts []rule.Mounts

	err := json.Unmarshal(MockMountsDatabase, &mounts)
	if err != nil {
		fmt.Println(err)
	}

	return mounts
}

func (db MockDatabase) GetHoardDatabase() []rule.TreasureHoard {
	var hoards []rule.TreasureHoard

	err := json.Unmarshal(MockHoardDatabase, &hoards)
	if err != nil {
		fmt.Println(err)
	}

	return hoards
}

func (db MockDatabase) GetServicesDatabase() []rule.Services {
	var hoards []rule.Services
	err := json.Unmarshal(MockServicesDatabase, &hoards)
	if err != nil {
		fmt.Println(err)
	}

	return hoards
}

func lazyMaster() appcontext.Component {
	return &MockDatabase{
		MonsterDatabase:          MockMonsterDatabase,
		SpellDescriptionDatabase: MockSpellDescriptionDatabase,
		SpellListByClass:         MockSpellListByClass,
		MagicItemDatabase:        MockMagicItemDatabase,
		ArmorDatabase:            MockArmorDatabase,
		WeaponDatabase:           MockWeaponDatabase,
		GearDatabase:             MockGearDatabase,
		PacksDatabase:            MockPacksDatabase,
		ToolsDatabase:            MockToolsDatabase,
		MountsDatabase:           MockMountsDatabase,
		HoardDatabase:            MockHoardDatabase,
		ServicesDatabase:         MockServicesDatabase,
	}
}

// GetDatabaseRepository func return DatabaseRepository interface
func GetDatabaseRepository() DatabaseRepository {
	return appcontext.Current.Get(appcontext.Database).(DatabaseRepository)
}

// InitDatabaseMock func returns a RepositoryDatabaseMock interface
func InitDatabaseMock() appcontext.Component {
	return lazyMaster()
}
