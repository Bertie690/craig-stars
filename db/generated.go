// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package db

import (
	cs "github.com/sirgwain/craig-stars/cs"
	"time"
)

type GameConverter struct{}

func (c *GameConverter) ConvertFleet(source *Fleet) *cs.Fleet {
	var pCsFleet *cs.Fleet
	if source != nil {
		var csFleet cs.Fleet
		csFleet.MapObject = ExtendFleetMapObject((*source))
		csFleet.FleetOrders = ExtendFleetFleetOrders((*source))
		csFleet.PlanetNum = (*source).PlanetNum
		csFleet.BaseName = (*source).BaseName
		csFleet.Cargo = c.dbFleetToCsCargo((*source))
		csFleet.Fuel = (*source).Fuel
		csFleet.Age = (*source).Age
		csFleet.Tokens = ShipTokensToGameShipTokens((*source).Tokens)
		csFleet.Heading = ExtendFleetHeading((*source))
		csFleet.WarpSpeed = (*source).WarpSpeed
		csFleet.PreviousPosition = ExtendFleetPreviousPosition((*source))
		csFleet.OrbitingPlanetNum = (*source).OrbitingPlanetNum
		csFleet.Starbase = (*source).Starbase
		csFleet.Spec = FleetSpecToGameFleetSpec((*source).Spec)
		pCsFleet = &csFleet
	}
	return pCsFleet
}
func (c *GameConverter) ConvertGame(source Game) cs.Game {
	var csGame cs.Game
	csGame.DBObject = c.dbGameToCsDBObject(source)
	csGame.HostID = source.HostID
	csGame.Name = source.Name
	csGame.State = cs.GameState(source.State)
	csGame.Public = source.Public
	csGame.Hash = source.Hash
	csGame.Size = cs.Size(source.Size)
	csGame.Density = cs.Density(source.Density)
	csGame.PlayerPositions = cs.PlayerPositions(source.PlayerPositions)
	csGame.RandomEvents = source.RandomEvents
	csGame.ComputerPlayersFormAlliances = source.ComputerPlayersFormAlliances
	csGame.PublicPlayerScores = source.PublicPlayerScores
	csGame.MaxMinerals = source.MaxMinerals
	csGame.AcceleratedPlay = source.AcceleratedPlay
	csGame.StartMode = cs.GameStartMode(source.StartMode)
	csGame.QuickStartTurns = source.QuickStartTurns
	csGame.OpenPlayerSlots = source.OpenPlayerSlots
	csGame.NumPlayers = source.NumPlayers
	csGame.VictoryConditions = ExtendVictoryConditions(source)
	csGame.Seed = source.Seed
	csGame.Rules = ExtendDefaultRules(source)
	csGame.Area = ExtendArea(source)
	csGame.Year = source.Year
	csGame.VictorDeclared = source.VictorDeclared
	csGame.Archived = source.Archived
	return csGame
}
func (c *GameConverter) ConvertGameFleet(source *cs.Fleet) *Fleet {
	var pDbFleet *Fleet
	if source != nil {
		var dbFleet Fleet
		dbFleet.ID = (*source).MapObject.GameDBObject.ID
		dbFleet.GameID = (*source).MapObject.GameDBObject.GameID
		dbFleet.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbFleet.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbFleet.X = (*source).MapObject.Position.X
		dbFleet.Y = (*source).MapObject.Position.Y
		dbFleet.Name = (*source).MapObject.Name
		dbFleet.Num = (*source).MapObject.Num
		dbFleet.PlayerNum = (*source).MapObject.PlayerNum
		dbFleet.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbFleet.Tokens = GameShipTokensToShipTokens((*source).Tokens)
		dbFleet.Waypoints = GameWaypointsToWaypoints((*source).FleetOrders.Waypoints)
		dbFleet.RepeatOrders = (*source).FleetOrders.RepeatOrders
		dbFleet.PlanetNum = (*source).PlanetNum
		dbFleet.BaseName = (*source).BaseName
		dbFleet.Ironium = (*source).Cargo.Ironium
		dbFleet.Boranium = (*source).Cargo.Boranium
		dbFleet.Germanium = (*source).Cargo.Germanium
		dbFleet.Colonists = (*source).Cargo.Colonists
		dbFleet.Fuel = (*source).Fuel
		dbFleet.Age = (*source).Age
		dbFleet.BattlePlanNum = (*source).FleetOrders.BattlePlanNum
		dbFleet.HeadingX = (*source).Heading.X
		dbFleet.HeadingY = (*source).Heading.Y
		dbFleet.WarpSpeed = (*source).WarpSpeed
		var pFloat64 *float64
		if (*source).PreviousPosition != nil {
			pFloat64 = &(*source).PreviousPosition.X
		}
		var pFloat642 *float64
		if pFloat64 != nil {
			xfloat64 := *pFloat64
			pFloat642 = &xfloat64
		}
		dbFleet.PreviousPositionX = pFloat642
		var pFloat643 *float64
		if (*source).PreviousPosition != nil {
			pFloat643 = &(*source).PreviousPosition.Y
		}
		var pFloat644 *float64
		if pFloat643 != nil {
			xfloat642 := *pFloat643
			pFloat644 = &xfloat642
		}
		dbFleet.PreviousPositionY = pFloat644
		dbFleet.OrbitingPlanetNum = (*source).OrbitingPlanetNum
		dbFleet.Starbase = (*source).Starbase
		dbFleet.Purpose = cs.FleetPurpose((*source).FleetOrders.Purpose)
		dbFleet.Spec = GameFleetSpecToFleetSpec((*source).Spec)
		pDbFleet = &dbFleet
	}
	return pDbFleet
}
func (c *GameConverter) ConvertGameGame(source *cs.Game) *Game {
	var pDbGame *Game
	if source != nil {
		var dbGame Game
		dbGame.ID = (*source).DBObject.ID
		dbGame.CreatedAt = TimeToTime((*source).DBObject.CreatedAt)
		dbGame.UpdatedAt = TimeToTime((*source).DBObject.UpdatedAt)
		dbGame.HostID = (*source).HostID
		dbGame.Name = (*source).Name
		dbGame.State = cs.GameState((*source).State)
		dbGame.Public = (*source).Public
		dbGame.Hash = (*source).Hash
		dbGame.Size = cs.Size((*source).Size)
		dbGame.Density = cs.Density((*source).Density)
		dbGame.PlayerPositions = cs.PlayerPositions((*source).PlayerPositions)
		dbGame.RandomEvents = (*source).RandomEvents
		dbGame.ComputerPlayersFormAlliances = (*source).ComputerPlayersFormAlliances
		dbGame.PublicPlayerScores = (*source).PublicPlayerScores
		dbGame.MaxMinerals = (*source).MaxMinerals
		dbGame.AcceleratedPlay = (*source).AcceleratedPlay
		dbGame.StartMode = cs.GameStartMode((*source).StartMode)
		dbGame.QuickStartTurns = (*source).QuickStartTurns
		dbGame.OpenPlayerSlots = (*source).OpenPlayerSlots
		dbGame.NumPlayers = (*source).NumPlayers
		dbGame.VictoryConditionsConditions = cs.Bitmask((*source).VictoryConditions.Conditions)
		dbGame.VictoryConditionsNumCriteriaRequired = (*source).VictoryConditions.NumCriteriaRequired
		dbGame.VictoryConditionsYearsPassed = (*source).VictoryConditions.YearsPassed
		dbGame.VictoryConditionsOwnPlanets = (*source).VictoryConditions.OwnPlanets
		dbGame.VictoryConditionsAttainTechLevel = (*source).VictoryConditions.AttainTechLevel
		dbGame.VictoryConditionsAttainTechLevelNumFields = (*source).VictoryConditions.AttainTechLevelNumFields
		dbGame.VictoryConditionsExceedsScore = (*source).VictoryConditions.ExceedsScore
		dbGame.VictoryConditionsExceedsSecondPlaceScore = (*source).VictoryConditions.ExceedsSecondPlaceScore
		dbGame.VictoryConditionsProductionCapacity = (*source).VictoryConditions.ProductionCapacity
		dbGame.VictoryConditionsOwnCapitalShips = (*source).VictoryConditions.OwnCapitalShips
		dbGame.VictoryConditionsHighestScoreAfterYears = (*source).VictoryConditions.HighestScoreAfterYears
		dbGame.Seed = (*source).Seed
		dbGame.Rules = GameRulesToRules((*source).Rules)
		dbGame.AreaX = (*source).Area.X
		dbGame.AreaY = (*source).Area.Y
		dbGame.Year = (*source).Year
		dbGame.VictorDeclared = (*source).VictorDeclared
		dbGame.Archived = (*source).Archived
		pDbGame = &dbGame
	}
	return pDbGame
}
func (c *GameConverter) ConvertGameMineField(source *cs.MineField) *MineField {
	var pDbMineField *MineField
	if source != nil {
		var dbMineField MineField
		dbMineField.ID = (*source).MapObject.GameDBObject.ID
		dbMineField.GameID = (*source).MapObject.GameDBObject.GameID
		dbMineField.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbMineField.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbMineField.X = (*source).MapObject.Position.X
		dbMineField.Y = (*source).MapObject.Position.Y
		dbMineField.Name = (*source).MapObject.Name
		dbMineField.Num = (*source).MapObject.Num
		dbMineField.PlayerNum = (*source).MapObject.PlayerNum
		dbMineField.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbMineField.MineFieldType = cs.MineFieldType((*source).MineFieldType)
		dbMineField.NumMines = (*source).NumMines
		dbMineField.Detonate = (*source).MineFieldOrders.Detonate
		dbMineField.Spec = GameMineFieldSpecToMineFieldSpec((*source).Spec)
		pDbMineField = &dbMineField
	}
	return pDbMineField
}
func (c *GameConverter) ConvertGameMineralPacket(source *cs.MineralPacket) *MineralPacket {
	var pDbMineralPacket *MineralPacket
	if source != nil {
		var dbMineralPacket MineralPacket
		dbMineralPacket.ID = (*source).MapObject.GameDBObject.ID
		dbMineralPacket.GameID = (*source).MapObject.GameDBObject.GameID
		dbMineralPacket.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbMineralPacket.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbMineralPacket.X = (*source).MapObject.Position.X
		dbMineralPacket.Y = (*source).MapObject.Position.Y
		dbMineralPacket.Name = (*source).MapObject.Name
		dbMineralPacket.Num = (*source).MapObject.Num
		dbMineralPacket.PlayerNum = (*source).MapObject.PlayerNum
		dbMineralPacket.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbMineralPacket.TargetPlanetNum = (*source).TargetPlanetNum
		dbMineralPacket.Ironium = (*source).Cargo.Ironium
		dbMineralPacket.Boranium = (*source).Cargo.Boranium
		dbMineralPacket.Germanium = (*source).Cargo.Germanium
		dbMineralPacket.SafeWarpSpeed = (*source).SafeWarpSpeed
		dbMineralPacket.WarpSpeed = (*source).WarpSpeed
		dbMineralPacket.ScanRange = (*source).ScanRange
		dbMineralPacket.ScanRangePen = (*source).ScanRangePen
		dbMineralPacket.HeadingX = (*source).Heading.X
		dbMineralPacket.HeadingY = (*source).Heading.Y
		pDbMineralPacket = &dbMineralPacket
	}
	return pDbMineralPacket
}
func (c *GameConverter) ConvertGameMysteryTrader(source *cs.MysteryTrader) *MysteryTrader {
	var pDbMysteryTrader *MysteryTrader
	if source != nil {
		var dbMysteryTrader MysteryTrader
		dbMysteryTrader.ID = (*source).MapObject.GameDBObject.ID
		dbMysteryTrader.GameID = (*source).MapObject.GameDBObject.GameID
		dbMysteryTrader.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbMysteryTrader.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbMysteryTrader.X = (*source).MapObject.Position.X
		dbMysteryTrader.Y = (*source).MapObject.Position.Y
		dbMysteryTrader.Name = (*source).MapObject.Name
		dbMysteryTrader.Num = (*source).MapObject.Num
		dbMysteryTrader.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbMysteryTrader.HeadingX = (*source).Heading.X
		dbMysteryTrader.HeadingY = (*source).Heading.Y
		dbMysteryTrader.WarpSpeed = (*source).WarpSpeed
		dbMysteryTrader.RequestedBoon = (*source).RequestedBoon
		dbMysteryTrader.DestinationX = (*source).Destination.X
		dbMysteryTrader.DestinationY = (*source).Destination.Y
		dbMysteryTrader.RewardType = cs.MysteryTraderRewardType((*source).RewardType)
		dbMysteryTrader.PlayersRewarded = GameMysteryTraderPlayersRewardedToMysteryTraderPlayersRewarded((*source).PlayersRewarded)
		dbMysteryTrader.Spec = GameMysteryTraderSpecToMysteryTraderSpec((*source).Spec)
		pDbMysteryTrader = &dbMysteryTrader
	}
	return pDbMysteryTrader
}
func (c *GameConverter) ConvertGamePlanet(source *cs.Planet) *Planet {
	var pDbPlanet *Planet
	if source != nil {
		var dbPlanet Planet
		dbPlanet.ID = (*source).MapObject.GameDBObject.ID
		dbPlanet.GameID = (*source).MapObject.GameDBObject.GameID
		dbPlanet.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbPlanet.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbPlanet.X = (*source).MapObject.Position.X
		dbPlanet.Y = (*source).MapObject.Position.Y
		dbPlanet.Name = (*source).MapObject.Name
		dbPlanet.Num = (*source).MapObject.Num
		dbPlanet.PlayerNum = (*source).MapObject.PlayerNum
		dbPlanet.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbPlanet.Grav = (*source).Hab.Grav
		dbPlanet.Temp = (*source).Hab.Temp
		dbPlanet.Rad = (*source).Hab.Rad
		dbPlanet.BaseGrav = (*source).BaseHab.Grav
		dbPlanet.BaseTemp = (*source).BaseHab.Temp
		dbPlanet.BaseRad = (*source).BaseHab.Rad
		dbPlanet.TerraformedAmountGrav = (*source).TerraformedAmount.Grav
		dbPlanet.TerraformedAmountTemp = (*source).TerraformedAmount.Temp
		dbPlanet.TerraformedAmountRad = (*source).TerraformedAmount.Rad
		dbPlanet.MineralConcIronium = (*source).MineralConcentration.Ironium
		dbPlanet.MineralConcBoranium = (*source).MineralConcentration.Boranium
		dbPlanet.MineralConcGermanium = (*source).MineralConcentration.Germanium
		dbPlanet.MineYearsIronium = (*source).MineYears.Ironium
		dbPlanet.MineYearsBoranium = (*source).MineYears.Boranium
		dbPlanet.MineYearsGermanium = (*source).MineYears.Germanium
		dbPlanet.Ironium = (*source).Cargo.Ironium
		dbPlanet.Boranium = (*source).Cargo.Boranium
		dbPlanet.Germanium = (*source).Cargo.Germanium
		dbPlanet.Colonists = (*source).Cargo.Colonists
		dbPlanet.Mines = (*source).Mines
		dbPlanet.Factories = (*source).Factories
		dbPlanet.Defenses = (*source).Defenses
		dbPlanet.Homeworld = (*source).Homeworld
		dbPlanet.ContributesOnlyLeftoverToResearch = (*source).PlanetOrders.ContributesOnlyLeftoverToResearch
		dbPlanet.Scanner = (*source).Scanner
		dbPlanet.RouteTargetType = cs.MapObjectType((*source).PlanetOrders.RouteTargetType)
		dbPlanet.RouteTargetNum = (*source).PlanetOrders.RouteTargetNum
		dbPlanet.RouteTargetPlayerNum = (*source).PlanetOrders.RouteTargetPlayerNum
		dbPlanet.PacketTargetNum = (*source).PlanetOrders.PacketTargetNum
		dbPlanet.PacketSpeed = (*source).PlanetOrders.PacketSpeed
		dbPlanet.RandomArtifact = (*source).RandomArtifact
		dbPlanet.ProductionQueue = GameProductionQueueItemsToProductionQueueItems((*source).PlanetOrders.ProductionQueue)
		dbPlanet.Spec = GamePlanetSpecToPlanetSpec((*source).Spec)
		pDbPlanet = &dbPlanet
	}
	return pDbPlanet
}
func (c *GameConverter) ConvertGamePlayer(source *cs.Player) *Player {
	var pDbPlayer *Player
	if source != nil {
		var dbPlayer Player
		dbPlayer.ID = (*source).GameDBObject.ID
		dbPlayer.CreatedAt = TimeToTime((*source).GameDBObject.CreatedAt)
		dbPlayer.UpdatedAt = TimeToTime((*source).GameDBObject.UpdatedAt)
		dbPlayer.GameID = (*source).GameDBObject.GameID
		dbPlayer.UserID = (*source).UserID
		dbPlayer.Name = (*source).Name
		dbPlayer.Num = (*source).Num
		dbPlayer.Ready = (*source).Ready
		dbPlayer.AIControlled = (*source).AIControlled
		dbPlayer.AIDifficulty = cs.AIDifficulty((*source).AIDifficulty)
		dbPlayer.Guest = (*source).Guest
		dbPlayer.SubmittedTurn = (*source).SubmittedTurn
		dbPlayer.Color = (*source).Color
		dbPlayer.DefaultHullSet = (*source).DefaultHullSet
		dbPlayer.TechLevelsEnergy = (*source).TechLevels.Energy
		dbPlayer.TechLevelsWeapons = (*source).TechLevels.Weapons
		dbPlayer.TechLevelsPropulsion = (*source).TechLevels.Propulsion
		dbPlayer.TechLevelsConstruction = (*source).TechLevels.Construction
		dbPlayer.TechLevelsElectronics = (*source).TechLevels.Electronics
		dbPlayer.TechLevelsBiotechnology = (*source).TechLevels.Biotechnology
		dbPlayer.TechLevelsSpentEnergy = (*source).TechLevelsSpent.Energy
		dbPlayer.TechLevelsSpentWeapons = (*source).TechLevelsSpent.Weapons
		dbPlayer.TechLevelsSpentPropulsion = (*source).TechLevelsSpent.Propulsion
		dbPlayer.TechLevelsSpentConstruction = (*source).TechLevelsSpent.Construction
		dbPlayer.TechLevelsSpentElectronics = (*source).TechLevelsSpent.Electronics
		dbPlayer.TechLevelsSpentBiotechnology = (*source).TechLevelsSpent.Biotechnology
		dbPlayer.ResearchAmount = (*source).PlayerOrders.ResearchAmount
		dbPlayer.ResearchSpentLastYear = (*source).ResearchSpentLastYear
		dbPlayer.NextResearchField = cs.NextResearchField((*source).PlayerOrders.NextResearchField)
		dbPlayer.Researching = cs.TechField((*source).PlayerOrders.Researching)
		dbPlayer.BattlePlans = GameBattlePlansToBattlePlans((*source).PlayerPlans.BattlePlans)
		dbPlayer.ProductionPlans = GameProductionPlansToProductionPlans((*source).PlayerPlans.ProductionPlans)
		dbPlayer.TransportPlans = GameTransportPlansToTransportPlans((*source).PlayerPlans.TransportPlans)
		dbPlayer.Relations = GamePlayerRelationshipsToPlayerRelationships((*source).Relations)
		dbPlayer.Messages = GamePlayerMessagesToPlayerMessages((*source).Messages)
		dbPlayer.BattleRecords = GameBattleRecordsToBattleRecords((*source).PlayerIntels.BattleRecords)
		dbPlayer.PlayerIntels = GamePlayerIntelsToPlayerIntels((*source).PlayerIntels.PlayerIntels)
		dbPlayer.ScoreIntels = GameScoreIntelsToScoreIntels((*source).PlayerIntels.ScoreIntels)
		dbPlayer.PlanetIntels = GamePlanetIntelsToPlanetIntels((*source).PlayerIntels.PlanetIntels)
		dbPlayer.FleetIntels = GameFleetIntelsToFleetIntels((*source).PlayerIntels.FleetIntels)
		dbPlayer.StarbaseIntels = GameFleetIntelsToFleetIntels((*source).PlayerIntels.StarbaseIntels)
		dbPlayer.ShipDesignIntels = GameShipDesignIntelsToShipDesignIntels((*source).PlayerIntels.ShipDesignIntels)
		dbPlayer.MineralPacketIntels = GameMineralPacketIntelsToMineralPacketIntels((*source).PlayerIntels.MineralPacketIntels)
		dbPlayer.MineFieldIntels = GameMineFieldIntelsToMineFieldIntels((*source).PlayerIntels.MineFieldIntels)
		dbPlayer.WormholeIntels = GameWormholeIntelsToWormholeIntels((*source).PlayerIntels.WormholeIntels)
		dbPlayer.MysteryTraderIntels = GameMysteryTraderIntelsToMysteryTraderIntels((*source).PlayerIntels.MysteryTraderIntels)
		dbPlayer.SalvageIntels = GameSalvageIntelsToSalvageIntels((*source).PlayerIntels.SalvageIntels)
		dbPlayer.Race = GameRaceToPlayerRace((*source).Race)
		dbPlayer.Stats = GamePlayerStatsToPlayerStats((*source).Stats)
		dbPlayer.ScoreHistory = GamePlayerScoresToPlayerScores((*source).ScoreHistory)
		dbPlayer.AcquiredTechs = GameAcquiredTechsToAcquiredTechs((*source).AcquiredTechs)
		dbPlayer.AchievedVictoryConditions = cs.Bitmask((*source).AchievedVictoryConditions)
		dbPlayer.Victor = (*source).Victor
		dbPlayer.Archived = (*source).Archived
		dbPlayer.Spec = GamePlayerSpecToPlayerSpec((*source).Spec)
		pDbPlayer = &dbPlayer
	}
	return pDbPlayer
}
func (c *GameConverter) ConvertGameRace(source *cs.Race) *Race {
	var pDbRace *Race
	if source != nil {
		var dbRace Race
		dbRace.ID = (*source).DBObject.ID
		dbRace.CreatedAt = TimeToTime((*source).DBObject.CreatedAt)
		dbRace.UpdatedAt = TimeToTime((*source).DBObject.UpdatedAt)
		dbRace.UserID = (*source).UserID
		dbRace.Name = (*source).Name
		dbRace.PluralName = (*source).PluralName
		dbRace.SpendLeftoverPointsOn = cs.SpendLeftoverPointsOn((*source).SpendLeftoverPointsOn)
		dbRace.PRT = cs.PRT((*source).PRT)
		dbRace.LRTs = cs.Bitmask((*source).LRTs)
		dbRace.HabLowGrav = (*source).HabLow.Grav
		dbRace.HabLowTemp = (*source).HabLow.Temp
		dbRace.HabLowRad = (*source).HabLow.Rad
		dbRace.HabHighGrav = (*source).HabHigh.Grav
		dbRace.HabHighTemp = (*source).HabHigh.Temp
		dbRace.HabHighRad = (*source).HabHigh.Rad
		dbRace.GrowthRate = (*source).GrowthRate
		dbRace.PopEfficiency = (*source).PopEfficiency
		dbRace.FactoryOutput = (*source).FactoryOutput
		dbRace.FactoryCost = (*source).FactoryCost
		dbRace.NumFactories = (*source).NumFactories
		dbRace.FactoriesCostLess = (*source).FactoriesCostLess
		dbRace.ImmuneGrav = (*source).ImmuneGrav
		dbRace.ImmuneTemp = (*source).ImmuneTemp
		dbRace.ImmuneRad = (*source).ImmuneRad
		dbRace.MineOutput = (*source).MineOutput
		dbRace.MineCost = (*source).MineCost
		dbRace.NumMines = (*source).NumMines
		dbRace.ResearchCostEnergy = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Energy)
		dbRace.ResearchCostWeapons = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Weapons)
		dbRace.ResearchCostPropulsion = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Propulsion)
		dbRace.ResearchCostConstruction = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Construction)
		dbRace.ResearchCostElectronics = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Electronics)
		dbRace.ResearchCostBiotechnology = c.csResearchCostLevelToCsResearchCostLevel((*source).ResearchCost.Biotechnology)
		dbRace.TechsStartHigh = (*source).TechsStartHigh
		dbRace.Spec = GameRaceSpecToRaceSpec((*source).Spec)
		pDbRace = &dbRace
	}
	return pDbRace
}
func (c *GameConverter) ConvertGameSalvage(source *cs.Salvage) *Salvage {
	var pDbSalvage *Salvage
	if source != nil {
		var dbSalvage Salvage
		dbSalvage.ID = (*source).MapObject.GameDBObject.ID
		dbSalvage.GameID = (*source).MapObject.GameDBObject.GameID
		dbSalvage.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbSalvage.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbSalvage.X = (*source).MapObject.Position.X
		dbSalvage.Y = (*source).MapObject.Position.Y
		dbSalvage.Name = (*source).MapObject.Name
		dbSalvage.Num = (*source).MapObject.Num
		dbSalvage.PlayerNum = (*source).MapObject.PlayerNum
		dbSalvage.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbSalvage.Ironium = (*source).Cargo.Ironium
		dbSalvage.Boranium = (*source).Cargo.Boranium
		dbSalvage.Germanium = (*source).Cargo.Germanium
		pDbSalvage = &dbSalvage
	}
	return pDbSalvage
}
func (c *GameConverter) ConvertGameShipDesign(source *cs.ShipDesign) *ShipDesign {
	var pDbShipDesign *ShipDesign
	if source != nil {
		var dbShipDesign ShipDesign
		dbShipDesign.ID = (*source).GameDBObject.ID
		dbShipDesign.GameID = (*source).GameDBObject.GameID
		dbShipDesign.UpdatedAt = TimeToNullTime((*source).GameDBObject.UpdatedAt)
		dbShipDesign.CreatedAt = TimeToNullTime((*source).GameDBObject.CreatedAt)
		dbShipDesign.Num = (*source).Num
		dbShipDesign.PlayerNum = (*source).PlayerNum
		dbShipDesign.OriginalPlayerNum = (*source).OriginalPlayerNum
		dbShipDesign.Name = (*source).Name
		dbShipDesign.Version = (*source).Version
		dbShipDesign.Hull = (*source).Hull
		dbShipDesign.HullSetNumber = (*source).HullSetNumber
		dbShipDesign.CannotDelete = (*source).CannotDelete
		dbShipDesign.Slots = GameShipDesignSlotsToShipDesignSlots((*source).Slots)
		dbShipDesign.Purpose = cs.ShipDesignPurpose((*source).Purpose)
		dbShipDesign.MysteryTrader = BoolToNullBool((*source).MysteryTrader)
		dbShipDesign.Spec = GameShipDesignSpecToShipDesignSpec((*source).Spec)
		pDbShipDesign = &dbShipDesign
	}
	return pDbShipDesign
}
func (c *GameConverter) ConvertGameUser(source *cs.User) *User {
	var pDbUser *User
	if source != nil {
		var dbUser User
		dbUser.ID = (*source).DBObject.ID
		dbUser.CreatedAt = TimeToTime((*source).DBObject.CreatedAt)
		dbUser.UpdatedAt = TimeToTime((*source).DBObject.UpdatedAt)
		dbUser.GameID = (*source).GameID
		dbUser.PlayerNum = (*source).PlayerNum
		dbUser.Username = (*source).Username
		dbUser.Password = (*source).Password
		dbUser.Email = (*source).Email
		dbUser.Role = cs.UserRole((*source).Role)
		dbUser.Banned = (*source).Banned
		dbUser.Verified = (*source).Verified
		dbUser.LastLogin = c.pTimeTimeToPTimeTime((*source).LastLogin)
		var pString *string
		if (*source).DiscordID != nil {
			xstring := *(*source).DiscordID
			pString = &xstring
		}
		dbUser.DiscordID = pString
		var pString2 *string
		if (*source).DiscordAvatar != nil {
			xstring2 := *(*source).DiscordAvatar
			pString2 = &xstring2
		}
		dbUser.DiscordAvatar = pString2
		pDbUser = &dbUser
	}
	return pDbUser
}
func (c *GameConverter) ConvertGameWormhole(source *cs.Wormhole) *Wormhole {
	var pDbWormhole *Wormhole
	if source != nil {
		var dbWormhole Wormhole
		dbWormhole.ID = (*source).MapObject.GameDBObject.ID
		dbWormhole.GameID = (*source).MapObject.GameDBObject.GameID
		dbWormhole.CreatedAt = TimeToTime((*source).MapObject.GameDBObject.CreatedAt)
		dbWormhole.UpdatedAt = TimeToTime((*source).MapObject.GameDBObject.UpdatedAt)
		dbWormhole.X = (*source).MapObject.Position.X
		dbWormhole.Y = (*source).MapObject.Position.Y
		dbWormhole.Name = (*source).MapObject.Name
		dbWormhole.Num = (*source).MapObject.Num
		dbWormhole.Tags = GameTagsToTags((*source).MapObject.Tags)
		dbWormhole.DestinationNum = (*source).DestinationNum
		dbWormhole.Stability = cs.WormholeStability((*source).Stability)
		dbWormhole.YearsAtStability = (*source).YearsAtStability
		dbWormhole.Spec = GameWormholeSpecToWormholeSpec((*source).Spec)
		pDbWormhole = &dbWormhole
	}
	return pDbWormhole
}
func (c *GameConverter) ConvertGames(source []Game) []cs.Game {
	var csGameList []cs.Game
	if source != nil {
		csGameList = make([]cs.Game, len(source))
		for i := 0; i < len(source); i++ {
			csGameList[i] = c.ConvertGame(source[i])
		}
	}
	return csGameList
}
func (c *GameConverter) ConvertMineField(source *MineField) *cs.MineField {
	var pCsMineField *cs.MineField
	if source != nil {
		var csMineField cs.MineField
		csMineField.MapObject = ExtendMineFieldMapObject((*source))
		csMineField.MineFieldOrders = c.dbMineFieldToCsMineFieldOrders((*source))
		csMineField.MineFieldType = cs.MineFieldType((*source).MineFieldType)
		csMineField.NumMines = (*source).NumMines
		csMineField.Spec = MineFieldSpecToGameMineFieldSpec((*source).Spec)
		pCsMineField = &csMineField
	}
	return pCsMineField
}
func (c *GameConverter) ConvertMineralPacket(source *MineralPacket) *cs.MineralPacket {
	var pCsMineralPacket *cs.MineralPacket
	if source != nil {
		var csMineralPacket cs.MineralPacket
		csMineralPacket.MapObject = ExtendMineralPacketMapObject((*source))
		csMineralPacket.TargetPlanetNum = (*source).TargetPlanetNum
		csMineralPacket.Cargo = c.mineralPaketCargo((*source))
		csMineralPacket.WarpSpeed = (*source).WarpSpeed
		csMineralPacket.SafeWarpSpeed = (*source).SafeWarpSpeed
		csMineralPacket.Heading = ExtendMineralPacketHeading((*source))
		csMineralPacket.ScanRange = (*source).ScanRange
		csMineralPacket.ScanRangePen = (*source).ScanRangePen
		pCsMineralPacket = &csMineralPacket
	}
	return pCsMineralPacket
}
func (c *GameConverter) ConvertMysteryTrader(source *MysteryTrader) *cs.MysteryTrader {
	var pCsMysteryTrader *cs.MysteryTrader
	if source != nil {
		var csMysteryTrader cs.MysteryTrader
		csMysteryTrader.MapObject = ExtendMysteryTraderMapObject((*source))
		csMysteryTrader.WarpSpeed = (*source).WarpSpeed
		csMysteryTrader.Destination = ExtendMysteryTraderDestination((*source))
		csMysteryTrader.RequestedBoon = (*source).RequestedBoon
		csMysteryTrader.RewardType = cs.MysteryTraderRewardType((*source).RewardType)
		csMysteryTrader.Heading = ExtendMysteryTraderHeading((*source))
		csMysteryTrader.PlayersRewarded = MysteryTraderPlayersRewardedToGameMysteryTraderPlayersRewarded((*source).PlayersRewarded)
		csMysteryTrader.Spec = MysteryTraderSpecToGameMysteryTraderSpec((*source).Spec)
		pCsMysteryTrader = &csMysteryTrader
	}
	return pCsMysteryTrader
}
func (c *GameConverter) ConvertPlanet(source *Planet) *cs.Planet {
	var pCsPlanet *cs.Planet
	if source != nil {
		var csPlanet cs.Planet
		csPlanet.MapObject = ExtendPlanetMapObject((*source))
		csPlanet.PlanetOrders = c.dbPlanetToCsPlanetOrders((*source))
		csPlanet.Hab = c.dbPlanetToCsHab((*source))
		csPlanet.BaseHab = ExtendBaseHab((*source))
		csPlanet.TerraformedAmount = ExtendTerraformedAmount((*source))
		csPlanet.MineralConcentration = ExtendMineralConcentration((*source))
		csPlanet.MineYears = ExtendMineYears((*source))
		csPlanet.Cargo = c.dbPlanetToCsCargo((*source))
		csPlanet.Mines = (*source).Mines
		csPlanet.Factories = (*source).Factories
		csPlanet.Defenses = (*source).Defenses
		csPlanet.Homeworld = (*source).Homeworld
		csPlanet.Scanner = (*source).Scanner
		csPlanet.Spec = PlanetSpecToGamePlanetSpec((*source).Spec)
		csPlanet.RandomArtifact = (*source).RandomArtifact
		pCsPlanet = &csPlanet
	}
	return pCsPlanet
}
func (c *GameConverter) ConvertPlayer(source Player) cs.Player {
	var csPlayer cs.Player
	csPlayer.GameDBObject = c.dbPlayerToCsGameDBObject(source)
	csPlayer.PlayerOrders = c.dbPlayerToCsPlayerOrders(source)
	csPlayer.PlayerIntels = c.dbPlayerToCsPlayerIntels(source)
	csPlayer.PlayerPlans = c.dbPlayerToCsPlayerPlans(source)
	csPlayer.UserID = source.UserID
	csPlayer.Name = source.Name
	csPlayer.Num = source.Num
	csPlayer.Ready = source.Ready
	csPlayer.AIControlled = source.AIControlled
	csPlayer.AIDifficulty = cs.AIDifficulty(source.AIDifficulty)
	csPlayer.Guest = source.Guest
	csPlayer.SubmittedTurn = source.SubmittedTurn
	csPlayer.Color = source.Color
	csPlayer.DefaultHullSet = source.DefaultHullSet
	csPlayer.Race = PlayerRaceToGameRace(source.Race)
	csPlayer.TechLevels = ExtendTechLevels(source)
	csPlayer.TechLevelsSpent = ExtendTechLevelsSpent(source)
	csPlayer.ResearchSpentLastYear = source.ResearchSpentLastYear
	csPlayer.Relations = PlayerRelationshipsToGamePlayerRelationships(source.Relations)
	csPlayer.Messages = PlayerMessagesToGamePlayerMessages(source.Messages)
	csPlayer.ScoreHistory = PlayerScoresToGamePlayerScores(source.ScoreHistory)
	csPlayer.AcquiredTechs = AcquiredTechsToGameAcquiredTechs(source.AcquiredTechs)
	csPlayer.AchievedVictoryConditions = cs.Bitmask(source.AchievedVictoryConditions)
	csPlayer.Victor = source.Victor
	csPlayer.Archived = source.Archived
	csPlayer.Stats = PlayerStatsToGamePlayerStats(source.Stats)
	csPlayer.Spec = PlayerSpecToGamePlayerSpec(source.Spec)
	return csPlayer
}
func (c *GameConverter) ConvertPlayers(source []Player) []cs.Player {
	var csPlayerList []cs.Player
	if source != nil {
		csPlayerList = make([]cs.Player, len(source))
		for i := 0; i < len(source); i++ {
			csPlayerList[i] = c.ConvertPlayer(source[i])
		}
	}
	return csPlayerList
}
func (c *GameConverter) ConvertRace(source Race) cs.Race {
	var csRace cs.Race
	csRace.DBObject = c.dbRaceToCsDBObject(source)
	csRace.UserID = source.UserID
	csRace.Name = source.Name
	csRace.PluralName = source.PluralName
	csRace.SpendLeftoverPointsOn = cs.SpendLeftoverPointsOn(source.SpendLeftoverPointsOn)
	csRace.PRT = cs.PRT(source.PRT)
	csRace.LRTs = cs.Bitmask(source.LRTs)
	csRace.HabLow = ExtendHabLow(source)
	csRace.HabHigh = ExtendHabHigh(source)
	csRace.GrowthRate = source.GrowthRate
	csRace.PopEfficiency = source.PopEfficiency
	csRace.FactoryOutput = source.FactoryOutput
	csRace.FactoryCost = source.FactoryCost
	csRace.NumFactories = source.NumFactories
	csRace.FactoriesCostLess = source.FactoriesCostLess
	csRace.ImmuneGrav = source.ImmuneGrav
	csRace.ImmuneTemp = source.ImmuneTemp
	csRace.ImmuneRad = source.ImmuneRad
	csRace.MineOutput = source.MineOutput
	csRace.MineCost = source.MineCost
	csRace.NumMines = source.NumMines
	csRace.ResearchCost = ExtendResearchCost(source)
	csRace.TechsStartHigh = source.TechsStartHigh
	csRace.Spec = RaceSpecToGameRaceSpec(source.Spec)
	return csRace
}
func (c *GameConverter) ConvertRaces(source []Race) []cs.Race {
	var csRaceList []cs.Race
	if source != nil {
		csRaceList = make([]cs.Race, len(source))
		for i := 0; i < len(source); i++ {
			csRaceList[i] = c.ConvertRace(source[i])
		}
	}
	return csRaceList
}
func (c *GameConverter) ConvertSalvage(source *Salvage) *cs.Salvage {
	var pCsSalvage *cs.Salvage
	if source != nil {
		var csSalvage cs.Salvage
		csSalvage.MapObject = ExtendSalvageMapObject((*source))
		csSalvage.Cargo = c.salvageCargo((*source))
		pCsSalvage = &csSalvage
	}
	return pCsSalvage
}
func (c *GameConverter) ConvertShipDesign(source *ShipDesign) *cs.ShipDesign {
	var pCsShipDesign *cs.ShipDesign
	if source != nil {
		var csShipDesign cs.ShipDesign
		csShipDesign.GameDBObject = c.dbShipDesignToCsGameDBObject((*source))
		csShipDesign.Num = (*source).Num
		csShipDesign.PlayerNum = (*source).PlayerNum
		csShipDesign.OriginalPlayerNum = (*source).OriginalPlayerNum
		csShipDesign.Name = (*source).Name
		csShipDesign.Version = (*source).Version
		csShipDesign.Hull = (*source).Hull
		csShipDesign.HullSetNumber = (*source).HullSetNumber
		csShipDesign.CannotDelete = (*source).CannotDelete
		csShipDesign.MysteryTrader = NullBoolToBool((*source).MysteryTrader)
		csShipDesign.Slots = ShipDesignSlotsToGameShipDesignSlots((*source).Slots)
		csShipDesign.Purpose = cs.ShipDesignPurpose((*source).Purpose)
		csShipDesign.Spec = ShipDesignSpecToGameShipDesignSpec((*source).Spec)
		pCsShipDesign = &csShipDesign
	}
	return pCsShipDesign
}
func (c *GameConverter) ConvertUser(source User) cs.User {
	var csUser cs.User
	csUser.DBObject = c.dbUserToCsDBObject(source)
	csUser.Username = source.Username
	csUser.Password = source.Password
	csUser.Email = source.Email
	csUser.Role = cs.UserRole(source.Role)
	csUser.Banned = source.Banned
	csUser.Verified = source.Verified
	csUser.GameID = source.GameID
	csUser.PlayerNum = source.PlayerNum
	csUser.LastLogin = c.pTimeTimeToPTimeTime(source.LastLogin)
	var pString *string
	if source.DiscordID != nil {
		xstring := *source.DiscordID
		pString = &xstring
	}
	csUser.DiscordID = pString
	var pString2 *string
	if source.DiscordAvatar != nil {
		xstring2 := *source.DiscordAvatar
		pString2 = &xstring2
	}
	csUser.DiscordAvatar = pString2
	return csUser
}
func (c *GameConverter) ConvertUsers(source []User) []cs.User {
	var csUserList []cs.User
	if source != nil {
		csUserList = make([]cs.User, len(source))
		for i := 0; i < len(source); i++ {
			csUserList[i] = c.ConvertUser(source[i])
		}
	}
	return csUserList
}
func (c *GameConverter) ConvertWormhole(source *Wormhole) *cs.Wormhole {
	var pCsWormhole *cs.Wormhole
	if source != nil {
		var csWormhole cs.Wormhole
		csWormhole.MapObject = c.wormHoleMapObject((*source))
		csWormhole.DestinationNum = (*source).DestinationNum
		csWormhole.Stability = cs.WormholeStability((*source).Stability)
		csWormhole.YearsAtStability = (*source).YearsAtStability
		csWormhole.Spec = WormholeSpecToGameWormholeSpec((*source).Spec)
		pCsWormhole = &csWormhole
	}
	return pCsWormhole
}
func (c *GameConverter) csResearchCostLevelToCsResearchCostLevel(source cs.ResearchCostLevel) cs.ResearchCostLevel {
	return cs.ResearchCostLevel(source)
}
func (c *GameConverter) dbFleetToCsCargo(source Fleet) cs.Cargo {
	var csCargo cs.Cargo
	csCargo.Ironium = source.Ironium
	csCargo.Boranium = source.Boranium
	csCargo.Germanium = source.Germanium
	csCargo.Colonists = source.Colonists
	return csCargo
}
func (c *GameConverter) dbGameToCsDBObject(source Game) cs.DBObject {
	var csDBObject cs.DBObject
	csDBObject.ID = source.ID
	csDBObject.CreatedAt = TimeToTime(source.CreatedAt)
	csDBObject.UpdatedAt = TimeToTime(source.UpdatedAt)
	return csDBObject
}
func (c *GameConverter) dbMineFieldToCsMineFieldOrders(source MineField) cs.MineFieldOrders {
	var csMineFieldOrders cs.MineFieldOrders
	csMineFieldOrders.Detonate = source.Detonate
	return csMineFieldOrders
}
func (c *GameConverter) dbPlanetToCsCargo(source Planet) cs.Cargo {
	var csCargo cs.Cargo
	csCargo.Ironium = source.Ironium
	csCargo.Boranium = source.Boranium
	csCargo.Germanium = source.Germanium
	csCargo.Colonists = source.Colonists
	return csCargo
}
func (c *GameConverter) dbPlanetToCsHab(source Planet) cs.Hab {
	var csHab cs.Hab
	csHab.Grav = source.Grav
	csHab.Temp = source.Temp
	csHab.Rad = source.Rad
	return csHab
}
func (c *GameConverter) dbPlanetToCsPlanetOrders(source Planet) cs.PlanetOrders {
	var csPlanetOrders cs.PlanetOrders
	csPlanetOrders.ContributesOnlyLeftoverToResearch = source.ContributesOnlyLeftoverToResearch
	csPlanetOrders.ProductionQueue = ProductionQueueItemsToGameProductionQueueItems(source.ProductionQueue)
	csPlanetOrders.RouteTargetType = cs.MapObjectType(source.RouteTargetType)
	csPlanetOrders.RouteTargetNum = source.RouteTargetNum
	csPlanetOrders.RouteTargetPlayerNum = source.RouteTargetPlayerNum
	csPlanetOrders.PacketTargetNum = source.PacketTargetNum
	csPlanetOrders.PacketSpeed = source.PacketSpeed
	return csPlanetOrders
}
func (c *GameConverter) dbPlayerToCsGameDBObject(source Player) cs.GameDBObject {
	var csGameDBObject cs.GameDBObject
	csGameDBObject.ID = source.ID
	csGameDBObject.GameID = source.GameID
	csGameDBObject.CreatedAt = TimeToTime(source.CreatedAt)
	csGameDBObject.UpdatedAt = TimeToTime(source.UpdatedAt)
	return csGameDBObject
}
func (c *GameConverter) dbPlayerToCsPlayerIntels(source Player) cs.PlayerIntels {
	var csPlayerIntels cs.PlayerIntels
	csPlayerIntels.BattleRecords = BattleRecordsToGameBattleRecords(source.BattleRecords)
	csPlayerIntels.PlayerIntels = PlayerIntelsToGamePlayerIntels(source.PlayerIntels)
	csPlayerIntels.ScoreIntels = ScoreIntelsToGameScoreIntels(source.ScoreIntels)
	csPlayerIntels.PlanetIntels = PlanetIntelsToGamePlanetIntels(source.PlanetIntels)
	csPlayerIntels.FleetIntels = FleetIntelsToGameFleetIntels(source.FleetIntels)
	csPlayerIntels.StarbaseIntels = FleetIntelsToGameFleetIntels(source.StarbaseIntels)
	csPlayerIntels.ShipDesignIntels = ShipDesignIntelsToGameShipDesignIntels(source.ShipDesignIntels)
	csPlayerIntels.MineralPacketIntels = MineralPacketIntelsToGameMineralPacketIntels(source.MineralPacketIntels)
	csPlayerIntels.MineFieldIntels = MineFieldIntelsToGameMineFieldIntels(source.MineFieldIntels)
	csPlayerIntels.WormholeIntels = WormholeIntelsToGameWormholeIntels(source.WormholeIntels)
	csPlayerIntels.MysteryTraderIntels = MysteryTraderIntelsToGameMysteryTraderIntels(source.MysteryTraderIntels)
	csPlayerIntels.SalvageIntels = SalvageIntelsToGameSalvageIntels(source.SalvageIntels)
	return csPlayerIntels
}
func (c *GameConverter) dbPlayerToCsPlayerOrders(source Player) cs.PlayerOrders {
	var csPlayerOrders cs.PlayerOrders
	csPlayerOrders.Researching = cs.TechField(source.Researching)
	csPlayerOrders.NextResearchField = cs.NextResearchField(source.NextResearchField)
	csPlayerOrders.ResearchAmount = source.ResearchAmount
	return csPlayerOrders
}
func (c *GameConverter) dbPlayerToCsPlayerPlans(source Player) cs.PlayerPlans {
	var csPlayerPlans cs.PlayerPlans
	csPlayerPlans.ProductionPlans = ProductionPlansToGameProductionPlans(source.ProductionPlans)
	csPlayerPlans.BattlePlans = BattlePlansToGameBattlePlans(source.BattlePlans)
	csPlayerPlans.TransportPlans = TransportPlansToGameTransportPlans(source.TransportPlans)
	return csPlayerPlans
}
func (c *GameConverter) dbRaceToCsDBObject(source Race) cs.DBObject {
	var csDBObject cs.DBObject
	csDBObject.ID = source.ID
	csDBObject.CreatedAt = TimeToTime(source.CreatedAt)
	csDBObject.UpdatedAt = TimeToTime(source.UpdatedAt)
	return csDBObject
}
func (c *GameConverter) dbShipDesignToCsGameDBObject(source ShipDesign) cs.GameDBObject {
	var csGameDBObject cs.GameDBObject
	csGameDBObject.ID = source.ID
	csGameDBObject.GameID = source.GameID
	csGameDBObject.CreatedAt = NullTimeToTime(source.CreatedAt)
	csGameDBObject.UpdatedAt = NullTimeToTime(source.UpdatedAt)
	return csGameDBObject
}
func (c *GameConverter) dbUserToCsDBObject(source User) cs.DBObject {
	var csDBObject cs.DBObject
	csDBObject.ID = source.ID
	csDBObject.CreatedAt = TimeToTime(source.CreatedAt)
	csDBObject.UpdatedAt = TimeToTime(source.UpdatedAt)
	return csDBObject
}
func (c *GameConverter) dbWormholeToCsGameDBObject(source Wormhole) cs.GameDBObject {
	var csGameDBObject cs.GameDBObject
	csGameDBObject.ID = source.ID
	csGameDBObject.GameID = source.GameID
	csGameDBObject.CreatedAt = TimeToTime(source.CreatedAt)
	csGameDBObject.UpdatedAt = TimeToTime(source.UpdatedAt)
	return csGameDBObject
}
func (c *GameConverter) dbWormholeToCsVector(source Wormhole) cs.Vector {
	var csVector cs.Vector
	csVector.X = source.X
	csVector.Y = source.Y
	return csVector
}
func (c *GameConverter) mineralPaketCargo(source MineralPacket) cs.Cargo {
	var csCargo cs.Cargo
	csCargo.Ironium = source.Ironium
	csCargo.Boranium = source.Boranium
	csCargo.Germanium = source.Germanium
	return csCargo
}
func (c *GameConverter) pTimeTimeToPTimeTime(source *time.Time) *time.Time {
	var pTimeTime *time.Time
	if source != nil {
		timeTime := TimeToTime((*source))
		pTimeTime = &timeTime
	}
	return pTimeTime
}
func (c *GameConverter) salvageCargo(source Salvage) cs.Cargo {
	var csCargo cs.Cargo
	csCargo.Ironium = source.Ironium
	csCargo.Boranium = source.Boranium
	csCargo.Germanium = source.Germanium
	return csCargo
}
func (c *GameConverter) wormHoleMapObject(source Wormhole) cs.MapObject {
	var csMapObject cs.MapObject
	csMapObject.GameDBObject = c.dbWormholeToCsGameDBObject(source)
	csMapObject.Type = MapObjectTypeWormhole()
	csMapObject.Position = c.dbWormholeToCsVector(source)
	csMapObject.Num = source.Num
	csMapObject.Name = source.Name
	csMapObject.Tags = TagsToGameTags(source.Tags)
	return csMapObject
}
