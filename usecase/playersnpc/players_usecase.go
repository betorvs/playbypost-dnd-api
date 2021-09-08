package playersnpc

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/mongodb"
	"github.com/betorvs/playbypost-dnd/domain/player"
	"github.com/betorvs/playbypost-dnd/domain/rule"
	ruleUsecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"github.com/betorvs/playbypost-dnd/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetAllPlayers func
func GetAllPlayers(queryParameters url.Values) ([]*player.Players, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetPlayers(queryParameters)
}

//GetOnePlayer func
func GetOnePlayer(playerID primitive.ObjectID) (*player.Players, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetOnePlayer(playerID)
}

//CreatePlayer func
func CreatePlayer(player *player.Players) (res string, err error) {
	repo := mongodb.GetMongoRepository()
	if player.AbilityWithoutMagic == nil {
		player.AbilityWithoutMagic = player.Ability
	}
	// if player.Notes == nil {
	//      player.Notes = make(map[string]string)
	// }
	if player.SpellListLevel == nil {
		player.SpellListLevel = make(map[string][]int)
	}
	if player.SpellList == nil {
		player.SpellList = make([]string, 0)
	}
	if player.Advantages == nil {
		player.Advantages = make([]string, 0)
	}
	if player.Disvantages == nil {
		player.Disvantages = make([]string, 0)
	}
	if player.AutoFail == nil {
		player.AutoFail = make([]string, 0)
	}
	if config.Values.RuleBuiltin {
		player, err = newPlayer(player)
		if err != nil {
			return "", err
		}
	}
	id, err := repo.SavePlayer(player)
	if err != nil {
		return "", err
	}
	err = newInventory(id.Hex())
	if err != nil {
		return "inventory not created", err
	}
	return fmt.Sprintf("%v", id), nil
}

// UpdatePlayerRecalc func
func UpdatePlayerRecalc(playerID primitive.ObjectID, player *player.Players) (string, error) {
	fmt.Println("update with recalc")
	if player.AbilityWithoutMagic == nil {
		player.AbilityWithoutMagic = player.Ability
	}
	// if player.Notes == nil {
	//      player.Notes = make(map[string]string)
	// }
	if player.SpellListLevel == nil {
		player.SpellListLevel = make(map[string][]int)
	}
	if player.SpellList == nil {
		player.SpellList = make([]string, 0)
	}
	if player.Advantages == nil {
		player.Advantages = make([]string, 0)
	}
	if player.Disvantages == nil {
		player.Disvantages = make([]string, 0)
	}
	if player.AutoFail == nil {
		player.AutoFail = make([]string, 0)
	}
	player, err := newPlayer(player)
	if err != nil {
		return "", err
	}
	fmt.Printf("%#v", player)
	return UpdatePlayer(playerID, player)
}

//UpdatePlayer func
func UpdatePlayer(playerID primitive.ObjectID, player *player.Players) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.UpdatePlayer(playerID, player)
	if err != nil {
		return "", err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("Player updated Level %v, changes number %v", player.Level, id)
	}
	return message, nil
}

//CheckPlayerExist func
func CheckPlayerExist(playerID primitive.ObjectID) bool {
	repo := mongodb.GetMongoRepository()
	player, err := repo.GetOnePlayer(playerID)
	if err != nil {
		return false
	}
	if playerID == player.ID {
		return true
	}
	return false
}

// DeletePlayer func
func DeletePlayer(playerID, inventoryID primitive.ObjectID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.DeletePlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	result2, err := repo.DeleteInventoryByID(inventoryID)
	if err != nil {
		return 0, err
	}
	res := result + result2
	return res, nil
}

//AddCampaignToPlayer func
func AddCampaignToPlayer(playerID primitive.ObjectID, campaign *player.AddCampaign) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.AddCampaignToPlayer(playerID, campaign.CampaignID, campaign.CampaignTitle, campaign.SlackChannelID)
	if err != nil {
		return "Cannot save campaign change to player", err
	}
	message := fmt.Sprintf("Player save to campaign %s", campaign.CampaignTitle)
	if id > 0 {
		message += fmt.Sprintf(" %v", id)
	}
	return message, nil
}

//AddOrRemoveHP func
func AddOrRemoveHP(playerID primitive.ObjectID, action string, hit int) (string, error) {
	var hp int
	repo := mongodb.GetMongoRepository()
	hp = hit
	if action == "remove" {
		hp = hit * -1
	}
	id, err := repo.AddOrRemovePlayerHP(playerID, hp)
	if err != nil {
		return fmt.Sprintf("NOT Player %s %v HP: %v", action, hit, err), err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("Player %s %v HP", action, hit)
	}
	return message, nil
}

