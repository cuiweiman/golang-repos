package sqlapi

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoDB 创建连接 */
func MongoConnect() {
	// connectTest()
}

/* 全局变量 */
var Global_Mongo_Client *mongo.Client

// func connectTest() {
func init() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// 连接到 MongoDB
	var err error
	Global_Mongo_Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = Global_Mongo_Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connect to MongoDB Successful!")
	}
}
