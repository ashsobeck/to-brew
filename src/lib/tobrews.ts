import { writable } from "svelte/store";
import type { ToBrew } from "./types";

export function makeStore() {
    const initToBrew: ToBrew = {
        id: crypto.randomUUID(),
        done: false,
        time: new Date(),
        bean: '',
        data: { coffee: 'Little Wolf Mystery' }
    };
    const tobrews: ToBrew[] = [initToBrew];
    const { subscribe, update } = writable(tobrews);

    return {
        subscribe,
        add: (brew: ToBrew) => {
            update(($tobrews) => [...$tobrews, brew]);
        },
        remove: (id: string) => {
            update(($tobrews) => $tobrews.filter((tobrew) => tobrew.id != id));
        },
        mark: (tobrew: ToBrew, done: boolean) => {
            update(($tobrews) => [...$tobrews.filter((t) => t !== tobrew), { ...tobrew, done }]);
        }
    };
}
