<script lang="ts">
    import { page } from "$app/stores";
    import { createEventDispatcher } from "svelte";

    export let link = "";

    const dispatch = createEventDispatcher();
    function onClick() {
        dispatch("click");
    }
</script>

<a href={link} class="hover:cursor-pointer">
    <div on:click={onClick} on:keypress={onClick} tabindex="0" role="button" class="{$page.route.id?.includes(link) && link ? 'rounded-full bg-indigo-600' : ''} p-2">
        <slot/>
    </div>
</a>