<script lang="ts">
    import { page } from "$app/stores";

    export let link: string;
</script>

<a href={link}>
<div class="{$page?.route.id.includes(link) ? 'bg-indigo-800' : 'bg-indigo-900'} py-4 pl-6 w-full flex flex-row items-center gap-2.5">
        <slot name="icon"/>
        <slot name="label"/>
    </div>
</a>

