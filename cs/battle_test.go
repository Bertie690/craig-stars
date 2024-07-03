package cs

import (
	"reflect"
	"testing"

	"github.com/sirgwain/craig-stars/test"
	"github.com/stretchr/testify/assert"
)

func testStalwartDefender(player *Player) *Fleet {
	fleet := &Fleet{
		MapObject: MapObject{
			PlayerNum: player.Num,
		},
		BaseName: "Stalwart Defender",
		Tokens: []ShipToken{
			{
				DesignNum: 1,
				Quantity:  1,
				design: NewShipDesign(player, 1).
					WithHull(Destroyer.Name).
					WithSlots([]ShipDesignSlot{
						{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
						{HullComponent: BetaTorpedo.Name, HullSlotIndex: 2, Quantity: 1},
						{HullComponent: XRayLaser.Name, HullSlotIndex: 3, Quantity: 1},
						{HullComponent: RhinoScanner.Name, HullSlotIndex: 4, Quantity: 1},
						{HullComponent: Crobmnium.Name, HullSlotIndex: 5, Quantity: 1},
						{HullComponent: Overthruster.Name, HullSlotIndex: 6, Quantity: 1},
						{HullComponent: BattleComputer.Name, HullSlotIndex: 7, Quantity: 1},
					}).
					WithSpec(&rules, player)},
		},
		battlePlan:        &player.BattlePlans[0],
		OrbitingPlanetNum: None,
		FleetOrders: FleetOrders{
			Waypoints: []Waypoint{
				NewPositionWaypoint(Vector{}, 5),
			},
		},
	}
	fleet.Spec = ComputeFleetSpec(&rules, player, fleet)
	fleet.Fuel = fleet.Spec.FuelCapacity
	return fleet
}

func testJihadCruiser(player *Player) *Fleet {
	fleet := &Fleet{
		MapObject: MapObject{
			PlayerNum: player.Num,
		},
		BaseName: "Jihad Cruiser",
		Tokens: []ShipToken{
			{
				DesignNum: 1,
				Quantity:  1,
				design: NewShipDesign(player, 1).
					WithHull(Cruiser.Name).
					WithSlots([]ShipDesignSlot{
						{HullComponent: TransStar10.Name, HullSlotIndex: 1, Quantity: 2},
						{HullComponent: Overthruster.Name, HullSlotIndex: 2, Quantity: 1},
						{HullComponent: BattleNexus.Name, HullSlotIndex: 3, Quantity: 1},
						{HullComponent: JihadMissile.Name, HullSlotIndex: 4, Quantity: 2},
						{HullComponent: JihadMissile.Name, HullSlotIndex: 5, Quantity: 2},
						{HullComponent: ElephantScanner.Name, HullSlotIndex: 6, Quantity: 2},
						{HullComponent: Kelarium.Name, HullSlotIndex: 5, Quantity: 2},
					}).
					WithSpec(&rules, player)},
		},
		battlePlan:        &player.BattlePlans[0],
		OrbitingPlanetNum: None,
		FleetOrders: FleetOrders{
			Waypoints: []Waypoint{
				NewPositionWaypoint(Vector{}, 5),
			},
		},
	}
	fleet.Spec = ComputeFleetSpec(&rules, player, fleet)
	fleet.Fuel = fleet.Spec.FuelCapacity
	return fleet
}

// create a new small freighter (with cargo pod) fleet for testing
func testTeamster(player *Player) *Fleet {
	fleet := &Fleet{
		MapObject: MapObject{
			PlayerNum: player.Num,
		},
		BaseName: "Teamster",
		Tokens: []ShipToken{
			{
				Quantity:  1,
				DesignNum: 1,
				design: NewShipDesign(player, 1).
					WithHull(MediumFreighter.Name).
					WithSlots([]ShipDesignSlot{
						{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
						{HullComponent: Crobmnium.Name, HullSlotIndex: 2, Quantity: 1},
						{HullComponent: RhinoScanner.Name, HullSlotIndex: 3, Quantity: 1},
					}).
					WithSpec(&rules, player)},
		},
		battlePlan:        &player.BattlePlans[0],
		OrbitingPlanetNum: None,
	}

	fleet.Spec = ComputeFleetSpec(&rules, player, fleet)
	fleet.Fuel = fleet.Spec.FuelCapacity
	return fleet

}

// create a new small freighter (with cargo pod) fleet for testing
func testPrivateer(player *Player, quantity int) *Fleet {
	fleet := &Fleet{
		MapObject: MapObject{
			PlayerNum: player.Num,
		},
		BaseName: "Privateer",
		Tokens: []ShipToken{
			{
				Quantity:  quantity,
				DesignNum: 1,
				design: NewShipDesign(player, 1).
					WithHull(Privateer.Name).
					WithSlots([]ShipDesignSlot{
						{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
						{HullComponent: Crobmnium.Name, HullSlotIndex: 2, Quantity: 1},
						{HullComponent: CargoPod.Name, HullSlotIndex: 3, Quantity: 1},
						{HullComponent: CargoPod.Name, HullSlotIndex: 4, Quantity: 1},
						{HullComponent: CargoPod.Name, HullSlotIndex: 5, Quantity: 1},
					}).
					WithSpec(&rules, player)},
		},
		battlePlan:        &player.BattlePlans[0],
		OrbitingPlanetNum: None,
	}

	fleet.Spec = ComputeFleetSpec(&rules, player, fleet)
	fleet.Fuel = fleet.Spec.FuelCapacity
	return fleet

}

func Test_battle_regenerateShields(t *testing.T) {
	type args struct {
		player *Player
		token  battleToken
	}
	tests := []struct {
		name        string
		args        args
		wantShields int
	}{
		{
			name: "no regen",
			args: args{
				player: testPlayer().WithNum(1),
				token: battleToken{
					BattleRecordToken: BattleRecordToken{
						PlayerNum: 1,
					},
					stackShields:      50,
					totalStackShields: 100,
				},
			},
			wantShields: 50,
		},
		{
			name: "regen",
			args: args{
				player: NewPlayer(1, NewRace().WithLRT(RS).WithSpec(&rules)).WithNum(1),
				token: battleToken{
					BattleRecordToken: BattleRecordToken{
						PlayerNum: 1,
					},
					stackShields:      50,
					totalStackShields: 100,
				},
			},
			wantShields: 60,
		},
		{
			name: "no regen when shields gone",
			args: args{
				player: NewPlayer(1, NewRace().WithLRT(RS).WithSpec(&rules)).WithNum(1),
				token: battleToken{
					BattleRecordToken: BattleRecordToken{
						PlayerNum: 1,
					},
					stackShields:      0,
					totalStackShields: 100,
				},
			},
			wantShields: 0,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			battle := battle{
				players: map[int]*Player{tt.args.player.Num: tt.args.player},
			}
			battle.regenerateShields(&tt.args.token)

			got := tt.args.token.stackShields
			if got != tt.wantShields {
				t.Errorf("battle.regenerateShields() = %v, want %v", got, tt.wantShields)
			}

		})
	}
}

func Test_battle_willTarget(t *testing.T) {

	type args struct {
		target BattleTarget
		token  battleToken
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// if our token has armed/starbase attributes, it should only target armed or starbases
		{args: args{BattleTargetAny, battleToken{}}, want: true},
		{args: args{BattleTargetStarbase, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: true},
		{args: args{BattleTargetArmedShips, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: true},
		{args: args{BattleTargetNone, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: false},
		{args: args{BattleTargetBombersFreighters, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: false},
		{args: args{BattleTargetUnarmedShips, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: false},
		{args: args{BattleTargetFuelTransports, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: false},
		{args: args{BattleTargetFreighters, battleToken{attributes: battleTokenAttributeArmed | battleTokenAttributeStarbase}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &battle{}
			if got := b.willTarget(tt.args.target, &tt.args.token); got != tt.want {
				t.Errorf("battle.willTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battleToken_getDistanceAway(t *testing.T) {

	type args struct {
		position BattleVector
	}
	tests := []struct {
		name string
		bt   battleToken
		args args
		want int
	}{
		{"no distance", battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{0, 0}}}, args{BattleVector{0, 0}}, 0},
		{"x distance greatest", battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{2, 1}}}, args{BattleVector{4, 2}}, 2},
		{"y distance greatest", battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{1, 2}}}, args{BattleVector{2, 5}}, 3},
		{"negative distance (token behind)", battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{1, 1}}}, args{BattleVector{0, 0}}, 1},
		{"3,4 to 7,4", battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{3, 4}}}, args{BattleVector{7, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.getDistanceAway(tt.args.position); got != tt.want {
				t.Errorf("battleToken.getDistanceAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battleWeaponSlot_isInRangePosition(t *testing.T) {
	type args struct {
		position BattleVector
	}
	tests := []struct {
		name   string
		weapon battleWeaponSlot
		args   args
		want   bool
	}{
		{"no distance, in range", battleWeaponSlot{token: &battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{0, 0}}}}, args{BattleVector{0, 0}}, true},
		{"distance 1, in range", battleWeaponSlot{token: &battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{0, 0}}}, weaponRange: 1}, args{BattleVector{1, 1}}, true},
		{"distance 2, out of range", battleWeaponSlot{token: &battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{0, 0}}}, weaponRange: 1}, args{BattleVector{1, 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.weapon.isInRangePosition(tt.args.position); got != tt.want {
				t.Errorf("battleWeaponSlot.isInRangePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battle_getBestMove(t *testing.T) {
	type args struct {
		token *battleToken
	}
	tests := []struct {
		name string
		args args
		want BattleVector
	}{
		{
			name: "token at 0,0 move towards target at 1,0",
			args: args{
				token: &battleToken{
					BattleRecordToken: BattleRecordToken{Position: BattleVector{0, 0}},
					moveTarget:        &battleToken{BattleRecordToken: BattleRecordToken{Position: BattleVector{1, 0}}}},
			},
			want: BattleVector{1, 0},
		},
		{
			name: "token at 3,3 move away from targetedBy at 3,3",
			args: args{
				token: &battleToken{
					BattleRecordToken: BattleRecordToken{Position: BattleVector{3, 3}, Tactic: BattleTacticDisengage},
					targetedBy:        []*battleToken{{BattleRecordToken: BattleRecordToken{Position: BattleVector{3, 3}}}}},
			},
			want: BattleVector{4, 4},
		},
		// this depends on a random number. TODO: mock the random
		// {
		// 	name: "token at 8,5 move away from targetedBy at 1,4",
		// 	args: args{
		// 		token: &battleToken{
		// 			BattleRecordToken: BattleRecordToken{Position: BattleVector{8, 5}, Tactic: BattleTacticDisengage},
		// 			targetedBy:        []*battleToken{{BattleRecordToken: BattleRecordToken{Position: BattleVector{1, 4}}}}},
		// 	},
		// 	want: BattleVector{7, 5},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &battle{
				tokens: []*battleToken{tt.args.token},
				rules:  &rules,
			}
			if got := b.getBestMove(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("battle.getBestMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battle_fireBeamWeapon(t *testing.T) {

	type weapon struct {
		weaponSlot   *battleWeaponSlot
		shipQuantity int
		position     BattleVector
	}
	type args struct {
		weapon  weapon
		targets []*battleToken
	}
	type want struct {
		damage            float64
		quantityDamaged   int
		quantityRemaining int
		stackShields      int
	}
	tests := []struct {
		name string
		args args
		want []want
	}{
		{name: "Single weapon, do 10 damage, no kills",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 beam weapon
						power:        10,
						weaponRange:  1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"}, // for logging
						},
						armor: 20,
					},
				},
			},
			want: []want{{damage: 10, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "Single weapon, do 30 damage, to a ship stack with two ships, one damaged",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 beam weapon
						power:        30,
						weaponRange:  1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity:        2,
							Damage:          5,
							QuantityDamaged: 1,
							design:          &ShipDesign{Name: "defender"}, // for logging
						},
						armor: 20,
					},
				},
			},
			want: []want{{damage: 15, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "Single weapon, do 10 damage reduced to 9 for range",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 beam weapon
						power:        10,
						weaponRange:  2,
					},
					shipQuantity: 1,
					position:     BattleVector{2, 0}, // 1 away from target
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"}, // for logging
						},
						BattleRecordToken: BattleRecordToken{
							Position: BattleVector{0, 0},
						},
						armor: 20,
					},
				},
			},
			want: []want{{damage: 9, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "two weapons, do 30 damage total, one (over)kill",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2,  // 2 beam weapons
						power:        15, // 15 damage per beam
						weaponRange:  2,
					},
					shipQuantity: 1, // one ship in the attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 20, // 20 armor, will be destroyed
					},
				},
			},
			want: []want{{damage: 0, quantityDamaged: 0, quantityRemaining: 0}},
		},
		{name: "two weapons, two ships, do 40 damage total, one kill, one damaged",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2, // 2 beam weapons
						power:        10,
					},
					shipQuantity: 2, // 2 ships in attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 2, // two ships in defender
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 30,
					},
				},
			},
			want: []want{{damage: 10, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "two weapons, two stacks, do 20 damage total, kill both",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2, // 2 beam weapons
						power:        10,
					},
					shipQuantity: 1, // 1 ships in attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 10,
					},
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 10,
					},
				},
			},
			// both stacks gone
			want: []want{
				{damage: 0, quantityDamaged: 0, quantityRemaining: 0},
				{damage: 0, quantityDamaged: 0, quantityRemaining: 0},
			},
		},
		{name: "two weapons, two stacks, do 20 damage total, don't get through shield of the first stack",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2, // 2 beam weapons
						power:        10,
					},
					shipQuantity: 1, // 1 ships in attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 3,
							design:   &ShipDesign{Name: "defender"},
						},
						armor:        10,
						shields:      10,
						stackShields: 30,
					},
					{
						ShipToken: &ShipToken{
							Quantity: 3,
							design:   &ShipDesign{Name: "defender"},
						},
						armor:        10,
						shields:      10,
						stackShields: 30,
					},
				},
			},
			// both stacks alive, but first stack with 20 less stackShields
			want: []want{
				{damage: 0, quantityDamaged: 0, quantityRemaining: 3, stackShields: 10},
				{damage: 0, quantityDamaged: 0, quantityRemaining: 3, stackShields: 30},
			},
		},
		{name: "one weapon, do 10 damage to shields, no damage",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1,
						power:        10,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor:        30,
						shields:      20,
						stackShields: 20,
					},
				},
			},
			want: []want{{damage: 0, quantityDamaged: 0, quantityRemaining: 1, stackShields: 10}},
		},
		{name: "one super beam, do 100 damage destroy one stack and damage another",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1,
						power:        100,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender1"},
						},
						armor: 10,
					},
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender2"},
						},
						armor: 100,
					},
				},
			},
			want: []want{
				{damage: 0, quantityDamaged: 0, quantityRemaining: 0},
				{damage: 90, quantityDamaged: 1, quantityRemaining: 1},
			},
		},
		{name: "one minigun, do 10 damage to all targets",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity:   1,
						power:          10,
						hitsAllTargets: true,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 2,
							design:   &ShipDesign{Name: "defender1"},
						},
						armor: 10,
					},
					{
						ShipToken: &ShipToken{
							Quantity: 2,
							design:   &ShipDesign{Name: "defender2"},
						},
						armor: 100,
					},
				},
			},
			want: []want{
				{damage: 0, quantityDamaged: 0, quantityRemaining: 1}, // destroy one token
				{damage: 5, quantityDamaged: 2, quantityRemaining: 2}, // damage both tokens
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &battle{rules: &rules,
				record: newBattleRecord(1, None, Vector{}, []BattleRecordToken{})}
			b.record.recordNewRound()

			// setup this weapon's token based on shipQuantity and position
			tt.args.weapon.weaponSlot.token = &battleToken{
				ShipToken: &ShipToken{
					Quantity: tt.args.weapon.shipQuantity,
					design:   &ShipDesign{Name: "attacker"}, // for logging
				},
				BattleRecordToken: BattleRecordToken{
					Position: tt.args.weapon.position,
				},
			}

			// fire the beam weapon!
			b.fireBeamWeapon(tt.args.weapon.weaponSlot, tt.args.targets)

			for i, target := range tt.args.targets {
				if target.Quantity != tt.want[i].quantityRemaining {
					t.Errorf("battleWeaponSlot.fireBeamWeapon() target: %d quantityRemaining = %v, want %v", i, target.Quantity, tt.want[i].quantityRemaining)
				}
				if target.Damage != tt.want[i].damage {
					t.Errorf("battleWeaponSlot.fireBeamWeapon() target: %d damage = %v, want %v", i, target.Damage, tt.want[i].damage)
				}
				if target.QuantityDamaged != tt.want[i].quantityDamaged {
					t.Errorf("battleWeaponSlot.fireBeamWeapon() target: %d quantityDamaged = %v, want %v", i, target.QuantityDamaged, tt.want[i].quantityDamaged)
				}
				if target.stackShields != tt.want[i].stackShields {
					t.Errorf("battleWeaponSlot.fireBeamWeapon() target: %d stackShields = %v, want %v", i, target.stackShields, tt.want[i].stackShields)
				}
			}

		})
	}
}

