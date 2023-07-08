<script lang="ts">
  import Input from "$lib/components/Input.svelte";
  import CardHeader from "$lib/components/CardHeader.svelte";
  import AuthCard from "$lib/components/AuthCard.svelte";
  import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";
  import { api } from "../../../lib/services/api_service";
  import { isEmailValid } from "../../../lib/services/validation_service";
  import Toast from "$lib/components/Toast.svelte";
  import { LoginUserErrors } from "@encedeus/js-api";

  let name = "";
  let password = "";

  let errorLabel = "Wrong Password";

  function signIn() {
    if (isEmailValid(name)) {
      api.authService.loginUser({
        email: name,
        password: password,
      });
    }

    api.authService.loginUser({
      username: name,
      password: password,
    });
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
      <Input bind:value={name} error="gaga" placeholder="Enter Username or E-Mail" size="lg" label="Username/E-Mail"/>
      <Input bind:value={password} error="gaghahgas" placeholder="Enter Password" size="lg" label="Password" type="password"/>
    </div>
  </AuthCard>
</main>

{#if errorLabel}
  <aside class="absolute bottom-10 left-10 come-up-animation">
    <Toast mode="error" size="md">
      <p slot="label">{errorLabel}</p>
    </Toast>
  </aside>
{/if}

<style lang="postcss">
  @keyframes come-up {
    from {
      @apply -bottom-16;
    }
    to {
      @apply bottom-10;
    }
  }

  .come-up-animation {
    animation-duration: 0.75s;
    animation-name: come-up;
  }
</style>