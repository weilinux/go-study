package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// ConnectToDB 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}

func main() {
	//TODO 使用mongo-driver操作
	//go get go.mongodb.org/mongo-driver/mongo

	//BSON：MongoDB中的JSON文档存储在名为BSON(二进制编码的JSON)的二进制表示中
	//D：一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongoDB命令。
	//M：一张无序的map。它和D是一样的，只是它不保持顺序。
	//A：一个BSON数组。
	//E：D里面的一个元素。

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://192.168.17.128:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// 指定获取要操作的数据集
	collection := client.Database("cbec_oc").Collection("inventory_sku")

	//CURD...
	//插入单条
	insertOneResult, err := collection.InsertOne(context.TODO(), "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertOneResult.InsertedID)
	//插入多条
	//insertManyResult, err := collection.InsertMany()

	//更新单条
	filter := bson.D{{"sku", "sku1"}}
	update := bson.D{
		{"$inc", bson.D{
			{"classifyData.initial", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	//根据id更新单条
	//updateResult, err := collection.UpdateByID()
	//更新多条
	//updateResult, err := collection.UpdateMany()

	//查询单条
	type InventorySku struct{}
	var inventorySku InventorySku
	//将结果绑定到结构体中
	err = collection.FindOne(context.TODO(), filter).Decode(&inventorySku)
	if err != nil {
		log.Fatal(err)
	}
	//查询选项
	findOptions := options.Find()
	findOptions.SetLimit(2) //更多配置需详查源代码
	//查询多条
	//把bson.D{{}}作为一个filter可以匹配所有文档
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var inventorySku InventorySku
		var results []*InventorySku
		err := cursor.Decode(&inventorySku)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &inventorySku)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	err = cursor.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//删除单条
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	//删除多条
	//deleteResult, err := collection.DeleteMany()
}
