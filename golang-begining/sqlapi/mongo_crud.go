package sqlapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func MongoCrudDemo() {
	// fmt.Println("----------------- InsertOne ---------------------")
	// insertOne()
	// fmt.Println("----------------- InsertMany ---------------------")
	// insertMany()
	// fmt.Println("----------------- find ---------------------")
	// findStu()
	// fmt.Println("----------------- 文档更新 ---------------------")
	// updateStu()
	fmt.Println("----------------- 文档删除 ---------------------")
	delDemo()
}

type Student struct {
	Name string
	Age  int
}

/* 插入文档，即插入 bson */
func insertOne() {
	stu := Student{
		Name: "张三",
		Age:  18,
	}
	collection := Global_Mongo_Client.Database("go_db").Collection("student")

	insertRes, err := collection.InsertOne(context.TODO(), stu)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Inserted a single document: ", insertRes.InsertedID)
	}
}

/* 插入多条文档 */
func insertMany() {
	stu1 := Student{
		Name: "张龙",
		Age:  37,
	}
	stu2 := Student{
		Name: "赵虎",
		Age:  35,
	}
	documents := []interface{}{stu1, stu2}

	collection := Global_Mongo_Client.Database("go_db").Collection("student")
	imr, err := collection.InsertMany(context.TODO(), documents)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("[InsertMany] imr.InsertedIDs: %v\n", imr.InsertedIDs)
	}
}

/* 搜索 文档 */
func findStu() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := Global_Mongo_Client.Database("go_db").Collection("student")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	// 最终释放 mongo 连接
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
		fmt.Printf("result.Map()[\"name\"]: %v\n", result.Map()["name"])
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

/* 更新文档 */
func updateStu() {
	ctx := context.TODO()
	defer Global_Mongo_Client.Disconnect(ctx)

	c := Global_Mongo_Client.Database("go_db").Collection("student")

	updateBson := bson.D{{"$set", bson.D{{"name", "张山-2"}, {"age", 8}}}}

	ur, err := c.UpdateOne(ctx, bson.D{{"name", "张山"}}, updateBson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("更新条数 ur.ModifiedCount: %v\n", ur.ModifiedCount)
}

/* 删除文档 */
func delDemo() {
	ctx := context.TODO()
	c := Global_Mongo_Client.Database("go_db").Collection("student")

	dr, err := c.DeleteOne(ctx, bson.D{{"name", "张山-2"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("删除条数 dr.DeletedCount: %v\n", dr.DeletedCount)
}
