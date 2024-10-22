package cs

import (
	"fmt"
	"math"
)

// A CostFloat64 is otherwise identical to a regular Cost struct, but uses float64s instead of ints
// used for internal calculations before being cast back into a regular Cost
type CostFloat64 struct {
	Ironium   float64 `json:"ironium,omitempty"`
	Boranium  float64 `json:"boranium,omitempty"`
	Germanium float64 `json:"germanium,omitempty"`
	Resources float64 `json:"resources,omitempty"`
}

func NewCostFloat64(
	ironium float64,
	boranium float64,
	germanium float64,
	resources float64,
) CostFloat64 {
	return CostFloat64{ironium, boranium, germanium, resources}
}

func (c CostFloat64) GetAmount(costType CostType) float64 {
	switch costType {
	case Ironium:
		return c.Ironium
	case Boranium:
		return c.Boranium
	case Germanium:
		return c.Germanium
	case Resources:
		return c.Resources
	}
	panic(fmt.Sprintf("GetAmount called with invalid CostType %s", costType))
}

// convert a CostFloat64 to an int using the specified rounding method
func (c CostFloat64) ToCost(roundFunc func(float64) float64) Cost {
	return Cost{
		Ironium:   int(roundFunc(c.Ironium)),
		Boranium:  int(roundFunc(c.Boranium)),
		Germanium: int(roundFunc(c.Germanium)),
		Resources: int(roundFunc(c.Resources)),
	}
}

func (c CostFloat64) AddFloat64(costType CostType, amount float64) CostFloat64 {
	switch costType {
	case Ironium:
		c.Ironium += amount
	case Boranium:
		c.Boranium += amount
	case Germanium:
		c.Germanium += amount
	case Resources:
		c.Resources += amount
	default:
		panic(fmt.Sprintf("AddFloat64 called with invalid CostType %s", costType))
	}
	return c
}

func (c CostFloat64) Add(other CostFloat64) CostFloat64 {
	return CostFloat64{
		Ironium:   c.Ironium + other.Ironium,
		Boranium:  c.Boranium + other.Boranium,
		Germanium: c.Germanium + other.Germanium,
		Resources: c.Resources + other.Resources,
	}
}

func (c CostFloat64) Subtract(other CostFloat64) CostFloat64 {
	return CostFloat64{
		Ironium:   c.Ironium - other.Ironium,
		Boranium:  c.Boranium - other.Boranium,
		Germanium: c.Germanium - other.Germanium,
		Resources: c.Resources - other.Resources,
	}
}

func (c CostFloat64) Multiply(factor float64) CostFloat64 {
	return CostFloat64{
		Ironium:   c.Ironium * factor,
		Boranium:  c.Boranium * factor,
		Germanium: c.Germanium * factor,
		Resources: c.Resources * factor,
	}
}

func (c CostFloat64) Divide(divisor float64) CostFloat64 {
	return CostFloat64{
		Ironium:   c.Ironium / divisor,
		Boranium:  c.Boranium / divisor,
		Germanium: c.Germanium / divisor,
		Resources: c.Resources / divisor,
	}
}

// Return greater of 2 cost structs for all ResourceTypes separately
func (c CostFloat64) Max(other CostFloat64) CostFloat64 {
	return CostFloat64{
		Ironium:   math.Max(c.Ironium, other.Ironium),
		Boranium:  math.Max(c.Boranium, other.Boranium),
		Germanium: math.Max(c.Germanium, other.Germanium),
		Resources: math.Max(c.Resources, other.Resources),
	}
}

// round a cost struct's values with passed in function
func (c CostFloat64) Round(roundFunc func(float64) float64) CostFloat64 {
	return CostFloat64{
		Ironium:   roundFunc(c.Ironium),
		Boranium:  roundFunc(c.Boranium),
		Germanium: roundFunc(c.Germanium),
		Resources: roundFunc(c.Resources),
	}
}
