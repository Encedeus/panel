<script lang="ts">
    import { onMount } from "svelte";
    import { getSignedInUser } from "$lib/services/auth_service";

    onMount(async () => {
        await getSignedInUser();
    });
</script>

<slot/>