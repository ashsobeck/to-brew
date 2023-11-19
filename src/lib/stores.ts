import { writable, type Writable } from "svelte/store";
import type { Bean, ToBrew } from "./types";

export const tobrews: Writable<ToBrew[]> = writable([]);
export const beans: Writable<Bean[]> = writable([]);
