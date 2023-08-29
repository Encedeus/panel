import { accessTokenStore, userStore } from "../store";
import { decodeJwt } from "jose";
import { api } from "./api";
import { goto } from "$app/navigation";
import { User, UUID, UserFindOneRequest } from "@encedeus/js-api";

export async function refreshAccessToken(): Promise<string | undefined> {
    const { response, error } = await api.authService.refreshAccessToken();
    if (error) {
        await signOut();
        return;
    }

    if (response?.accessToken) {
        saveAccessToken(response.accessToken);
        return response.accessToken;
    }

    return;
}

export async function getAccessToken(): Promise<string | undefined> {
    let accessToken = "";
    accessTokenStore.subscribe(token => (accessToken = token))();
    if (!accessToken) {
        return await refreshAccessToken();
    }

    const payload = decodeJwt(accessToken);

    if (Date.now() >= payload.exp! * 1000) {
        return await refreshAccessToken();
    }

    return accessToken;
}

export async function isUserSignedIn(): Promise<boolean> {
    return !(await getAccessToken());
}

export async function getSignedInUser(): Promise<User> {
    let user = User.create();
    const unsubscribe = userStore.subscribe(v => user = v);
    if (user) {
        return user;
    }

    const accessToken = await getAccessToken();
    if (!accessToken) {
        await signOut();
    }

    console.log(accessToken);
    const tokenPayload = decodeJwt(accessToken!);
    const userId = UUID.create({
        value: (tokenPayload.user_id as any).value as string,
    });

    const { response, error } = await api.usersService.findUserById(UserFindOneRequest.create({
        userId,
    }));
    if (error) {
        await signOut();
    }
    userStore.set(response?.user!);

    unsubscribe();
    return user;
}

export function saveAccessToken(accessToken: string) {
    accessTokenStore.set(accessToken);
    api.accessToken = accessToken;
}

export async function signOut(): Promise<void> {
    saveAccessToken("");
    await api.authService.signOut();
    await goto("/auth/signin");
}
