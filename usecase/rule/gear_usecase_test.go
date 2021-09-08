package rule

import (
	"net/url"
	"testing"

	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	"github.com/betorvs/playbypost-dnd/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllWeapons(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllWeapons(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"longsword"}
	res1 := GetAllWeapons(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["kind"] = []string{"martial-weapon"}
	res2 := GetAllWeapons(value2)
	assert.NotEmpty(t, res2)
	value3 := make(url.Values)
	value3["damage_type"] = []string{"slashing"}
	res3 := GetAllWeapons(value3)
	assert.NotEmpty(t, res3)
}

func TestWeaponsList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := WeaponsList()
	assert.NotEmpty(t, res)
}

func TestWeaponsByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := WeaponsByName("longsword")
	assert.NotEmpty(t, res)
}

func TestGetAllArmor(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllArmor(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"padded"}
	res1 := GetAllArmor(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["kind"] = []string{"light-armor"}
	res2 := GetAllArmor(value2)
	assert.NotEmpty(t, res2)
}

func TestArmorList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ArmorList()
	assert.NotEmpty(t, res)
}

func TestArmorByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ArmorByName("padded")
	assert.NotEmpty(t, res)
}

func TestGetAllGears(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllGears(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"bottle glass"}
	res1 := GetAllGears(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["kind"] = []string{"other"}
	res2 := GetAllGears(value2)
	assert.NotEmpty(t, res2)
}

func TestGearByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := GearByName("bottle glass")
	assert.NotEmpty(t, res)
}

func TestGetAllPacks(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllPacks(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"explorers pack"}
	res1 := GetAllPacks(value1)
	assert.NotEmpty(t, res1)
}

func TestPacksByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := PacksByName("explorers pack")
	assert.NotEmpty(t, res)
}

func TestPacksList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := PacksList()
	assert.NotEmpty(t, res)
}

func TestGetAllTools(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllTools(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"cooks utensils"}
	res1 := GetAllTools(value1)
	assert.NotEmpty(t, res1)
	value2 := make(url.Values)
	value2["kind"] = []string{"artisans tools"}
	res2 := GetAllTools(value2)
	assert.NotEmpty(t, res2)
}

func TestToolsByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ToolsByName("cooks utensils")
	assert.NotEmpty(t, res)
}

func TestToolsList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := ToolsList()
	assert.NotEmpty(t, res)
}

func TestGetAllMounts(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	value := make(url.Values)
	res := GetAllMounts(value)
	assert.NotEmpty(t, res)
	value1 := make(url.Values)
	value1["name"] = []string{"mule"}
	res1 := GetAllMounts(value1)
	assert.NotEmpty(t, res1)
}

func TestMountsByName(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := MountsByName("donkey or mule")
	assert.NotEmpty(t, res)
}

func TestMountsList(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := MountsList()
	assert.NotEmpty(t, res)
}

func TestCalcShoppingCart(t *testing.T) {
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	testList := new(rule.SimpleList)
	// Added some fake gear to test calcCoinExpenses too
	testList.List = []string{"longsword", "padded", "bottle glass", "explorers pack", "cooks utensils", "donkey or mule", "lifestyle poor", "unknown", "bucket", "fake-electrum", "fake-platinum"}
	res := CalcShoppingCart(testList)
	assert.NotEmpty(t, res)
}

func TestCalcRandomTreasureByChallengeLevel(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res, err := CalcRandomTreasureByChallengeLevel(-1, true)
	assert.NotNil(t, res)
	assert.Error(t, err)
}

func TestCalcRandomTreasureByChallengeLevelTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res, err := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCalcRandomTreasureByChallengeLevelTest35(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e35Mock)
	res, err := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCalcRandomTreasureByChallengeLevelTest65(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e65Mock)
	res, err := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCalcRandomTreasureByChallengeLevelTest80(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e80Mock)
	res, err := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

func TestCalcRandomTreasureByChallengeLevelTest96(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e96Mock)
	res, err := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res)
	assert.NoError(t, err)
}

// TestCheckTableBelowFive

func TestIndividualPercentageByLevelCheckTableBelowFiveTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
	res3, res4 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, res4)
}

func TestIndividualPercentageByLevelCheckTableBelowFiveTest35(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e35Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowFiveTest65(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e65Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowFiveTest80(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e80Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowFiveTest96(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e96Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowFiveTestZero(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e0Mock)
	res1, res2 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

// TestCheckTableBelowTen

func TestIndividualPercentageByLevelCheckTableBelowTenTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
	res3, res4 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, res4)
}

func TestIndividualPercentageByLevelCheckTableBelowTenTest35(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e35Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowTenTest65(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e65Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowTenTest80(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e80Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowTenTest96(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e96Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowTenTestZero(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e0Mock)
	res1, res2 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

//TestCheckTableBelowSeventeen

func TestIndividualPercentageByLevelCheckTableBelowSeventeenTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res1, res2 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
	res3, res4 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, res4)
}

func TestIndividualPercentageByLevelCheckTableBelowSeventeenTest35(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e35Mock)
	res1, res2 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowSeventeenTest65(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e65Mock)
	res1, res2 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowSeventeenTest80(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e80Mock)
	res1, res2 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableBelowSeventeenTestZero(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e0Mock)
	res1, res2 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

// TestCheckTableAboveSeventeen

func TestIndividualPercentageByLevelCheckTableAboveSeventeenTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100b20Mock)
	res1, res2 := individualPercentageByLevel(18, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
	res3, res4 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, res4)
}

func TestIndividualPercentageByLevelCheckTableAboveSeventeenTest35(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e35Mock)
	res1, res2 := individualPercentageByLevel(18, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableAboveSeventeenTest80(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e80Mock)
	res1, res2 := individualPercentageByLevel(18, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestIndividualPercentageByLevelCheckTableAboveSeventeenTestZero(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e0Mock)
	res1, res2 := individualPercentageByLevel(18, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, res2)
}

func TestCalcHoardPercentageByLevel(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDice100e96HoardMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	res := CalcHoardPercentageByLevel(4)
	assert.NotNil(t, res)
}

func TestCheckHoardBelowFive(t *testing.T) {
	test := []int{1, 7, 17, 27, 37, 45, 53, 61, 66, 71, 76, 79, 81, 86, 98, 100}
	for _, v := range test {
		res, _, _, _ := checkHoardBelowFive(v)
		assert.NotNil(t, res)
	}
}
