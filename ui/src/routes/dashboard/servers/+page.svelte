<script lang="ts">
    import ServersCardList from "$lib/components/internal/servers/ServersCardList.svelte";
    import type {ServersListData} from "$lib/interfaces/serversListData";

    export let data: ServersListData;

    $: servers = data.servers;

    let ff =() => {
        console.log("data", data);
    };

    $: data && ff();
</script>

<div class="absolute bg-slate-900 w-screen h-screen top-0 -z-10"></div>

<main class="flex justify-center mt-14">
    <ServersCardList serverList={servers}/>
</main>
