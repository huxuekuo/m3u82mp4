package utils

import (
	"errors"
	"fmt"
	"m3u82mp4/consts"

	"github.com/bwmarrin/snowflake"
)

func ID() (error, int64) {
	node, err := snowflake.NewNode(consts.SYS_SEQUENCE_ID)
	if err != nil {
		fmt.Println(err)
		return errors.New("获取ID失败"), 0
	}
	return nil, node.Generate().Int64()
}
