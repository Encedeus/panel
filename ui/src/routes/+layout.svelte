<script lang="ts">
  import "../app.css";
  import { isUserSignedIn } from "../lib/services/auth_service";
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  onMount(() => {
    if (!isUserSignedIn()) {
      goto("/auth/signin");
    }
  });
</script>

<slot/>