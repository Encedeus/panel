<script lang="ts">
    import { createEventDispatcher } from "svelte";

    export let open = false;

    const dispatch = createEventDispatcher();
    function onClose() {
        dispatch("close");
    }
</script>

{#if open}
<div role="button" tabindex="0" on:keydown={onClose} on:click={onClose} class="cursor-default absolute left-0 right-0 top-0 bottom-0 m-0 w-screen h-screen bg-slate-950 bg-opacity-40 flex justify-center">
    <div class="bg-indigo-900 rounded-2xl w-4/12 h-64 absolute top-1/4 flex flex-col justify-between items-center p-8 z-0">
        <span class="text-white text-3xl font-semibold z-10">
            <slot name="title"/>
        </span>
        <span class="text-white text-lg text-opacity-40">
            <slot name="description"/>
        </span>
        <div>
            <slot name="content"/>
        </div>
    </div>
</div>
{/if}

