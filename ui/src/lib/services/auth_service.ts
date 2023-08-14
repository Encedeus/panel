import { accessTokenStore } from "../store";
import { decodeJwt } from "jose";
import { api } from "./api_service";
import {
    User,
} from "@encedeus/js-api";
import { goto } from "$app/navigation";

export async function refreshAccessToken(): Promise<string> {
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

export async function getAccessToken(): Promise<string> {
    let accessToken = "";
    accessTokenStore.subscribe(token => (accessToken = token))();
    if (!accessToken) {
        return await refreshAccessToken();
    }

    const payload = decodeJwt(accessToken);
    if (Date.now() >= payload.exp * 1000) {
        return await refreshAccessToken();
    }

    return null;
}

export async function isUserSignedIn(): Promise<boolean> {
    return !(await getAccessToken());
}

export async function getSignedInUser(): Promise<User> {
    const accessToken = await getAccessToken();
    if (!accessToken) {
        return null;
    }

    const tokenPayload = decodeJwt(accessToken);
    const resp = await api.usersService.findUserById(tokenPayload.userId as string);
    if (resp.error) {
        await signOut();
    }

    return resp.user;
}

export function saveAccessToken(accessToken: string) {
    accessTokenStore.set(accessToken);
    api.accessToken = accessToken;
}

export async function signOut() {
    saveAccessToken("");
    await api.authService.signOut();
    await goto("/auth/signin");
}
