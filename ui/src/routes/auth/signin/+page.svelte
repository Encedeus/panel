<script lang="ts">
    import Input from "$lib/components/generic/Input.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import AuthCard from "$lib/components/generic/AuthCard.svelte";
    import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";
    import { api } from "$lib/services/api";
    import Toast from "$lib/components/generic/Toast.svelte";
    import type { HttpError, SignInUserResponse } from "@encedeus/js-api";
    import { isBadRequestError, isWrongEmailOrUsernameError, isWrongPasswordError } from "@encedeus/js-api";
    import Button from "$lib/components/generic/Button.svelte";
    import { saveAccessToken } from "$lib/services/auth_service";
    import { goto } from "$app/navigation";

    let uid = "";
    let password = "";
    let errorLabel: string | null = null;

    let passwordError = false;
    let uidError = false;

    async function signIn() {
        const {error, accessToken} = await sendAuthenticationRequest(uid, password);
        checkForErrors(error);
        if (error) {
            return;
        }
        saveAccessToken(accessToken as string);
        await goto("/dashboard/servers");
    }

    async function sendAuthenticationRequest(uid: string, password: string): Promise<SignInUserResponse> {
        return await api.authService.signInUser({
            uid,
            password,
        });
    }

    function checkForErrors(error: HttpError | null | undefined) {
        if (!error) {
            return;
        }
        errorLabel = error.message;

        if (isWrongEmailOrUsernameError(error) || isBadRequestError(error)) {
            uidError = true;
        }
        if (isWrongPasswordError(error)) {
            passwordError = true;
        }
    }

    function clearError() {
        if (errorLabel) {
            errorLabel = "";
            passwordError = false;
            uidError = false;
        }
    }

</script>

<div class="overflow-hidden">
    <aside class="absolute top-0 right-0 mt-5 mr-7">
        <span class="drop-shadow-xl text-white text-sm font-bold tracking-wide">Don't have an account?&nbsp; • &nbsp;<a
                class="text-indigo-600" href="/auth/signup">Sign Up&nbsp;<SmallArrowRight/></a></span>
    </aside>

    <div class="w-screen h-screen bg-image"></div>

    <main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <AuthCard height="[16rem]">
            <CardHeader slot="title">
                Sign In
            </CardHeader>
            <div class="flex flex-col gap-5" slot="inputs">
                <Input bind:value={uid} error={uidError} label="Username/E-Mail" on:input={clearError}
                       placeholder="Enter Username or E-Mail"
                       size="lg"/>
                <Input bind:value={password} error={passwordError} label="Password" on:input={clearError}
                       placeholder="Enter Password" size="lg"
                       type="password"/>
            </div>
            <Button on:click={async () => await signIn()} slot="button">Sign In</Button>
        </AuthCard>
    </main>

    {#if errorLabel !== null}
        <aside class="absolute left-10 {errorLabel ? 'come-up-animation' : 'come-down-animation'}">
            <Toast mode="error" size="md">
                {errorLabel}
            </Toast>
        </aside>
    {/if}
</div>


<style lang="postcss">
    .bg-image {
        background: url("$lib/assets/auth-bg.svg");
        background-size: contain;
    }

    :global(body) {
        overflow: hidden;
    }

    :root {
        --animation-delay: 0.25s;
    }

    @keyframes come-up {
        from {
            @apply -bottom-16;
        }
        to {
            @apply bottom-10 block;
        }
    }

    @keyframes come-down {
        from {
            @apply bottom-10;
        }
        to {
            @apply -bottom-16 hidden;
        }
    }

    .come-up-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-up;
        animation-fill-mode: forwards;
    }

    .come-down-animation {
        animation-duration: var(--animation-delay);
        animation-name: come-down;
        animation-fill-mode: forwards;
    }
</style>