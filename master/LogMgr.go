package master

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/owenliang/crontab/common"
	"time"
)


var (
	G_logMgr *LogMgr
)

// mongodb日志管理
type LogMgr struct {
	client *mongo.Client
	logCollection *mongo.Collection

}


func InitLogMgr()(err error){
	var (
		client *mongo.Client

	)
	if client, err = mongo.Connect(context.TODO(),
		G_config.MongodbUri,
		clientopt.ConnectTimeout(time.Duration(G_config.MongodbConnectTimeout) * time.Millisecond));err!=nil{
		return
	}

	G_logMgr = &LogMgr{
		client:client,
	    logCollection:client.Database("cron").Collection("log"),

	}


	return
}



// 查看任务日志
func (logMgr *LogMgr) ListLog(name string, skip int,limit int)(logArr []*common.JobLog, err error){

	var (
		filter *common.JobLogFilter
		logSort *common.SortLogByStartTime
		cursor mongo.Cursor
		jobLog *common.JobLog



	)

	// len(logArr)
	logArr = make([]*common.JobLog, 0)



	// 过滤条件
	filter = &common.JobLogFilter{JobName:name}

	// 按照任务开始时间倒排
	logSort = &common.SortLogByStartTime{SortOrder:-1}


	if cursor, err = logMgr.logCollection.Find(context.TODO(), filter, findopt.Sort(logSort), findopt.Skip(int64(skip)), findopt.Skip(int64(limit))); err!=nil{
		return
	}

	// 延迟释放游标
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()){
		jobLog = &common.JobLog{}

		// 反序列化bson
		if err = cursor.Decode(jobLog); err!=nil{
			// 有日志不合法
			continue
		}

		logArr = append(logArr, jobLog)
	}








	return

}