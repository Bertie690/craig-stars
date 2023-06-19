CREATE TABLE
  games (
    id INTEGER PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    name TEXT NOT NULL,
    hostId INTEGER,
    quickStartTurns INTEGER,
    size TEXT,
    density TEXT,
    playerPositions TEXT,
    randomEvents NUMERIC,
    computerPlayersFormAlliances NUMERIC,
    publicPlayerScores NUMERIC,
    startMode TEXT,
    year INTEGER,
    state TEXT,
    openPlayerSlots INTEGER,
    numPlayers INTEGER,
    victoryConditionsConditions TEXT,
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
    victorDeclared NUMERIC,
    seed INTEGER,
    rules TEXT,
    areaX REAL,
    areaY REAL
  );

CREATE TABLE
  rules (
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

CREATE TABLE
  players (
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
    playerIntels TEXT,
    planetIntels TEXT,
    fleetIntels TEXT,
    shipDesignIntels TEXT,
    mineralPacketIntels TEXT,
    mineFieldIntels TEXT,
    wormholeIntels TEXT,
    race TEXT,
    stats TEXT,
    spec TEXT,
    UNIQUE (gameId, num),
    CONSTRAINT fkGamesPlayers FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
  );

CREATE TABLE
  fleets (
    id INTEGER PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    gameId INTEGER NOT NULL,
    battlePlanName TEXT NOT NULL,
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
    damage INTEGER,
    headingX REAL,
    headingY REAL,
    warpSpeed INTEGER,
    previousPositionX REAL,
    previousPositionY REAL,
    orbitingPlanetNum INTEGER,
    starbase NUMERIC,
    spec TEXT,
    UNIQUE (gameId, playerNum, num),
    CONSTRAINT fkPlayersFleets FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num) ON DELETE CASCADE,
    CONSTRAINT fkGamesFleets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
  );

CREATE TABLE
  shipDesigns (
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

CREATE TABLE
  planets (
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
    packetSpeed INTEGER,
    productionQueue TEXT,
    spec TEXT,
    UNIQUE (gameId, num),
    CONSTRAINT fkGamesPlanets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
  );

CREATE TABLE
  mineralPackets (
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
    warpFactor INTEGER,
    distanceTravelled REAL,
    headingX REAL,
    headingY REAL,
    UNIQUE (gameId, playerNum, num),
    CONSTRAINT fkPlayersMineralPackets FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num),
    CONSTRAINT fkGamesMineralPackets FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
  );

CREATE TABLE
  salvages (
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

CREATE TABLE
  wormholes (
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

CREATE TABLE
  mysteryTraders (
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

CREATE TABLE
  mineFields (
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
    type TEXT,
    UNIQUE (gameId, playerNum, num),
    CONSTRAINT fkPlayersMineFields FOREIGN KEY (gameId, playerNum) REFERENCES players (gameId, num),
    CONSTRAINT fkGamesMineFields FOREIGN KEY (gameId) REFERENCES games (id) ON DELETE CASCADE
  );

CREATE TABLE
  techStores (
    id INTEGER PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENTTIMESTAMP,
    rulesId INTEGER
  );

CREATE TABLE
  techEngines (
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

CREATE TABLE
  techPlanetaryScanners (
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

CREATE TABLE
  techDefenses (
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

CREATE TABLE
  techHullComponents (
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

CREATE TABLE
  techHulls (
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