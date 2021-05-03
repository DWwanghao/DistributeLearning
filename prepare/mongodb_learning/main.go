package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//定义时间格式
type TimePoint struct {
	StartTime int64 `bson:"start_time"`
	EndTime   int64 `bson:"end_time"`
}

//定义日志格式
type LogRecond struct {
	JobName   string    `bson:"job_name"`
	Command   string    `bson:"command"`
	Err       error     `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"time_point"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
		return
	}
	database := client.Database("cron")
	//选择表
	collection := database.Collection("log")

	//插入记录
	recond := &LogRecond{
		JobName: "job1",
		Command: "echo hello",
		Err:     nil,
		Content: "hello",
		TimePoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix() + 10,
		},
	}
	i := []interface{}{recond, recond, recond}

	result, err := collection.InsertMany(context.TODO(), i)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, insertedID := range result.InsertedIDs {
		objectID := insertedID.(primitive.ObjectID)
		fmt.Println("自增id:", objectID.Hex())

	}

}