//AddPlayerXP func
func AddPlayerXP(playerID primitive.ObjectID, xp int) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.AddPlayerXP(playerID, xp)
	if err != nil {
		return "Cannot give XP.", err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("Player receive  %v XP", xp)
	}
	return message, nil
}

//UseSpellByLevel func
func UseSpellByLevel(playerID primitive.ObjectID, level, number int) (string, error) {
	doc := make(map[string]int)
	name := fmt.Sprintf("spells_used.level%v", level)
	doc[name] = number
	repo := mongodb.GetMongoRepository()
	id, err := repo.UsageSpellByLevel(playerID, doc)
	if err != nil {
		return "Cannot use a spell, sorry", err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("Spell Level %v was used", level)
	}
	return message, nil
}

//FullRestPlayer func
func FullRestPlayer(playerID primitive.ObjectID) (string, error) {
	repo := mongodb.GetMongoRepository()
	player, err := repo.GetOnePlayer(playerID)
	if err != nil {
		return "Cannot find a player", err
	}

	doc := make(map[string]int)
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("spells_used.level%v", i)
		doc[name] = 0
	}
	hp := player.HPMax

	hpRes, err := repo.SetHPTempPlayerByID(playerID, hp)
	if err != nil {
		return "its impossible. Player cannot rest and recover HP.", err
	}
	spellRes, err := repo.SetSpellUsedPlayerByID(playerID, doc)
	if err != nil {
		return "its impossible. Player cannot rest and recover it spell used.", err
	}
	var message string
	if hpRes >= 1 {
		message += "HP Recovery"
	}
	if spellRes >= 1 {
		message += " Spells Recovery\n"
	}

	return message, nil
}

//ChangeCondition func
func ChangeCondition(playerID primitive.ObjectID, player *player.Condition) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.ChangePlayerCondition(playerID, player)
	if err != nil {
		return "", err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("Player updated Advantages %v, Conditions %v, Disvantages %v and Auto Fail list %v, changes number %v", player.Advantages, player.Conditions, player.Disvantages, player.AutoFail, id)
	}
	return message, nil
}

//InventoryIDByPlayerID func
func InventoryIDByPlayerID(playerID primitive.ObjectID) (primitive.ObjectID, error) {
	repo := mongodb.GetMongoRepository()
	inventoryID, err := repo.GetInventoryID(playerID.Hex())
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return inventoryID, nil
}

//GetInventoryByPlayerID func
func GetInventoryByPlayerID(playerID primitive.ObjectID) (*player.Inventory, error) {
	repo := mongodb.GetMongoRepository()
	inventory, err := repo.GetInventoryByID(playerID)
	if err != nil {
		return &player.Inventory{}, err
	}
	return inventory, nil
}

//AddTreasure func
func AddTreasure(inventoryID primitive.ObjectID, treasure *player.Treasure) (string, error) {
	doc := make(map[string]int)
	doc["treasure.copper"] = treasure.Copper
	doc["treasure.silver"] = treasure.Silver
	doc["treasure.electrum"] = treasure.Electrum
	doc["treasure.gold"] = treasure.Gold
	doc["treasure.platinum"] = treasure.Platinum
	repo := mongodb.GetMongoRepository()
	id, err := repo.AddTreasurePlayer(inventoryID, doc)
	if err != nil {
		return "Cannot add any coin, sorry", err
	}
	var message string
	if id == 0 {
		message = "add treasure successfully"
	}
	return message, nil
}

//AddOrRemoveOtherItems func
func AddOrRemoveOtherItems(inventoryID primitive.ObjectID, action string, item []string) (string, error) {
	repo := mongodb.GetMongoRepository()
	var res string
	switch action {
	case "add":
		id, err := repo.AddOthersItems(inventoryID, item, false)
		if err != nil {
			return "Cannot add any item.", err
		}
		res = fmt.Sprintf("%v", id)
	case "remove":
		id, err := repo.RemoveOthersItems(inventoryID, item, false)
		if err != nil {
			return "Cannot add any item.", err
		}
		res = fmt.Sprintf("%v", id)
	}

	return res, nil
}

