<script lang="ts">
    import PencilSquareIcon from "$lib/components/heroicons/PencilSquareIcon.svelte";

    export let name: string;
    export let description: string;
    export let permissions: string[];
    export let userList: string[];
</script>

<tr class="border-y-2 border-indigo-900">
    <td>
        {name}
    </td>
    <td>
        {#if description.length > 30}
            {description.substring(0, 29)}...
        {:else }
            {description}
        {/if}

    </td>
    <td>
        {#if permissions.length > 3}
            {permissions.slice(0, 3)}...
        {:else }
            {permissions}
        {/if}
    </td>
    <td>
        {#if userList.length > 3}
            {userList.slice(0, 3)}...
        {:else }
            {userList}
        {/if}
    </td>
    <td class="min-w-full flex flex-row justify-end">
        <PencilSquareIcon/>
    </td>
</tr>

<style>
    td {
        padding-left: 1.25rem;
        padding-right: 1.25rem;
    }

    td:not(:first-child) {
        padding-top: 1rem;
        padding-bottom: 1rem;
    }
</style>