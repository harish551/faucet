package db

import (
	"github.com/kataras/golog"
	"gopkg.in/mgo.v2"
	"os"
	"time"
)

var (
	DB_NAME       = "vWallet"
	TxnCollection = "transactions"

	MongoDbUrl = &mgo.DialInfo{
		Addrs:    []string{string("localhost")},
		Timeout:  30 * time.Second,
		Username: "",
		Password: "",
		Database: "vWallet",
	}
)

var MongoSession *mgo.Session
var err error

func InitDB() {

	MongoSession, err = mgo.DialWithInfo(MongoDbUrl)

	if err != nil {
		golog.Error("Error while connecting to DB: ", err)
		os.Exit(1)
	}

	if err = MongoSession.Ping(); err != nil {
		golog.Error("Error while connecting to Database: ", err)
		defer MongoSession.Close()
		os.Exit(1)
	}
	golog.Info("Database connected successfully ")

}
