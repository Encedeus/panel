import { accessTokenStore, userStore } from "../store";
import { decodeJwt } from "jose";
import { api } from "./api";
import {
    User,
} from "@encedeus/js-api";
import { goto } from "$app/navigation";

export async function refreshAccessToken(): Promise<string | undefined> {
    const { accessToken, error } = await api.authService.refreshAccessToken();
    if (error) {
        await signOut();
        return;
    }

    if (accessToken) {
        saveAccessToken(accessToken);
    }
    return accessToken;
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
    let user = new User();
    const unsubscribe = userStore.subscribe(v => user = v);
    if (user) {
        return user;
    }

    const accessToken = await getAccessToken();
    if (!accessToken) {
        await signOut();
    }

    const tokenPayload = decodeJwt(accessToken!);
    const userId = tokenPayload.userId as string;

    const resp = await api.usersService.findUserById(userId);
    if (resp.error) {
        await signOut();
    }
    userStore.set(resp.user!);

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
