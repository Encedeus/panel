<script lang="ts">
    import { page } from "$app/stores";

    export let link: string;
</script>

<a href={link}>
    <div class="{$page.route.id.includes(link) ? 'rounded-full bg-indigo-600' : ''} p-2">
        <slot/>
    </div>
</a>