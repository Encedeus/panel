<script lang="ts">
    import Input from "$lib/components/generic/Input.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import AuthCard from "$lib/components/generic/AuthCard.svelte";
    import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";
    import { api } from "../../../lib/services/api_service";
    import { isEmailValid } from "../../../lib/services/validation_service";
    import Toast from "$lib/components/generic/Toast.svelte";
    import type { LoginUserResponse } from "@encedeus/js-api";
    import { LoginUserErrors } from "@encedeus/js-api";
    import Button from "$lib/components/generic/Button.svelte";
    import { saveAccessToken } from "../../../lib/services/auth_service";
    import { goto } from "$app/navigation";

    let name = "";
    let password = "";

    let errorLabel = "";
    let usernameError = false;
    let passwordError = false;

    async function signIn() {
        const { error, accessToken } = await sendAuthenticationRequest(name, password);
        checkForErrors(error);
        if (error) {
            signIn.called = true;
            return;
        }
        saveAccessToken(accessToken);
        await goto("/dashboard/servers");
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
            return;
        case LoginUserErrors.WRONG_EMAIL_OR_USERNAME as LoginUserErrors:
            errorLabel = "Wrong Email or Username";
            usernameError = true;
            return;
        case LoginUserErrors.USERNAME_OR_EMAIL_NOT_SPECIFIED as LoginUserErrors:
            errorLabel = "Email or Username not specified";
            usernameError = true;
            return;
        case LoginUserErrors.INTERNAL_SERVER_ERROR as LoginUserErrors:
            errorLabel = "Internal Server Error";
            return;
        }
    }

</script>

<aside class="absolute top-0 right-0 mt-5 mr-7">
    <span class="drop-shadow-xl text-white text-sm font-bold tracking-wide">Don't have an account?&nbsp; • &nbsp;<a href="/auth/signup" class="text-indigo-600">Sign Up&nbsp;<SmallArrowRight/></a></span>
</aside>

<div class="w-screen h-screen bg-image"></div>

<main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
    <AuthCard height="[16rem]" buttonLabel="Sign In">
        <CardHeader slot="title">
            Sign In
        </CardHeader>
        <div class="flex flex-col gap-5" slot="inputs">
            <Input on:input={() => {
                if(usernameError) {
                    usernameError = false;
                }
            }} bind:error={usernameError} bind:value={name} placeholder="Enter Username or E-Mail" size="lg" label="Username/E-Mail"/>
            <Input on:input={() => {
                if(passwordError) {
                    passwordError = false;
                }
            }} bind:error={passwordError} bind:value={password} placeholder="Enter Password" size="lg" label="Password" type="password"/>
        </div>
        <Button on:click={async () => await signIn()} slot="button">Sign In</Button>
    </AuthCard>
</main>

{#if signIn.called}
    <aside class="absolute left-10 {(passwordError || usernameError) ? 'come-up-animation bottom-10' : 'come-down-animation -bottom-16'}">
        <Toast mode="error" size="md">
            <p slot="label">{errorLabel}</p>
        </Toast>
    </aside>
{/if}

<style lang="postcss">
    .bg-image {
        background: url("$lib/assets/auth-bg.svg");
        background-size: 450%;
    }

    :root {
        --animation-delay: 0.25s;
    }

    @keyframes come-up {
        from {
            @apply -bottom-16;
        }
        to {
            @apply bottom-10;
        }
    }

    @keyframes come-down {
        from {
            @apply bottom-10;
        }
        to {
            @apply -bottom-16;
        }
    }

    .come-up-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-up;
    }

    .come-down-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-down;
    }
</style>