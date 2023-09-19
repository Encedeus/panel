<script lang="ts">
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
    import TrashCanIcon from "$lib/components/heroicons/TrashCanIcon.svelte";
    import { createEventDispatcher, onMount } from "svelte";

    export let name: string;
    export let key: string;
    export let lastUsed = "";
    export let className = "";
    export let id: string;

    let notification: HTMLElement;

    onMount(() => {
        notification.style.opacity = "0";
    });

    function copyKeyToClipboard() {
        navigator.clipboard.writeText(key);

        notification.classList.remove("fade-out-animation");
        notification.classList.add("fade-in-animation");

        setTimeout(() => {
            notification.classList.add("fade-out-animation");
            notification.classList.remove("fade-in-animation");
        }, 950);
    }

    const dispatch = createEventDispatcher();

    function onDelete() {
        dispatch("delete", {
            keyId: id,
        });
    }
</script>

<div class="flex items-center justify-between py-5 px-6 bg-indigo-900 rounded-xl {className}">
    <div class="flex items-center justify-center gap-3">
        <KeyIcon height={34} width={34}/>
        <div class="flex flex-col items-start justify-center">
            <span class="text-white text-lg font-semibold">{name}</span>
            {#if lastUsed}
                <span class="text-white text-opacity-25 text-[9px] font-semibold -mt-0.5">{`Last Used: ${lastUsed.toUpperCase()}`}</span>
            {/if}
        </div>
    </div>
    <div class="flex items-center gap-6">
        <div class="flex-col justify-center items-center relative">
            <div class="absolute left-[46.875%] -translate-y-6">
                <div bind:this={notification} class="flex flex-col gap-0 justify-center items-center opacity-0">
                    <div class="absolute w-40 h-6 bg-indigo-500 rounded-md text-xs text-white font-semibold flex justify-center items-center">
                        <span>Copied to clipboard</span>
                    </div>
                    <span class="relative w-0 h-0 top-4 border-l-8 border-l-transparent border-r-8 border-r-transparent border-b-8 border-b-indigo-500 rotate-180">
                    </span>
                </div>
            </div>
            <span class="rounded-xl bg-indigo-950 text-white text-sm py-1.5 px-7 cursor-pointer flex"
                  on:click={copyKeyToClipboard} on:keydown={copyKeyToClipboard} role="button"
                  tabindex="0">{key?.slice(0, 24)}...</span>
        </div>
        <span class="hover:cursor-pointer" on:click={onDelete} on:keydown={onDelete} role="button" tabindex="0">
            <TrashCanIcon height={34} width={34}/>
        </span>
    </div>
</div>

<style lang="postcss">

    @keyframes fade-in {
        0% {
            @apply opacity-0;
        }
        100% {
            @apply opacity-100;
        }
    }

    @keyframes fade-out {
        0% {
            @apply opacity-100;
        }
        100% {
            @apply opacity-0;
        }
    }

    .fade-in-animation {
        animation-name: fade-in;
        animation-fill-mode: forwards;
        animation-duration: 0.5s;
    }

    .fade-out-animation {
        animation-name: fade-out;
        animation-fill-mode: forwards;
        animation-duration: 0.75s;
    }
</style>

