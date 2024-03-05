import { convertBrew, type Brew, type ToBrew, type BrewedResponse } from "$lib/types";
import type { Actions } from "@sveltejs/kit";

export async function load() {
    // get data from api here
    // map it to ToBrew data type
    const res = await fetch(`http://localhost:3333/tobrews/`, {
        method: 'GET',
        headers: {
            Origin: 'http://localhost:5173/'
        }
    });
    // TODO: make this one call
    const beanRes = await fetch(`http://localhost:3333/beans/`, {
        method: 'GET',
        headers: {
            Origin: 'http://localhost:5173/'
        }
    });
    const brews = await res.json();
    const beans = await beanRes.json();

    return {
        brews: brews?.map((b: Brew) => {
            return convertBrew(b)
        }) ?? [],
        beans: beans ?? []
    }
}

export const actions: Actions = {
    create: async ({ cookies, request }) => {
        const data = await request.formData()
        console.log("creating...")
        console.log(data)
        const time = data.get('time') as string === '' ? new Date() : new Date(data.get('time') as string)
        console.log(time)
        console.log(data.get('weight') as unknown as number)
        const response = await fetch('http://localhost:3333/tobrews/', {
            method: 'POST',
            body: JSON.stringify({
                // name: data.get('name'),
                bean: data.get('bean'),
                // roaster: { String: data.get('roaster') as string, Valid: true },
                // link: { String: data.get('link') as string, Valid: true },
                weight: parseFloat(data.get('weight') as string),
                brewed: false,
                timeToBrew: new Date().toISOString()
            }),
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        });

        const brew = (await response.json()) as Brew;
        console.log(brew);

    },
    update: async ({ cookies, request }): Promise<{ success: boolean; brew: ToBrew; }> => {
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
                // name: data.get('name')?.toString(),
                // roaster: { String: data.get('roaster')?.toString(), Valid: true },
                // link: { String: data.get('link')?.toString(), Valid: true },
                bean: data.get('bean') as string,
                brewed: !strToBool(data.get('brewed')?.toString()),
                timeToBrew: time
            }),
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        });
        const brew = (await response.json()) as Brew;
        console.log("updated")
        console.log(brew)
        console.log(convertBrew(brew))
        return { success: brew !== null, brew: convertBrew(brew) }
    },
    brewed: async ({ cookies, request }): Promise<{ success: boolean; brew: ToBrew, newWeight: number; }> => {
        const data = await request.formData();
        console.log(data)
        const time = data.get('time')?.toString() ?? new Date().toISOString()
        console.log(time)
        const response = await fetch(`http://localhost:3333/tobrews/complete/${data.get('id')?.toString()}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                Origin: 'http://localhost:5173/'
            }
        });
        console.log("response:")
        console.log(response);
        const brew = (await response.json()) as BrewedResponse
        console.log("marked as brewed:")
        console.log(brew)
        console.log(convertBrew(brew.brew))
        return { success: brew !== null, brew: convertBrew(brew.brew), newWeight: brew.newWeight }
    }
}
