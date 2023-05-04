/*
@Time : 2022/1/19 16:44
@File : sclock
@Description: sclock主要用于将该节点时间同步为config.toml文件中配置的基准节点的时间，来保证联盟链中各节点时间一致。另，提供可调用当前时间的接口。
*/

package mclock

import (
	"time"
	"zkjg/lattice/common/config"
)

var (
	Delay    time.Duration
	StartNTP = false
)

type syncClock struct {
	time.Time
}

// GenesisClock gets current system time
func GenesisClock() time.Time {
	return syncClock{time.Now().Add(Delay)}.Time
}

// GetCurrentClock 联盟链使用的是同步时钟(创世节点)的时间。该系统中所有涉及的获取当前时间，都调用该函数。
func GetCurrentClock() time.Time {
	if config.IsConsortiumBlockchain || StartNTP {
		return GenesisClock()
	}
	return time.Now()
}
