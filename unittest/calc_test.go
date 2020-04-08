package unittest

import (
	"routerWeb/calc"
	"testing"
)

func TestCalc(t *testing.T)  {
	result := calc.Add(4, 5)
	println(result)
}
