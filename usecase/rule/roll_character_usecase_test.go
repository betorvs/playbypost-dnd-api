package rule

import (
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
				test1 := new(rule.NewCharacter)
				test1.Level = 1
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
				_, backgroundSkills := BackgroundStatistics(background)
				_, _, _, skillNumber := ClassStatistics(class)
				tmpSkills := []string{}
				for _, v := range skillListByClass(class) {
					if len(tmpSkills) >= skillNumber {
						continue
					}
					if race == "half-orc" && v == "intimidation" {
						continue
					}
					if !utils.StringInSlice(v, backgroundSkills) {
						tmpSkills = append(tmpSkills, v)
					}
				}
				test1.ChosenLanguages = languages
				test1.ChosenAbility = abilities
				test1.ChosenSkills = tmpSkills
				// levels := []int{1, 3}
				// for _,
				switch class {
				case "barbarian":
					classFeatures = append(classFeatures, "berseker")
				case "bard":
					classFeatures = append(classFeatures, "lore")
				case "cleric":
					classFeatures = append(classFeatures, "knowledge")
					test1.ChosenLanguagesByFeatures = append(test1.ChosenLanguagesByFeatures, "primordial", "sylvan")
					count := 0
					for _, v := range blessingsOfKnowledge() {
						if count >= 2 {
							break
						}
						if !utils.StringInSlice(v, tmpSkills) && !utils.StringInSlice(v, backgroundSkills) {
							test1.ChosenSkillsByFeatures = append(test1.ChosenSkillsByFeatures, v)
							count++
						}
					}
				case "druid":
					classFeatures = append(classFeatures, "land")
				case "fighter":
					classFeatures = append(classFeatures, "champion")
				case "monk":
					classFeatures = append(classFeatures, "open-hand")
				case "ranger":
					classFeatures = append(classFeatures, "hunter", "forest", "giants")
					test1.ChosenLanguagesByFeatures = []string{"giant"}
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
				case "paladin":
					classFeatures = append(classFeatures, "devotion")
				case "sorcerer":
					classFeatures = append(classFeatures, "draconic-bloodline", "black")
				case "warlock":
					classFeatures = append(classFeatures, "archfey")
				case "wizard":
					classFeatures = append(classFeatures, "abjuration")
				}

				// fmt.Println(race, class, background, tmpSkills)
				test1.ChosenClassFeatures = classFeatures
				res1, err1 := CalculateCharacter(test1)
				assert.NotEmpty(t, res1)
				assert.NoError(t, err1)
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
