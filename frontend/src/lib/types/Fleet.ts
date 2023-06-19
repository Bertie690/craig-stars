// Generated by https://quicktype.io

import type { Cargo } from './Cargo';
import type { Cost } from './Cost';
import type { MapObject, MapObjectType } from './MapObject';
import type { Vector } from './Vector';

export interface Fleet extends MapObject {
	baseName: string;
	fuel?: number;
	cargo?: Cargo;
	damage?: number;
	battlePlanName?: string;
	tokens?: ShipToken[];
	waypoints?: Waypoint[];

	repeatOrders?: boolean;
	heading?: Vector;
	warpSpeed?: number;
	orbitingPlanetNum?: number;
	starbase?: boolean;
	spec?: Spec;
}

export interface ShipToken {
	id?: number;
	createdAt?: string;
	updatedAt?: string;

	gameId?: number;
	designUuid: string;
	quantity: number;
}

export interface Waypoint {
	position: Vector;
	warpFactor: number;
	estFuelUsage?: number;
	task?: WaypointTask;
	waitAtWaypoint?: boolean;
	targetType?: MapObjectType;
	targetNum?: number;
	targetName?: string;
	targetPlayerNum?: number;
	transferToPlayer?: number;
	partiallyComplete?: boolean;
}

export enum WaypointTask {
	None = '',
	Transport = 'Transport',
	Colonize = 'Colonize',
	RemoteMining = 'RemoteMining',
	MergeWithFleet = 'MergeWithFleet',
	ScrapFleet = 'ScrapFleet',
	LayMineField = 'LayMineField',
	Patrol = 'Patrol',
	Route = 'Route',
	TransferFleet = 'TransferFleet'
}

export interface Spec {
	engine: string;
	cost: Cost;
	mass: number;
	armor: number;
	fuelCapacity: number;
	immuneToOwnDetonation: boolean;
	mineLayingRateByMineType?: null;
	weaponSlots?: null;
	purposes?: any;
	totalShips: number;
	massEmpty: number;
	basePacketSpeed: number;
	safePacketSpeed: number;
	baseCloakedCargo: number;
	stargate?: string;

	idealSpeed?: number;
	numEngines?: number;
	cargoCapacity?: number;
	cloakUnits?: number;
	scanRange?: number;
	scanRangePen?: number;
	repairBonus?: number;
	torpedoInaccuracyFactor?: number;
	initiative?: number;
	movement?: number;
	powerRating?: number;
	bomber?: number;
	bombs?: number;
	smartBombs?: number;
	retroBombs?: number;
	scanner?: boolean;
	shield?: number;
	colonizer?: number;
	canLayMines?: number;
	spaceDock?: number;
	miningRate?: number;
	terraformRate?: number;
	mineSweep?: number;
	cloakPercent?: number;
	reduceCloaking?: number;
	canStealFleetCargo?: number;
	canStealPlanetCargo?: number;
	orbitalConstructionModule?: number;
	hasWeapons?: boolean;
	hasStargate?: boolean;
	hasMassDriver?: boolean;
}
