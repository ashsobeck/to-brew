import { writable, type Writable } from "svelte/store";
import type { ToBrew } from "./types";

const initToBrew: ToBrew = {
    id: crypto.randomUUID(),
    done: false,
    time: new Date(),
    bean: '',
    data: { coffee: 'Little Wolf Mystery' }
};

export const tobrews: Writable<ToBrew[]> = writable([initToBrew]);
