<script>
    import NodeCardList from "$lib/components/internal/nodes/NodeCardList.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import ComputerDesktopIcon from "$lib/components/heroicons/ComputerDesktopIcon.svelte";
</script>

<!--<div class="absolute bg-slate-900 w-screen h-screen top-0 -z-10"></div>-->

<main class="flex flex-col p-8 gap-3">
    <CardHeader size="lg">
        Nodes
    </CardHeader>
    <Card className="p-0">
        <span slot="title" class="flex flex-row items-center">
            Node List
        </span>
        <span slot="icon">
            <ComputerDesktopIcon/>
        </span>
        <div slot="content" class="text-white">
            <table>
                <tr>
                    <td></td>
                    <td>Name</td>
                    <td>Location</td>
                    <td>Memory</td>
                    <td>Disk</td>
                    <td>Servers</td>
                    <td>TLS</td>
                    <td>Public</td>
                </tr>
            </table>
        </div>
    </Card>
<!--    <NodeCardList/>-->
</main>