//AddArmorWeaponPlayerByID func
func AddArmorWeaponPlayerByID(playerID, inventoryID primitive.ObjectID, ac *player.Armory) (string, error) {
	repo := mongodb.GetMongoRepository()
	armory, err := repo.GetArmory(inventoryID)
	if err != nil {
		return "get armory", err
	}
	if ac.ArmorName != "" && ac.ArmorName != armory.ArmorName {
		armory.ArmorName = ac.ArmorName
	}
	if ac.ShieldName != "" && ac.ShieldName != armory.ShieldName {
		armory.ShieldName = ac.ShieldName
	}
	if ac.WeaponName != "" && ac.WeaponName != armory.WeaponName {
		armory.WeaponName = ac.WeaponName
	}
	extra := ""
	if config.Values.RuleBuiltin {
		player, err := GetOnePlayer(playerID)
		if err != nil {
			return "get player", err
		}
		if armory.ArmorName != "" && utils.StringInSlice(armory.ArmorName, ruleUsecase.ArmorList()) {
			armorClass := new(rule.ArmorClass)
			armorClass.Ability = player.Ability
			armorClass.Armor = armory.ArmorName
			armorClass.ArmorMagicBonus = ac.ArmorMagicBonus
			armorClass.ArmorProficiency = player.ArmorProficiency
			armorClass.ClassFeatures = player.ClassFeatures
			armorClass.Shield = armory.ShieldName
			armorClass.ShieldMagicBonus = ac.ShieldMagicBonus
			result := ruleUsecase.CheckArmorClass(armorClass)
			player.ArmorClass = result.ArmorClass
			player.ArmorClassBonus = ac.ArmorMagicBonus
			player.AutoFail = result.ArmorClassAutomaticallyFails
			player.Disvantages = result.ArmorClassDisvantages
			player.ArmorName = armory.ArmorName
			player.ShieldName = armory.ShieldName
			player.WeaponName = armory.WeaponName
			res, err := UpdatePlayer(playerID, player)
			if err != nil {
				return "update player", err
			}
			extra = res
		}
	}
	id, err := repo.SetArmorWeaponPlayerByID(inventoryID, armory)
	if err != nil {
		return "set armory", err
	}
	value := fmt.Sprintf("Armory configured ( Changes %v ) %s", id, extra)
	return value, nil
}

//AddOrRemoveMagicItems func
func AddOrRemoveMagicItems(inventoryID primitive.ObjectID, action string, item []string) (string, error) {
	repo := mongodb.GetMongoRepository()
	var res string
	switch action {
	case "add":
		id, err := repo.AddOthersItems(inventoryID, item, true)
		if err != nil {
			return "Cannot add any magic item.", err
		}
		res = fmt.Sprintf("%v", id)
	case "remove":
		id, err := repo.RemoveOthersItems(inventoryID, item, true)
		if err != nil {
			return "Cannot add any magic item.", err
		}
		res = fmt.Sprintf("%v", id)
	case "attune":
		player, _ := GetOnePlayer(inventoryID)
		inventory, _ := GetInventoryByPlayerID(inventoryID)
		if len(inventory.AttunedItems) >= 3 {
			return "number of attunement is 3", fmt.Errorf("limit of attunement items reached")
		}
		magicalEffect := []string{}
		var err error
		for _, v := range item {
			magic := ruleUsecase.GetMagicItemByName(v)
			if magic.RequiredAttunement {
				if len(magic.AttunementRestriction) != 0 && !utils.StringInSlice(player.Class, magic.AttunementRestriction) {
					// check class restrictions
					return "magic item require a specific class", fmt.Errorf("class invalid for this item")
				}
				magicalEffect = append(magicalEffect, magic.Name)
				player, err = addMagicalEffects(player, magic)
				if err != nil {
					return "cannot add any magical effect into player", err
				}
			}
		}
		// update player directly
		id, err := repo.SetMagicalEffect(player.ID, magicalEffect, true)
		if err != nil {
			return "Cannot add any magic item.", err
		}
		res = fmt.Sprintf("%v", id)
	}
	return res, nil
}

func addMagicalEffects(player *player.Players, magicItem rule.MagicItem) (*player.Players, error) {
	if magicItem.Feature == nil {
		return player, nil
	}
	// verify all CoreFeatures here
	// switch {
	//      case magicItem.Feature.AbilityBonus
	// }

	return player, nil
}

