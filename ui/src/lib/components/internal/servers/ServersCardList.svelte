<script lang="ts">
    import Card from "$lib/components/generic/Card.svelte";
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
    import ServerCard from "$lib/components/internal/servers/ServerCard.svelte";
    import type {Server} from "@encedeus/js-api";

    export let className = "";
    export let serverList: Server[];
</script>

<div>
    <Card headerHeight="sm" fixedWidth={true} fixedHeight={true} height="lg" width="screen" className={className}>
        <span slot="title">Servers</span>
        <ServerStackIcon slot="icon"/>
        <div slot="content" class="flex flex-col items-stretch gap-3 p-6">
            {#each serverList as server}
                <ServerCard server={server}/>
            {/each}
        </div>
    </Card>
</div>
