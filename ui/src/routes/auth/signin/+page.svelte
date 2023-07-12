<script lang="ts">
  import Input from "$lib/components/Input.svelte";
  import CardHeader from "$lib/components/CardHeader.svelte";
  import AuthCard from "$lib/components/AuthCard.svelte";
  import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";
  import { api } from "../../../lib/services/api_service";
  import { isEmailValid } from "../../../lib/services/validation_service";
  import Toast from "$lib/components/Toast.svelte";
  import type { LoginUserResponse } from "@encedeus/js-api";
  import { LoginUserErrors } from "@encedeus/js-api";
  import Button from "$lib/components/Button.svelte";
  import {
    saveRefreshToken,
    saveAccessToken,
  } from "../../../lib/services/auth_service";

  let name = "";
  let password = "";

  let errorLabel = "";
  let usernameError = false;
  let passwordError = false;

  let alert: HTMLElement;

  async function signIn() {
    const { error, accessToken, refreshToken } = await sendAuthenticationRequest(name, password);

    checkForErrors(error);
    saveTokens(accessToken, refreshToken);
    signIn.called = true;

    console.log(errorLabel);
  }

  async function sendAuthenticationRequest(name: string, password: string): Promise<LoginUserResponse> {
    let resp: LoginUserResponse;
    if (isEmailValid(name)) {
      resp = await api.authService.loginUser({
        email: name,
        password: password,
      });
    } else {
      resp = await api.authService.loginUser({
        username: name,
        password: password,
      });
    }

    return resp;
  }

  function saveTokens(accessToken: string, refreshToken: string) {
    saveRefreshToken(refreshToken);
    saveAccessToken(accessToken);
  }

  function checkForErrors(error: LoginUserErrors) {
    if(error) {
      usernameError = false;
      passwordError = false;
      errorLabel = "";
    }

    switch (LoginUserErrors[error]) {
      case LoginUserErrors.WRONG_PASSWORD as LoginUserErrors:
        errorLabel = "Wrong Password";
        passwordError = true;
        break;
      case LoginUserErrors.WRONG_EMAIL_OR_USERNAME as LoginUserErrors:
        errorLabel = "Wrong Email or Username";
        usernameError = true;
        break;
      case LoginUserErrors.USERNAME_OR_EMAIL_NOT_SPECIFIED as LoginUserErrors:
        errorLabel = "Email or Username not specified";
        usernameError = true;
        break;
      case LoginUserErrors.INTERNAL_SERVER_ERROR as LoginUserErrors:
        errorLabel = "Internal Server Error";
        break;
    }
  }

  function gracefulAlertShutdown() {
    setTimeout(() => {
      alert.classList.add("hidden");
      errorLabel = "";
    }, 240)
  }
</script>

<aside class="absolute top-0 right-0 mt-5 mr-7">
  <span class="drop-shadow-xl text-white text-sm font-bold tracking-wide">Don't have an account?&nbsp; • &nbsp;<a href="/auth/signup" class="text-indigo-600">Sign Up&nbsp;<SmallArrowRight/></a></span>
</aside>

<main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
  <AuthCard height="[16rem]" buttonLabel="Sign In">
    <CardHeader slot="title">
      Sign In
    </CardHeader>
    <div class="flex flex-col gap-5" slot="inputs">
      <Input on:input={() => {
        if(usernameError) {
          usernameError = false;
          gracefulAlertShutdown();        }
       }} bind:error={usernameError} bind:value={name} placeholder="Enter Username or E-Mail" size="lg" label="Username/E-Mail"/>
      <Input on:input={() => {
        if(passwordError) {
          passwordError = false;
          gracefulAlertShutdown();
        }
      }} bind:error={passwordError} bind:value={password} placeholder="Enter Password" size="lg" label="Password" type="password"/>
    </div>
    <Button on:click={async () => await signIn()} slot="button">Sign In</Button>
  </AuthCard>
</main>

<aside bind:this={alert} class="absolute bottom-10 left-10 {signIn.called ? '' : 'hidden'} {(passwordError || usernameError) ? 'come-up-animation' : 'come-down-animation'}">
  <Toast mode="error" size="md">
    <p slot="label">{errorLabel}</p>
  </Toast>
</aside>

<style lang="postcss">
  :global(body) {
    overflow: hidden;
  }

  @keyframes come-up {
    from {
      @apply -bottom-16 block;
    }
    to {
      @apply bottom-10;
    }
  }

  @keyframes come-down {
    from {
      @apply bottom-10 block;
    }
    to {
      @apply -bottom-16;
    }
  }

  .come-up-animation {
    animation-duration: 0.5s;
    animation-name: come-up;
  }

  .come-down-animation {
    animation-duration: 0.25s;
    animation-name: come-down;
  }
</style>