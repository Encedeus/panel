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
  const resp = await api.authService.refreshAccessToken(getRefreshToken());
  if (resp.error !== RefreshAccessTokenErrors.OK) {
    signOut();
    return;
  }

  accessTokenStore.set(resp.accessToken);
  return resp.accessToken || "";
}

export function getAccessToken(): string {
  let accessToken = "";
  const unsubscribe = accessTokenStore.subscribe(
    token => (accessToken = token),
  );
  if (!accessToken) {
    if (getRefreshToken()) {
      refreshAccessToken().then(v => (accessToken = v));
    }
    return accessToken;
  }

  const payload = decodeJwt(accessToken);
  if (Date.now() >= payload.exp * 1000) {
    refreshAccessToken().then(v => (accessToken = v));
  }

  unsubscribe();
  return accessToken;
}

export function getRefreshToken(): string {
  const refreshToken: string = localStorage.getItem("encedeus_refreshToken");
  if (!refreshToken) {
    signOut();
  }

  return refreshToken;
}

export function saveRefreshToken(refreshToken: string) {
  localStorage.setItem("encedeus_refreshToken", refreshToken);
}

export async function isUserSignedIn(): Promise<boolean> {
  if (!getAccessToken()) {
    return false;
  }

  return true;
}

export async function getSignedInUser(): Promise<User> {
  const accessToken = getAccessToken();
  if (!accessToken) {
    return null;
  }
  const tokenPayload = decodeJwt(accessToken);

  const resp = await api.usersService.getUserById(<string>tokenPayload.userId);
  if (resp.error && resp.error !== GetUserErrors.OK) {
    signOut();
  }

  return resp.user;
}

export function saveAccessToken(accessToken: string) {
  accessTokenStore.set(accessToken);
  api.accessToken = accessToken;
}

export function signOut() {
  saveAccessToken("");
  saveRefreshToken("");
  goto("/auth/signin");
}
