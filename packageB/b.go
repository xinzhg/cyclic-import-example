package packageB

import (
	"cyclic-import-example/common"
	"cyclic-import-example/packageA"
	"fmt"
)

func FunctionB() {
	fmt.Println("FunctionB called")
	packageA.FunctionA()
	common.CallFunctionAFromCommon()
}
