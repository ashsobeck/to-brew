export type ToBrew = {
	id: string,
	name: string,
	brewed: boolean,
	timeToBrew: string,
	bean: string,
	roaster: string,
	link: string,
	created: string,
}

export type Brew = ToBrew & {
	link: NullString;
	roaster: NullString;
	created: string;
	timeToBrew: string;
}

export type NullString = {
	String: string;
	Valid: boolean;
};

export const convertBrew = (brew: Brew): ToBrew => {
	return ({
		...brew,
		link: brew.link.String,
		roaster: brew.roaster.String
	});
};

// <!-- Id string `json:"id"` -->
// <!-- Name     string `json:"name"` -->
// <!-- Roaster  string `json:"roaster"` -->
// <!-- Country  string `json:"country"` -->
// <!-- Varietal string `json:"varietal"` -->
// <!-- Process  string `json:"process"` -->
// <!-- Altitude string `json:"altitude"` -->
// <!-- Notes    string `json:"notes"` -->
// <!-- Weight   float32 `json:"weight"` -->

export type Bean = {
	id: string;
	name: string;
	roaster: string;
	country: string;
	varietal: string;
	process: string;
	altitude: string;
	notes: string;
	weight: number;
}
