package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
)

// User 模型
type User struct {
	Id  primitive.ObjectID `bson:"_id"`
	Ips []Ip               `bson:"ips"`
	Key string             `bson:"key"`
}

type Ip struct {
	Ip string `bson:"ip"`
}

var ctx = context.Background()

func main() {
	//TODO 使用mongo-driver操作
	//go get go.mongodb.org/mongo-driver/mongo

	//BSON：MongoDB中的JSON文档存储在名为BSON(二进制编码的JSON)的二进制表示中
	//D：一个BSON文档。这种类型应该在顺序重要的情况下使用，比如MongoDB命令。
	//M：一张无序的map。它和D是一样的，只是它不保持顺序。
	//A：一个BSON数组。
	//E：D里面的一个元素。

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://192.168.17.128:27017").SetAuth(options.Credential{
		AuthSource: "database", //用于身份验证的数据库名 （admin）
		Username:   "username", //用户名
		Password:   "password", //密码
	})

	// 设置连接池最大连接数
	clientOptions.SetMaxPoolSize(100)

	// 设置连接池同时建立的最大连接数
	clientOptions.SetMaxConnecting(10)

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 断开连接
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 指定操作的数据库名
	database := client.Database("cbec_oc")

	// 指定获取要操作的数据集
	collection := database.Collection("inventory_sku")

	// 设置索引
	idx := mongo.IndexModel{
		Keys:    bsonx.Doc{{"name", bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
	idxRet, err := collection.Indexes().CreateOne(context.Background(), idx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.Indexes().CreateOne:", idxRet)

	// 事务 只有闭包方式才能使用事务
	if err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		// 开启事务
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		// 事务内使用读写操作必须传入一个带session-value的context对象 （sessionContext）
		_, err = collection.InsertOne(sessionContext, bson.M{"name": 13, "password": "123456"})
		if err != nil {
			return err
		}

		_, err = collection.InsertOne(sessionContext, bson.M{"name": 14, "password": "123456"})
		if err != nil {
			// 事务回滚
			err := sessionContext.AbortTransaction(sessionContext)
			if err != nil {
				return err
			}
			return err
		} else {
			// 事务提交
			err := sessionContext.CommitTransaction(sessionContext)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	//CURD...
	//插入单条
	insertOneResult, err := collection.InsertOne(ctx, "")
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
	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	//根据id更新单条
	//updateResult, err := collection.UpdateByID()
	//更新多条
	//updateResult, err := collection.UpdateMany()

	//查询单条
	var user *User
	//将结果绑定到结构体中
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	//查询选项
	findOptions := options.Find()
	findOptions.SetLimit(2) //更多配置需详查源代码
	//查询多条
	//把bson.D{{}}作为一个filter可以匹配所有文档
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(ctx) {
		// 创建一个值，将单个文档解码为该值
		var results []*User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	err = cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//删除单条
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	//删除多条
	//deleteResult, err := collection.DeleteMany()
}
