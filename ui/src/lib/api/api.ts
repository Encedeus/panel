import {get} from 'svelte/store';
import {EncedeusRegistryApi, type User} from "@encedeus/registry-js-api";


export function getApi(refreshCookie: string | undefined = undefined): EncedeusRegistryApi {
    return new EncedeusRegistryApi("http://localhost:3001", "", {
        axiosConfig: {
            withCredentials: true,
            headers: refreshCookie ? {Cookie: `encedeus_plugins_refreshToken=${refreshCookie};`} : {}
        },
        callbacks: {
            // eslint-disable-next-line @typescript-eslint/no-empty-function
            onAuth: () => {
            },
            // eslint-disable-next-line @typescript-eslint/no-empty-function
            onUser: () => {
            }
        }
    });
}