import { readable, writable } from "svelte/store";
import type { User } from "@encedeus/js-api";

export const accessTokenStore = writable("");
export const userStore = writable<User | undefined>();
export const api = readable();

export const fileManagerHistory = writable<Array<string>>([""]);

export const fileManagerPathIndex = writable(0);