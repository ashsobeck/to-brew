import { writable, type Writable } from "svelte/store";
import type { ToBrew } from "./types";

export function makeStore(): Writable<ToBrew> {
    let uid = 1;
    const initToBrew: ToBrew = {
        id: uid++,
        done: false,
        time: new Date(),
        data: { coffee: 'Little Wolf Mystery' }
    };
    const tobrews: ToBrew[] = [initToBrew];
    const { subscribe, update } = writable(tobrews);

    return {
        subscribe,
        add: (data: unknown) => {
            const tobrew: ToBrew = {
                id: uid++,
                done: false,
                time: new Date(),
                data: data
            };
            update(($tobrews) => [...$tobrews, tobrew]);
        },
        remove: (id: number) => {
            update(($tobrews) => $tobrews.filter((tobrew) => tobrew.id != id));
        },
        mark: (tobrew: ToBrew, done: boolean) => {
            update(($tobrews) => [...$tobrews.filter((t) => t !== tobrew), { ...tobrew, done }]);
        }
    };
}
