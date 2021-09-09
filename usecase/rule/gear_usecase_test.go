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
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test.DiceResult = 14
	res, err := CalcRandomTreasureByChallengeLevel(-1, true)
	assert.NotNil(t, res)
	assert.Error(t, err)
	res1, err1 := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res1)
	assert.NoError(t, err1)
	test.DiceResult = 35
	res2, err2 := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res2)
	assert.NoError(t, err2)
	test.DiceResult = 65
	res3, err3 := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res3)
	assert.NoError(t, err3)
	test.DiceResult = 80
	res4, err4 := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res4)
	assert.NoError(t, err4)
	test.DiceResult = 96
	res5, err5 := CalcRandomTreasureByChallengeLevel(3, true)
	assert.NotNil(t, res5)
	assert.NoError(t, err5)
}

// TestCheckTableBelowFive

func TestIndividualPercentageByLevelCheckTableBelowFive(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test.DiceResult = 14
	res1, err1 := individualPercentageByLevel(3, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, err1)
	res2, err2 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res2)
	assert.NotNil(t, err2)
	test.DiceResult = 35
	res3, err3 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, err3)
	test.DiceResult = 65
	res4, err4 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res4)
	assert.NotNil(t, err4)
	test.DiceResult = 80
	res5, err5 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res5)
	assert.NotNil(t, err5)
	test.DiceResult = 96
	res6, err6 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res6)
	assert.NotNil(t, err6)
	test.DiceResult = 101
	res7, err7 := individualPercentageByLevel(3, false)
	assert.NotNil(t, res7)
	assert.NotNil(t, err7)
}

// TestCheckTableBelowTen

func TestIndividualPercentageByLevelCheckTableBelowTenTest14(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test.DiceResult = 14
	res1, err1 := individualPercentageByLevel(9, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, err1)
	res2, err2 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res2)
	assert.NotNil(t, err2)
	test.DiceResult = 35
	res3, err3 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, err3)
	test.DiceResult = 65
	res4, err4 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res4)
	assert.NotNil(t, err4)
	test.DiceResult = 80
	res5, err5 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res5)
	assert.NotNil(t, err5)
	test.DiceResult = 96
	res6, err6 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res6)
	assert.NotNil(t, err6)
	test.DiceResult = 101
	res7, err7 := individualPercentageByLevel(9, false)
	assert.NotNil(t, res7)
	assert.NotNil(t, err7)
}

//TestCheckTableBelowSeventeen

func TestIndividualPercentageByLevelCheckTableBelowSeventeen(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test.DiceResult = 14
	res1, err1 := individualPercentageByLevel(15, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, err1)
	res2, err2 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res2)
	assert.NotNil(t, err2)
	test.DiceResult = 35
	res3, err3 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, err3)
	test.DiceResult = 65
	res4, err4 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res4)
	assert.NotNil(t, err4)
	test.DiceResult = 80
	res5, err5 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res5)
	assert.NotNil(t, err5)
	test.DiceResult = 96
	res6, err6 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res6)
	assert.NotNil(t, err6)
	test.DiceResult = 101
	res7, err7 := individualPercentageByLevel(15, false)
	assert.NotNil(t, res7)
	assert.NotNil(t, err7)
}

// TestCheckTableAboveSeventeen

func TestIndividualPercentageByLevelCheckTableAboveSeventeen(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	test.DiceResult = 14
	res1, err1 := individualPercentageByLevel(18, true)
	assert.NotNil(t, res1)
	assert.NotNil(t, err1)
	res2, err2 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res2)
	assert.NotNil(t, err2)
	test.DiceResult = 35
	res3, err3 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res3)
	assert.NotNil(t, err3)
	test.DiceResult = 65
	res4, err4 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res4)
	assert.NotNil(t, err4)
	test.DiceResult = 80
	res5, err5 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res5)
	assert.NotNil(t, err5)
	test.DiceResult = 96
	res6, err6 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res6)
	assert.NotNil(t, err6)
	test.DiceResult = 101
	res7, err7 := individualPercentageByLevel(18, false)
	assert.NotNil(t, res7)
	assert.NotNil(t, err7)
}

func TestCalcHoardPercentageByLevel(t *testing.T) {
	appcontext.Current.Add(appcontext.Dice, test.InitDiceMock)
	appcontext.Current.Add(appcontext.Database, test.InitDatabaseMock)
	list := []int{17, 96}
	for _, v := range list {
		test.DiceResult = v
		res := CalcHoardPercentageByLevel(4)
		assert.NotNil(t, res)
	}

}

func TestCheckHoardBelowFive(t *testing.T) {
	test := []int{1, 7, 17, 27, 37, 45, 53, 61, 66, 71, 76, 79, 81, 86, 98, 100}
	for _, v := range test {
		res, _, _, _ := checkHoardBelowFive(v)
		assert.NotNil(t, res)
	}
}