func Test_battle_fireTorpedo(t *testing.T) {

	type weapon struct {
		weaponSlot   *battleWeaponSlot
		shipQuantity int
		position     BattleVector
	}
	type args struct {
		weapon  weapon
		targets []*battleToken
	}
	type want struct {
		damage            float64
		quantityDamaged   int
		quantityRemaining int
		stackShields      int
	}
	tests := []struct {
		name string
		args args
		want []want
	}{
		{name: "Single torpedo, do 10 damage, no kills",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 torpedo
						power:        10,
						accuracy:     1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"}, // for logging
						},
						armor: 20,
					},
				},
			},
			want: []want{{damage: 10, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "Single torpedo, do 10 damage to a 2 ship stack with 1@5 damage",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 torpedo
						power:        10,
						accuracy:     1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity:        2,
							QuantityDamaged: 1,
							Damage:          5,
							design:          &ShipDesign{Name: "defender"}, // for logging
						},
						armor: 20,
					},
				},
			},
			// TODO: not sure about this. It doesn't make sense for a torpedo to splash damage at the end...
			want: []want{{damage: 15 / 2., quantityDamaged: 2, quantityRemaining: 2}},
		},
		{name: "Single torpedo, do 30 damage to a stack with two ships, destroy one, other undamaged",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1, // 1 torpedo
						power:        30,
						accuracy:     1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 2,
							design:   &ShipDesign{Name: "defender"}, // for logging
						},
						armor: 20,
					},
				},
			},
			want: []want{{damage: 0, quantityDamaged: 0, quantityRemaining: 1}},
		},
		{name: "two torpedos, do 15 damage each, kill ship with first hit",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2,  // 2 torpedos
						power:        15, // 15 damage per torpedo
						accuracy:     1,
					},
					shipQuantity: 1, // one ship in the attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 10, // 10 armor, will be destroyed
					},
				},
			},
			want: []want{{damage: 0, quantityDamaged: 0, quantityRemaining: 0}},
		},
		{name: "two capital missiles, do 10 damage each, take down shields with first hit, double damage with second",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity:       2,  // 2 torpedos
						power:              10, // 15 damage per torpedo
						accuracy:           1,
						capitalShipMissile: true,
					},
					shipQuantity: 1, // one ship in the attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						stackShields: 5,  // 5 shields will be gone first hit
						armor:        35, // 10 armor gone first hit, then 10x2=20 damage on second hit
					},
				},
			},
			want: []want{{damage: 30, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "two torpedos, two attacker ships, 4x torpedos do 40 damage total, one kill, one damaged",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2, // 2 torpedos
						power:        10,
						accuracy:     1,
					},
					shipQuantity: 2, // 2 ships in attacker stack
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 2, // two ships in defender
							design:   &ShipDesign{Name: "defender"},
						},
						armor: 30,
					},
				},
			},
			want: []want{{damage: 10, quantityDamaged: 1, quantityRemaining: 1}},
		},
		{name: "from testbed, two omega torps w 300 power, 2 1700dp1300 damage",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 2,   // 2 torpedos
						power:        300, // 600 damage total
						accuracy:     1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity:        3, // three defenders, 3@5 damaged
							QuantityDamaged: 3,
							Damage:          1300, // 400dp left
							design:          &ShipDesign{Name: "defender"},
						},
						armor: 1700,
					},
				},
			},
			// 600 damage total, first ship takes 400, 200 split between remaining ships
			want: []want{{damage: 1400, quantityDamaged: 2, quantityRemaining: 2}},
		},
		{name: "one torpedo, do 5 damage to shields, 5 damage to hull",
			args: args{
				weapon: weapon{
					weaponSlot: &battleWeaponSlot{
						slotQuantity: 1,
						power:        10,
						accuracy:     1,
					},
					shipQuantity: 1,
				},
				targets: []*battleToken{
					{
						ShipToken: &ShipToken{
							Quantity: 1,
							design:   &ShipDesign{Name: "defender"},
						},
						armor:        20,
						shields:      20,
						stackShields: 20,
					},
				},
			},
			want: []want{{damage: 5, quantityDamaged: 1, quantityRemaining: 1, stackShields: 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &battle{rules: &rules,
				record: newBattleRecord(1, None, Vector{}, []BattleRecordToken{})}
			b.record.recordNewRound()

			// setup this weapon's token based on shipQuantity and position
			tt.args.weapon.weaponSlot.token = &battleToken{
				ShipToken: &ShipToken{
					Quantity: tt.args.weapon.shipQuantity,
					design:   &ShipDesign{Name: "attacker"}, // for logging
				},
				BattleRecordToken: BattleRecordToken{
					Position: tt.args.weapon.position,
				},
			}

			// fire the beam weapon!
			b.fireTorpedo(tt.args.weapon.weaponSlot, tt.args.targets)

			for i, target := range tt.args.targets {
				if target.Quantity != tt.want[i].quantityRemaining {
					t.Errorf("battleWeaponSlot.fireTorpedo() target: %d quantityRemaining = %v, want %v", i, target.Quantity, tt.want[i].quantityRemaining)
				}
				if target.Damage != tt.want[i].damage {
					t.Errorf("battleWeaponSlot.fireTorpedo() target: %d damage = %v, want %v", i, target.Damage, tt.want[i].damage)
				}
				if target.QuantityDamaged != tt.want[i].quantityDamaged {
					t.Errorf("battleWeaponSlot.fireTorpedo() target: %d quantityDamaged = %v, want %v", i, target.QuantityDamaged, tt.want[i].quantityDamaged)
				}
				if target.stackShields != tt.want[i].stackShields {
					t.Errorf("battleWeaponSlot.fireTorpedo() target: %d stackShields = %v, want %v", i, target.stackShields, tt.want[i].stackShields)
				}
			}

		})
	}
}

