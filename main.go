package main

import (
    "cyclic-import-example/packageA"
    "cyclic-import-example/packageB"
)

func main() {
    packageA.FunctionA()
    packageB.FunctionB()
}
