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
