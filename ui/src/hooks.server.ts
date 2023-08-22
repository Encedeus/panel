import type { Handle } from "@sveltejs/kit";

export const handle = (async ({ event, resolve }) => {
    (event.locals as any).isUserSignedIn = event.cookies.get("encedeus_refreshToken") !== undefined;

    return resolve(event);
}) satisfies Handle;
