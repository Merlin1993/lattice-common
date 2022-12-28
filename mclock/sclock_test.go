package mclock

import (
	"fmt"
	"testing"
	"time"
)

func TestSyncTime_Now(t *testing.T) {
	for i := 1; i < 10; i++ {
		Delay = 4 * time.Second
		fmt.Println(GenesisClock().Unix())
		fmt.Println(time.Now().Unix())
	}
}
