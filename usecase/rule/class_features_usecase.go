package rule

import (
	"fmt"

	"github.com/betorvs/playbypost-dnd/utils"
)

//
func humanoidList() []string {
	return []string{"bugbear", "gnoll", "goblin", "grimlock", "hobgoblin", "kobold", "lizardfolk", "merfolk", "orc", "sahuagin", "thug", "werebear", "wereboar", "wererat", "weretiger", "werewolf", "gnome", "human", "elf", "dwarf"}
}

func monsterTypeList() []string {
	return []string{"aberrations", "beasts", "celestials", "constructs", "dragons", "elementals", "fey", "fiends", "giants", "monstrosities", "oozes", "plants", "undead"}
}

// "natural-explorer", "natural-explorer-improvement", "favored-enemy", "favored-enemy-improvement",

func featuresListRename() []string {
	return []string{"fighting-style", "primal-path", "path-feature", "bard-college", "bard-college-feature", "divine-domain", "divine-domain-feature", "druid-circle", "druid-circle-feature", "martial-archetype", "martial-archetype-feature", "sacred-oath", "sacred-oath-feature", "monastic-tradition", "monastic-tradition-feature", "ranger-archetype", "ranger-archetype-feature", "roguish-archetype", "roguish-archetype-feature", "sorcerous-origin", "sorcerous-origin-feature", "otherworldly-patron", "otherworldly-patron-feature", "pact-boon", "arcane-tradition", "arcane-tradition-feature"}
}

func terrainList() []string {
	return []string{"arctic", "coast", "desert", "forest", "grassland", "mountain", "swamp", "underdark"}
}

func dragonKinds() []string {
	return []string{"black", "blue", "brass", "bronze", "copper", "gold", "green", "red", "silver", "white"}
}

func fightingStyleList() []string {
	return []string{"archery", "defense", "dueling", "great-weapon-fighting", "protection", "two-weapon-fighting"}
}

func fightingStyleFullNameList() []string {
	return []string{"fighting-style-defense", "fighting-style-dueling", "fighting-style-great-weapon-fighting", "fighting-style-protection", "fighting-style-two-weapon-fighting"}
}

func barbarianPrimalPathList() []string {
	return []string{"berseker", "totem-warrior"}
}

func bardCollege() []string {
	return []string{"lore", "valor"}
}

func clericDivineDomainList() []string {
	return []string{"knowledge", "life", "light", "nature", "tempest", "trickery", "war"}
}

func druidCircleList() []string {
	return []string{"land", "moon"}
}

func fighterArchetypeList() []string {
	return []string{"champion", "battle-master", "eldritch-knight"}
}

func monkMonasticTraditionList() []string {
	return []string{"open-hand", "shadow", "four-elements"}
}

func paladinSacredOaths() []string {
	return []string{"devotion", "ancients", "vengeance"}
}

func rangerArchetypesList() []string {
	return []string{"hunter", "beast-master"}
}

func rogueArchetypeList() []string {
	return []string{"thief", "assassin", "arcane-trickster"}
}

func sorcererOriginList() []string {
	return []string{"draconic-bloodline", "wild-magic"}
}

func warlockOtherworldlyPatronList() []string {
	return []string{"archfey", "fiend", "great-old-one"}
}

func warlockPactBoonList() []string {
	return []string{"chain", "blade", "tome"}
}

func wizardArcaneTraditionList() []string {
	return []string{"abjuration", "conjuration", "divination", "enchantment", "evocation", "illusion", "necromancy", "transmutation"}
}

// func rangerHuntersFullList() (fullList []string) {
// 	fullList = append(fullList, huntersPreyList()...)
// 	fullList = append(fullList, defensiveTacticsList()...)
// 	fullList = append(fullList, multiattackList()...)
// 	fullList = append(fullList, huntersDefenseList()...)
// 	return fullList
// }

func huntersPreyList() []string {
	return []string{"colossus-slayer", "giant-killer", "horde-breaker"}
}
func defensiveTacticsList() []string {
	return []string{"escape-the-horde", "multiattack-defense", "steel-will"}
}
func multiattackList() []string {
	return []string{"volley", "whirlwind-attack"}
}
func huntersDefenseList() []string {
	return []string{"evasion", "stand-against-the-tide", "uncanny-dodge"}
}

func turnMonstersFeatureList() []string {
	return []string{
		"sacred-oath-of-devotion-channel-divinity",
		"sacred-oath-of-ancients-channel-divinity",
		"channel-divinity",
	}
}

// func archetypeListByClass(class string) []string {
// 	switch class {
// 	case "barbarian":
// 		return barbarianPrimalPathList()
// 	case "bard":
// 		return bardCollege()
// 	case "cleric":
// 		return clericDivineDomainList()
// 	case "druid":
// 		return druidCircleList()
// 	case "fighter":
// 		return fighterArchetypeList()
// 	case "monk":
// 		return monkMonasticTraditionList()
// 	case "paladin":
// 		return paladinSacredOaths()
// 	case "ranger":
// 		return rangerArchetypesList()
// 	case "rogue":
// 		return rogueArchetypeList()
// 	case "sorcerer":
// 		return sorcererOriginList()
// 	case "warlock":
// 		return warlockOtherworldlyPatronList()
// 	case "wizard":
// 		return wizardArcaneTraditionList()

// 	default:
// 		return []string{}
// 	}
// }

