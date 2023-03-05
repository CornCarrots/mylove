package packed

import (
	"github.com/bwmarrin/snowflake"
)

var IDGenerator *snowflake.Node

func init() {
	var err error
	IDGenerator, err = snowflake.NewNode(1)
	if err != nil {
		println(err)
	}
}
