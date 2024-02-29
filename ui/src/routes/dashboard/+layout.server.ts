import type { LayoutServerLoad } from "./$types";
import { api } from "$lib/services/api";
import type { ModulesFindAllRequest } from "@encedeus/js-api";

export const load: LayoutServerLoad = async () => {
    const resp = await api.modulesService.findAllModules({
        backendOnly: true,
        frontendOnly: true,
    } as ModulesFindAllRequest);

    return {
        modules: resp.response?.modules,
    };
};