package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	//TODO 使用go-redis操作

	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.17.128:6379",
		Password: "123456", // no password set
		DB:       15,       // use default DB
	})

	//心跳
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	//哨兵模式
	/*rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})*/

	//集群模式
	/*rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})*/

	//关闭redis
	defer func(rdb *redis.Client) {
		err = rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	//TODO string操作
	//过期时间(time.Second*100/100秒) 0：则为永久
	err = rdb.Set(ctx, "name", "jerry", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "name").Result()
	//获取key不存在则用redis.Nil校验
	if err == redis.Nil {
		fmt.Println("name2 不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("name 值为：", val)
	}

	set, err := rdb.SetNX(ctx, "name", "tom", 0).Result()
	if err != nil {
		panic(err)
	}
	if set {
		fmt.Println("setNX success")
	} else {
		fmt.Println("setNX fail")
	}

	//TODO hash操作
	//批量添加
	userInfo := make(map[string]interface{})
	userInfo["name"] = "小明"
	userInfo["age"] = 18
	userInfo["hobby"] = "篮球"
	err = rdb.HSet(ctx, "user", userInfo).Err()
	if err != nil {
		panic(err)
	}
	//单个添加
	err = rdb.HSet(ctx, "user2", "name", "小红").Err()
	if err != nil {
		panic(err)
	}

	//单个取值
	userName, err := rdb.HGet(ctx, "user", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user name：", userName)

	//多个取值
	userHash, err := rdb.HGetAll(ctx, "user").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user name：", userHash["name"])
	fmt.Println("user age：", userHash["age"])
	fmt.Println("user hobby：", userHash["hobby"])

	//TODO list操作
	//批量入队 right
	err = rdb.RPush(ctx, "userList", "jerry", "tom", "小明").Err()
	if err != nil {
		panic(err)
	}

	//查询队列元素
	userListNum, err := rdb.LLen(ctx, "userList").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userList sum num:", userListNum)

	//出队 left
	popUser, err := rdb.LPop(ctx, "userList").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userList lPop user:", popUser)

	//TODO set操作
	//添加成员
	err = rdb.SAdd(ctx, "userList2", "jerry", "tom", "小明").Err()
	if err != nil {
		panic(err)
	}

	//获取所有元素
	userInfo2, err := rdb.SMembers(ctx, "userList2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userList2 user:", userInfo2)

	//判断元素是否是该集合中的成员
	isMember, err := rdb.SIsMember(ctx, "userList2", "jerry").Result()
	if err != nil {
		panic(err)
	}
	if isMember {
		fmt.Println("是集合userList2中的成员")
	} else {
		fmt.Println("不是集合userList2中的成员")
	}

	//移除指定的单个或多个元素
	remInt, err := rdb.SRem(ctx, "userList2", "jerry", "jerry2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("移除元素的个数为：", remInt)

	//TODO zSet操作
	//添加分数
	tom := redis.Z{Member: "tom", Score: 60}
	jerry := redis.Z{Member: "jerry", Score: 95}
	lucy := redis.Z{Member: "lucy", Score: 78}
	zAddNum, err := rdb.ZAdd(ctx, "userList3", &tom, &jerry, &lucy).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("ZAdd添加成功的个数", zAddNum)

	//给指定的成员分数增量increment
	score, err := rdb.ZIncrBy(ctx, "userList3", 10, "tom").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("tom 分数变化后为：", score)

	//获取集合中的成员数
	userList3Num, err := rdb.ZCard(ctx, "userList3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("userList3 成员数为：", userList3Num)

	//获取指定范围的成员 分数从高到低
	userRange, err := rdb.ZRevRange(ctx, "userList3", 0, 9).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(userRange)

	//获取指定成员的分数
	UserScore, err := rdb.ZScore(ctx, "userList3", "tom").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("集合userList3 成员tom 的分数为：", UserScore)
}
