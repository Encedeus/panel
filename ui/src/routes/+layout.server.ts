import type { LayoutLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

export const load = (({ locals, route }) => {
    if (!locals.isUserSignedIn && !route.id.startsWith("/auth")) {
        throw redirect(307, "/auth/signin");
    }
}) satisfies LayoutLoad;
