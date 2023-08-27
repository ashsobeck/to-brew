import { writable, type Writable } from "svelte/store";
import type { ToBrew } from "./types";


export const tobrews: Writable<ToBrew[]> = writable([]);
