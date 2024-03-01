import {api} from "$lib/services/api";
import {error} from "@sveltejs/kit";
import type {RoleListData} from "$lib/interfaces/roleListData";
import type {Role} from "@encedeus/js-api";

export async function load() {
    const resp = await api.roleService.findAllRoles();
    console.log("token", api.accessToken);
    if (resp.error) {
        error(500, resp.error.message);
        return {
            rolesData: []
        } as RoleListData;
    }

    if (!resp.response?.roles) {
        return {
            rolesData: []
        } as RoleListData;
    }

    const roles: Role[] = resp.response.roles;
    const usersWithRoles: string[][] = [];

    for (const role of roles) {
        const resp = await api.usersService.findAllWithRole(<string>role.id?.value);

        if (resp.error) {
            error(500, resp.error.message);
            return;
        }

        const usernames: string[] = [];

        resp.response!.users.forEach((v) => {
            usernames.push(v.name);
        });

        usersWithRoles.push(usernames);
    }

    const rolesData: Array<{ role: Role, userList: string[] }> = [];

    for (let i = 0; i < roles?.length; i++) {
        rolesData.push({
            role: roles[i],
            userList: usersWithRoles[i]
        });
    }

    return {
        rolesData: rolesData
    } as RoleListData;
}