func Test_battle_runBattle1(t *testing.T) {
	player1 := testPlayer().WithNum(1)
	player2 := testPlayer().WithNum(2)
	player1.Relations = []PlayerRelationship{{Relation: PlayerRelationFriend}, {Relation: PlayerRelationEnemy}}
	player2.Relations = []PlayerRelationship{{Relation: PlayerRelationEnemy}, {Relation: PlayerRelationFriend}}

	fleets := []*Fleet{
		testStalwartDefender(player1),
		testLongRangeScout(player1),
		testTeamster(player2),
	}

	designNum := 1
	for _, fleet := range fleets {
		for _, token := range fleet.Tokens {
			token.design.Num = designNum
			designNum += 1
		}
	}

	battle := newBattler(&rules, &StaticTechStore, 1, map[int]*Player{1: player1, 2: player2}, fleets, nil)

	record := battle.runBattle()

	// ran some number of turns
	assert.Greater(t, len(record.ActionsPerRound), 1)
	assert.Equal(t, 2, record.Stats.NumShipsByPlayer[player1.Num])
	assert.Equal(t, 1, record.Stats.NumShipsByPlayer[player2.Num])
}

func Test_battle_runBattle2(t *testing.T) {
	player1 := NewPlayer(0, NewRace()).WithNum(1)
	player2 := NewPlayer(0, NewRace()).WithNum(2)
	player1.Name = AINames[0] + "s"
	player2.Name = AINames[1] + "s"
	player1.Race.PluralName = AINames[0] + "s"
	player2.Race.PluralName = AINames[1] + "s"
	player1.Relations = []PlayerRelationship{{Relation: PlayerRelationFriend}, {Relation: PlayerRelationEnemy}}
	player2.Relations = []PlayerRelationship{{Relation: PlayerRelationEnemy}, {Relation: PlayerRelationFriend}}
	player1.PlayerIntels.PlayerIntels = []PlayerIntel{{Num: player1.Num}, {Num: player2.Num}}
	player2.PlayerIntels.PlayerIntels = []PlayerIntel{{Num: player1.Num}, {Num: player2.Num}}

	player1.Designs = append(player1.Designs,
		NewShipDesign(player1, 1).
			WithName("Battle Cruiser").
			WithHull(BattleCruiser.Name).
			WithSlots([]ShipDesignSlot{
				{HullComponent: TransStar10.Name, HullSlotIndex: 1, Quantity: 2},
				{HullComponent: Overthruster.Name, HullSlotIndex: 2, Quantity: 2},
				{HullComponent: BattleSuperComputer.Name, HullSlotIndex: 3, Quantity: 2},
				{HullComponent: ColloidalPhaser.Name, HullSlotIndex: 4, Quantity: 3},
				{HullComponent: DeltaTorpedo.Name, HullSlotIndex: 5, Quantity: 3},
				{HullComponent: Overthruster.Name, HullSlotIndex: 6, Quantity: 3},
				{HullComponent: GorillaDelagator.Name, HullSlotIndex: 7, Quantity: 4},
			}),
	)

	player2.Designs = append(player2.Designs,
		NewShipDesign(player2, 1).
			WithName("Teamster").
			WithHull(SmallFreighter.Name).
			WithSlots([]ShipDesignSlot{
				{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
				{HullComponent: Crobmnium.Name, HullSlotIndex: 2, Quantity: 1},
				{HullComponent: RhinoScanner.Name, HullSlotIndex: 3, Quantity: 1},
			}),
		NewShipDesign(player2, 2).
			WithName("Long Range Scout").
			WithHull(Scout.Name).
			WithSlots([]ShipDesignSlot{
				{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
				{HullComponent: RhinoScanner.Name, HullSlotIndex: 2, Quantity: 1},
				{HullComponent: CompletePhaseShield.Name, HullSlotIndex: 3, Quantity: 1},
			}),
		NewShipDesign(player2, 3).
			WithName("Jammed&Fluxed Defender").
			WithHull(Destroyer.Name).
			WithSlots([]ShipDesignSlot{
				{HullComponent: TransStar10.Name, HullSlotIndex: 1, Quantity: 1},
				{HullComponent: ColloidalPhaser.Name, HullSlotIndex: 2, Quantity: 1},
				{HullComponent: ColloidalPhaser.Name, HullSlotIndex: 3, Quantity: 1},
				{HullComponent: RhinoScanner.Name, HullSlotIndex: 4, Quantity: 1},
				{HullComponent: Superlatanium.Name, HullSlotIndex: 5, Quantity: 1},
				{HullComponent: Jammer30.Name, HullSlotIndex: 6, Quantity: 1},
				{HullComponent: FluxCapacitor.Name, HullSlotIndex: 7, Quantity: 1},
			}),
		NewShipDesign(player2, 4).
			WithName("Stalwart Sapper").
			WithHull(Destroyer.Name).
			WithSlots([]ShipDesignSlot{
				{HullComponent: LongHump6.Name, HullSlotIndex: 1, Quantity: 1},
				{HullComponent: PulsedSapper.Name, HullSlotIndex: 2, Quantity: 1},
				{HullComponent: PulsedSapper.Name, HullSlotIndex: 3, Quantity: 1},
				{HullComponent: RhinoScanner.Name, HullSlotIndex: 4, Quantity: 1},
				{HullComponent: Superlatanium.Name, HullSlotIndex: 5, Quantity: 1},
				{HullComponent: Overthruster.Name, HullSlotIndex: 6, Quantity: 1},
				{HullComponent: Overthruster.Name, HullSlotIndex: 7, Quantity: 1},
			}),
	)

	fleets := []*Fleet{
		{
			MapObject: MapObject{
				PlayerNum: player1.Num,
			},
			BaseName: "Battle Cruiser",
			Tokens: []ShipToken{
				{
					DesignNum: player1.Designs[0].Num,
					Quantity:  2,
				},
			},
		},
		// player2's teamster
		{
			MapObject: MapObject{
				PlayerNum: player2.Num,
			},
			BaseName: "Teamster+",
			Tokens: []ShipToken{
				{
					Quantity:  5,
					DesignNum: player2.Designs[0].Num,
				},
				{
					Quantity:  2,
					DesignNum: player2.Designs[1].Num,
				},
				{
					Quantity:  3,
					DesignNum: player2.Designs[2].Num,
				},
				{
					Quantity:  4,
					DesignNum: player2.Designs[3].Num,
				},
			},
		}}

	record := RunTestBattle([]*Player{player1, player2}, fleets)
	// ran some number of turns
	assert.Less(t, 5, len(record.ActionsPerRound))
}

func Test_updateBestPositions(t *testing.T) {
	type args struct {
		better      bool
		newPosition BattleVector
		bestMoves   []BattleVector
	}
	tests := []struct {
		name string
		args args
		want []BattleVector
	}{
		{
			name: "better move 1,0, pick it",
			args: args{better: true, newPosition: BattleVector{1, 0}, bestMoves: []BattleVector{{0, 0}}},
			want: []BattleVector{{1, 0}},
		},
		{
			name: "better move away from center, pick it",
			args: args{better: true, newPosition: BattleVector{3, 3}, bestMoves: []BattleVector{{4, 4}, {4, 5}}},
			want: []BattleVector{{3, 3}},
		},
		{
			name: "equivalent damage move, but newPosition is closer to center",
			args: args{better: false, newPosition: BattleVector{4, 4}, bestMoves: []BattleVector{{4, 3}}},
			want: []BattleVector{{4, 4}},
		},
		{
			name: "equivalent damage move, newPosition is closer to center",
			args: args{better: false, newPosition: BattleVector{4, 5}, bestMoves: []BattleVector{{4, 3}}},
			want: []BattleVector{{4, 5}},
		},
		{
			name: "equivalent damage move, newPosition is same distance to center",
			args: args{better: false, newPosition: BattleVector{4, 5}, bestMoves: []BattleVector{{4, 4}}},
			want: []BattleVector{{4, 4}, {4, 5}},
		},
		{
			name: "equivalent damage move, newPosition is same distance to center",
			args: args{better: false, newPosition: BattleVector{5, 5}, bestMoves: []BattleVector{{4, 4}, {4, 5}}},
			want: []BattleVector{{4, 4}, {4, 5}, {5, 5}},
		},
		{
			name: "equivalent damage move, newPosition is farther from center, discard it",
			args: args{better: false, newPosition: BattleVector{6, 5}, bestMoves: []BattleVector{{4, 4}, {4, 5}}},
			want: []BattleVector{{4, 4}, {4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBestMoves(tt.args.better, tt.args.newPosition, tt.args.bestMoves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBestPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBattleMovement(t *testing.T) {
	type args struct {
		idealEngineSpeed int
		mass             int
		numEngines       int
		movementBonus    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Destroyer + Trans Galactic Drive + thruster", args{idealEngineSpeed: 9, mass: 244, numEngines: 1, movementBonus: 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBattleMovement(tt.args.idealEngineSpeed, tt.args.movementBonus, tt.args.mass, tt.args.numEngines); got != tt.want {
				t.Errorf("getBattleMovement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battle_buildMovementOrder(t *testing.T) {
	tokenMovement2Mass1 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 1,
			Num:       1,
			Movement:  2,
			Mass:      1,
		},
	}
	tokenMovement4Mass1 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 2,
			Num:       2,
			Movement:  4,
			Mass:      1,
		},
	}
	tokenMovement4Mass2 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 3,
			Num:       3,
			Movement:  4,
			Mass:      2,
		},
	}
	tokenMovement5Mass2 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 4,
			Num:       4,
			Movement:  5,
			Mass:      2,
		},
	}

	tokenMovement6Mass190 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 5,
			Num:       5,
			Movement:  6,
			Mass:      190,
		},
	}
	tokenMovement4Mass22 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 6,
			Num:       6,
			Movement:  4,
			Mass:      22,
		},
	}
	tokenMovement4Mass41 := &battleToken{
		ShipToken: &ShipToken{},
		BattleRecordToken: BattleRecordToken{
			PlayerNum: 6,
			Num:       7,
			Movement:  4,
			Mass:      41,
		},
	}

	type args struct {
		tokens []*battleToken
	}
	tests := []struct {
		name          string
		args          args
		wantMoveOrder [4][]*battleToken
	}{
		{name: "one token, move 2", args: args{[]*battleToken{tokenMovement2Mass1}},
			wantMoveOrder: [4][]*battleToken{
				{tokenMovement2Mass1},
				nil,
				{tokenMovement2Mass1},
				nil,
			},
		},
		{name: "two tokens, same mass, different movements", args: args{[]*battleToken{tokenMovement2Mass1, tokenMovement4Mass1}},
			wantMoveOrder: [4][]*battleToken{
				{tokenMovement2Mass1, tokenMovement4Mass1},
				{tokenMovement4Mass1},
				{tokenMovement2Mass1, tokenMovement4Mass1},
				{tokenMovement4Mass1},
			},
		},
		// higher mass moves first
		{name: "two tokens, diff mass, same movement", args: args{[]*battleToken{tokenMovement4Mass1, tokenMovement4Mass2}},
			wantMoveOrder: [4][]*battleToken{
				{tokenMovement4Mass2, tokenMovement4Mass1},
				{tokenMovement4Mass2, tokenMovement4Mass1},
				{tokenMovement4Mass2, tokenMovement4Mass1},
				{tokenMovement4Mass2, tokenMovement4Mass1},
			},
		},
		// higher move/mass moves twice
		{name: "two tokens, one higher move/mass", args: args{[]*battleToken{tokenMovement4Mass1, tokenMovement5Mass2}},
			wantMoveOrder: [4][]*battleToken{
				{tokenMovement5Mass2, tokenMovement5Mass2, tokenMovement4Mass1},
				{tokenMovement5Mass2, tokenMovement4Mass1},
				{tokenMovement5Mass2, tokenMovement4Mass1},
				{tokenMovement5Mass2, tokenMovement4Mass1},
			},
		},
		// higher mass should move twice, then other tokens move
		{name: "three tokens", args: args{[]*battleToken{tokenMovement6Mass190, tokenMovement4Mass22, tokenMovement4Mass41}},
			wantMoveOrder: [4][]*battleToken{
				{tokenMovement6Mass190, tokenMovement6Mass190, tokenMovement4Mass41, tokenMovement4Mass22},
				{tokenMovement6Mass190, tokenMovement4Mass41, tokenMovement4Mass22},
				{tokenMovement6Mass190, tokenMovement6Mass190, tokenMovement4Mass41, tokenMovement4Mass22},
				{tokenMovement6Mass190, tokenMovement4Mass41, tokenMovement4Mass22},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &battle{rules: &rules}
			if gotMoveOrder := b.buildMovementOrder(tt.args.tokens); !reflect.DeepEqual(gotMoveOrder, tt.wantMoveOrder) {
				for _, r := range gotMoveOrder {
					for _, t := range r {
						s := t.String()
						_ = s
					}
				}
				t.Errorf("battle.buildMovementOrder() = %v, want %v", gotMoveOrder, tt.wantMoveOrder)
			}
		})
	}
}

