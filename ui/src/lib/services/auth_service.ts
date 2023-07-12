import { accessTokenStore } from "../store";
import { decodeJwt } from "jose";
import { api } from "./api_service";
import {
  GetUserErrors,
  RefreshAccessTokenErrors,
  User,
} from "@encedeus/js-api";
import { goto } from "$app/navigation";

export async function refreshAccessToken(): Promise<string> {
  const { accessToken, error } = await api.authService.refreshAccessToken(
    await getRefreshToken(),
  );
  if (error !== RefreshAccessTokenErrors.OK) {
    await signOut();
    return;
  }

  if (accessToken) {
    saveAccessToken(accessToken);
  }
  return accessToken || null;
}

export async function getAccessToken(): Promise<string> {
  let accessToken = "";
  accessTokenStore.subscribe(token => (accessToken = token))();
  if (!accessToken) {
    if (await getRefreshToken()) {
      return await refreshAccessToken();
    }

    await signOut();
    return null;
  }

  const payload = decodeJwt(accessToken);
  if (Date.now() >= payload.exp * 1000) {
    return await refreshAccessToken();
  }

  return null;
}

export async function getRefreshToken(): Promise<string> {
  const refreshToken: string = localStorage.getItem("encedeus_refreshToken");
  if (!refreshToken) {
    await signOut();
  }

  return refreshToken;
}

export function saveRefreshToken(refreshToken: string) {
  if (refreshToken) {
    localStorage.setItem("encedeus_refreshToken", refreshToken);
  }
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
  const resp = await api.usersService.getUserById(<string>tokenPayload.userId);
  if (resp.error && resp.error !== GetUserErrors.OK) {
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
  localStorage.removeItem("encedeus_refreshToken");
  await goto("/auth/signin");
}
