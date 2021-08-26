package test

// import (
// 	"fmt"
// 	"net/url"

// 	"github.com/betorvs/playbypost-dnd/appcontext"
// 	"github.com/betorvs/playbypost-dnd/domain/game"
// 	"github.com/betorvs/playbypost-dnd/domain/player"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// var (
// 	// MongoDBCalls int
// 	MongoDBCalls int
// )

// // MongoRepositoryMock struct is used for Mock MongoRepositoryMock requests
// type MongoRepositoryMock struct {
// }

// //SaveGame func
// func (repo MongoRepositoryMock) SaveGame(game *game.Game) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// // UpdateGame func
// func (repo MongoRepositoryMock) UpdateGame(gameID primitive.ObjectID, game *game.Game) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // AddPlayerToGame func
// func (repo MongoRepositoryMock) AddPlayerToGame(gameID primitive.ObjectID, playerID, playerName string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetGame func
// func (repo MongoRepositoryMock) GetGame(queryParameters url.Values) ([]*game.Game, error) {
// 	games := new(game.Game)
// 	games.MasterID = "testmaster"
// 	var sliceGames []*game.Game
// 	sliceGames = append(sliceGames, games)
// 	MongoDBCalls++
// 	if len(queryParameters) != 0 {
// 		if queryParameters["master_id"][0] == "testerror" {
// 			return sliceGames, fmt.Errorf("error")
// 		}
// 	}
// 	return sliceGames, nil
// }

// // GetOneGame func
// func (repo MongoRepositoryMock) GetOneGame(gameID primitive.ObjectID) (*game.Game, error) {
// 	games := new(game.Game)
// 	MongoDBCalls++
// 	return games, nil
// }

// // DeleteGameByID func
// func (repo MongoRepositoryMock) DeleteGameByID(gameID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //AddGameDay func
// func (repo MongoRepositoryMock) AddGameDay(gameday *game.Day) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// // UpdateGameDay func
// func (repo MongoRepositoryMock) UpdateGameDay(gameDayID primitive.ObjectID, gameday *game.Day) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetGameDay func
// func (repo MongoRepositoryMock) GetGameDay(queryParameters url.Values) ([]*game.Day, error) {
// 	day := new(game.Day)
// 	var sliceGameDay []*game.Day
// 	sliceGameDay = append(sliceGameDay, day)
// 	MongoDBCalls++
// 	return sliceGameDay, nil
// }

// // GetOneGameDay func
// func (repo MongoRepositoryMock) GetOneGameDay(gameDayID primitive.ObjectID) (*game.Day, error) {
// 	day := new(game.Day)
// 	MongoDBCalls++
// 	return day, nil
// }

// // AddFightSceneToGameDay func
// func (repo MongoRepositoryMock) AddFightSceneToGameDay(gameDayID primitive.ObjectID, fight string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //ChangeGameDayStatusByID func
// func (repo MongoRepositoryMock) ChangeGameDayStatusByID(gameDayID primitive.ObjectID, status string, playersID *game.AddPlayersID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // DeleteGameDayByID func
// func (repo MongoRepositoryMock) DeleteGameDayByID(gameDayID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //AddJournals func
// func (repo MongoRepositoryMock) AddJournals(journals *game.Journals) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// // GetJournals func
// func (repo MongoRepositoryMock) GetJournals(queryParameters url.Values) ([]*game.Journals, error) {
// 	journal := new(game.Journals)
// 	var sliceJournal []*game.Journals
// 	sliceJournal = append(sliceJournal, journal)
// 	MongoDBCalls++
// 	return sliceJournal, nil
// }

// // DeleteJournalsByID func
// func (repo MongoRepositoryMock) DeleteJournalsByID(journalsID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //AddFightScene func
// func (repo MongoRepositoryMock) AddFightScene(fight *game.FightScene) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// //UpdateFightScene func
// func (repo MongoRepositoryMock) UpdateFightScene(fightID primitive.ObjectID, fight *game.FightScene) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetFightScene func
// func (repo MongoRepositoryMock) GetFightScene(queryParameters url.Values) ([]*game.FightScene, error) {
// 	fight := new(game.FightScene)
// 	var sliceFightScene []*game.FightScene
// 	sliceFightScene = append(sliceFightScene, fight)
// 	MongoDBCalls++
// 	return sliceFightScene, nil
// }

// // GetOneFightScene func
// func (repo MongoRepositoryMock) GetOneFightScene(fightID primitive.ObjectID) (*game.FightScene, error) {
// 	fight := new(game.FightScene)
// 	MongoDBCalls++
// 	return fight, nil
// }

// // AddNPCTOFightScene func
// func (repo MongoRepositoryMock) AddNPCTOFightScene(fightID primitive.ObjectID, npc string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //ChangeFightSceneStatusByID func
// func (repo MongoRepositoryMock) ChangeFightSceneStatusByID(fightID primitive.ObjectID, status string, playersID *game.AddPlayersID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // DeleteFightSceneByID func
// func (repo MongoRepositoryMock) DeleteFightSceneByID(fightSceneID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetPlayers func
// func (repo MongoRepositoryMock) GetPlayers(queryParameters url.Values) ([]*player.Players, error) {
// 	play := new(player.Players)
// 	play.SlackID = "testplayer"
// 	var slicePlayers []*player.Players
// 	slicePlayers = append(slicePlayers, play)
// 	MongoDBCalls++
// 	if len(queryParameters) != 0 {
// 		if queryParameters["slack_id"][0] == "testerror" {
// 			return slicePlayers, fmt.Errorf("error")
// 		}
// 	}
// 	return slicePlayers, nil
// }

