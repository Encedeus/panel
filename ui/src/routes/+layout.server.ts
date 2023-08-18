import type { LayoutLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load = (({ locals, route }) => {
    if (!locals?.isUserSignedIn) {
        if (!route.id.includes("/auth")) {
            throw redirect(307, "/auth/signin");
        }
    } else if (!route.id.includes("/dashboard")) {
        throw redirect(307, "/dashboard/servers");
    }
}) satisfies LayoutLoad;
