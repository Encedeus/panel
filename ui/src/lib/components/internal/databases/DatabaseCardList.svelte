<script lang="ts">
    import Card from "$lib/components/generic/Card.svelte";
    import NodeCard from "$lib/components/internal/nodes/NodeCard.svelte";
    import ComputerDesktop from "$lib/components/heroicons/ComputerDesktopIcon.svelte";

    export let className = "";
</script>

<div>
    <Card headerHeight="sm" fixedWidth={true} fixedHeight={true} height="lg" width="screen" className={className}>
        <span slot="title">Nodes</span>
        <ComputerDesktop slot="icon"/>
        <div slot="content" class="flex flex-col items-stretch gap-3 p-6">
            <NodeCard/>
        </div>
    </Card>
</div>
