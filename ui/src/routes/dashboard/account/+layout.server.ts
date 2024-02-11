import type { LayoutServerLoad } from "./$types";
import { api } from "$lib/services/api";
import { ModulesFindAllRequest, ModulesFindAllResponse } from "@encedeus/js-api";

export const load: LayoutServerLoad = async () => {
    const resp = await api.modulesService.findAllModules(ModulesFindAllRequest.create({
        backendOnly: true,
        frontendOnly: true,
    }));

    return {
        modules: resp.response?.modules,
    };
};