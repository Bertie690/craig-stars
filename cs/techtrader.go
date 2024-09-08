package cs

import "math"

// Check for tech level increases
type techTrader interface {
	// Return the tech field gained for this tech trade event, if any
	techLevelGained(rules *Rules, current, target TechLevel) TechField

	// Return if an MT part was gained for this tech trade event
	//
	// Does not take a part as an argument as the only way MT parts can be transferred
	// is if a ship with one dies in battle or is scrapped
	mtPartGained(rules *Rules, numItems int) bool
}

type techTrade struct {
}

func newTechTrader() techTrader {
	return &techTrade{}
}

// check for a tech level bonus from a player tech level and some target we scrapped, invaded, etc
// https://wiki.starsautohost.org/wiki/Guts_of_Tech_Trading
func (t *techTrade) techLevelGained(rules *Rules, current, target TechLevel) TechField {
	diff := target.Minus(current).MinZero()
	if diff.Sum() <= 0 {
		return TechFieldNone
	}

	for _, field := range TechFields {
		level := diff.Get(field)
		if level > 0 {
			chance := techTradeChance(rules.TechTradeChance, level)
			// check if our random number between 0 and 1 is under the above, i.e. < .375 for 2 levels above
			if rules.random.Float64() <= chance {
				return field
			}
		}
	}

	return TechFieldNone
}

func (t *techTrade) mtPartGained(rules *Rules, numItems int) bool {
	if numItems <= 0 {
		return false
	} else if numItems > 25 {
		// maximum of 25 items per check
		numItems = 25
	}

	chance := rules.MysteryTraderRules.MysteryPartTradeChance * rules.TechTradeChance * float64(numItems)
	if chance > rules.random.Float64() {
		return true
	}
	return false
}

// get the chance of a tech trade. If we are one level above this is:
// .5 * (1 - .5) = .25
// if we are two levels above this is:
// .5 * (1 - .5*.5) = .375
func techTradeChance(baseChance float64, level int) float64 {
	return baseChance * (1 - math.Pow(baseChance, float64(level)))
}
