package assistant

import (
	"fmt"
	"time"
)

func LoopFor(number int) string{
	index := 0
	for index < number {
		index++
		fmt.Println(index)
		time.Sleep(time.Second)
	}
	return "Finish"
}