import {api} from "$lib/services/api";
import {error} from "@sveltejs/kit";
import type {ServersListData} from "$lib/interfaces/serversListData";

export async function load() {
    const resp = await api.serversService.findAllServers();

    if (resp.error) {
        error(500, resp.error.message);
        return;
    }
    if (!resp.response?.servers) {
        return {
            servers: []
        } as ServersListData;
    }

    return {
        servers: resp.response?.servers
    } as ServersListData;
}