import type { Handle } from "@sveltejs/kit";
import { getUserIdFromToken } from "$lib/services/auth_service";

export const handle: Handle = async ({ event, resolve }) => {
    const token = event.cookies.get("encedeus_refreshToken");
    const isUserSignedIn = token !== undefined;

    (event.locals as any).isUserSignedIn = event.cookies.get("encedeus_refreshToken") !== undefined;
    if (isUserSignedIn) {
        (event.locals as any).userId = getUserIdFromToken(token);
    }

    return resolve(event);
};
