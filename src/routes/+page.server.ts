import type { Brew } from "$lib/types";

export async function load() {
    // get data from api here
    // map it to ToBrew data type
    const res = await fetch(`http://localhost:3333/tobrews/`, {
        method: 'GET', headers: {
            'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, PATCH, OPTIONS',
            'Access-Control-Allow-Origin': 'http://localhost',
            'Access-Control-Allow-Headers': '*'
        }
    });
    const brews = await res.json();
    console.log(brews);
    return {
        brews: brews.map((b: Brew) => {
            console.log(b);
            return {
                id: b.id,
                name: b.name,
                brewed: b.brewed ?? false,
                bean: b.bean,
                time: b.timeToBrew
            }
        })
    }
}
