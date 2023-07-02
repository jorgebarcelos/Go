package assistant
import (
	"fmt"
	"time"
)

func RangeValues(names []string) string{
	for _, value := range(names){
		fmt.Println(value)
		time.Sleep(time.Second)
	}
	return "Cabou-se"
}