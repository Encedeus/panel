<script lang="ts">
    import { page } from "$app/stores";

    export let link: string | null = null;
</script>

<a href={link} class="hover:cursor-pointer">
    <div on:click class="{$page.route.id?.includes(String(link)) ? 'rounded-full bg-indigo-600' : ''} p-2">
        <slot/>
    </div>
</a>
