package packageA

import (
	"cyclic-import-example/common"
	"cyclic-import-example/packageB"
	"fmt"
)

func FunctionA() {
	fmt.Println("FunctionA called")
	packageB.FunctionB()
	common.CommonFunction()
}