// // GetOnePlayer func
// func (repo MongoRepositoryMock) GetOnePlayer(playerID primitive.ObjectID) (*player.Players, error) {
// 	play := new(player.Players)
// 	MongoDBCalls++
// 	return play, nil
// }

// //SavePlayer func
// func (repo MongoRepositoryMock) SavePlayer(player *player.Players) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// //UpdatePlayer func
// func (repo MongoRepositoryMock) UpdatePlayer(playerID primitive.ObjectID, player *player.Players) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // AddGameToPlayer func
// func (repo MongoRepositoryMock) AddGameToPlayer(playerID primitive.ObjectID, gameID, gameTitle, slackChannelID string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // ChangePlayerCondition func
// func (repo MongoRepositoryMock) ChangePlayerCondition(playerID primitive.ObjectID, player *player.Condition) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // AddPlayerXP func
// func (repo MongoRepositoryMock) AddPlayerXP(playerID primitive.ObjectID, xp int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // AddOrRemovePlayerHP func
// func (repo MongoRepositoryMock) AddOrRemovePlayerHP(playerID primitive.ObjectID, hit int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //CreateInventory func
// func (repo MongoRepositoryMock) CreateInventory(inventory *player.Inventory) error {
// 	MongoDBCalls++
// 	return nil
// }

// //DeleteInventoryByID delete Inventory by ID
// func (repo MongoRepositoryMock) DeleteInventoryByID(inventoryID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetInventoryID func
// func (repo MongoRepositoryMock) GetInventoryID(playerID string) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// //AddTreasurePlayer func
// func (repo MongoRepositoryMock) AddTreasurePlayer(playerID primitive.ObjectID, coins map[string]int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //AddOthersItems func
// func (repo MongoRepositoryMock) AddOthersItems(playerID primitive.ObjectID, item []string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //RemoveOthersItems func
// func (repo MongoRepositoryMock) RemoveOthersItems(playerID primitive.ObjectID, item []string) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // GetArmory func
// func (repo MongoRepositoryMock) GetArmory(inventoryID primitive.ObjectID) (*player.Armory, error) {
// 	MongoDBCalls++
// 	armory := new(player.Armory)
// 	return armory, nil
// }

// //UsageSpellByLevel func spells_used.level0
// func (repo MongoRepositoryMock) UsageSpellByLevel(playerID primitive.ObjectID, spellByLevel map[string]int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // SetHPTempPlayerByID func
// func (repo MongoRepositoryMock) SetHPTempPlayerByID(playerID primitive.ObjectID, hp int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // SetSpellUsedPlayerByID func
// func (repo MongoRepositoryMock) SetSpellUsedPlayerByID(playerID primitive.ObjectID, spell map[string]int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // SetArmorWeaponPlayerByID func
// func (repo MongoRepositoryMock) SetArmorWeaponPlayerByID(playerID primitive.ObjectID, armor *player.Armory) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // DeletePlayerByID func
// func (repo MongoRepositoryMock) DeletePlayerByID(playerID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //SaveNPC func
// func (repo MongoRepositoryMock) SaveNPC(npc *player.NPC) (primitive.ObjectID, error) {
// 	MongoDBCalls++
// 	return primitive.ObjectID{}, nil
// }

// // GetNPCS func
// func (repo MongoRepositoryMock) GetNPCS(queryParameters url.Values) ([]*player.NPC, error) {
// 	npc := new(player.NPC)
// 	var sliceNPC []*player.NPC
// 	sliceNPC = append(sliceNPC, npc)
// 	MongoDBCalls++
// 	return sliceNPC, nil
// }

// // SetDamageNPCByID func
// func (repo MongoRepositoryMock) SetDamageNPCByID(npcID primitive.ObjectID, hit int) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // SetContitionNPCByID func
// func (repo MongoRepositoryMock) SetContitionNPCByID(npcID primitive.ObjectID, npc *player.NPCCondition) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// // DeleteNPCByID func
// func (repo MongoRepositoryMock) DeleteNPCByID(npcID primitive.ObjectID) (int64, error) {
// 	MongoDBCalls++
// 	return 1, nil
// }

// //SaveMonsters func
// // func (repo MongoRepositoryMock) SaveMonsters(npc *player.Monster) (primitive.ObjectID, error) {
// // 	MongoDBCalls++
// // }
// // func (repo MongoRepositoryMock) GetMonsters() ([]*player.Monster, error) {
// // 	MongoDBCalls++
// // }
// // func (repo MongoRepositoryMock) DeleteMonsterByID(monsterID primitive.ObjectID) (int64, error) {
// // 	MongoDBCalls++
// // }

// // InitMongoMock func returns a RepositoryMongoMock interface
// func InitMongoMock() appcontext.Component {

// 	return MongoRepositoryMock{}
// }
