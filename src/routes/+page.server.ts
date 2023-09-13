import type { Brew, ToBrew } from "$lib/types";
import type { Actions } from "@sveltejs/kit";

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
        brews: brews?.map((b: Brew) => {
            console.log(b);
            return {
                id: b.id,
                name: b.name,
                brewed: b.brewed ?? false,
                bean: b.bean,
                time: b.timeToBrew
            }
        }) ?? []
    }
}

export const actions: Actions = {
    create: async ({ cookies, request }) => {
        const data = await request.formData()
        const response = await fetch('http://localhost:3333/tobrews/', {
            method: 'POST',
            body: JSON.stringify({
                name: data.get('name'),
                roaster: { String: data.get('roaster'), Valid: true },
                link: { String: data.get('link'), Valid: true },
                brewed: false,
                timeToBrew: data.get('time')
            }),
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        });

        const brew = (await response.json()) as Brew;
        console.log(brew);

    },
}