func newInventory(playerID string) error {
	inventory := new(player.Inventory)
	inventory.PlayerID = playerID
	inventory.Items = make([]string, 0)
	inventory.Armory = player.Armory{
		ArmorMagicBonus:  0,
		ArmorName:        "",
		ShieldName:       "",
		ShieldMagicBonus: 0,
		WeaponName:       "",
		WeaponMagicBonus: 0,
	}
	inventory.Treasure = player.Treasure{
		Copper:   0,
		Silver:   0,
		Electrum: 0,
		Gold:     0,
		Platinum: 0,
	}
	inventory.AttunedItems = []string{}
	inventory.MagicItems = []string{}
	// inventory.MagicItems = player.MagicItems{
	//      UnderSpellEffect:  []string{},
	//      UnderPotionEffect: []string{},
	//      AttunedItems:      []string{},
	//      Foot:              "",
	//      Hands:             "",
	//      Waist:             "",
	//      Core:              "",
	//      Neck:              "",
	//      Head:              "",
	//      Back:              "",
	//      Weapon:            "",
	//      Shield:            "",
	// }
	repo := mongodb.GetMongoRepository()
	err := repo.CreateInventory(inventory)
	if err != nil {
		return err
	}

	return nil
}

func newPlayer(player *player.Players) (*player.Players, error) {
	purpose := new(rule.NewCharacter)
	purpose.Level = player.Level
	purpose.Class = player.Class
	purpose.Race = player.Race
	purpose.Subrace = player.Subrace
	purpose.Background = player.Background
	purpose.Ability = player.Ability
	purpose.ChosenLanguages = player.ChosenLanguages
	purpose.ChosenSkills = player.ChosenSkills
	purpose.ChosenAbility = player.ChosenAbility
	purpose.ChosenAbilityByLevel = player.ChosenAbilityByLevel
	purpose.ChosenClassFeatures = player.ChosenClassFeatures
	purpose.ChosenSkillsByFeatures = player.ChosenSkillsByFeatures
	purpose.ChosenLanguagesByFeatures = player.ChosenLanguagesByFeatures
	character, err := ruleUsecase.CalculateCharacter(purpose)
	if err != nil {
		return player, err
	}
	player.Level = character.Level
	player.Class = character.Class
	player.ClassFeatures = character.ClassFeatures
	player.Race = character.Race
	player.Subrace = character.Subrace
	player.RaceFeatures = character.RaceFeatures
	player.Background = character.Background
	player.Proficiency = character.Proficiency
	player.HitDice = character.HitDice
	player.Size = character.Size
	player.SpellKnown = character.SpellKnown
	player.SpellListLevel = character.SpellListLevel
	player.SpellMaxLevel = character.SpellMaxLevel
	player.CantripsKnown = character.CantripsKnown
	player.XPNextLevel = character.XPNextLevel
	player.BarbarianRage = character.BarbarianRage
	player.BarbarianDamage = character.BarbarianDamage
	player.MonkMartial = character.MonkMartial
	player.MonkKi = character.MonkKi
	player.MonkMovement = character.MonkMovement
	player.RogueSneak = character.RogueSneak
	player.SorceryPoints = character.SorceryPoints
	player.WarlockSpellSlots = character.WarlockSpellSlots
	player.WarlockSlotLevel = character.WarlockSlotLevel
	player.WarlockInvocationsKnown = character.WarlockInvocationsKnown
	player.Speed = character.Speed
	player.Ability = character.Ability
	player.AbilityModifier = character.AbilityModifier
	player.ChosenLanguages = character.ChosenLanguages
	player.ChosenSkills = character.ChosenSkills
	player.ChosenAbility = character.ChosenAbility
	player.ChosenAbilityByLevel = character.ChosenAbilityByLevel
	player.ChosenClassFeatures = character.ChosenClassFeatures
	player.ChosenSkillsByFeatures = character.ChosenSkillsByFeatures
	player.ChosenLanguagesByFeatures = character.ChosenLanguagesByFeatures
	player.HPMax = character.HPMax
	player.HPTemp = character.HPTemp
	player.Language = character.Language
	player.Savings = character.Savings
	player.ArmorProficiency = character.ArmorProficiency
	player.Skills = character.Skills
	player.Disvantages = character.Disvantages
	player.Advantages = character.Advantages
	player.AutoFail = character.AutoFail
	player.DamageResistence = character.DamageResistence
	player.DamageVulnerabilities = character.DamageVulnerabilities
	player.DamageImmunities = character.DamageImmunities
	player.ConditionImmunities = character.ConditionImmunities
	player.MagicalEffect = character.MagicalEffect
	return player, nil
}
