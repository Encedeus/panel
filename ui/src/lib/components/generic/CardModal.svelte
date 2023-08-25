<script lang="ts">
    import Card from "$lib/components/generic/Card.svelte";
    import { createEventDispatcher } from "svelte";

    export let open = false;

    export let className = "";
    export let height: "sm" | "md" | "lg" = "sm";
    export let headerHeight: "sm" | "md" = "sm";
    export let width: "sm" | "md" | "lg" | "screen" = "sm";
    export let headerTextSize: "sm" | "md" | "lg" = "sm";
    export let fixedHeight = false;
    export let fixedWidth = false;

    const dispatch = createEventDispatcher();
    function onClose() {
        dispatch("close");
    }
</script>

{#if open}
    <div role="button" tabindex="0" on:keydown={onClose} on:click={onClose} class="cursor-default absolute left-0 right-0 top-0 bottom-0 m-0 w-screen h-screen bg-slate-950 bg-opacity-40">
    </div>
    <div class="absolute left-1/2 top-1/2 mr-[-50%] -translate-x-1/2 -translate-y-1/2 {className}">
        <Card {height} {headerHeight} {width} {headerTextSize} {fixedHeight} {fixedWidth}>
                <span slot="icon">
                    <slot name="icon"/>
                </span>
            <span slot="title">
                    <slot name="title"/>
                </span>
            <div slot="content" class="w-full h-full">
                <slot name="content"/>
            </div>
        </Card>
    </div>
{/if}