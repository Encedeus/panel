<script lang="ts">

    import Checkbox from "$lib/components/generic/Checkbox.svelte";

    export let crumbs: Array<{ name: string; path: string; }>;
</script>

<div class="rounded-xl py-4 px-6 bg-indigo-700 flex flex-row gap-4 items-center min-w-[42rem] max-w-[80rem] h-14">
    <Checkbox/>
    <span>
        {#each crumbs.toSpliced(crumbs.length - 1, 1) as crumb, idx}
            <span class="opacity-25">/</span>
            <span class="{idx < crumbs.length - 1 ? "opacity-25" : "opacity-100"} opacity-25">{crumb.name}</span>
        {/each}
        <span class="opacity-25">/</span>
        <span>{crumbs[crumbs.length - 1].name}</span>
    </span>
</div>

<style>

</style>