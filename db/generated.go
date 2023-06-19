// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package db

import game "github.com/sirgwain/craig-stars/game"

type GameConverter struct{}

func (c *GameConverter) ConvertFleet(source *Fleet) *game.Fleet {
	var pGameFleet *game.Fleet
	if source != nil {
		gameFleet := c.dbFleetToGameFleet(*source)
		pGameFleet = &gameFleet
	}
	return pGameFleet
}
func (c *GameConverter) ConvertGame(source Game) game.Game {
	var gameGame game.Game
	gameGame.ID = source.ID
	gameGame.CreatedAt = TimeToTime(source.CreatedAt)
	gameGame.UpdatedAt = TimeToTime(source.UpdatedAt)
	gameGame.Name = source.Name
	gameGame.HostID = source.HostID
	gameGame.QuickStartTurns = source.QuickStartTurns
	gameGame.Size = game.Size(source.Size)
	gameGame.Density = game.Density(source.Density)
	gameGame.PlayerPositions = game.PlayerPositions(source.PlayerPositions)
	gameGame.RandomEvents = source.RandomEvents
	gameGame.ComputerPlayersFormAlliances = source.ComputerPlayersFormAlliances
	gameGame.PublicPlayerScores = source.PublicPlayerScores
	gameGame.StartMode = game.GameStartMode(source.StartMode)
	gameGame.Year = source.Year
	gameGame.State = game.GameState(source.State)
	gameGame.OpenPlayerSlots = source.OpenPlayerSlots
	gameGame.NumPlayers = source.NumPlayers
	gameGame.VictoryConditions = ExtendVictoryConditions(source)
	gameGame.VictorDeclared = source.VictorDeclared
	gameGame.Seed = source.Seed
	gameGame.Rules = ExtendDefaultRules(source)
	gameGame.Area = ExtendArea(source)
	return gameGame
}
func (c *GameConverter) ConvertGameFleet(source *game.Fleet) *Fleet {
	var pDbFleet *Fleet
	if source != nil {
		dbFleet := c.gameFleetToDbFleet(*source)
		pDbFleet = &dbFleet
	}
	return pDbFleet
}
func (c *GameConverter) ConvertGameGame(source *game.Game) *Game {
	var pDbGame *Game
	if source != nil {
		dbGame := c.gameGameToDbGame(*source)
		pDbGame = &dbGame
	}
	return pDbGame
}
func (c *GameConverter) ConvertGamePlanet(source *game.Planet) *Planet {
	var pDbPlanet *Planet
	if source != nil {
		dbPlanet := c.gamePlanetToDbPlanet(*source)
		pDbPlanet = &dbPlanet
	}
	return pDbPlanet
}
func (c *GameConverter) ConvertGamePlayer(source *game.Player) *Player {
	var pDbPlayer *Player
	if source != nil {
		dbPlayer := c.gamePlayerToDbPlayer(*source)
		pDbPlayer = &dbPlayer
	}
	return pDbPlayer
}
func (c *GameConverter) ConvertGameRace(source *game.Race) *Race {
	var pDbRace *Race
	if source != nil {
		dbRace := c.gameRaceToDbRace(*source)
		pDbRace = &dbRace
	}
	return pDbRace
}
func (c *GameConverter) ConvertGameSalvage(source *game.Salvage) *Salvage {
	var pDbSalvage *Salvage
	if source != nil {
		dbSalvage := c.gameSalvageToDbSalvage(*source)
		pDbSalvage = &dbSalvage
	}
	return pDbSalvage
}
func (c *GameConverter) ConvertGameShipDesign(source *game.ShipDesign) *ShipDesign {
	var pDbShipDesign *ShipDesign
	if source != nil {
		dbShipDesign := c.gameShipDesignToDbShipDesign(*source)
		pDbShipDesign = &dbShipDesign
	}
	return pDbShipDesign
}
func (c *GameConverter) ConvertGameShipToken(source game.ShipToken) ShipToken {
	var dbShipToken ShipToken
	dbShipToken.ID = source.ID
	dbShipToken.CreatedAt = TimeToNullTime(source.CreatedAt)
	dbShipToken.UpdatedAt = TimeToNullTime(source.UpdatedAt)
	dbShipToken.FleetID = source.FleetID
	dbShipToken.DesignUUID = UUIDToUUID(source.DesignUUID)
	dbShipToken.Quantity = source.Quantity
	dbShipToken.Damage = source.Damage
	dbShipToken.QuantityDamaged = source.QuantityDamaged
	return dbShipToken
}
func (c *GameConverter) ConvertGameUser(source *game.User) *User {
	var pDbUser *User
	if source != nil {
		dbUser := c.gameUserToDbUser(*source)
		pDbUser = &dbUser
	}
	return pDbUser
}
func (c *GameConverter) ConvertGameWormhole(source *game.Wormhole) *Wormhole {
	var pDbWormhole *Wormhole
	if source != nil {
		dbWormhole := c.gameWormholeToDbWormhole(*source)
		pDbWormhole = &dbWormhole
	}
	return pDbWormhole
}
func (c *GameConverter) ConvertGames(source []Game) []game.Game {
	gameGameList := make([]game.Game, len(source))
	for i := 0; i < len(source); i++ {
		gameGameList[i] = c.ConvertGame(source[i])
	}
	return gameGameList
}
func (c *GameConverter) ConvertPlanet(source *Planet) *game.Planet {
	var pGamePlanet *game.Planet
	if source != nil {
		gamePlanet := c.dbPlanetToGamePlanet(*source)
		pGamePlanet = &gamePlanet
	}
	return pGamePlanet
}
func (c *GameConverter) ConvertPlayer(source Player) game.Player {
	var gamePlayer game.Player
	gamePlayer.ID = source.ID
	gamePlayer.CreatedAt = TimeToTime(source.CreatedAt)
	gamePlayer.UpdatedAt = TimeToTime(source.UpdatedAt)
	gamePlayer.GameID = source.GameID
	gamePlayer.UserID = source.UserID
	gamePlayer.Name = source.Name
	gamePlayer.Num = source.Num
	gamePlayer.Ready = source.Ready
	gamePlayer.AIControlled = source.AIControlled
	gamePlayer.SubmittedTurn = source.SubmittedTurn
	gamePlayer.Color = source.Color
	gamePlayer.DefaultHullSet = source.DefaultHullSet
	gamePlayer.Race = PlayerRaceToGameRace(source.Race)
	gamePlayer.TechLevels = ExtendTechLevels(source)
	gamePlayer.TechLevelsSpent = ExtendTechLevelsSpent(source)
	gamePlayer.ResearchAmount = source.ResearchAmount
	gamePlayer.ResearchSpentLastYear = source.ResearchSpentLastYear
	gamePlayer.NextResearchField = game.NextResearchField(source.NextResearchField)
	gamePlayer.Researching = game.TechField(source.Researching)
	gamePlayer.BattlePlans = BattlePlansToGameBattlePlans(source.BattlePlans)
	gamePlayer.ProductionPlans = ProductionPlansToGameProductionPlans(source.ProductionPlans)
	gamePlayer.TransportPlans = TransportPlansToGameTransportPlans(source.TransportPlans)
	gamePlayer.Messages = PlayerMessagesToGamePlayerMessages(source.Messages)
	gamePlayer.PlanetIntels = PlanetIntelsToGamePlanetIntels(source.PlanetIntels)
	gamePlayer.FleetIntels = FleetIntelsToGameFleetIntels(source.FleetIntels)
	gamePlayer.ShipDesignIntels = ShipDesignIntelsToGameShipDesignIntels(source.ShipDesignIntels)
	gamePlayer.MineralPacketIntels = MineralPacketIntelsToGameMineralPacketIntels(source.MineralPacketIntels)
	gamePlayer.MineFieldIntels = MineFieldIntelsToGameMineFieldIntels(source.MineFieldIntels)
	gamePlayer.Stats = PlayerStatsToGamePlayerStats(source.Stats)
	gamePlayer.Spec = PlayerSpecToGamePlayerSpec(source.Spec)
	return gamePlayer
}
func (c *GameConverter) ConvertPlayers(source []Player) []game.Player {
	gamePlayerList := make([]game.Player, len(source))
	for i := 0; i < len(source); i++ {
		gamePlayerList[i] = c.ConvertPlayer(source[i])
	}
	return gamePlayerList
}
func (c *GameConverter) ConvertRace(source Race) game.Race {
	var gameRace game.Race
	gameRace.ID = source.ID
	gameRace.CreatedAt = TimeToTime(source.CreatedAt)
	gameRace.UpdatedAt = TimeToTime(source.UpdatedAt)
	gameRace.UserID = source.UserID
	gameRace.Name = source.Name
	gameRace.PluralName = source.PluralName
	gameRace.PRT = game.PRT(source.PRT)
	gameRace.LRTs = game.Bitmask(source.LRTs)
	gameRace.HabLow = ExtendHabLow(source)
	gameRace.HabHigh = ExtendHabHigh(source)
	gameRace.GrowthRate = source.GrowthRate
	gameRace.PopEfficiency = source.PopEfficiency
	gameRace.FactoryOutput = source.FactoryOutput
	gameRace.FactoryCost = source.FactoryCost
	gameRace.NumFactories = source.NumFactories
	gameRace.FactoriesCostLess = source.FactoriesCostLess
	gameRace.ImmuneGrav = source.ImmuneGrav
	gameRace.ImmuneTemp = source.ImmuneTemp
	gameRace.ImmuneRad = source.ImmuneRad
	gameRace.MineOutput = source.MineOutput
	gameRace.MineCost = source.MineCost
	gameRace.NumMines = source.NumMines
	gameRace.ResearchCost = ExtendResearchCost(source)
	gameRace.TechsStartHigh = source.TechsStartHigh
	gameRace.Spec = RaceSpecToGameRaceSpec(source.Spec)
	return gameRace
}
func (c *GameConverter) ConvertRaces(source []Race) []game.Race {
	gameRaceList := make([]game.Race, len(source))
	for i := 0; i < len(source); i++ {
		gameRaceList[i] = c.ConvertRace(source[i])
	}
	return gameRaceList
}
func (c *GameConverter) ConvertSalvage(source *Salvage) *game.Salvage {
	var pGameSalvage *game.Salvage
	if source != nil {
		gameSalvage := c.dbSalvageToGameSalvage(*source)
		pGameSalvage = &gameSalvage
	}
	return pGameSalvage
}
func (c *GameConverter) ConvertShipDesign(source *ShipDesign) *game.ShipDesign {
	var pGameShipDesign *game.ShipDesign
	if source != nil {
		gameShipDesign := c.dbShipDesignToGameShipDesign(*source)
		pGameShipDesign = &gameShipDesign
	}
	return pGameShipDesign
}
func (c *GameConverter) ConvertShipToken(source ShipToken) game.ShipToken {
	var gameShipToken game.ShipToken
	gameShipToken.ID = source.ID
	gameShipToken.CreatedAt = NullTimeToTime(source.CreatedAt)
	gameShipToken.UpdatedAt = NullTimeToTime(source.UpdatedAt)
	gameShipToken.FleetID = source.FleetID
	gameShipToken.DesignUUID = UUIDToUUID(source.DesignUUID)
	gameShipToken.Quantity = source.Quantity
	gameShipToken.Damage = source.Damage
	gameShipToken.QuantityDamaged = source.QuantityDamaged
	return gameShipToken
}
func (c *GameConverter) ConvertUser(source User) game.User {
	var gameUser game.User
	gameUser.ID = source.ID
	gameUser.CreatedAt = TimeToTime(source.CreatedAt)
	gameUser.UpdatedAt = TimeToTime(source.UpdatedAt)
	gameUser.Username = source.Username
	gameUser.Password = source.Password
	gameUser.Role = game.Role(source.Role)
	return gameUser
}
func (c *GameConverter) ConvertUsers(source []User) []game.User {
	gameUserList := make([]game.User, len(source))
	for i := 0; i < len(source); i++ {
		gameUserList[i] = c.ConvertUser(source[i])
	}
	return gameUserList
}
func (c *GameConverter) ConvertWormhole(source *Wormhole) *game.Wormhole {
	var pGameWormhole *game.Wormhole
	if source != nil {
		gameWormhole := c.dbWormholeToGameWormhole(*source)
		pGameWormhole = &gameWormhole
	}
	return pGameWormhole
}
func (c *GameConverter) dbFleetToGameFleet(source Fleet) game.Fleet {
	var gameFleet game.Fleet
	gameFleet.MapObject = ExtendFleetMapObject(source)
	gameFleet.FleetOrders = ExtendFleetFleetOrders(source)
	gameFleet.PlanetNum = source.PlanetNum
	gameFleet.BaseName = source.BaseName
	gameFleet.Cargo = ExtendFleetCargo(source)
	gameFleet.Fuel = source.Fuel
	gameFleet.Damage = source.Damage
	gameFleet.BattlePlanName = source.BattlePlanName
	gameFleet.Heading = ExtendFleetHeading(source)
	gameFleet.WarpSpeed = source.WarpSpeed
	gameFleet.PreviousPosition = ExtendFleetPreviousPosition(source)
	gameFleet.OrbitingPlanetNum = source.OrbitingPlanetNum
	gameFleet.Starbase = source.Starbase
	gameFleet.Spec = FleetSpecToGameFleetSpec(source.Spec)
	return gameFleet
}
func (c *GameConverter) dbPlanetToGamePlanet(source Planet) game.Planet {
	var gamePlanet game.Planet
	gamePlanet.MapObject = ExtendPlanetMapObject(source)
	gamePlanet.Hab = ExtendHab(source)
	gamePlanet.BaseHab = ExtendBaseHab(source)
	gamePlanet.TerraformedAmount = ExtendTerraformedAmount(source)
	gamePlanet.MineralConcentration = ExtendMineralConcentration(source)
	gamePlanet.MineYears = ExtendMineYears(source)
	gamePlanet.Cargo = ExtendPlanetCargo(source)
	gamePlanet.Mines = source.Mines
	gamePlanet.Factories = source.Factories
	gamePlanet.Defenses = source.Defenses
	gamePlanet.Homeworld = source.Homeworld
	gamePlanet.ContributesOnlyLeftoverToResearch = source.ContributesOnlyLeftoverToResearch
	gamePlanet.Scanner = source.Scanner
	gamePlanet.PacketSpeed = source.PacketSpeed
	gamePlanet.BonusResources = source.BonusResources
	gamePlanet.ProductionQueue = ProductionQueueItemsToGameProductionQueueItems(source.ProductionQueue)
	gamePlanet.Spec = PlanetSpecToGamePlanetSpec(source.Spec)
	return gamePlanet
}
func (c *GameConverter) dbSalvageToGameSalvage(source Salvage) game.Salvage {
	var gameSalvage game.Salvage
	gameSalvage.MapObject = ExtendSalvageMapObject(source)
	gameSalvage.Cargo = ExtendSalvageCargo(source)
	return gameSalvage
}
func (c *GameConverter) dbShipDesignToGameShipDesign(source ShipDesign) game.ShipDesign {
	var gameShipDesign game.ShipDesign
	gameShipDesign.ID = source.ID
	gameShipDesign.CreatedAt = TimeToTime(source.CreatedAt)
	gameShipDesign.UpdatedAt = TimeToTime(source.UpdatedAt)
	gameShipDesign.PlayerID = source.PlayerID
	gameShipDesign.PlayerNum = source.PlayerNum
	gameShipDesign.UUID = UUIDToUUID(source.UUID)
	gameShipDesign.Name = source.Name
	gameShipDesign.Version = source.Version
	gameShipDesign.Hull = source.Hull
	gameShipDesign.HullSetNumber = source.HullSetNumber
	gameShipDesign.CanDelete = source.CanDelete
	gameShipDesign.Slots = ShipDesignSlotsToGameShipDesignSlots(source.Slots)
	gameShipDesign.Purpose = game.ShipDesignPurpose(source.Purpose)
	gameShipDesign.Spec = ShipDesignSpecToGameShipDesignSpec(source.Spec)
	return gameShipDesign
}
func (c *GameConverter) dbWormholeToGameWormhole(source Wormhole) game.Wormhole {
	var gameWormhole game.Wormhole
	gameWormhole.MapObject = ExtendWormholeMapObject(source)
	gameWormhole.DestinationNum = source.DestinationNum
	gameWormhole.Stability = game.WormholeStability(source.Stability)
	gameWormhole.YearsAtStability = source.YearsAtStability
	return gameWormhole
}
func (c *GameConverter) gameFleetToDbFleet(source game.Fleet) Fleet {
	var dbFleet Fleet
	dbFleet.ID = source.MapObject.ID
	dbFleet.GameID = source.MapObject.GameID
	dbFleet.CreatedAt = TimeToTime(source.MapObject.CreatedAt)
	dbFleet.UpdatedAt = TimeToTime(source.MapObject.UpdatedAt)
	dbFleet.PlayerID = source.MapObject.PlayerID
	dbFleet.X = source.MapObject.Position.X
	dbFleet.Y = source.MapObject.Position.Y
	dbFleet.Name = source.MapObject.Name
	dbFleet.Num = source.MapObject.Num
	dbFleet.PlayerNum = source.MapObject.PlayerNum
	dbFleet.Waypoints = GameWaypointsToWaypoints(source.FleetOrders.Waypoints)
	dbFleet.RepeatOrders = source.FleetOrders.RepeatOrders
	dbFleet.PlanetNum = source.PlanetNum
	dbFleet.BaseName = source.BaseName
	dbFleet.Ironium = source.Cargo.Ironium
	dbFleet.Boranium = source.Cargo.Boranium
	dbFleet.Germanium = source.Cargo.Germanium
	dbFleet.Colonists = source.Cargo.Colonists
	dbFleet.Fuel = source.Fuel
	dbFleet.Damage = source.Damage
	dbFleet.BattlePlanName = source.BattlePlanName
	dbFleet.HeadingX = source.Heading.X
	dbFleet.HeadingY = source.Heading.Y
	dbFleet.WarpSpeed = source.WarpSpeed
	var pFloat64 *float64
	if source.PreviousPosition != nil {
		pFloat64 = &source.PreviousPosition.X
	}
	var pFloat642 *float64
	if pFloat64 != nil {
		xfloat64 := *pFloat64
		pFloat642 = &xfloat64
	}
	dbFleet.PreviousPositionX = pFloat642
	var pFloat643 *float64
	if source.PreviousPosition != nil {
		pFloat643 = &source.PreviousPosition.Y
	}
	var pFloat644 *float64
	if pFloat643 != nil {
		xfloat642 := *pFloat643
		pFloat644 = &xfloat642
	}
	dbFleet.PreviousPositionY = pFloat644
	dbFleet.OrbitingPlanetNum = source.OrbitingPlanetNum
	dbFleet.Starbase = source.Starbase
	dbFleet.Spec = GameFleetSpecToFleetSpec(source.Spec)
	return dbFleet
}
func (c *GameConverter) gameGameToDbGame(source game.Game) Game {
	var dbGame Game
	dbGame.ID = source.ID
	dbGame.CreatedAt = TimeToTime(source.CreatedAt)
	dbGame.UpdatedAt = TimeToTime(source.UpdatedAt)
	dbGame.Name = source.Name
	dbGame.HostID = source.HostID
	dbGame.QuickStartTurns = source.QuickStartTurns
	dbGame.Size = game.Size(source.Size)
	dbGame.Density = game.Density(source.Density)
	dbGame.PlayerPositions = game.PlayerPositions(source.PlayerPositions)
	dbGame.RandomEvents = source.RandomEvents
	dbGame.ComputerPlayersFormAlliances = source.ComputerPlayersFormAlliances
	dbGame.PublicPlayerScores = source.PublicPlayerScores
	dbGame.StartMode = game.GameStartMode(source.StartMode)
	dbGame.Year = source.Year
	dbGame.State = game.GameState(source.State)
	dbGame.OpenPlayerSlots = source.OpenPlayerSlots
	dbGame.NumPlayers = source.NumPlayers
	dbVictoryConditions := c.gameVictoryConditionListToDbVictoryConditions(source.VictoryConditions.Conditions)
	dbGame.VictoryConditionsConditions = &dbVictoryConditions
	dbGame.VictoryConditionsNumCriteriaRequired = source.VictoryConditions.NumCriteriaRequired
	dbGame.VictoryConditionsYearsPassed = source.VictoryConditions.YearsPassed
	dbGame.VictoryConditionsOwnPlanets = source.VictoryConditions.OwnPlanets
	dbGame.VictoryConditionsAttainTechLevel = source.VictoryConditions.AttainTechLevel
	dbGame.VictoryConditionsAttainTechLevelNumFields = source.VictoryConditions.AttainTechLevelNumFields
	dbGame.VictoryConditionsExceedsScore = source.VictoryConditions.ExceedsScore
	dbGame.VictoryConditionsExceedsSecondPlaceScore = source.VictoryConditions.ExceedsSecondPlaceScore
	dbGame.VictoryConditionsProductionCapacity = source.VictoryConditions.ProductionCapacity
	dbGame.VictoryConditionsOwnCapitalShips = source.VictoryConditions.OwnCapitalShips
	dbGame.VictoryConditionsHighestScoreAfterYears = source.VictoryConditions.HighestScoreAfterYears
	dbGame.VictorDeclared = source.VictorDeclared
	dbGame.Seed = source.Seed
	dbGame.Rules = GameRulesToRules(source.Rules)
	dbGame.AreaX = source.Area.X
	dbGame.AreaY = source.Area.Y
	return dbGame
}
func (c *GameConverter) gamePlanetToDbPlanet(source game.Planet) Planet {
	var dbPlanet Planet
	dbPlanet.ID = source.MapObject.ID
	dbPlanet.GameID = source.MapObject.GameID
	dbPlanet.CreatedAt = TimeToTime(source.MapObject.CreatedAt)
	dbPlanet.UpdatedAt = TimeToTime(source.MapObject.UpdatedAt)
	dbPlanet.PlayerID = source.MapObject.PlayerID
	dbPlanet.X = source.MapObject.Position.X
	dbPlanet.Y = source.MapObject.Position.Y
	dbPlanet.Name = source.MapObject.Name
	dbPlanet.Num = source.MapObject.Num
	dbPlanet.PlayerNum = source.MapObject.PlayerNum
	dbPlanet.Grav = source.Hab.Grav
	dbPlanet.Temp = source.Hab.Temp
	dbPlanet.Rad = source.Hab.Rad
	dbPlanet.BaseGrav = source.BaseHab.Grav
	dbPlanet.BaseTemp = source.BaseHab.Temp
	dbPlanet.BaseRad = source.BaseHab.Rad
	dbPlanet.TerraformedAmountGrav = source.TerraformedAmount.Grav
	dbPlanet.TerraformedAmountTemp = source.TerraformedAmount.Temp
	dbPlanet.TerraformedAmountRad = source.TerraformedAmount.Rad
	dbPlanet.MineralConcIronium = source.MineralConcentration.Ironium
	dbPlanet.MineralConcBoranium = source.MineralConcentration.Boranium
	dbPlanet.MineralConcGermanium = source.MineralConcentration.Germanium
	dbPlanet.MineYearsIronium = source.MineYears.Ironium
	dbPlanet.MineYearsBoranium = source.MineYears.Boranium
	dbPlanet.MineYearsGermanium = source.MineYears.Germanium
	dbPlanet.Ironium = source.Cargo.Ironium
	dbPlanet.Boranium = source.Cargo.Boranium
	dbPlanet.Germanium = source.Cargo.Germanium
	dbPlanet.Colonists = source.Cargo.Colonists
	dbPlanet.Mines = source.Mines
	dbPlanet.Factories = source.Factories
	dbPlanet.Defenses = source.Defenses
	dbPlanet.Homeworld = source.Homeworld
	dbPlanet.ContributesOnlyLeftoverToResearch = source.ContributesOnlyLeftoverToResearch
	dbPlanet.Scanner = source.Scanner
	dbPlanet.PacketSpeed = source.PacketSpeed
	dbPlanet.BonusResources = source.BonusResources
	dbPlanet.ProductionQueue = GameProductionQueueItemsToProductionQueueItems(source.ProductionQueue)
	dbPlanet.Spec = GamePlanetSpecToPlanetSpec(source.Spec)
	return dbPlanet
}
func (c *GameConverter) gamePlayerToDbPlayer(source game.Player) Player {
	var dbPlayer Player
	dbPlayer.ID = source.ID
	dbPlayer.CreatedAt = TimeToTime(source.CreatedAt)
	dbPlayer.UpdatedAt = TimeToTime(source.UpdatedAt)
	dbPlayer.GameID = source.GameID
	dbPlayer.UserID = source.UserID
	dbPlayer.Name = source.Name
	dbPlayer.Num = source.Num
	dbPlayer.Ready = source.Ready
	dbPlayer.AIControlled = source.AIControlled
	dbPlayer.SubmittedTurn = source.SubmittedTurn
	dbPlayer.Color = source.Color
	dbPlayer.DefaultHullSet = source.DefaultHullSet
	dbPlayer.TechLevelsEnergy = source.TechLevels.Energy
	dbPlayer.TechLevelsWeapons = source.TechLevels.Weapons
	dbPlayer.TechLevelsPropulsion = source.TechLevels.Propulsion
	dbPlayer.TechLevelsConstruction = source.TechLevels.Construction
	dbPlayer.TechLevelsElectronics = source.TechLevels.Electronics
	dbPlayer.TechLevelsBiotechnology = source.TechLevels.Biotechnology
	dbPlayer.TechLevelsSpentEnergy = source.TechLevelsSpent.Energy
	dbPlayer.TechLevelsSpentWeapons = source.TechLevelsSpent.Weapons
	dbPlayer.TechLevelsSpentPropulsion = source.TechLevelsSpent.Propulsion
	dbPlayer.TechLevelsSpentConstruction = source.TechLevelsSpent.Construction
	dbPlayer.TechLevelsSpentElectronics = source.TechLevelsSpent.Electronics
	dbPlayer.TechLevelsSpentBiotechnology = source.TechLevelsSpent.Biotechnology
	dbPlayer.ResearchAmount = source.ResearchAmount
	dbPlayer.ResearchSpentLastYear = source.ResearchSpentLastYear
	dbPlayer.NextResearchField = game.NextResearchField(source.NextResearchField)
	dbPlayer.Researching = game.TechField(source.Researching)
	dbPlayer.BattlePlans = GameBattlePlansToBattlePlans(source.BattlePlans)
	dbPlayer.ProductionPlans = GameProductionPlansToProductionPlans(source.ProductionPlans)
	dbPlayer.TransportPlans = GameTransportPlansToTransportPlans(source.TransportPlans)
	dbPlayer.Messages = GamePlayerMessagesToPlayerMessages(source.Messages)
	dbPlayer.PlanetIntels = GamePlanetIntelsToPlanetIntels(source.PlanetIntels)
	dbPlayer.FleetIntels = GameFleetIntelsToFleetIntels(source.FleetIntels)
	dbPlayer.ShipDesignIntels = GameShipDesignIntelsToShipDesignIntels(source.ShipDesignIntels)
	dbPlayer.MineralPacketIntels = GameMineralPacketIntelsToMineralPacketIntels(source.MineralPacketIntels)
	dbPlayer.MineFieldIntels = GameMineFieldIntelsToMineFieldIntels(source.MineFieldIntels)
	dbPlayer.Race = GameRaceToPlayerRace(source.Race)
	dbPlayer.Stats = GamePlayerStatsToPlayerStats(source.Stats)
	dbPlayer.Spec = GamePlayerSpecToPlayerSpec(source.Spec)
	return dbPlayer
}
func (c *GameConverter) gameRaceToDbRace(source game.Race) Race {
	var dbRace Race
	dbRace.ID = source.ID
	dbRace.CreatedAt = TimeToTime(source.CreatedAt)
	dbRace.UpdatedAt = TimeToTime(source.UpdatedAt)
	dbRace.UserID = source.UserID
	dbRace.Name = source.Name
	dbRace.PluralName = source.PluralName
	dbRace.PRT = game.PRT(source.PRT)
	dbRace.LRTs = game.Bitmask(source.LRTs)
	dbRace.HabLowGrav = source.HabLow.Grav
	dbRace.HabLowTemp = source.HabLow.Temp
	dbRace.HabLowRad = source.HabLow.Rad
	dbRace.HabHighGrav = source.HabHigh.Grav
	dbRace.HabHighTemp = source.HabHigh.Temp
	dbRace.HabHighRad = source.HabHigh.Rad
	dbRace.GrowthRate = source.GrowthRate
	dbRace.PopEfficiency = source.PopEfficiency
	dbRace.FactoryOutput = source.FactoryOutput
	dbRace.FactoryCost = source.FactoryCost
	dbRace.NumFactories = source.NumFactories
	dbRace.FactoriesCostLess = source.FactoriesCostLess
	dbRace.ImmuneGrav = source.ImmuneGrav
	dbRace.ImmuneTemp = source.ImmuneTemp
	dbRace.ImmuneRad = source.ImmuneRad
	dbRace.MineOutput = source.MineOutput
	dbRace.MineCost = source.MineCost
	dbRace.NumMines = source.NumMines
	dbRace.ResearchCostEnergy = game.ResearchCostLevel(source.ResearchCost.Energy)
	dbRace.ResearchCostWeapons = game.ResearchCostLevel(source.ResearchCost.Weapons)
	dbRace.ResearchCostPropulsion = game.ResearchCostLevel(source.ResearchCost.Propulsion)
	dbRace.ResearchCostConstruction = game.ResearchCostLevel(source.ResearchCost.Construction)
	dbRace.ResearchCostElectronics = game.ResearchCostLevel(source.ResearchCost.Electronics)
	dbRace.ResearchCostBiotechnology = game.ResearchCostLevel(source.ResearchCost.Biotechnology)
	dbRace.TechsStartHigh = source.TechsStartHigh
	dbRace.Spec = GameRaceSpecToRaceSpec(source.Spec)
	return dbRace
}
func (c *GameConverter) gameSalvageToDbSalvage(source game.Salvage) Salvage {
	var dbSalvage Salvage
	dbSalvage.ID = source.MapObject.ID
	dbSalvage.GameID = source.MapObject.GameID
	dbSalvage.CreatedAt = TimeToTime(source.MapObject.CreatedAt)
	dbSalvage.UpdatedAt = TimeToTime(source.MapObject.UpdatedAt)
	dbSalvage.PlayerID = source.MapObject.PlayerID
	dbSalvage.X = source.MapObject.Position.X
	dbSalvage.Y = source.MapObject.Position.Y
	dbSalvage.Name = source.MapObject.Name
	dbSalvage.Num = source.MapObject.Num
	dbSalvage.PlayerNum = source.MapObject.PlayerNum
	dbSalvage.Ironium = source.Cargo.Ironium
	dbSalvage.Boranium = source.Cargo.Boranium
	dbSalvage.Germanium = source.Cargo.Germanium
	return dbSalvage
}
func (c *GameConverter) gameShipDesignToDbShipDesign(source game.ShipDesign) ShipDesign {
	var dbShipDesign ShipDesign
	dbShipDesign.ID = source.ID
	dbShipDesign.CreatedAt = TimeToTime(source.CreatedAt)
	dbShipDesign.UpdatedAt = TimeToTime(source.UpdatedAt)
	dbShipDesign.PlayerID = source.PlayerID
	dbShipDesign.PlayerNum = source.PlayerNum
	dbShipDesign.UUID = UUIDToUUID(source.UUID)
	dbShipDesign.Name = source.Name
	dbShipDesign.Version = source.Version
	dbShipDesign.Hull = source.Hull
	dbShipDesign.HullSetNumber = source.HullSetNumber
	dbShipDesign.CanDelete = source.CanDelete
	dbShipDesign.Slots = GameShipDesignSlotsToShipDesignSlots(source.Slots)
	dbShipDesign.Purpose = game.ShipDesignPurpose(source.Purpose)
	dbShipDesign.Spec = GameShipDesignSpecToShipDesignSpec(source.Spec)
	return dbShipDesign
}
func (c *GameConverter) gameUserToDbUser(source game.User) User {
	var dbUser User
	dbUser.ID = source.ID
	dbUser.CreatedAt = TimeToTime(source.CreatedAt)
	dbUser.UpdatedAt = TimeToTime(source.UpdatedAt)
	dbUser.Username = source.Username
	dbUser.Password = source.Password
	dbUser.Role = string(source.Role)
	return dbUser
}
func (c *GameConverter) gameVictoryConditionListToDbVictoryConditions(source []game.VictoryCondition) VictoryConditions {
	dbVictoryConditions := make(VictoryConditions, len(source))
	for i := 0; i < len(source); i++ {
		dbVictoryConditions[i] = game.VictoryCondition(source[i])
	}
	return dbVictoryConditions
}
func (c *GameConverter) gameWormholeToDbWormhole(source game.Wormhole) Wormhole {
	var dbWormhole Wormhole
	dbWormhole.ID = source.MapObject.ID
	dbWormhole.GameID = source.MapObject.GameID
	dbWormhole.CreatedAt = TimeToTime(source.MapObject.CreatedAt)
	dbWormhole.UpdatedAt = TimeToTime(source.MapObject.UpdatedAt)
	dbWormhole.X = source.MapObject.Position.X
	dbWormhole.Y = source.MapObject.Position.Y
	dbWormhole.Name = source.MapObject.Name
	dbWormhole.Num = source.MapObject.Num
	dbWormhole.DestinationNum = source.DestinationNum
	dbWormhole.Stability = game.WormholeStability(source.Stability)
	dbWormhole.YearsAtStability = source.YearsAtStability
	return dbWormhole
}
