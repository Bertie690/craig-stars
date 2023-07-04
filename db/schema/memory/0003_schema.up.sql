CREATE TABLE IF NOT EXISTS games (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  hostId INTEGER,
  name TEXT NOT NULL,
  state TEXT,
  public NUMERIC,
  hash TEXT,
  size TEXT,
  density TEXT,
  playerPositions TEXT,
  randomEvents NUMERIC,
  computerPlayersFormAlliances NUMERIC,
  publicPlayerScores NUMERIC,
  startMode TEXT,
  quickStartTurns INTEGER,
  openPlayerSlots INTEGER,
  numPlayers INTEGER,
  victoryConditionsConditions INTEGER,
  victoryConditionsNumCriteriaRequired INTEGER,
  victoryConditionsYearsPassed INTEGER,
  victoryConditionsOwnPlanets INTEGER,
  victoryConditionsAttainTechLevel INTEGER,
  victoryConditionsAttainTechLevelNumFields INTEGER,
  victoryConditionsExceedsScore INTEGER,
  victoryConditionsExceedsSecondPlaceScore INTEGER,
  victoryConditionsProductionCapacity INTEGER,
  victoryConditionsOwnCapitalShips INTEGER,
  victoryConditionsHighestScoreAfterYears INTEGER,
  seed INTEGER,
  rules TEXT,
  areaX REAL,
  areaY REAL,
  year INTEGER,
  victorDeclared NUMERIC
);
CREATE TABLE IF NOT EXISTS rules (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  seed INTEGER,
  tachyonCloakReduction INTEGER,
  maxPopulation INTEGER,
  fleetsScanWhileMoving NUMERIC,
  populationScannerError REAL,
  smartDefenseCoverageFactor REAL,
  invasionDefenseCoverageFactor REAL,
  numBattleRounds INTEGER,
  movesToRunAway INTEGER,
  beamRangeDropoff REAL,
  torpedoSplashDamage REAL,
  salvageDecayRate REAL,
  salvageDecayMin INTEGER,
  mineFieldCloak INTEGER,
  stargateMaxRangeFactor INTEGER,
  stargateMaxHullMassFactor INTEGER,
  randomEventChances TEXT,
  randomMineralDepositBonusRange TEXT,
  wormholeCloak INTEGER,
  wormholeMinDistance INTEGER,
  wormholeStatsByStability TEXT,
  wormholePairsForSize TEXT,
  mineFieldStatsByType TEXT,
  repairRates TEXT,
  maxPlayers INTEGER,
  startingYear INTEGER,
  showPublicScoresAfterYears INTEGER,
  planetMinDistance INTEGER,
  maxExtraWorldDistance INTEGER,
  minExtraWorldDistance INTEGER,
  minHomeworldMineralConcentration INTEGER,
  minExtraPlanetMineralConcentration INTEGER,
  minMineralConcentration INTEGER,
  minStartingMineralConcentration INTEGER,
  maxStartingMineralConcentration INTEGER,
  highRadGermaniumBonus INTEGER,
  highRadGermaniumBonusThreshold INTEGER,
  maxStartingMineralSurface INTEGER,
  minStartingMineralSurface INTEGER,
  mineralDecayFactor INTEGER,
  startingMines INTEGER,
  startingFactories INTEGER,
  startingDefenses INTEGER,
  raceStartingPoints INTEGER,
  scrapMineralAmount REAL,
  scrapResourceAmount REAL,
  factoryCostGermanium INTEGER,
  defenseCost TEXT,
  mineralAlchemyCost INTEGER,
  terraformCost TEXT,
  starbaseComponentCostFactor REAL,
  packetDecayRate TEXT,
  maxTechLevel INTEGER,
  techBaseCost TEXT,
  prtSpecs TEXT,
  lrtSpecs TEXT,
  techsId INTEGER
);
CREATE TABLE IF NOT EXISTS players (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  userId INTEGER,
  name TEXT NOT NULL,
  num INTEGER NOT NULL,
  ready NUMERIC,
  aiControlled NUMERIC,
  submittedTurn NUMERIC,
  color TEXT,
  defaultHullSet INTEGER,
  techLevelsEnergy INTEGER,
  techLevelsWeapons INTEGER,
  techLevelsPropulsion INTEGER,
  techLevelsConstruction INTEGER,
  techLevelsElectronics INTEGER,
  techLevelsBiotechnology INTEGER,
  techLevelsSpentEnergy INTEGER,
  techLevelsSpentWeapons INTEGER,
  techLevelsSpentPropulsion INTEGER,
  techLevelsSpentConstruction INTEGER,
  techLevelsSpentElectronics INTEGER,
  techLevelsSpentBiotechnology INTEGER,
  researchAmount INTEGER,
  researchSpentLastYear INTEGER,
  nextResearchField TEXT,
  researching TEXT,
  battlePlans TEXT,
  productionPlans TEXT,
  transportPlans TEXT,
  relations TEXT,
  messages TEXT,
  battleRecords TEXT,
  playerIntels TEXT,
  scoreIntels TEXT,
  planetIntels TEXT,
  fleetIntels TEXT,
  starbaseIntels TEXT,
  shipDesignIntels TEXT,
  mineralPacketIntels TEXT,
  mineFieldIntels TEXT,
  wormholeIntels TEXT,
  mysteryTraderIntels TEXT,
  salvageIntels TEXT,
  race TEXT,
  stats TEXT,
  scoreHistory TEXT,
  achievedVictoryConditions INTEGER,
  victor NUMERIC,
  spec TEXT,
  UNIQUE (gameId, num),
  CONSTRAINT fkGamesPlayers FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS fleets (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  battlePlanNum INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  playerNum INTEGER,
  tokens TEXT,
  waypoints TEXT,
  repeatOrders NUMERIC,
  planetNum INTEGER,
  baseName TEXT NOT NULL,
  ironium INTEGER,
  boranium INTEGER,
  germanium INTEGER,
  colonists INTEGER,
  fuel INTEGER,
  age INTEGER,
  headingX REAL,
  headingY REAL,
  warpSpeed INTEGER,
  previousPositionX REAL,
  previousPositionY REAL,
  orbitingPlanetNum INTEGER,
  starbase NUMERIC,
  spec TEXT,
  CONSTRAINT fkPlayersFleets FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num) ON DELETE CASCADE,
  CONSTRAINT fkGamesFleets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX IF NOT EXISTS fleetNum on fleets(gameId, playerNum, num) WHERE starbase = 0;
CREATE TABLE IF NOT EXISTS shipDesigns (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  num INTEGER NOT NULL,
  playerNum INTEGER NOT NULL,
  name TEXT NOT NULL,
  version INTEGER,
  hull TEXT,
  hullSetNumber INTEGER,
  canDelete NUMERIC,
  slots TEXT,
  purpose TEXT,
  spec TEXT,
  UNIQUE (gameId, playerNum, num),
  UNIQUE (gameId, playerNum, name),
  CONSTRAINT fkPlayersDesigns FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS planets (
  id INTEGER PRIMARY KEY,
  gameId INTEGER NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  playerNum INTEGER,
  grav INTEGER,
  temp INTEGER,
  rad INTEGER,
  baseGrav INTEGER,
  baseTemp INTEGER,
  baseRad INTEGER,
  terraformedAmountGrav INTEGER,
  terraformedAmountTemp INTEGER,
  terraformedAmountRad INTEGER,
  mineralConcIronium INTEGER,
  mineralConcBoranium INTEGER,
  mineralConcGermanium INTEGER,
  mineYearsIronium INTEGER,
  mineYearsBoranium INTEGER,
  mineYearsGermanium INTEGER,
  ironium INTEGER,
  boranium INTEGER,
  germanium INTEGER,
  colonists INTEGER,
  mines INTEGER,
  factories INTEGER,
  defenses INTEGER,
  homeworld NUMERIC,
  contributesOnlyLeftoverToResearch NUMERIC,
  scanner NUMERIC,
  routeTargetType TEXT,
  routeTargetNum INTEGER,
  routeTargetPlayerNum INTEGER,
  packetTargetNum INTEGER,
  packetSpeed INTEGER,
  productionQueue TEXT,
  spec TEXT,
  UNIQUE (gameId, num),
  CONSTRAINT fkGamesPlanets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS mineralPackets (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  playerNum INTEGER,
  targetPlanetNum INTEGER,
  ironium INTEGER,
  boranium INTEGER,
  germanium INTEGER,
  safeWarpSpeed INTEGER,
  warpSpeed INTEGER,
  scanRange INTEGER,
  scanRangePen INTEGER,
  headingX REAL,
  headingY REAL,
  UNIQUE (gameId, playerNum, num),
  CONSTRAINT fkPlayersMineralPackets FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num),
  CONSTRAINT fkGamesMineralPackets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS salvages (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  playerNum INTEGER,
  ironium INTEGER,
  boranium INTEGER,
  germanium INTEGER,
  UNIQUE (gameId, num),
  CONSTRAINT fkPlayersSalvages FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num),
  CONSTRAINT fkGamesSalvages FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS wormholes (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  destinationNum INTEGER,
  stability TEXT,
  yearsAtStability INTEGER,
  spec TEXT,
  UNIQUE (gameId, num),
  CONSTRAINT fkGamesWormholes FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS mysteryTraders (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  headingX REAL,
  headingY REAL,
  warpSpeed INTEGER,
  spec TEXT,
  UNIQUE (gameId, num),
  CONSTRAINT fkGamesMysteryTraders FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS mineFields (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  gameId INTEGER NOT NULL,
  x REAL,
  y REAL,
  name TEXT NOT NULL,
  num INTEGER,
  playerNum INTEGER,
  numMines INTEGER,
  detonate NUMERIC,
  mineFieldType TEXT,
  spec TEXT,
  UNIQUE (gameId, playerNum, num),
  CONSTRAINT fkPlayersMineFields FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num),
  CONSTRAINT fkGamesMineFields FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS techStores (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  rulesId INTEGER
);
CREATE TABLE IF NOT EXISTS techEngines (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  techStoreId INTEGER NOT NULL,
  name TEXT NOT NULL,
  costIronium INTEGER,
  costBoranium INTEGER,
  costGermanium INTEGER,
  costResources INTEGER,
  requirementsEnergy INTEGER,
  requirementsWeapons INTEGER,
  requirementsPropulsion INTEGER,
  requirementsConstruction INTEGER,
  requirementsElectronics INTEGER,
  requirementsBiotechnology INTEGER,
  requirementsPrtDenied TEXT,
  requirementsLrTsRequired INTEGER,
  requirementsLrTsDenied INTEGER,
  requirementsPrtRequired TEXT,
  ranking INTEGER,
  category TEXT,
  hullSlotType INTEGER,
  mass INTEGER,
  scanRange INTEGER,
  scanRangePen INTEGER,
  safeHullMass INTEGER,
  safeRange INTEGER,
  maxHullMass INTEGER,
  maxRange INTEGER,
  radiating NUMERIC,
  packetSpeed INTEGER,
  cloakUnits INTEGER,
  terraformRate INTEGER,
  miningRate INTEGER,
  killRate REAL,
  minKillRate INTEGER,
  structureDestroyRate REAL,
  unterraformRate INTEGER,
  smart NUMERIC,
  canStealFleetCargo NUMERIC,
  canStealPlanetCargo NUMERIC,
  armor INTEGER,
  shield INTEGER,
  torpedoBonus REAL,
  initiativeBonus INTEGER,
  beamBonus REAL,
  reduceMovement INTEGER,
  torpedoJamming REAL,
  reduceCloaking NUMERIC,
  cloakUnarmedOnly NUMERIC,
  mineFieldType TEXT,
  mineLayingRate INTEGER,
  beamDefense INTEGER,
  cargoBonus INTEGER,
  colonizationModule NUMERIC,
  fuelBonus INTEGER,
  movementBonus INTEGER,
  orbitalConstructionModule NUMERIC,
  power INTEGER,
  range INTEGER,
  initiative INTEGER,
  gattling NUMERIC,
  hitsAllTargets NUMERIC,
  damageShieldsOnly NUMERIC,
  fuelRegenerationRate INTEGER,
  accuracy INTEGER,
  capitalShipMissile NUMERIC,
  idealSpeed INTEGER,
  freeSpeed INTEGER,
  fuelUsage TEXT,
  CONSTRAINT fkTechStoresEngines FOREIGN KEY (techStoreId) REFERENCES techStores (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS techPlanetaryScanners (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  techStoreId INTEGER NOT NULL,
  name TEXT NOT NULL,
  costIronium INTEGER,
  costBoranium INTEGER,
  costGermanium INTEGER,
  costResources INTEGER,
  requirementsEnergy INTEGER,
  requirementsWeapons INTEGER,
  requirementsPropulsion INTEGER,
  requirementsConstruction INTEGER,
  requirementsElectronics INTEGER,
  requirementsBiotechnology INTEGER,
  requirementsPrtDenied TEXT,
  requirementsLrTsRequired INTEGER,
  requirementsLrTsDenied INTEGER,
  requirementsPrtRequired TEXT,
  ranking INTEGER,
  category TEXT,
  scanRange INTEGER,
  scanRangePen INTEGER,
  CONSTRAINT fkTechStoresPlanetaryScanners FOREIGN KEY (techStoreId) REFERENCES techStores (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS techDefenses (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  techStoreId INTEGER NOT NULL,
  name TEXT NOT NULL,
  costIronium INTEGER,
  costBoranium INTEGER,
  costGermanium INTEGER,
  costResources INTEGER,
  requirementsEnergy INTEGER,
  requirementsWeapons INTEGER,
  requirementsPropulsion INTEGER,
  requirementsConstruction INTEGER,
  requirementsElectronics INTEGER,
  requirementsBiotechnology INTEGER,
  requirementsPrtDenied TEXT,
  requirementsLrTsRequired INTEGER,
  requirementsLrTsDenied INTEGER,
  requirementsPrtRequired TEXT,
  ranking INTEGER,
  category TEXT,
  defenseCoverage REAL,
  CONSTRAINT fkTechStoresDefenses FOREIGN KEY (techStoreId) REFERENCES techStores (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS techHullComponents (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  techStoreId INTEGER NOT NULL,
  name TEXT NOT NULL,
  costIronium INTEGER,
  costBoranium INTEGER,
  costGermanium INTEGER,
  costResources INTEGER,
  requirementsEnergy INTEGER,
  requirementsWeapons INTEGER,
  requirementsPropulsion INTEGER,
  requirementsConstruction INTEGER,
  requirementsElectronics INTEGER,
  requirementsBiotechnology INTEGER,
  requirementsPrtDenied TEXT,
  requirementsLrTsRequired INTEGER,
  requirementsLrTsDenied INTEGER,
  requirementsPrtRequired TEXT,
  ranking INTEGER,
  category TEXT,
  hullSlotType INTEGER,
  mass INTEGER,
  scanRange INTEGER,
  scanRangePen INTEGER,
  safeHullMass INTEGER,
  safeRange INTEGER,
  maxHullMass INTEGER,
  maxRange INTEGER,
  radiating NUMERIC,
  packetSpeed INTEGER,
  cloakUnits INTEGER,
  terraformRate INTEGER,
  miningRate INTEGER,
  killRate REAL,
  minKillRate INTEGER,
  structureDestroyRate REAL,
  unterraformRate INTEGER,
  smart NUMERIC,
  canStealFleetCargo NUMERIC,
  canStealPlanetCargo NUMERIC,
  armor INTEGER,
  shield INTEGER,
  torpedoBonus REAL,
  initiativeBonus INTEGER,
  beamBonus REAL,
  reduceMovement INTEGER,
  torpedoJamming REAL,
  reduceCloaking NUMERIC,
  cloakUnarmedOnly NUMERIC,
  mineFieldType TEXT,
  mineLayingRate INTEGER,
  beamDefense INTEGER,
  cargoBonus INTEGER,
  colonizationModule NUMERIC,
  fuelBonus INTEGER,
  movementBonus INTEGER,
  orbitalConstructionModule NUMERIC,
  power INTEGER,
  range INTEGER,
  initiative INTEGER,
  gattling NUMERIC,
  hitsAllTargets NUMERIC,
  damageShieldsOnly NUMERIC,
  fuelRegenerationRate INTEGER,
  accuracy INTEGER,
  capitalShipMissile NUMERIC,
  CONSTRAINT fkTechStoresHullComponents FOREIGN KEY (techStoreId) REFERENCES techStores (id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS techHulls (
  id INTEGER PRIMARY KEY,
  createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
  techStoreId INTEGER NOT NULL,
  name TEXT NOT NULL,
  costIronium INTEGER,
  costBoranium INTEGER,
  costGermanium INTEGER,
  costResources INTEGER,
  requirementsEnergy INTEGER,
  requirementsWeapons INTEGER,
  requirementsPropulsion INTEGER,
  requirementsConstruction INTEGER,
  requirementsElectronics INTEGER,
  requirementsBiotechnology INTEGER,
  requirementsPrtDenied TEXT,
  requirementsLrTsRequired INTEGER,
  requirementsLrTsDenied INTEGER,
  requirementsPrtRequired TEXT,
  ranking INTEGER,
  category TEXT,
  type TEXT,
  mass INTEGER,
  armor INTEGER,
  fuelCapacity INTEGER,
  cargoCapacity INTEGER,
  spaceDock INTEGER,
  mineLayingFactor INTEGER,
  builtInScanner NUMERIC,
  initiative INTEGER,
  repairBonus REAL,
  immuneToOwnDetonation NUMERIC,
  rangeBonus INTEGER,
  starbase NUMERIC,
  orbitalConstructionHull NUMERIC,
  doubleMineEfficiency NUMERIC,
  innateScanRangePenFactor REAL,
  slots TEXT,
  CONSTRAINT fkTechStoresHulls FOREIGN KEY (techStoreId) REFERENCES techStores (id) ON DELETE CASCADE
);