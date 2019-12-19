package types

import "gopkg.in/mgo.v2/bson"

type Transactions struct {
	Id          bson.ObjectId `json:"id" bson:"id"`
	Type        string        `json:"type" bson:"type"`
	From        string        `json:"from" bson:"from"`
	To          string        `json:"to" bson:"to"`
	Amount      string        `json:"amount" bson:"amount"`
	Denom       string        `json:"denom" bson:"denom"`
	Transfer    TxnRes        `json:"transfer" bson:"transfer"`
	Receive     TxnRes        `json:"receive" bson:"receive"`
	Client1     string        `json:"client1" bson:"client1"`
	Client2     string        `json:"client2" bson:"client2"`
	Connection1 string        `json:"connection1" bson:"connection1"`
	Connection2 string        `json:"connection2" bson:"connection2"`
	Channel1    string        `json:"channel1" bson:"channel1"`
	Channel2    string        `json:"channel2" bson: "channel2"`
	FromNode    string        `json:"fromNode" bson:"fromNode"`
	ToNode      string        `json:"toNode" bson:"toNode"`
	FromChain   string        `json:"fromChain" bson:"fromChain"`
	ToChain     string        `json:"toChain" bson:"toChain"`
}

type TxnRes struct {
	Success   bool   `json:"success" bson:"success"`
	Message   string `json:"message" bson:"message"`
	TxHash    string `json:"txHash" bson:"txHash"`
	Height    string `json:"height" bson:"height"`
	Timestamp string `json:"timestamp" bson:"timestamp"`
}

type ErrorResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}