package cs

import "math"

// handles terraforming planets
type Terraformer interface {
	PermaformOneStep(planet *Planet, player *Player, habType HabType) TerraformResult
	TerraformOneStep(planet *Planet, player *Player, terraformer *Player, reverse bool) TerraformResult
	GetBestTerraform(planet *Planet, player *Player, terraformer *Player) *HabType
	getTerraformAmount(planet *Planet, player, terraformer *Player) Hab
	getMinTerraformAmount(planet *Planet, player *Player, terraformer *Player) Hab
}

type TerraformResult struct {
	Type      HabType
	Direction int
}

func (t TerraformResult) Terraformed() bool {
	return t.Direction != 0
}

type terraform struct {
}

func NewTerraformer() Terraformer {
	return &terraform{}
}

// getTerraformAbility returns the terraform ability of a player taking into account total terraform and hab terraform
func (t *terraform) getTerraformAbility(player *Player) Hab {
	bestTotalTerraform := player.Spec.Terraform[TerraformHabTypeAll]
	totalTerraformAbility := 0
	if bestTotalTerraform != nil {
		totalTerraformAbility = bestTotalTerraform.Ability
	}

	terraformAbility := Hab{totalTerraformAbility, totalTerraformAbility, totalTerraformAbility}

	for _, habType := range HabTypes {
		// get the two ways we can terraform
		bestHabTerraform := player.Spec.Terraform[FromHabType(habType)]

		// find out which terraform tech has the greater terraform ability
		ability := totalTerraformAbility
		if bestHabTerraform != nil {
			ability = maxInt(ability, bestHabTerraform.Ability)
			terraformAbility = terraformAbility.Set(habType, ability)
		}

		if ability == 0 {
			continue
		}
	}

	return terraformAbility
}

// getTerraformAmount returns the total amount we can terraform this planet
func (t *terraform) getTerraformAmount(planet *Planet, player, terraformer *Player) Hab {
	terraformAmount := Hab{}
	if player == nil || planet == nil {
		// can't terraform, return an empty Hab
		return terraformAmount
	}

	// default the terraformer to the player
	if terraformer == nil {
		terraformer = player
	}

	terraformAbility := t.getTerraformAbility(terraformer)
	enemy := terraformer.IsEnemy(player.Num)
	habCenter := player.Race.HabCenter()

	for _, habType := range HabTypes {
		if player.Race.IsImmune(habType) {
			continue
		}

		ability := terraformAbility.Get(habType)
		if ability == 0 {
			continue
		}

		// The distance from the starting hab of this planet
		fromIdealBase := habCenter.Get(habType) - planet.BaseHab.Get(habType)

		// the distance from the current hab of this planet
		fromIdeal := habCenter.Get(habType) - planet.Hab.Get(habType)

		// if we have any left to terraform
		if fromIdeal > 0 {
			// i.e. our ideal is 50 and the planet hab is 47
			if enemy {
				alreadyTerraformed := absInt(fromIdealBase - fromIdeal)
				terraformAmount = terraformAmount.Set(habType, -(ability - alreadyTerraformed))
			} else {
				// we can either terrform up to our full ability, or however much
				// we have left to terraform on this
				alreadyTerraformed := fromIdealBase - fromIdeal
				terraformAmount = terraformAmount.Set(habType, minInt(ability-alreadyTerraformed, fromIdeal))
			}
		} else if fromIdeal < 0 {
			if enemy {
				alreadyTerraformed := absInt(fromIdealBase - fromIdeal)
				terraformAmount = terraformAmount.Set(habType, ability-alreadyTerraformed)
			} else {
				// i.e. our ideal is 50 and the planet hab is 53
				alreadyTerraformed := fromIdeal - fromIdealBase
				terraformAmount = terraformAmount.Set(habType, maxInt(-(ability-alreadyTerraformed), fromIdeal))
			}
		} else if enemy {
			// the terrformer is enemies with the player, terraform away from ideal
			alreadyTerraformed := absInt(fromIdealBase - fromIdeal)
			terraformAmount = terraformAbility.Set(habType, ability-alreadyTerraformed)
		}
	}

	return terraformAmount
}

