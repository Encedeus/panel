<!--
<script lang="ts">
  import Input from "$lib/components/Input.svelte";
  import CardHeader from "$lib/components/CardHeader.svelte";
  import AuthCard from "$lib/components/AuthCard.svelte";
  import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";
  import { api } from "../../../lib/services/api_service";
  import { User } from "@encedeus/js-api";

  let username = "";
  let email = "";
  let password = "";
  let confirmPassword = "";

/*  function signUp() {


    api.usersService.createUser({
      name: username,
      email,
      password, roleId: 0, roleName: ""
    });
  }*/
</script>

<aside class="absolute top-0 right-0 mt-5 mr-7">
  <span class="text-white text-sm font-bold tracking-wide">Already have an account?&nbsp; • &nbsp;<a href="/auth/signin" class="text-indigo-600">Sign In&nbsp;<SmallArrowRight/></a></span>
</aside>
<main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
  <AuthCard height="[16rem]" buttonLabel="Sign Up">
    <CardHeader slot="title">
      Sign Up
    </CardHeader>
    <div class="flex flex-col gap-5" slot="inputs">
      <Input bind:value={username} placeholder="Enter Username" size="lg" label="Username"/>
      <Input bind:value={email} placeholder="Enter E-Mail" size="lg" label="E-Mail"/>
      <Input bind:value={password} placeholder="Enter Password" size="lg" label="Password" type="password"/>
      <Input bind:value={confirmPassword} placeholder="Enter Password" size="lg" label="Confirm Password" type="password"/>
    </div>
  </AuthCard>
</main>
-->