func Test_battleWeaponSlot_getAttractiveness(t *testing.T) {
	type fields struct {
		weaponType         battleWeaponType
		accuracy           float64
		damagesShieldsOnly bool
		capitalShipMissile bool
	}
	type args struct {
		cost           Cost
		armor          int
		shields        int
		beamDefense    float64
		torpedoJamming float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name:   "beam, attractiveness 1",
			fields: fields{weaponType: battleWeaponTypeBeam},
			args: args{
				cost:    Cost{Boranium: 1, Resources: 1},
				armor:   1,
				shields: 1,
			},
			want: 1,
		},
		{
			name:   "torpedo, more shields than armor",
			fields: fields{weaponType: battleWeaponTypeTorpedo, accuracy: .45},
			args: args{
				cost:    Cost{Boranium: 1, Resources: 1},
				armor:   1,
				shields: 2,
			},
			want: .45,
		},
		{
			name:   "torpedo, more armor than shields",
			fields: fields{weaponType: battleWeaponTypeTorpedo, accuracy: .45},
			args: args{
				cost:    Cost{Boranium: 1, Resources: 1},
				armor:   2,
				shields: 1,
			},
			want: .3,
		},
		{
			name:   "torpedo, attractiveness 1",
			fields: fields{weaponType: battleWeaponTypeTorpedo, accuracy: .45},
			args: args{
				cost:    Cost{Boranium: 1, Resources: 1},
				armor:   1,
				shields: 1,
			},
			want: .45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weapon := &battleWeaponSlot{
				weaponType:         tt.fields.weaponType,
				damagesShieldsOnly: tt.fields.damagesShieldsOnly,
				accuracy:           tt.fields.accuracy,
				capitalShipMissile: tt.fields.capitalShipMissile,
			}
			target := &battleToken{
				// we only care about the design cost on the target shiptoken
				ShipToken: &ShipToken{
					design: &ShipDesign{
						Spec: ShipDesignSpec{
							Cost: tt.args.cost,
						},
					},
				},
				armor:          tt.args.armor,
				shields:        tt.args.shields,
				beamDefense:    tt.args.beamDefense,
				torpedoJamming: tt.args.torpedoJamming,
			}
			if got := weapon.getAttractiveness(target); !test.WithinTolerance(got, tt.want, .01) {
				t.Errorf("battleWeaponSlot.getAttractiveness() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battleWeaponSlot_getAccuracy(t *testing.T) {
	type fields struct {
		accuracy     float64
		torpedoBonus float64
	}
	type args struct {
		torpedoJamming float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"beta torpedo", fields{accuracy: .45}, args{}, .45},
		{"beta torpedo, 1 BC", fields{accuracy: .45, torpedoBonus: .2}, args{}, .56},
		{"beta torpedo, 1 BC, 1 jammer 20", fields{accuracy: .45, torpedoBonus: .2}, args{torpedoJamming: .2}, .45},
		{"beta torpedo, 1 BC, 1 jammer 10", fields{accuracy: .45, torpedoBonus: .2}, args{torpedoJamming: .1}, .505},
		{"beta torpedo, 1 BC, 1 jammer 30", fields{accuracy: .45, torpedoBonus: .1}, args{torpedoJamming: .2}, .405},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weapon := &battleWeaponSlot{
				accuracy:     tt.fields.accuracy,
				torpedoBonus: tt.fields.torpedoBonus,
			}
			if got := weapon.getAccuracy(tt.args.torpedoJamming); got != tt.want {
				t.Errorf("battleWeaponSlot.getAccuracy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battleWeaponSlot_getDamage(t *testing.T) {
	type fields struct {
		weaponType  battleWeaponType
		weaponRange int
		power       int
	}
	type args struct {
		dist        int
		beamDefense float64
		beamDropoff float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"torpedo, base damage", fields{weaponType: battleWeaponTypeTorpedo, power: 10}, args{}, 10},
		{
			name:   "laser, 0 range",
			fields: fields{weaponType: battleWeaponTypeBeam, weaponRange: 1, power: 10},
			args:   args{dist: 0, beamDropoff: .1},
			want:   10,
		},
		{
			name:   "laser, 1 range",
			fields: fields{weaponType: battleWeaponTypeBeam, weaponRange: 1, power: 10},
			args:   args{dist: 1, beamDropoff: .1},
			want:   9,
		},
		{
			name:   "ColloidalPhaser, 3 range",
			fields: fields{weaponType: battleWeaponTypeBeam, weaponRange: 3, power: 26},
			args:   args{dist: 3, beamDropoff: .1},
			want:   24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weapon := &battleWeaponSlot{
				weaponType:  tt.fields.weaponType,
				weaponRange: tt.fields.weaponRange,
				power:       tt.fields.power,
			}
			if got := weapon.getDamage(tt.args.dist, tt.args.beamDefense, tt.args.beamDropoff); got != tt.want {
				t.Errorf("battleWeaponSlot.getDamage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_battleWeaponSlot_getTargetBeamDamage(t *testing.T) {
	type fields struct {
		position           BattleVector
		shipQuantity       int
		slotQuantity       int
		weaponRange        int
		damagesShieldsOnly bool
	}
	type args struct {
		damage               int
		position             BattleVector
		armor                int
		shields              int
		beamDefense          float64
		tokenQuantity        int
		tokenDamage          float64
		tokenQuantityDamaged int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   battleWeaponDamage
	}{
		{
			name:   "1 laser, 20dp target, 10 damage done",
			fields: fields{shipQuantity: 1, slotQuantity: 1, weaponRange: 1},
			args: args{
				damage:        10,
				tokenQuantity: 1,
				armor:         20,
				shields:       0,
			},
			want: battleWeaponDamage{armorDamage: 10, damage: 10, quantityDamaged: 1},
		},
		{
			name:   "1 laser, 1 range away, 20dp target, 9 damage done",
			fields: fields{shipQuantity: 1, slotQuantity: 1, weaponRange: 1},
			args: args{
				position:      BattleVector{1, 0},
				damage:        10,
				tokenQuantity: 1,
				armor:         20,
				shields:       0,
			},
			want: battleWeaponDamage{armorDamage: 9, damage: 9, quantityDamaged: 1},
		},
		{
			name:   "1 laser, 1 range away, 1 defelctor, 20dp target, 8 damage done",
			fields: fields{shipQuantity: 1, slotQuantity: 1, weaponRange: 1},
			args: args{
				position:      BattleVector{1, 0},
				damage:        10,
				tokenQuantity: 1,
				armor:         20,
				shields:       0,
				beamDefense:   .1,
			},
			want: battleWeaponDamage{armorDamage: 8, damage: 8, quantityDamaged: 1},
		},
		{
			name:   "1 laser, 20 shields 20dp target, 10 damage done",
			fields: fields{shipQuantity: 1, slotQuantity: 1, weaponRange: 1},
			args: args{
				damage:        10,
				tokenQuantity: 1,
				armor:         20,
				shields:       20,
			},
			want: battleWeaponDamage{shieldDamage: 10},
		},
		{
			name:   "3 lasers, 20 shields 20dp target, 20 shield, 10 armor damage done",
			fields: fields{shipQuantity: 1, slotQuantity: 3, weaponRange: 1},
			args: args{
				damage:        30, // 3 lasers * 10 damage each
				tokenQuantity: 1,
				armor:         20,
				shields:       20,
			},
			want: battleWeaponDamage{shieldDamage: 20, armorDamage: 10, damage: 10, quantityDamaged: 1},
		},
		{
			name:   "2 ships, 3 lasers, 20 shields 20dp target, destroyed",
			fields: fields{shipQuantity: 2, slotQuantity: 3, weaponRange: 1},
			args: args{
				damage:        60, // 3 lasers * 2 ships * 10 damage each
				tokenQuantity: 1,
				armor:         20,
				shields:       20,
			},
			want: battleWeaponDamage{shieldDamage: 20, armorDamage: 20, numDestroyed: 1, leftoverDamage: 20, damage: 0, quantityDamaged: 0},
		},
		{
			name:   "2 ships, 3 lasers, 2 targets with 20x2 shields 20dp, destroy one",
			fields: fields{shipQuantity: 2, slotQuantity: 3, weaponRange: 1},
			args: args{
				damage:        60, // 3 lasers * 2 ships * 10 damage each
				tokenQuantity: 2,
				armor:         20,
				shields:       40,
			},
			want: battleWeaponDamage{shieldDamage: 40, armorDamage: 20, numDestroyed: 1, leftoverDamage: 0, damage: 0, quantityDamaged: 0},
		},
		{
			name:   "2 ships, 3 lasers, 2 targets 1sq away with 20x2 shields 20dp, damage both",
			fields: fields{shipQuantity: 2, slotQuantity: 3, weaponRange: 1},
			args: args{
				position:      BattleVector{1, 0}, // 1 away
				damage:        60,                 // 3 lasers * 2 ships * 10 damage each
				tokenQuantity: 2,
				armor:         20,
				shields:       40,
			},
			want: battleWeaponDamage{shieldDamage: 40, armorDamage: 14, numDestroyed: 0, leftoverDamage: 0, damage: 7, quantityDamaged: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			weapon := &battleWeaponSlot{
				token: &battleToken{
					BattleRecordToken: BattleRecordToken{Position: tt.fields.position},
					ShipToken: &ShipToken{
						Quantity: tt.fields.shipQuantity,
					},
				},
				slotQuantity:       tt.fields.slotQuantity,
				weaponType:         battleWeaponTypeBeam,
				weaponRange:        tt.fields.weaponRange,
				damagesShieldsOnly: tt.fields.damagesShieldsOnly,
			}

			target := &battleToken{
				BattleRecordToken: BattleRecordToken{Position: tt.args.position},
				ShipToken: &ShipToken{
					Quantity:        tt.args.tokenQuantity,
					Damage:          tt.args.tokenDamage,
					QuantityDamaged: tt.args.tokenQuantityDamaged,
				},
				armor:        tt.args.armor,
				stackShields: tt.args.shields,
				beamDefense:  tt.args.beamDefense,
			}
			if got := weapon.getTargetBeamDamage(tt.args.damage, target, rules.BeamRangeDropoff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("battleWeaponSlot.getTargetBeamDamage() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_getBeamDamageAtRange(t *testing.T) {
	type args struct {
		damage      int
		weaponRange int
		dist        int
		beamDefense float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 laser, 0 range", args{damage: 10, weaponRange: 1, dist: 0}, 10},
		{"1 laser, 1 range", args{damage: 10, weaponRange: 1, dist: 1}, 9},
		{"2 colloidal phasers, 3 range", args{damage: 52, weaponRange: 3, dist: 3}, 47}, // real Stars! is 48...
		{"1 laser, 0 range, 1 deflector", args{damage: 10, weaponRange: 1, dist: 0, beamDefense: .1}, 9},
		{"1 laser, 1 range, 1 deflector", args{damage: 10, weaponRange: 1, dist: 1, beamDefense: .1}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBeamDamageAtRange(tt.args.damage, tt.args.weaponRange, tt.args.dist, tt.args.beamDefense, rules.BeamRangeDropoff); got != tt.want {
				t.Errorf("getBeamDamageAtRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