func choosenClassFeatures(name, choosen string, level int) []string {
	switch name {
	case "fighting-style":
		switch choosen {
		case "archery":
			return []string{"fighting-style-archery"}
		case "defense":
			return []string{"fighting-style-defense"}
		case "dueling":
			return []string{"fighting-style-dueling"}
		case "great-weapon-fighting":
			return []string{"fighting-style-great-weapon-fighting"}
		case "protection":
			return []string{"fighting-style-protection"}
		case "two-weapon-fighting":
			return []string{"fighting-style-two-weapon-fighting"}

		}

	case "primal-path", "path-feature":
		switch choosen {
		case "berseker":
			// return []string{ "path-of-berseker"}
			if level == 3 {
				// extra attack as bonus action and exhaustion 1
				return []string{"path-of-berseker-frenzy"}
			}
			if level == 6 {
				return []string{"path-of-berseker-intimidating-presence"}
			}
			if level == 10 {
				return []string{"path-of-berseker-mindless-rage"}
			}
			if level == 14 {
				return []string{"path-of-berseker-retalion"}
			}
		case "totem", "totem-warrior":
			if level == 3 {
				// you gain the ability to cast the beast sense and speak with animals spells, but only as rituals,
				// when raging
				//   bear have resistance to all damage except psychic damage
				//   eagle other creatures have disadvantage on opportunity attack rolls against you, and you can use the Dash action as a bonus action on your turn.
				//   wolf your friends have advantage on melee attack rolls against any creature within 5 feet of you that is hostile to you
				return []string{"path-of-totem-warrior-totem-spirit", "path-of-totem-warrior-spirit-seeker"}
			}
			if level == 6 {
				return []string{"path-of-totem-warrior-aspect-of-the-beast"}
			}
			if level == 10 {
				return []string{"path-of-totem-warrior-spirit-walker"}
			}
			if level == 14 {
				return []string{"path-of-totem-warrior-totemic-attunement"}
			}
		}

	// bard colleges
	case "bard-college", "bard-college-feature":
		switch choosen {
		case "lore":
			if level == 3 {
				// +3 skills any in character creation OK
				// rolling
				// a Bardic Inspiration die and subtracting the number rolled from the creature's roll
				return []string{"college-of-lore-bonus-proficiencies", "college-of-lore-cutting-words"}
			}
			if level == 6 {
				return []string{"college-of-lore-additional-magical-secrets"}
			}
			if level == 14 {
				return []string{"college-of-lore-peerless-skill"}
			}
		case "valor":
			if level == 3 {
				// proficiency medium-armor, shields, martial-weapon
				// has a Bardic Inspiration die from you can roll that die and add the number rolled to a weapon damage roll it just made.
				// roll the Bardic Inspiration die and add the number rolled to its AC against that attack
				return []string{"college-of-valor-bonus-proficiencies", "college-of-valor-combat-inspiration"}
			}
			if level == 6 {
				return []string{"college-of-valor-extra-attack"}
			}
			if level == 14 {
				return []string{"college-of-valor-battle-magic"}
			}
		}

	// cleric domains
	case "divine-domain", "divine-domain-feature":
		switch choosen {
		case "knowledge":
			if level == 1 {
				// + languages + 2 skills with double proficiency: arcana, history, nature, religion
				// character creation done
				return []string{"domain-knowledge", "blessings-of-knowledge"}
			}
			if level == 2 {
				// add a proficiency when didnt have one
				// under activation
				return []string{"channel-divinity-knowledge-of-the-ages"}
			}
			if level == 6 {
				//read mind target saving wisdom and suggestion magic without using slot
				// under activation
				return []string{"channel-divinity-read-thoughts"}
			}
			if level == 8 {
				// add wisdom modifier in damage in cantrips
				// automatic calc in spell
				return []string{"potent-spellcasting"}
			}
			if level == 17 {
				// ritual to discovery things
				// under activation without rolls
				return []string{"visions-of-the-past"}
			}

		case "life":
			if level == 1 {
				// proficiency heavy-armor character creation ok
				// healing spells add 2 + spell level to healing value done
				// automatic calc in healing spell
				return []string{"domain-life", "domain-life-bonus-proficiency", "disciple-of-life"}
			}
			if level == 2 {
				// similar paladin lays-on-hands
				// can heal 5 * cleric level
				// under activation
				return []string{"channel-divinity-preserve-life"}
			}
			if level == 6 {
				// when use healing spell higher than 1 level
				// you recover 2 + spell level done
				// automatic calc when healing spell
				return []string{"blessed-healer"}
			}
			if level == 8 {
				// extra radiant damage 1d8
				//  on 14th level 2d8
				// under activation and in one attack
				return []string{"domain-life-divine-strike"}
			}
			if level == 17 {
				// all healing dices uses maximum value (1d6 is 6...)
				// special feature when calc spell
				// maybe a maximum effect instead rolling dices
				return []string{"supreme-healing"}
			}
		case "light":
			if level == 1 {
				// one cantrip extra done
				// impose disvantage on attack roll on you
				//  if target cannot be blinded
				//  wisdom mofifier * 1 number of times per day
				// under activation
				return []string{"domain-light", "domain-light-bonus-cantrip", "warding-flare"}
			}
			if level == 2 {
				// damage radiant 2d10 + cleric level to all creatures saving constitution
				//  if pass half of damage done
				// under activation
				return []string{"channel-divinity-radiance-of-the-dawn"}
			}
			if level == 6 {
				// use warding-flare in another target instead only when attack you
				// same 1 level but others
				return []string{"improved-flare"}
			}
			if level == 8 {
				// add wisdom modifier in damage in cantrips
				// automatic calc spells
				return []string{"potent-spellcasting"}
			}
			if level == 17 {
				// impose disvantage on saving for all enemies against fire and radiant spells
				// under activation
				return []string{"corona-of-light"}
			}
		case "nature":
			if level == 1 {
				// +1 druid cantrip maybe add spell druid in class feature to avoid non list check done
				// +1 skill: animal-handling, nature, survival. character creation done
				// heavy-armor character creation done
				return []string{"domain-nature", "acolyte-of-nature", "domain-nature-bonus-proficiency"}
			}
			if level == 2 {
				// charmed animals and plants
				// saving wisdom: fails charmed for 1 min or first damage
				//
				return []string{"channel-divinity-charm-animals-and-plants"}
			}
			if level == 6 {
				// you can add resistance in any creature near you
				// for: acid, cold, fire, lightning, or thunder
				return []string{"dampen-elements"}
			}
			if level == 8 {
				// extra one cold, fire, lightning damage 1d8
				//  on 14th level 2d8
				return []string{"domain-nature-divine-strike"}
			}
			if level == 17 {
				// your charmed animals and plants can be commanded
				return []string{"master-of-nature"}
			}
		case "tempest":
			if level == 1 {
				// proficiency martial-weapons, heavy-armor done
				// use reaction action to cause 2d8 lightning or thunder on creature that attacks you
				// target dexterity saving or half of damage
				// wisdom modifier / per day
				return []string{"domain-tempest", "domain-tempest-bonus-proficiencies", "wrath-of-the-storm"}
			}
			if level == 2 {
				// when lightning or thunder damage you can use your channel divinity to maximum damage
				// done
				return []string{"channel-divinity-destructive-wrath"}
			}
			if level == 6 {
				//  when you deal lightning damage to a Large or smaller creature, you can also push it up to 10 feet away from you.
				return []string{"thunderbolt-strike"}
			}
			if level == 8 {
				// extra thunder damage 1d8
				//  on 14th level 2d8
				return []string{"domain-tempest-divine-strike"}
			}
			if level == 17 {
				//  you have a flying speed equal to your current walking speed whenever you are not underground or indoors.
				return []string{"stormborn"}
			}
		case "trickery":
			if level == 1 {
				// to give it advantage on Dexterity (Stealth) checks. done
				return []string{"domain-trickery", "blessing-of-the-trickster"}
			}
			if level == 2 {
				// you can use your Channel Divinity to create an illusory duplicate of yourself.
				//
				return []string{"channel-divinity-invoke-duplicity"}
			}
			if level == 6 {
				// As an action, you become invisible until the end of your next turn.
				// You become visible ifyou attack or cast a spell.
				return []string{"channel-divinity-cloak-of-shadows"}
			}
			if level == 8 {
				// // extra poison damage 1d8
				//  on 14th level 2d8
				return []string{"domain-trickery-divine-strike"}
			}
			if level == 17 {
				// you can create up to four duplicates
				//  of yourself, instead of one, when you use Invoke Duplicity.
				return []string{"improved-duplicity"}
			}
		case "war":
			if level == 1 {
				// // proficiency martial-weapons, heavy-armor done
				// When you use the Attack action, you can make one weapon attack as a bonus action.
				// wisdom modifier / per day
				return []string{"domain-war", "domain-war-bonus-proficiencies", "war-priest"}
			}
			if level == 2 {
				// you can use your Channel Divinity to gain a +10 bonus to the roll.
				// use temporary bonus done
				return []string{"channel-divinity-guided-strike"}
			}
			if level == 6 {
				// you can use your reaction to grant that creature a +10 bonus to the roll
				return []string{"channel-divinity-war-gods-blessing"}
			}
			if level == 8 {
				// // extra weapon type damage 1d8
				//  on 14th level 2d8
				return []string{"domain-war-divine-strike"}
			}
			if level == 17 {
				// you gain resistance to bludgeoning, piercing, and slashing damage from nonmagical weapons.
				return []string{"avatar-of-battle"}
			}
		}

	// druid cicle
	case "druid-circle", "druid-circle-feature":
		switch choosen {
		case "arctic", "coast", "desert", "forest", "grassland", "mountain", "swamp", "underdark":
			var value string
			if level == 2 {
				if utils.StringInSlice(choosen, terrainList()) {
					value = fmt.Sprintf("circle-of-the-land-spells-%s", choosen)
				}
			}
			return []string{value}
		case "land":
			if level == 2 {
				// +1 cantrip done
				// short rest recover slot spell equal half druid level
				value := "circle-of-the-land-circle-spells"
				return []string{value, "circle-of-the-land-bonus-cantrip", "circle-of-the-land-natural-recovery"}
			}
			if level == 6 {
				return []string{"circle-of-the-land-lands-stride"}
			}
			if level == 10 {
				return []string{"circle-of-the-land-natures-ward"}
			}
			if level == 14 {
				return []string{"circle-of-the-land-natures-sanctuary"}
			}

		case "moon":
			if level == 2 {
				// wild shape as bonus action
				// bonus action when in wild shape spend spell slot to revover
				//  1d8 per spell level ok
				return []string{"circle-of-the-moon-combat-wild-shape", "circle-of-the-moon-forms"}
			}
			if level == 6 {
				return []string{"circle-of-the-moon-primal-strike"}
			}
			if level == 10 {
				return []string{"circle-of-the-moon-elemental-wild-shape"}
			}
			if level == 14 {
				return []string{"circle-of-the-moon-thousand-forms"}
			}
		}
		// fighter archetype
	case "martial-archetype", "martial-archetype-feature":
		switch choosen {
		case "champion":
			if level == 3 {
				// critical hit 19 and 20 done
				return []string{"archetype-champion-improved-critical"}
			}
			if level == 7 {
				return []string{"archetype-champion-remarkable-athlete"}
			}
			if level == 10 {
				return []string{"archetype-champion-additional-fighting-style"}
			}
			if level == 15 {
				// critical hit 18, 19 and 20
				return []string{"archetype-champion-superior-critical"}
			}
			if level == 18 {
				return []string{"archetype-champion-survivor"}
			}

		case "battle-master":

			if level == 3 {
				// 4 d8 superiority dice
				// +1 7th and +1 15th
				// saving throw  8 + your proficiency bonus + your Strength or Dexterity modifier (your choice)
				// 3 maneuvers +2 at 7th 10th 15th
				return []string{"archetype-battle-master-combat-superiority", "archetype-battle-master-student-of-war"}
			}
			if level == 7 {
				return []string{"archetype-battle-master-know-your-enemy"}
			}
			if level == 10 {
				return []string{"archetype-battle-master-improved-combat-superiority"}
			}
			if level == 15 {
				return []string{"archetype-battle-master-relentless"}
			}
			if level == 18 {
				return []string{"archetype-battle-master-improved-combat-superiority"}
			}
		case "eldritch-knight":
			if level == 3 {
				// spellcast ability using intelligence done
				// bond weapon
				return []string{"archetype-eldritch-knight-spellcasting", "archetype-eldritch-knight-weapon-bond"}
			}
			if level == 7 {
				return []string{"archetype-eldritch-knight-war-magic"}
			}
			if level == 10 {
				return []string{"archetype-eldritch-knight-eldritch-strike"}
			}
			if level == 15 {
				return []string{"archetype-eldritch-knight-arcane-charge"}
			}
			if level == 18 {
				return []string{"archetype-eldritch-knight-improved-war-magic"}
			}
		}
	case "monastic-tradition", "monastic-tradition-feature":
		// monk monastic tradition
		switch choosen {
		case "open-hand":
			if level == 3 {
				// after successfully attack using flurry-of-blows
				// It must succeed on a Dexterity saving throw or be knocked prone.
				// It must make a Strength saving throw. If it fails, you can push it up to 1 5 feet away from you.
				// It can't take reactions until the end ofyour next turn.
				return []string{"monastic-tradition-way-of-the-open-hand-technique"}
			}
			if level == 6 {
				return []string{"monastic-tradition-way-of-the-open-wholeness-of-body"}
			}
			if level == 11 {
				return []string{"monastic-tradition-way-of-the-open-tranquility"}
			}
			if level == 17 {
				return []string{"monastic-tradition-way-of-the-open-quivering-palm"}
			}
		case "shadow":
			if level == 3 {
				// As an action, you can spend 2 ki points to cast darkness, darkvision, pass without trace, or silence, without providing material components.
				// Additionally, you gain the minor-illusion cantrip if you don't already know it.
				return []string{"monastic-tradition-way-of-the-shadow-arts"}
			}
			if level == 6 {
				return []string{"monastic-tradition-way-of-the-shadow-step"}
			}
			if level == 11 {
				return []string{"monastic-tradition-way-of-the-shadow-cloak-of-shadows"}
			}
			if level == 17 {
				return []string{"monastic-tradition-way-of-the-shadow-opportunist"}
			}
		case "four-elements":
			if level == 3 {
				// list of magic using ki points
				// elemental-attunement
				// fangs-of-the-fire-snake A hit with such an attack deals fire damage instead of bludgeoning damage, and if you spend 1 ki point when the attack hits, it also deals an extra ldlO fire damage.
				// fist-of-four-thunders  can spend 2 ki points to cast thunderwave.
				// fist-of-unbroken-air As an action, you can spend 2 ki points and choose a creature within 30 feet ofyou. That creature must make a Strength saving throw. On a failed save, the creature takes 3d10 bludgeoning damage, plus an extra ldlO bludgeoning damage for each additional ki point you spend, and you can push the creature up to 20 feet away from you and knock it prone. On a successful save, the creature takes half as much damage, and you don't push it or knock it prone.
				// Rush ofthe Gale Spirits can spend 2 ki points to cast gust ofwind.
				// Shape the Flowing River
				// sweeping-cinder Strike spend 2 ki points to cast burning hands.
				// Water Whip can spend 2 ki points as an action to create a whip of water that shoves and pulls a creature to unbalance it. A creature that you can see that is within 30 feet of you must make a Dexterity saving throw. On a failed save, the creature takes 3d10 bludgeoning damage, plus an extra ldlO bludgeoning damage for each additional ki point you spend, and you can either knock it prone or pull it up to 25 feet closer to you. On a successful save, the creature takes half as much damage, and you don't pull it or knock it prone.

				return []string{"monastic-tradition-way-of-the-four-elements-disciple-of-the-elements"}
			}
			if level == 6 {
				return []string{"monastic-tradition-way-of-the-four-elements-disciple-of-the-elements-6"}
			}
			if level == 11 {
				return []string{"monastic-tradition-way-of-the-four-elements-disciple-of-the-elements-11"}
			}
			if level == 17 {
				return []string{"monastic-tradition-way-of-the-four-elements-disciple-of-the-elements-17"}
			}
		}
	case "sacred-oath", "sacred-oath-feature":
		// paladin sacred oaths
		switch choosen {
		case "devotion":
			if level == 3 {
				//
				return []string{"sacred-oath-of-devotion", "sacred-oath-of-devotion-channel-divinity"}
			}
			if level == 7 {
				return []string{"sacred-oath-of-devotion-aura-of-devotion"}
			}
			if level == 15 {
				return []string{"sacred-oath-of-devotion-purity-of-spirit"}
			}
			if level == 20 {
				return []string{"sacred-oath-of-devotion-holy-nimbus"}
			}

		case "ancients":
			if level == 3 {
				return []string{"sacred-oath-of-ancients", "sacred-oath-of-ancients-channel-divinity"}
			}
			if level == 7 {
				return []string{"sacred-oath-of-ancients-aura-of-warding"}
			}
			if level == 15 {
				return []string{"sacred-oath-of-ancients-undying-sentinel"}
			}
			if level == 20 {
				return []string{"sacred-oath-of-ancients-elder-champion"}
			}
		case "vengeance":
			if level == 3 {
				return []string{"sacred-oath-of-vengeance", "sacred-oath-of-vengeance-channel-divinity"}
			}
			if level == 7 {
				return []string{"sacred-oath-of-vengeance-relentless-avenger"}
			}
			if level == 15 {
				return []string{"sacred-oath-of-vengeance-soul-of-vengeance"}
			}
			if level == 20 {
				return []string{"sacred-oath-of-vengeance-avenging-angel"}
			}
		}

	case "ranger-archetype", "ranger-archetype-feature":
		switch choosen {
		// ranger archetypes hunters-prey
		case "hunter", "colossus-slayer", "giant-killer", "horde-breaker", "escape-the-horde", "multiattack-defense", "steel-will", "volley", "whirlwind-attack", "evasion", "stand-against-the-tide", "uncanny-dodge":
			if level == 3 {
				// huntersPreyList := []string{"colossus-slayer", "giant-killer", "horde-breaker"}
				value := "archetype-hunter-hunters-prey"
				if utils.StringInSlice(choosen, huntersPreyList()) {
					value = fmt.Sprintf("archetype-hunter-hunters-prey-%s", choosen)
				}
				return []string{value}
			}
			if level == 7 {
				// defensiveTacticsList := []string{"escape-the-horde", "multiattack-defense", "steel-will"}
				value := "archetype-hunter-defensive-tactics"
				if utils.StringInSlice(choosen, defensiveTacticsList()) {
					value = fmt.Sprintf("archetype-hunter-defensive-tactics-%s", choosen)
				}
				return []string{value}
			}
			if level == 11 {
				// multiattackList := []string{ "volley", "whirlwind-attack"}
				value := "archetype-hunter-multiattack"
				if utils.StringInSlice(choosen, multiattackList()) {
					value = fmt.Sprintf("archetype-hunter-multiattack-%s", choosen)
				}
				return []string{value}
			}
			if level == 15 {
				// huntersDefenseList := []string{"evasion", "stand-against-the-tide", "uncanny-dodge"}
				value := "archetype-hunter-superior-hunters-defense"
				if utils.StringInSlice(choosen, huntersDefenseList()) {
					value = fmt.Sprintf("archetype-hunter-superior-hunters-defense-%s", choosen)
				}
				return []string{value}
			}

		case "beast-master":
			if level == 3 {
				return []string{"archetype-beast-master-rangers-companion"}
			}
			if level == 7 {
				return []string{"archetype-beast-master-exceptional-training"}
			}
			if level == 11 {
				return []string{"archetype-beast-master-bestial-fury"}
			}
			if level == 15 {
				return []string{"archetype-beast-master-share-spells"}
			}
		}

	case "natural-explorer", "natural-explorer-improvement":
		switch choosen {
		// ranger natural explorer
		// arctic, coast, desert, forest, grassland, mountain, swamp, underdark.
		case "arctic":
			return []string{"natural-explorer-of-arctic"}
		case "coast":
			return []string{"natural-explorer-of-coast"}
		case "desert":
			return []string{"natural-explorer-of-desert"}
		case "forest":
			return []string{"natural-explorer-of-forest"}
		case "grassland":
			return []string{"natural-explorer-of-grassland"}
		case "mountain":
			return []string{"natural-explorer-of-mountain"}
		case "swamp":
			return []string{"natural-explorer-of-swamp"}
		case "underdark":
			return []string{"natural-explorer-of-underdark"}
		}

	// ranger favored enemy
	//  humanoid: bugbear gnoll goblin grimlock hobgoblin kobold lizardfolk merfolk orc sahuagin thug werebear wereboar wererat weretiger werewolf gnome human elf dwarf
	// aberrations, beasts, celestials, constructs, dragons, elementals, fey, fiends, giants, monstrosities, oozes, plants, undead
	case "favored-enemy", "favored-enemy-improvement":
		//
		switch choosen {
		case "aberrations":
			return []string{"favored-enemy-aberrations"}
		case "beasts":
			return []string{"favored-enemy-beasts"}
		case "celestials":
			return []string{"favored-enemy-celestials"}
		case "constructs":
			return []string{"favored-enemy-constructs"}
		case "dragons":
			return []string{"favored-enemy-dragons"}
		case "elementals":
			return []string{"favored-enemy-elementals"}
		case "fey":
			return []string{"favored-enemy-fey"}
		case "fiends":
			return []string{"favored-enemy-fiends"}
		case "giants":
			return []string{"favored-enemy-giants"}
		case "monstrosities":
			return []string{"favored-enemy-monstrosities"}
		case "oozes":
			return []string{"favored-enemy-oozes"}
		case "plants":
			return []string{"favored-enemy-plants"}
		case "undead":
			return []string{"favored-enemy-undead"}
		case "bugbear":
			return []string{"favored-enemy-bugbear"}
		case "gnoll":
			return []string{"favored-enemy-gnoll"}
		case "goblin":
			return []string{"favored-enemy-goblin"}
		case "grimlock":
			return []string{"favored-enemy-grimlock"}
		case "hobgoblin":
			return []string{"favored-enemy-hobgoblin"}
		case "kobold":
			return []string{"favored-enemy-kobold"}
		case "lizardfolk":
			return []string{"favored-enemy-lizardfolk"}
		case "merfolk":
			return []string{"favored-enemy-merfolk"}
		case "orc":
			return []string{"favored-enemy-orc"}
		case "sahuagin":
			return []string{"favored-enemy-sahuagin"}
		case "thug":
			return []string{"favored-enemy-thug"}
		case "werebear":
			return []string{"favored-enemy-werebear"}
		case "wereboar":
			return []string{"favored-enemy-wereboar"}
		case "wererat":
			return []string{"favored-enemy-wererat"}
		case "weretiger":
			return []string{"favored-enemy-weretiger"}
		case "werewolf":
			return []string{"favored-enemy-werewolf"}
		case "gnome":
			return []string{"favored-enemy-gnome"}
		case "human":
			return []string{"favored-enemy-human"}
		case "elf":
			return []string{"favored-enemy-elf"}
		case "dwarf":
			return []string{"favored-enemy-dwarf"}

		}
	case "roguish-archetype", "roguish-archetype-feature":
		switch choosen {
		//rogue archetype
		case "thief":
			if level == 3 {
				return []string{"archetype-thief-fast-hands-and-second-storywork"}
			}
			if level == 9 {
				return []string{"archetype-thief-supreme-sneak"}
			}
			if level == 13 {
				return []string{"archetype-thief-use-magic-device"}
			}
			if level == 17 {
				return []string{"archetype-thief-thiefs-reflexes"}
			}

		case "assassin":
			if level == 3 {
				// under approved by master
				return []string{"archetype-assassin-bonus-proficiencies", "archetype-assassin-assassinate"}
			}
			if level == 9 {
				return []string{"archetype-assassin-infiltration-expertise"}
			}
			if level == 13 {
				return []string{"archetype-assassin-impostor"}
			}
			if level == 17 {
				return []string{"archetype-assassin-death-strike"}
			}
		case "arcane-trickster":
			if level == 3 {
				return []string{"archetype-arcane-trickster-spellcasting", "archetype-arcane-trickster-mage-hand-legerdemain"}
			}
			if level == 9 {
				return []string{"archetype-arcane-trickster-magical-ambush"}
			}
			if level == 13 {
				return []string{"archetype-arcane-trickster-versatile-trickster"}
			}
			if level == 17 {
				return []string{"archetype-arcane-trickster-spell-thief"}
			}
		}

	case "sorcerous-origin", "sorcerous-origin-feature":
		switch choosen {
		// sorcerer origin done
		case "black", "blue", "brass", "bronze", "copper", "gold", "green", "red", "silver", "white":
			var value string
			if utils.StringInSlice(choosen, dragonKinds()) {
				value = fmt.Sprintf("sorcerous-origin-draconic-bloodline-%s-dragon-ancestor", choosen)
			}
			return []string{value}
		case "draconic", "draconic-bloodline":
			if level == 1 {
				value := "sorcerous-origin-draconic-bloodline-dragon-ancestor"
				return []string{value, "sorcerous-origin-draconic-resistance"}

			}
			if level == 6 {
				return []string{"sorcerous-origin-draconic-bloodline-elemental-affinity"}
			}
			if level == 14 {
				return []string{"sorcerous-origin-draconic-bloodline-dragon-wings"}
			}
			if level == 18 {
				return []string{"sorcerous-origin-draconic-bloodline-draconic-presence"}
			}

		case "wild", "wild-magic":
			if level == 1 {
				// ???
				return []string{"sorcerous-origin-wild-magic-wild-magic-surge-and-tides-of-chaos"}
			}
			if level == 6 {
				return []string{"sorcerous-origin-wild-magic-bend-luck"}
			}
			if level == 14 {
				return []string{"sorcerous-origin-wild-magic-controlled-chaos"}
			}
			if level == 18 {
				return []string{"sorcerous-origin-wild-magic-spell-bombardment"}
			}

		}
	case "otherworldly-patron", "otherworldly-patron-feature":
		switch choosen {
		// warlock otherworldly-patron
		case "archfey":
			if level == 1 {
				return []string{"otherworldly-patron-the-archfey-expanded-spell-list", "otherworldly-patron-the-archfey-fey-presence"}
			}
			if level == 6 {
				return []string{"otherworldly-patron-the-archfey-misty-escape"}
			}
			if level == 10 {
				return []string{"otherworldly-patron-the-archfey-beguiling-defenses"}
			}
			if level == 14 {
				return []string{"otherworldly-patron-the-archfey-dark-delirium"}
			}
		case "fiend":
			if level == 1 {
				return []string{"otherworldly-patron-the-fiend-expanded-spell-list", "otherworldly-patron-the-fiend-dark-ones-blessing"}
			}
			if level == 6 {
				return []string{"otherworldly-patron-the-fiend-dark-ones-own-luck"}
			}
			if level == 10 {
				return []string{"otherworldly-patron-the-fiend-fiendish-resilience"}
			}
			if level == 14 {
				return []string{"otherworldly-patron-the-fiend-hurl-through-hell"}
			}
		case "old-one", "great-old-one":
			if level == 1 {
				return []string{"otherworldly-patron-the-great-old-one-expanded-spell-list", "otherworldly-patron-the-great-old-one-awakened-mind"}
			}
			if level == 6 {
				return []string{"otherworldly-patron-the-great-old-one-entropic-ward"}
			}
			if level == 10 {
				return []string{"otherworldly-patron-the-great-old-one-thought-shield"}
			}
			if level == 14 {
				return []string{"otherworldly-patron-the-great-old-one-create-thrall"}
			}
		}
	case "pact-boon":
		switch choosen {
		// warlock pact boon re read !!
		case "chain":
			return []string{"pact-of-the-chain"}
		case "blade":
			return []string{"pact-of-the-blade"}
		case "tome":
			return []string{"pact-of-the-tome"}
		}

	case "arcane-tradition", "arcane-tradition-feature":
		switch choosen {
		case "abjuration":
			if level == 2 {
				return []string{"arcane-tradition-school-of-abjuration-savant", "arcane-tradition-school-of-abjuration-arcane-ward"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-abjuration-projected-ward"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-abjuration-improved-abjuration"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-abjuration-spell-resistance"}
			}

		case "conjuration":
			if level == 2 {
				return []string{"arcane-tradition-school-of-conjuration-savant", "arcane-tradition-school-of-conjuration-minor-conjuration"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-conjuration-benign-transposition"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-conjuration-focused-conjuration"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-conjuration-durable-summons"}
			}
		case "divination":
			if level == 2 {
				return []string{"arcane-tradition-school-of-divination-savant", "arcane-tradition-school-of-divination-portent"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-divination-expert-divination"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-divination-the-third-eye"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-divination-greater-portent"}
			}
		case "enchantment":
			if level == 2 {
				return []string{"arcane-tradition-school-of-enchantment-savant", "arcane-tradition-school-of-enchantment-hypnotic-gaze"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-enchantment-instinctive-charm"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-enchantment-split-enchantment"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-enchantment-alter-memories"}
			}
		case "evocation":
			if level == 2 {
				return []string{"arcane-tradition-school-of-evocation-savant", "arcane-tradition-school-of-evocation-sculpt-spells"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-evocation-potent-cantrip"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-evocation-empowered-evocation"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-evocation-overchannel"}
			}
		case "illusion":
			if level == 2 {
				return []string{"arcane-tradition-school-of-illusion-savant", "arcane-tradition-school-of-illusion-improved-minor-illusion"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-illusion-malleable-illusions"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-illusion-illusory-self"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-illusion-illusory-reality"}
			}
		case "necromancy":
			if level == 2 {
				return []string{"arcane-tradition-school-of-necromancy-savant", "arcane-tradition-school-of-necromancy-grim-harvest"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-necromancy-undead-thralls"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-necromancy-inured-to-undeath"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-necromancy-command-undead"}
			}
		case "transmutation":
			if level == 2 {
				return []string{"arcane-tradition-school-of-transmutation-savant", "arcane-tradition-school-of-transmutation-minor-alchemy"}
			}
			if level == 6 {
				return []string{"arcane-tradition-school-of-transmutation-transmuters-stone"}
			}
			if level == 10 {
				return []string{"arcane-tradition-school-of-transmutation-shapechanger"}
			}
			if level == 14 {
				return []string{"arcane-tradition-school-of-transmutation-master-transmuter"}
			}
		}

	}
	return []string{}
}

func featureImprovedByLevel(name string, level int) (dice string) {
	switch name {
	case "bardic-inspiration":
		if level < 5 {
			dice = "1d6"
		}
		if level >= 5 {
			dice = "1d8"
		}
		if level >= 10 {
			dice = "1d10"
		}
		if level >= 15 {
			dice = "1d12"
		}

	case "song-of-rest":
		if level < 9 {
			dice = "1d6"
		}
		if level >= 9 {
			dice = "1d8"
		}
		if level >= 13 {
			dice = "1d10"
		}
		if level >= 17 {
			dice = "1d12"
		}

	}

	return dice

}

func clericDestroyUndead(level int) (challenge float64) {
	if level <= 5 {
		challenge = 0.5
	}
	if level >= 8 {
		challenge = 1
	}
	if level >= 11 {
		challenge = 2
	}
	if level >= 14 {
		challenge = 3
	}
	if level >= 17 {
		challenge = 4
	}
	return challenge
}

func featuresWithExtraSpellList() []string {
	return []string{
		"domain-knowledge",
		"domain-life",
		"domain-nature",
		"domain-tempest",
		"domain-trickery",
		"domain-war",
		"sacred-oath-of-devotion",
		"sacred-oath-of-ancients",
		"sacred-oath-of-vengeance",
		"circle-of-the-land-spells-arctic",
		"circle-of-the-land-spells-coast",
		"circle-of-the-land-spells-desert",
		"circle-of-the-land-spells-forest",
		"circle-of-the-land-spells-grassland",
		"circle-of-the-land-spells-mountain",
		"circle-of-the-land-spells-swamp",
		"circle-of-the-land-spells-underdark",
	}
}

func extraSpellList(feature string, level int) (list []string) {
	switch feature {
	case "domain-knowledge":
		if level == 1 {
			list = []string{"command", "identify"}
		}
		if level == 3 {
			list = []string{"augury", "suggestion"}
		}
		if level == 5 {
			list = []string{"nondetection", "speak-with-dead"}
		}
		if level == 7 {
			list = []string{"arcane-eye", "confusion"}
		}
		if level == 9 {
			list = []string{"legend-lore", "scrying"}
		}

	case "domain-life":
		if level == 1 {
			list = []string{"bless", "cure wounds"}
		}
		if level == 3 {
			list = []string{"lesser-restoration", "spiritual-weapon"}
		}
		if level == 5 {
			list = []string{"beacon-of-hope", "revivify"}
		}
		if level == 7 {
			list = []string{"death-ward", "guardian-of-faith"}
		}
		if level == 9 {
			list = []string{"mass-cure-wounds", "raise-dead"}
		}

	case "domain-light":

		if level == 1 {
			list = []string{"burning-hands", "faerie-fire"}
		}
		if level == 3 {
			list = []string{"flaming-sphere", "scorching-ray"}
		}
		if level == 5 {
			list = []string{"daylight", "fireball"}
		}
		if level == 7 {
			list = []string{"guardian-offaith", "wall-of-fire"}
		}
		if level == 9 {
			list = []string{"flame-strike", "scrying"}
		}

	case "domain-nature":

		if level == 1 {
			list = []string{"animal-friendship", "speak-with-animals"}
		}
		if level == 3 {
			list = []string{"barkskin", "spike-growth"}
		}
		if level == 5 {
			list = []string{"plant-growth", "wind-wall"}
		}
		if level == 7 {
			list = []string{"dominate-beast", "grasping-vine"}
		}
		if level == 9 {
			list = []string{"insect-plague", "tree-stride"}
		}

	case "domain-tempest":

		if level == 1 {
			list = []string{"fog-cloud", "thunderwave"}
		}
		if level == 3 {
			list = []string{"gust-of-wind", "shatter"}
		}
		if level == 5 {
			list = []string{"call-lightning", "sleet-storm"}
		}
		if level == 7 {
			list = []string{"control-water", "ice-storm"}
		}
		if level == 9 {
			list = []string{"destructive-wave", "insect-plague"}
		}

	case "domain-trickery":

		if level == 1 {
			list = []string{"charm-person", "disguise-self"}
		}
		if level == 3 {
			list = []string{"mirror-image", "pass-without-trace"}
		}
		if level == 5 {
			list = []string{"blink", "dispel-magic"}
		}
		if level == 7 {
			list = []string{"dimension-door", "polymorph"}
		}
		if level == 9 {
			list = []string{"dominate-person", "modify -emory"}
		}

	case "domain-war":
		if level == 1 {
			list = []string{"divine-favor", "shield-of-faith"}
		}
		if level == 3 {
			list = []string{"magic-weapon", "spiritual-weapon"}
		}
		if level == 5 {
			list = []string{"crusaders-mantle", "spirit-guardians"}
		}
		if level == 7 {
			list = []string{"freedom-of-movement", "stoneskin"}
		}
		if level == 9 {
			list = []string{"flame-strike", "hold-monster"}
		}
	case "sacred-oath-of-devotion":
		if level == 3 {
			list = []string{"protection-from-evil-and-good", "sanctuary"}
		}
		if level == 5 {
			list = []string{"lesser-restoration", "zone-of-truth"}
		}
		if level == 9 {
			list = []string{"beacon-of-hope", "dispel-magic"}
		}
		if level == 13 {
			list = []string{"freedom-of-movement", "guardian-of-faith"}
		}
		if level == 17 {
			list = []string{"commune", "flame-strike"}
		}
	case "sacred-oath-of-ancients":
		if level == 3 {
			list = []string{"ensnaring-strike", "speak-with-animals"}
		}
		if level == 5 {
			list = []string{"misty-step", "moonbeam"}
		}
		if level == 9 {
			list = []string{"plant-growth", "protection-from-energy"}
		}
		if level == 13 {
			list = []string{"ice-storm", "stoneskin"}
		}
		if level == 17 {
			list = []string{"commune-with-nature", "tree-stride"}
		}

	case "sacred-oath-of-vengeance":
		if level == 3 {
			list = []string{"bane", "hunters-mark"}
		}
		if level == 5 {
			list = []string{"hold-person", "misty-step"}
		}
		if level == 9 {
			list = []string{"haste", "protection-from-energy"}
		}
		if level == 13 {
			list = []string{"banishment", "dimension-door"}
		}
		if level == 17 {
			list = []string{"hold-monster", "scrying"}
		}

	case "circle-of-the-land-spells-arctic":
		if level == 3 {
			list = []string{"hold-person", "spike-growth"}
		}
		if level == 5 {
			list = []string{"sleet-storm", "slow"}
		}
		if level == 7 {
			list = []string{"freedom-of-movement", "ice-storm"}
		}
		if level == 9 {
			list = []string{"commune-with-nature", "cone-of-cold"}
		}
	case "circle-of-the-land-spells-coast":
		if level == 3 {
			list = []string{"mirror-image", "misty-step"}
		}
		if level == 5 {
			list = []string{"water-breathing", "water-walk"}
		}
		if level == 7 {
			list = []string{"control water", "freedom-of-movement"}
		}
		if level == 9 {
			list = []string{"conjure-elemental", "scrying"}
		}
	case "circle-of-the-land-spells-desert":
		if level == 3 {
			list = []string{"blur", "silence"}
		}
		if level == 5 {
			list = []string{"create-food-and-water", "protection-from-energy"}
		}
		if level == 7 {
			list = []string{"blight", "hallucinatory-terrain"}
		}
		if level == 9 {
			list = []string{"insect-plague", "wall-of-stone"}
		}

	case "circle-of-the-land-spells-forest":
		if level == 3 {
			list = []string{"barkskin", "spider-climb"}
		}
		if level == 5 {
			list = []string{"call-lightning", "plant-growth"}
		}
		if level == 7 {
			list = []string{"divination", "freedom-of-movement"}
		}
		if level == 9 {
			list = []string{"commune-with-nature", "tree-stride"}
		}
	case "circle-of-the-land-spells-grassland":
		if level == 3 {
			list = []string{"invisibility", "pass-without-trace"}
		}
		if level == 5 {
			list = []string{"daylight", "haste"}
		}
		if level == 7 {
			list = []string{"divination", "freedom-of-movement"}
		}
		if level == 9 {
			list = []string{"dream", "insect-plague"}
		}
	case "circle-of-the-land-spells-mountain":
		if level == 3 {
			list = []string{"spider-climb", "spike-growth"}
		}
		if level == 5 {
			list = []string{"lightning-bolt", "meld-into-stone"}
		}
		if level == 7 {
			list = []string{"stone-shape", "stoneskin"}
		}
		if level == 9 {
			list = []string{"passwall", "wall-of-stone"}
		}
	case "circle-of-the-land-spells-swamp":
		if level == 3 {
			list = []string{"darkness", "melfs-acid-arrow"}
		}
		if level == 5 {
			list = []string{"water-walk", "stinking-cloud"}
		}
		if level == 7 {
			list = []string{"locate-creature", "freedom-of-movement"}
		}
		if level == 9 {
			list = []string{"insect-plague", "scrying"}
		}
	case "circle-of-the-land-spells-underdark":
		if level == 3 {
			list = []string{"spider-climb", "web"}
		}
		if level == 5 {
			list = []string{"gaseousform", "stinking-cloud"}
		}
		if level == 7 {
			list = []string{"greater-invisibility", "stone-shape"}
		}
		if level == 9 {
			list = []string{"cloudkill", "insect-plague"}
		}

	}
	return list
}

func extraDamageMeleeAttackFeature(name, choosen string, level, spellSlot int, paladinEnemy bool) (damage string, damageType string) {
	switch name {
	case "sneak-attack":
		if utils.StringInSlice(choosen, listOfWeaponDamageType()) {
			damageType = choosen
		}
		damage = RogueClass(level)
	case "archetype-hunter-hunters-prey-colossus-slayer":
		if utils.StringInSlice(choosen, listOfWeaponDamageType()) {
			damageType = choosen
		}
		damage = "1d8"

	case "domain-life-divine-strike":
		damageType = "radiant"
		if level >= 8 {
			damage = "1d8"
		}
		if level >= 14 {
			damage = "2d8"
		}
	case "domain-nature-divine-strike":
		allowedDamageTypes := []string{"cold", "fire", "lightning"}
		if utils.StringInSlice(choosen, allowedDamageTypes) {
			damageType = choosen
		}
		if level >= 8 {
			damage = "1d8"
		}
		if level >= 14 {
			damage = "2d8"
		}
	case "domain-tempest-divine-strike":
		damageType = "thunder"
		if level >= 8 {
			damage = "1d8"
		}
		if level >= 14 {
			damage = "2d8"
		}
	case "domain-trickery-divine-strike":
		damageType = "poison"
		if level >= 8 {
			damage = "1d8"
		}
		if level >= 14 {
			damage = "2d8"
		}
	case "domain-war-divine-strike":
		if utils.StringInSlice(choosen, listOfWeaponDamageType()) {
			damageType = choosen
		}
		if level >= 8 {
			damage = "1d8"
		}
		if level >= 14 {
			damage = "2d8"
		}
	case "divine-smite":
		damageType = "radiant"
		if spellSlot > 5 {
			spellSlot = 5
		}
		if spellSlot <= 5 {
			damage = fmt.Sprintf("%vd8", spellSlot)
		}
		if paladinEnemy {
			damage = fmt.Sprintf("%vd8", spellSlot+1)
		}

	default:
		damageType = ""
		damage = ""
	}
	return damage, damageType
}

// func featureHowToUse(name string) (message string) {
// 	switch name {
// 	case "martial-arts", "fighting-style-archery", "fighting-style-defense", "fighting-style-dueling", "fighting-style-great-weapon-fighting", "fighting-style-two-weapon-fighting":
// 		message = "calculate automatically in attack"

// 	case "unarmored-defense-barbarian", "unarmored-defense-monk":
// 		message = "calculate automatically in armor class"

// 	case "rage", "sneak-attack", "divine-smite", "archetype-hunter-hunters-prey-colossus-slayer", "domain-life-divine-strike", "domain-nature-divine-strike", "domain-tempest-divine-strike", "domain-trickery-divine-strike", "domain-war-divine-strike":
// 		message = "calculate automatically in attack, check, saving and monster attack, but need activation"

// 	case "spellcasting":
// 		message = "calculate automatically in spell"

// 	case "expertise", "jack-of-all-trades":
// 		message = "calculate automatically in check roll"

// 	case "bardic-inspiration":
// 		message = "need to add in one roll before use it"

// 	case "song-of-rest", "turn-undead", "destroy-undead", "second-wind", "lay-on-hands":
// 		message = "need a roll"

// 	case "favored-enemy", "natural-explorer", "fighting-style-protection", "divine-sense", "archetype-assassin-bonus-proficiencies-and-assassinate":
// 		message = "need to be approved by master and included in check/saving roll"
// 	}

// 	return message
// }
