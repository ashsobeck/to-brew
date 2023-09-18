export type ToBrew = {
    id: string,
    name: string,
    brewed: boolean,
    time: Date,
    bean: string,
    roaster: string,
    link: string
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
