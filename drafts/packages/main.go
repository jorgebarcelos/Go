package main

import (
	"fmt"
	pkg "module/pkg"
)

func main() {
	fmt.Println(pkg.UserData{Name: "Jorge", Age: 41,})
	fmt.Println(pkg.SumNumbers(16, 45))
    fmt.Println(pkg.MyList([]int{1, 2, 3, 4, 5}))
	fmt.Println(pkg.SumAll(1,2,3,4,8,5,6))
	fmt.Printf("The day is %s\n", pkg.DayWeek(10))
	fmt.Println(pkg.LoopFor(10))
	fmt.Println(pkg.RangeValues([]string{"A", "B", "C", "CÚ"}))
}