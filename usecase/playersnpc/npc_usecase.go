package playersnpc

import (
	"fmt"
	"net/url"

	"github.com/betorvs/playbypost-dnd/config"
	"github.com/betorvs/playbypost-dnd/domain/mongodb"
	"github.com/betorvs/playbypost-dnd/domain/player"
	ruleUsecase "github.com/betorvs/playbypost-dnd/usecase/rule"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetAllNPCS func
func GetAllNPCS(queryParameters url.Values) ([]*player.NPC, error) {
	repo := mongodb.GetMongoRepository()
	return repo.GetNPCS(queryParameters)
}

//CreateNPC func
func CreateNPC(npc *player.NPC) (string, error) {
	if npc.Conditions == nil {
		npc.Conditions = make([]string, 0)
	}
	if npc.Advantages == nil {
		npc.Advantages = make([]string, 0)
	}
	if npc.Disvantages == nil {
		npc.Disvantages = make([]string, 0)
	}
	if npc.AutoFail == nil {
		npc.AutoFail = make([]string, 0)
	}
	if config.Values.RuleBuiltin {
		monster := ruleUsecase.MosterByName(npc.Monster)
		npc.ArmorClass = monster.ArmorClass
		npc.HitPoints = monster.HitPoints
		npc.XP = monster.XP
	}
	repo := mongodb.GetMongoRepository()
	id, err := repo.SaveNPC(npc)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", id), nil
}

// DeleteNPC func
func DeleteNPC(npcID primitive.ObjectID) (int64, error) {
	repo := mongodb.GetMongoRepository()
	result, err := repo.DeleteNPCByID(npcID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

//DamageNPC func
func DamageNPC(npcID primitive.ObjectID, action string, hit int) (string, error) {
	var message string
	var hp int
	repo := mongodb.GetMongoRepository()
	if action == "remove" {
		hp = hit * -1
	}
	id, err := repo.SetDamageNPCByID(npcID, hp)
	if err != nil {
		return "its impossible. Cannot remove any HP.", err
	}
	if id >= 1 {
		message = fmt.Sprintf("NPC %s %v HP", action, hit)
	}

	// if npc.Condition != "" {
	//      res, err := repo.SetContitionNPCByID(npcID, npc.Condition)
	//      if err != nil {
	//              messageError := fmt.Sprintf("no contition added %v", err)
	//              return messageError, err
	//      }
	//      if res >= 1 {
	//              message += fmt.Sprintf(" contition %s added", npc.Condition)
	//      }

	// }
	return message, nil
}

//ChangeNPCCondition func
func ChangeNPCCondition(npcID primitive.ObjectID, npc *player.NPCCondition) (string, error) {
	repo := mongodb.GetMongoRepository()
	id, err := repo.SetContitionNPCByID(npcID, npc)
	if err != nil {
		return "", err
	}
	var message string
	if id >= 1 {
		message = fmt.Sprintf("NPC updated Advantages %v, Conditions %v, Disvantages %v and Auto Fail list %v, changes number %v", npc.Advantages, npc.Condition, npc.Disvantages, npc.AutoFail, id)
	}
	return message, nil
}
