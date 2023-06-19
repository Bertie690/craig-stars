import type { Vector } from './Vector';

export interface MapObject {
	id?: number;
	createdAt?: string;
	updatedAt?: string;
	type: MapObjectType;
	gameId?: number;
	position: Vector;
	name: string;
	num: number;
	playerNum: number | undefined;
}

export enum MapObjectType {
	None = '',
	Planet = 'Planet',
	Fleet = 'Fleet',
	Wormhole = 'Wormhole',
	MineField = 'MineField',
	MysteryTrader = 'MysteryTrader',
	Salvage = 'Salvage',
	MineralPacket = 'MineralPacket'
}

/**
 * Check if this MapObject is owned by a player
 * @param mo The MapObject to check
 * @param playerNum The player
 * @returns true if this mapobject is owned by the player
 */
export function ownedBy(mo: MapObject, playerNum: number): boolean {
	return mo.playerNum === playerNum;
}

/**
 * Check if this MapObject is owned by any player
 * @param mo The MapObject to check
 * @returns true if this mapobject is owned
 */
export function owned(mo: MapObject): boolean {
	return mo.playerNum != null;
}

export const positionKey = (mo: MapObject) => `${mo.position.x},${mo.position.y}`;
