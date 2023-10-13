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
    return {
        brews: brews?.map((b: Brew) => {
            return {
                id: b.id,
                name: b.name,
                brewed: b.brewed ?? false,
                bean: b.bean,
                link: b.link.String,
                roaster: b.roaster.String,
                time: b.timeToBrew
            }
        }) ?? []
    }
}

export const actions: Actions = {
    create: async ({ cookies, request }) => {
        const data = await request.formData()
        console.log("creating...")
        console.log(data)
        const response = await fetch('http://localhost:3333/tobrews/', {
            method: 'POST',
            body: JSON.stringify({
                name: data.get('name'),
                bean: data.get('bean'),
                roaster: { String: data.get('roaster') as string, Valid: true },
                link: { String: data.get('link') as string, Valid: true },
                brewed: false,
                timeToBrew: new Date(data.get('time') as string).toISOString()
            }),
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        });

        const brew = (await response.json()) as Brew;
        console.log(brew);

    },
    brewed: async ({ cookies, request }) => {
        const data = await request.formData();
        console.log(data)
        const time = data.get('time')?.toString() ?? new Date().toISOString()
        console.log(time)
        const strToBool = (str: string | undefined) => {
            if (str === "true")
                return true;
            return false;
        }
        const response = await fetch(`http://localhost:3333/tobrews/${data.get('id')?.toString()}`, {
            method: 'PUT',
            body: JSON.stringify({
                name: data.get('name')?.toString(),
                roaster: { String: data.get('roaster')?.toString(), Valid: true },
                link: { String: data.get('link')?.toString(), Valid: true },
                brewed: !strToBool(data.get('brewed')?.toString()),
                timeToBrew: time
            }),
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        })
        const brew = (await response.json()) as Brew;
        console.log("marked as brewed:")
        console.log(brew)
    }
}
