import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import { api } from "$lib/services/api";
import { UserFindOneRequest, UUID } from "@encedeus/js-api";

export const load: LayoutServerLoad = async ({ locals, route }) => {
    if (!(locals as any)?.isUserSignedIn) {
        if (!route.id?.includes("/auth")) {
            throw redirect(307, "/auth/signin");
        }
    } else if (!route.id?.includes("/dashboard")) {
        throw redirect(307, "/dashboard/servers");
    }



    return {
        user: (await api.usersService.findUserById(UserFindOneRequest.create({
            userId: UUID.create({
                value: (locals as any).userId,
            }),
        }))).response?.user,
    };
};