// getMinTerraformAmount gets the minimum amount we need to terraform this planet to make it habitable (if we can terraform it at all)
func (t *terraform) getMinTerraformAmount(planet *Planet, player *Player, terraformer *Player) Hab {
	terraformAmount := Hab{}
	if player == nil || planet == nil {
		// can't terraform, return an empty Hab
		return terraformAmount
	}

	// default the terraformer to the player
	if terraformer == nil {
		terraformer = player
	}

	// get how much this player can terraform each hab
	terraformAbility := t.getTerraformAbility(terraformer)

	habCenter := player.Race.HabCenter()

	for _, habType := range HabTypes {
		if player.Race.IsImmune(habType) {
			continue
		}

		ability := terraformAbility.Get(habType)

		if ability == 0 {
			continue
		}

		// the distance from the current hab of this planet to our minimum hab threshold
		// If this is positive, it means we need to terraform a certain percent to get it in range
		fromHabitableDistance := 0
		planetHabValue := planet.Hab.Get(habType)
		playerHabIdeal := habCenter.Get(habType)
		if planetHabValue > playerHabIdeal {
			// this planet is higher that we want, check the upper bound distance
			// if the player's high temp is 85, and this planet is 83, we are already in range
			// and don't need to min-terraform. If the planet is 87, we need to drop it 2 to be in range.
			fromHabitableDistance = planetHabValue - player.Race.HabHigh.Get(habType)
		} else {
			// this planet is lower than we want, check the lower bound distance
			// if the player's low temp is 15, and this planet is 17, we are already in range
			// and don't need to min-terraform. If the planet is 13, we need to increase it 2 to be in range.
			fromHabitableDistance = player.Race.HabLow.Get(habType) - planetHabValue
		}

		// if we are already in range, set this to 0 because we don't want to terraform anymore
		if fromHabitableDistance < 0 {
			fromHabitableDistance = 0
		}

		// if we have any left to terraform
		if fromHabitableDistance > 0 {
			// the distance from the current hab of this planet
			fromIdeal := playerHabIdeal - planetHabValue
			fromIdealDistance := absInt(fromIdeal)

			// The distance from the starting hab of this planet
			fromIdealBaseDistance := absInt(playerHabIdeal - planet.BaseHab.Get(habType))

			// we can either terrform up to our full ability, or however much
			// we have left to terraform on this
			alreadyTerraformed := fromIdealBaseDistance - fromIdealDistance
			terraformAmountPossible := minInt(ability-alreadyTerraformed, fromIdealDistance)

			// if we are in range for this hab type, we won't terraform at

			// if we are in range for this hab type, we won't terraform at all, otherwise return the max possible terraforming
			// left.
			terraformAmount = terraformAmount.Set(habType, minInt(fromHabitableDistance, terraformAmountPossible))
		}

	}
	return terraformAmount
}

// Get the best hab to terraform (the one with the most distance away from ideal that we can still terraform)
func (t *terraform) GetBestTerraform(planet *Planet, player *Player, terraformer *Player) *HabType {
	if player == nil || planet == nil {
		return nil
	}
	// if we can terraform any, return true
	largestDistance := 0
	var bestHabType *HabType

	// default the terraformer to the player
	if terraformer == nil {
		terraformer = player
	}

	// get how much this player can terraform each hab
	terraformAbility := t.getTerraformAbility(terraformer)

	habCenter := player.Race.HabCenter()

	for _, habType := range HabTypes {
		if player.Race.IsImmune(habType) {
			continue
		}

		ability := terraformAbility.Get(habType)
		if ability == 0 {
			continue
		}

		playerHabIdeal := habCenter.Get(habType)

		// The distance from the starting hab of this planet
		fromIdealBaseDistance := absInt(playerHabIdeal - planet.BaseHab.Get(habType))

		// figure out what our hab is without any instaforming
		habWithoutInstaforming := planet.BaseHab.Add(planet.TerraformedAmount)

		// the distance from the current hab of this planet
		fromIdeal := playerHabIdeal - habWithoutInstaforming.Get(habType)
		fromIdealDistance := absInt(fromIdeal)

		terraformAmount := 0
		// if we have any left to terraform
		if fromIdealDistance > 0 {
			// we can either terrform up to our full ability, or however much
			// we have left to terraform on this
			alreadyTerraformed := fromIdealBaseDistance - fromIdealDistance
			terraformAmount = minInt(ability-alreadyTerraformed, fromIdealDistance)
		}

		// we want to get the largest distance that we can terraform
		if fromIdealDistance > largestDistance && terraformAmount > 0 {
			largestDistance = fromIdealDistance
			newBest := habType
			bestHabType = &newBest
		}
	}

	return bestHabType
}

