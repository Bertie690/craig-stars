import type {
	TechStore,
	TechEngine,
	TechPlanetaryScanner,
	TechTerraform,
	TechDefense,
	TechHullComponent,
	TechHull,
	Tech,
	TechPlanetary
} from '$lib/types/Tech';
import techjson from '$lib/ssr/techs.json';
import { kebabCase } from 'lodash-es';

export class TechService implements TechStore {
	techs: Tech[] = [];
	engines: TechEngine[] = [];
	planetaryScanners: TechPlanetaryScanner[] = [];
	terraforms: TechTerraform[] = [];
	defenses: TechDefense[] = [];
	planetaries: TechPlanetary[] = [];
	hullComponents: TechHullComponent[] = [];
	hulls: TechHull[] = [];

	techsByName: Map<string, Tech> = new Map();
	hullsByName: Map<string, TechHull> = new Map();
	hullComponentsByName: Map<string, TechHullComponent> = new Map();

	constructor(store?: TechStore) {
		store = store ?? (techjson as TechStore);
		this.buildMaps(store);
	}

	async fetch() {
		const response = await fetch(`/api/techs`, {
			method: 'GET',
			headers: {
				accept: 'application/json'
			}
		});
		if (response.ok) {
			const store = (await response.json()) as TechStore;
			this.buildMaps(store);
		} else {
			console.error(response);
		}
	}

	private buildMaps(store: TechStore) {
		this.engines = store.engines ?? [];
		this.planetaryScanners = store.planetaryScanners ?? [];
		this.terraforms = store.terraforms ?? [];
		this.defenses = store.defenses ?? [];
		this.planetaries = store.planetaries ?? [];
		this.hullComponents = (store.engines ?? []).concat(store.hullComponents ?? []);
		this.hulls = store.hulls ?? [];

		this.techs = [];
		this.techs = this.techs.concat(store.engines);
		this.techs = this.techs.concat(store.planetaryScanners);
		this.techs = this.techs.concat(store.defenses);
		this.techs = this.techs.concat(store.planetaries);
		this.techs = this.techs.concat(store.hullComponents);
		this.techs = this.techs.concat(store.hulls);
		this.techs = this.techs.concat(store.terraforms);

		this.techsByName = new Map(this.techs.map((t) => [kebabCase(t.name), t]));
		this.hullComponentsByName = new Map(
			this.hullComponents.concat(this.engines).map((t) => [kebabCase(t.name), t])
		);
		this.hullsByName = new Map(this.hulls.map((t) => [kebabCase(t.name), t]));
	}

	getTech(name: string): Tech | undefined {
		return this.techsByName.get(kebabCase(name));
	}

	getHull(name: string): TechHull | undefined {
		return this.hullsByName.get(kebabCase(name));
	}

	getHullComponent(name: string): TechHullComponent | undefined {
		return this.hullComponentsByName.get(kebabCase(name));
	}
}
