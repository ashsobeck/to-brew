import type { ToBrew } from "$lib/types";

type NullString = {
    String: string;
    Valid: false;
};

type Brew = ToBrew & {
    link: NullString;
    roaster: NullString;
    created: string;
    timeToBrew: string;
}
export async function load() {
    // get data from api here
    // map it to ToBrew data type
    const res = await fetch(`http://localhost:3333/tobrews/`, { method: 'GET' });
    const brews = await res.json();
    console.log(brews);
    return {
        brews: brews.map((b: Brew) => {
            console.log(b);
            return {
                id: b.id,
                name: b.name,
                done: false,
                bean: b.bean,
                time: b.timeToBrew
            }
        })
    }
}
