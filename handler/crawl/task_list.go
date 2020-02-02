package crawl

import (
	"github.com/gin-gonic/gin"
	"gowatcher/go_monitor/consts"
	"gowatcher/go_monitor/exceptions"
	"gowatcher/go_monitor/model"
	"gowatcher/go_monitor/processor"
	"gowatcher/go_monitor/processor/crawl"
)

//TaskList 获取任务
func TaskList(c *gin.Context) {
	context := model.NewTaskListContext()

	preContextLoader := &processor.PreContextLoader{}
	preContextCode := processor.LoaderCommon(c, context, preContextLoader)
	if preContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(preContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	taskListContextLoader := &crawl.TaskLoader{}
	taskListContextCode := processor.LoaderCommon(c, context, taskListContextLoader)
	if taskListContextCode != nil {
		errNo, errTips := exceptions.ErrConvert(taskListContextCode)
		c.JSON(200, gin.H{
			"message":  consts.MsgError,
			"data":     map[string]interface{}{},
			"err_no":   errNo,
			"err_tips": errTips,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  consts.MsgSuccess,
		"data":     packTaskList(context),
		"err_no":   0,
		"err_tips": "成功",
	})
}

func packTaskList(context model.ITaskListContext) map[string]interface{} {
	taskList := context.GetTaskList()
	inputParameter := context.GetInputParameter()
	listMap := make([]map[string]interface{}, len(taskList))
	if inputParameter.CrawlParams.QueryType == consts.IdType {
		listMap = append(listMap, map[string]interface{}{
			"task_id":     taskList[0].ID,
			"app_id":      taskList[0].AppID,
			"app_name":    taskList[0].AppName,
			"status":      taskList[0].Status,
			"create_time": taskList[0].CreateTime,
			"modify_time": taskList[0].ModifyTime,
		})
	} else if inputParameter.CrawlParams.QueryType == consts.ListType {
		for key, val := range taskList {
			listMap[key] = map[string]interface{}{
				"task_id":  val.ID,
				"app_name": val.AppName,
				"status":   val.Status,
			}
		}
	}

	result := map[string]interface{}{
		"task_list": listMap,
	}

	return result
}
