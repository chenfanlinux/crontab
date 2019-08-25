package common

import (
	"encoding/json"
)

// 定时任务
type Job struct {
	Name string `json:"name"`  // 任务名
	Command string `json:"command"`   // shell命令
	CronExpr string `json:"cronExpr"`  // cron表达式
}


//type Job struct {
//	Name string `json:"name"`	//  任务名
//	Command string	`json:"command"` // shell命令
//	CronExpr string	`json:"cronExpr"`	// cron表达式
//}


// HTTP接口应答
type Response struct {
	Erron int `json:"errno"`
	Msg string `json:"msg"`
	Data interface{}  `json:"data"`
}


// 应答方法
func BuildResponse(errno int, msg string, data interface{})(resp []byte, err error){

	// 1. 定义一个response
	var (
		response Response
	)
	response.Erron = errno
	response.Data = data
	response.Msg = msg

	// 2. 序列化json
	resp, err = json.Marshal(response)

	return


}
