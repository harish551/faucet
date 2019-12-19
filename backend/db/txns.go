package db

import (
	"github.com/kataras/golog"
	"github.com/vitwit/faucet/backend/types"
	"gopkg.in/mgo.v2/bson"
)

func AddTransaction(txns types.Transactions) error {
	var TxCollection = MongoSession.DB(DB_NAME).C(TxnCollection)

	err := TxCollection.Insert(&txns)
	return err
}

func GetTransactions(query bson.M) ([]types.Transactions, error) {
	var txns []types.Transactions
	TxCollection := MongoSession.DB(DB_NAME).C(TxnCollection)
	err := TxCollection.Find(query).All(&txns)
	if err != nil {
		golog.Error("Error while fetching transactions ", err)
		return txns, err
	}
	return txns, nil
}
