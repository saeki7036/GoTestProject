package main

import (
	"fmt"

	"github.com/saeki7036/GoTestProject/internal/logic"
	"github.com/saeki7036/GoTestProject/internal/mathutil"
	"github.com/saeki7036/GoTestProject/internal/printin"
)

func main() {
	printin.PrintTest()
	fmt.Println("Result:", mathutil.Add(10, 5))
	logic.VarTest()
}
