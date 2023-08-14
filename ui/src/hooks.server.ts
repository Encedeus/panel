import type { Handle } from "@sveltejs/kit";

export const handle = (async ({ event, resolve }) => {
    event.locals.isUserSignedIn = event.cookies.get("encedeus_refreshToken") !== undefined;

    return await resolve(event);
}) satisfies Handle;
