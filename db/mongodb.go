package db

import "gopkg.in/mgo.v2"

var (
	mongoSession *mgo.Session
	mongo        *mgo.DialInfo
)

func MongoConnect() {
	var err error
	var url = "mongodb://businessdb:6TQbbC9rrGuL3c@10.10.108.86:27017,10.10.108.74:27017,10.10.108.89:27017/businessdb?replicaSet=mymongodb"
	mongo, err = mgo.ParseURL(url)
	if err != nil {
		panic(err)
	}
	mongoSession, err = mgo.Dial(url)
	if err != nil {
		panic(err.Error())
	}
	mongoSession.SetMode(mgo.Eventual, false)
	mongoSession.SetSafe(&mgo.Safe{})
}

type mgoDB struct {
	*mgo.Database
}

func (m *mgoDB) Close() {
	m.Session.Close()
}

func MongoDB() *mgoDB {
	clone := mongoSession.Clone()
	internal := clone.DB(mongo.Database)
	return &mgoDB{internal}
}
