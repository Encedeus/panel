<script lang="ts">
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
    import TrashCanIcon from "$lib/components/heroicons/TrashCanIcon.svelte";
    import { createEventDispatcher } from "svelte";

    export let name: string;
    export let key: string;
    export let lastUsed = "";
    export let className = "";
    export let id: string;

    const dispatch = createEventDispatcher();
    function onDelete() {
        dispatch("delete", {
            keyId: id,
        });
    }

    function copyKeyToClipboard() {
        navigator.clipboard.writeText(key);
    }
</script>

<div class="flex items-center justify-between py-4 px-6 bg-indigo-900 rounded-xl {className}">
   <div class="flex items-center justify-center gap-3">
        <KeyIcon width={34} height={34}/>
        <div class="flex flex-col items-start justify-center">
            <span class="text-white text-lg font-semibold">{name}</span>
            {#if lastUsed}
                <span class="text-white text-opacity-25 text-[9px] font-semibold -mt-0.5">{`Last Used: ${lastUsed.toUpperCase()}`}</span>
            {/if}
        </div>
   </div>
   <div class="flex items-center gap-6">
        <span role="button" tabindex="0" on:keydown={copyKeyToClipboard} on:click={copyKeyToClipboard} class="rounded-xl bg-indigo-950 text-white text-sm py-1.5 px-7 cursor-pointer">{key.slice(0, 24) + "..."}</span>
        <span role="button" tabindex="0" on:keydown={onDelete} on:click={onDelete} class="hover:cursor-pointer">
            <TrashCanIcon width={34} height={34}/>
        </span>
   </div>
</div>
