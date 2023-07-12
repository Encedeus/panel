<script lang="ts">
  import "../app.css";
  import { isUserSignedIn } from "../lib/services/auth_service";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  onMount(async () => {
    console.log(await isUserSignedIn());
    if (!(await isUserSignedIn())) {
      await goto("/auth/signin");
    }
  });
</script>

<slot/>