// getBestUnterraform returns the worst habitable parameter to terraform on the planet, or nil if no terraforming can be done.
func (t *terraform) getBestUnterraform(planet *Planet, player, terraformer *Player) *HabType {
	if player == nil || planet == nil {
		return nil
	}
	// if we can terraform any, return true
	farthestAmount := math.MinInt
	var farthest *HabType

	// default the terraformer to the player
	if terraformer == nil {
		terraformer = player
	}

	// get how much this player can terraform each hab
	terraformAbility := t.getTerraformAbility(terraformer)

	habCenter := player.Race.HabCenter()

	for _, habType := range HabTypes {
		if player.Race.IsImmune(habType) {
			continue
		}

		fromIdeal := habCenter.Get(habType) - planet.Hab.Get(habType)
		terraformedAlready := absInt(planet.TerraformedAmount.Get(habType))

		// if we can terraform this at all
		if terraformedAlready < terraformAbility.Get(habType) {
			// pick the farthest from ideal
			if absInt(fromIdeal) > farthestAmount {
				farthestAmount = absInt(fromIdeal)
				newFarthest := habType
				farthest = &newFarthest
			}
		}
	}

	return farthest
}

// TerraformOneStep terraforms the planet one step in whatever the best option is
// If reverse is true, this will terraform in the opposite direction making the planet less habitable
func (t *terraform) TerraformOneStep(planet *Planet, player *Player, terraformer *Player, reverse bool) TerraformResult {
	var bestHab *HabType
	if !reverse || terraformer == nil {
		bestHab = t.GetBestTerraform(planet, player, terraformer)
	} else {
		bestHab = t.getBestUnterraform(planet, player, terraformer)
	}

	if bestHab != nil {
		habType := *bestHab
		fromIdeal := player.Race.Spec.HabCenter.Get(habType) - planet.Hab.Get(habType)
		if fromIdeal > 0 {
			// for example, the planet has Grav 49, but our player wants Grav 50
			direction := 1
			if reverse {
				direction = -1
			}
			planet.Hab = planet.Hab.Set(habType, planet.Hab.Get(habType)+direction)
			planet.TerraformedAmount = planet.TerraformedAmount.Set(habType, planet.TerraformedAmount.Get(habType)+direction)
			return TerraformResult{Type: habType, Direction: direction}
		} else if fromIdeal < 0 {
			// for example, the planet has Grav 51, but our player wants Grav 50
			direction := -1
			if reverse {
				direction = 1
			}
			planet.Hab = planet.Hab.Set(habType, planet.Hab.Get(habType)+direction)
			planet.TerraformedAmount = planet.TerraformedAmount.Set(habType, planet.TerraformedAmount.Get(habType)+direction)
			return TerraformResult{Type: habType, Direction: direction}
		} else if fromIdeal == 0 && reverse {
			// this is a perfect planet, just make it hotter, higher rad, etc
			direction := 1
			planet.Hab = planet.Hab.Set(habType, planet.Hab.Get(habType)+direction)
			planet.TerraformedAmount = planet.TerraformedAmount.Set(habType, planet.TerraformedAmount.Get(habType)+direction)
			return TerraformResult{Type: habType, Direction: direction}
		}
	}
	return TerraformResult{}
}

// Permanently terraform this planet one step for a random habtype
// This adjusts the BaseHab as well as the hab
func (t *terraform) PermaformOneStep(planet *Planet, player *Player, habType HabType) TerraformResult {
	direction := 0

	habCenter := player.Race.HabCenter()
	playerHabIdeal := habCenter.Get(habType)
	baseHab := planet.BaseHab.Get(habType)
	fromIdealBaseHab := playerHabIdeal - baseHab
	if fromIdealBaseHab > 0 {
		// for example, the planet has Grav 49, but our player wants Grav 50
		planet.BaseHab = planet.BaseHab.Set(habType, baseHab+1)
		direction = 1
	} else if fromIdealBaseHab < 0 {
		// for example, the planet has Grav 51, but our player wants Grav 50
		planet.BaseHab = planet.BaseHab.Set(habType, baseHab-1)
		direction = -1
	}

	// if this means our terraformed hab is beter as well, improve it
	fromIdealHab := playerHabIdeal - planet.Hab.Get(habType)
	if fromIdealHab > 0 {
		planet.Hab = planet.Hab.Set(habType, planet.Hab.Get(habType)+1)
	} else if fromIdealHab < 0 {
		planet.Hab = planet.Hab.Set(habType, planet.Hab.Get(habType)-1)
	}

	return TerraformResult{Type: habType, Direction: direction}
}
