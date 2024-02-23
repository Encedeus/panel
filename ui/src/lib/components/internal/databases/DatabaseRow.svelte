<script lang="ts">
    import HeartIcon from "$lib/components/heroicons/HeartIcon.svelte";
    import LockIcon from "$lib/components/heroicons/LockIcon.svelte";

    export let status: "online" | "offline";
    export let name: string;
    export let location: string;
    export let sizeOnDisk: number;
    export let iops: number;
    export let tls: boolean;
</script>

<tr class="border-y-2 border-indigo-900">
    <td>
        {#if status === "online"}
            <HeartIcon color="#10B981"/>
        {:else}
            <HeartIcon color="#EF4444"/>
        {/if}
    </td>
    <td>
        {name}
    </td>
    <td>
        {location}
    </td>
    <td>
        {sizeOnDisk} MB
    </td>
    <td>
        {iops}
    </td>
    <td>
        {#if tls}
            <LockIcon/>
        {:else}
            <LockIcon color="red" unlocked={true}/>
        {/if}
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