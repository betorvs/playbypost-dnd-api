package mongodb

import (
	"github.com/betorvs/playbypost-dnd/appcontext"
	"github.com/betorvs/playbypost-dnd/config"
	"go.mongodb.org/mongo-driver/mongo"
)

//MongoRepository is the MongoDB implementation of the interface domain.sRepository
type MongoRepository struct {
	Conn *mongo.Client
}

//BuildMongoRepository func
func BuildMongoRepository() appcontext.Component {
	dbClient := appcontext.Current.Get(appcontext.DBClient).(*MongoClient)
	return MongoRepository{Conn: dbClient.Conn}
}

func init() {
	if config.Values.TestRun {
		return
	}
	appcontext.Current.Add(appcontext.MongoRepository, BuildMongoRepository)
	logLocal := config.GetLogger()
	logLocal.Info("MongoDB ready")
}
