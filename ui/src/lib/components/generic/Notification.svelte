<script lang="ts">
    import Toast from "$lib/components/generic/Toast.svelte";
    import type { NotificationMode } from "$lib/components/generic/Toast.svelte";
    import { fly } from "svelte/transition";

    export let show = false;
    export let mode: NotificationMode;
    export let notification: string;
</script>
{#if show}
    <aside class="absolute left-10 bottom-6" in:fly={{ y: 50, opacity: 100, duration: 250, }}
           out:fly={{ y: 50, opacity: 0, duration: 250, delay: 1500, }}>
        <Toast {mode} size="md">
            {notification}
        </Toast>
    </aside>
{/if}