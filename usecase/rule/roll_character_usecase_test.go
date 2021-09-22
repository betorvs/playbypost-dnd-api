package rule

import (
	"fmt"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/betorvs/playbypost-dnd/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculateCharacter(t *testing.T) {
	// appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)

	for _, class := range ClassList() {
		for _, race := range RaceList() {
			for _, background := range BackgroundList() {
				levels := []int{1, 3}
				for _, level := range levels {
					test1 := new(rule.NewCharacter)
					test1.Level = level
					ability := make(map[string]int)
					for _, v := range AbilityList() {
						ability[v] = 14
					}
					test1.Ability = ability
					test1.Class = class
					test1.Race = race
					test1.Background = background
					res, err := CalculateCharacter(test1)
					assert.NotEmpty(t, res)
					assert.Error(t, err)
					subrace := ""
					languages := []string{}
					abilities := []string{}
					classFeatures := []string{}
					switch race {
					case "dwarf":
						subrace = "hill-dwarf"

					case "elf":
						subrace = "high-elf"
						languages = append(languages, "celestial")

					case "hafling":
						subrace = "lightfoot"

					case "gnome":
						subrace = "rock-gnome"
					case "human":
						languages = append(languages, "celestial")
					case "half-elf":
						languages = append(languages, "celestial")
						abilities = append(abilities, "charisma", "constitution")
					}
					test1.Subrace = subrace

					switch background {
					case "acolyte":
						languages = append(languages, "deep-speech", "abyssal")
					case "sage":
						languages = append(languages, "deep-speech", "abyssal")
					}
					back := BackgroundStatistics(background)
					_, _, _, skillNumber := ClassDetails(class)
					tmpSkills := []string{}
					for _, v := range skillListByClass(class) {
						if len(tmpSkills) >= skillNumber {
							continue
						}
						if race == "half-orc" && v == "intimidation" {
							continue
						}
						if !utils.StringInSlice(v, back.Skills) {
							tmpSkills = append(tmpSkills, v)
						}
					}
					test1.ChosenLanguages = languages
					test1.ChosenAbility = abilities
					test1.ChosenSkills = tmpSkills
					switch class {
					case "barbarian":
						classFeatures = append(classFeatures, "berseker")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1.BarbarianRage)
						assert.NotEmpty(t, res1.BarbarianDamage)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "bard":
						classFeatures = append(classFeatures, "lore")
						if level >= 3 {
							countLore := 0
							for _, v := range skillListByClass(class) {
								if countLore >= 3 {
									break
								}
								if race == "half-orc" && v == "intimidation" {
									continue
								}
								if !utils.StringInSlice(v, tmpSkills) && !utils.StringInSlice(v, back.Skills) {
									test1.ChosenSkillsByFeatures = append(test1.ChosenSkillsByFeatures, v)
									countLore++
								}
							}
							countExpertise := 0
							for _, v := range tmpSkills {
								if countExpertise >= 2 {
									break
								}
								if utils.StringInSlice(v, tmpSkills) {
									classFeatures = append(classFeatures, v)
									countExpertise++
								}
							}
						}
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "cleric":
						classFeatures = append(classFeatures, "knowledge")
						test1.ChosenLanguagesByFeatures = append(test1.ChosenLanguagesByFeatures, "primordial", "sylvan")
						count := 0
						for _, v := range blessingsOfKnowledge() {
							if count >= 2 {
								break
							}
							if !utils.StringInSlice(v, tmpSkills) && !utils.StringInSlice(v, back.Skills) {
								test1.ChosenSkillsByFeatures = append(test1.ChosenSkillsByFeatures, v)
								count++
							}
						}
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "druid":
						classFeatures = append(classFeatures, "land")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "fighter":
						classFeatures = append(classFeatures, "champion")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "monk":
						classFeatures = append(classFeatures, "open-hand")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
						assert.NotEmpty(t, res1.MonkMartial)
						if level != 1 {
							assert.NotEmpty(t, res1.MonkKi)
						}
						assert.NotEmpty(t, res1.MonkMovement)
					case "ranger":
						classFeatures = append(classFeatures, "hunter", "forest", "giants")
						test1.ChosenLanguagesByFeatures = []string{"giant"}
						if level >= 3 {
							classFeatures = append(classFeatures, "colossus-slayer")
						}
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "rogue":
						classFeatures = append(classFeatures, "thief")
						count := 0
						for _, v := range tmpSkills {
							if count >= 2 {
								break
							}
							if utils.StringInSlice(v, tmpSkills) {
								classFeatures = append(classFeatures, v)
								count++
							}
						}
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
						assert.NotEmpty(t, res1.RogueSneak)
					case "paladin":
						classFeatures = append(classFeatures, "devotion")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "sorcerer":
						classFeatures = append(classFeatures, "draconic-bloodline", "black")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					case "warlock":
						classFeatures = append(classFeatures, "archfey")
						if level >= 3 {
							classFeatures = append(classFeatures, "chain")
						}
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
						assert.NotEmpty(t, res1.WarlockSlotLevel)
						assert.NotEmpty(t, res1.WarlockSpellSlots)
						if level != 1 {
							assert.NotEmpty(t, res1.WarlockInvocationsKnown)
						}

					case "wizard":
						classFeatures = append(classFeatures, "abjuration")
						test1.ChosenClassFeatures = classFeatures
						fmt.Println(race, class, background, tmpSkills, classFeatures)
						res1, err1 := CalculateCharacter(test1)
						assert.NotEmpty(t, res1)
						assert.NoError(t, err1)
					}
				}
			}
		}
	}
}

func TestLanguagesAdded(t *testing.T) {}

func TestSkillsAdded(t *testing.T) {}

func TestBlessingsOfKnowledge(t *testing.T) {
	test1 := "arcana"
	res := blessingsOfKnowledge()
	assert.GreaterOrEqual(t, len(res), 4)
	assert.Contains(t, blessingsOfKnowledge(), test1)
}
