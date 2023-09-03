export type ToBrew = {
    id: string,
    name: string,
    done: boolean,
    time: Date,
    bean: string,
    data: unknown
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
