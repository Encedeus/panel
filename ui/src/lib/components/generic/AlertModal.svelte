<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";

    export let open = false;

    const dispatch = createEventDispatcher();
    function onClose() {
        dispatch("close");
    }
</script>

{#if open}
    <div role="button" tabindex="0" on:keydown={onClose} on:click={onClose} id="modal"
         class="cursor-default absolute top-0 left-0 bottom-0 right-0 mr-0 flex justify-center bg-opacity-40 bg-slate-950">
    </div>
    <div class="bg-indigo-900 rounded-2xl w-4/12 h-64 top-1/4 text-center absolute flex flex-col justify-between items-center p-8 z-0">
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
{/if}

