package task

import (
	"Open_IM/pkg/common/config"
	"Open_IM/pkg/common/log"
	"Open_IM/pkg/common/tracelog"
	"Open_IM/pkg/utils"
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

const cronTaskOperationID = "cronTaskOperationID-"
const moduleName = "cron"

func StartCronTask() error {
	log.NewPrivateLog(moduleName)
	log.NewInfo(utils.OperationIDGenerator(), "start cron task", "cron config", config.Config.Mongo.ChatRecordsClearTime)
	fmt.Println("cron task start, config", config.Config.Mongo.ChatRecordsClearTime)
	clearCronTask := msgTool{}
	ctx := context.Background()
	operationID := clearCronTask.getCronTaskOperationID()
	tracelog.SetOperationID(ctx, operationID)
	c := cron.New()
	_, err := c.AddFunc(config.Config.Mongo.ChatRecordsClearTime, clearCronTask.ClearAll)
	if err != nil {
		fmt.Println("start cron failed", err.Error(), config.Config.Mongo.ChatRecordsClearTime)
		return err
	}
	c.Start()
	fmt.Println("start cron task success")
	for {
		time.Sleep(10 * time.Second)
	}
}

func FixSeq(userID, workingGroupID string, fixAllSeq bool) {
	log.NewPrivateLog(moduleName)
	log.NewInfo(utils.OperationIDGenerator(), "start cron task", "cron config", config.Config.Mongo.ChatRecordsClearTime)
	clearCronTask := msgTool{}
	ctx := context.Background()
	operationID := clearCronTask.getCronTaskOperationID()
	tracelog.SetOperationID(ctx, operationID)
	if userID != "" {
		clearCronTask.ClearUsersMsg(ctx, []string{userID})
	}
	if workingGroupID != "" {
		clearCronTask.ClearSuperGroupMsg(ctx, []string{workingGroupID})
	}
	if fixAllSeq {
		clearCronTask.FixAllSeq(ctx)
	}
	fmt.Println("fix seq finished")
}