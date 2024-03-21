<script lang="ts">
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import ComputerDesktopIcon from "$lib/components/heroicons/ComputerDesktopIcon.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import PlusIcon from "$lib/components/heroicons/PlusIcon.svelte";
    import NodeRow from "$lib/components/internal/nodes/NodeRow.svelte";
    import NodeCreateModal from "$lib/components/internal/nodes/NodeCreateModal.svelte";

    let isShowingModal = false;

    function displayCreateModal() {
        isShowingModal = true;
    }
</script>

<!--<div class="absolute bg-slate-900 w-screen h-screen top-0 -z-10"></div>-->

<main class="flex flex-col p-8 gap-3">
    <CardHeader size="lg">
        Nodes
    </CardHeader>
    <Card height="lg" fixedHeight={true} className="overflow-x-hidden overflow-y-auto">
        <span slot="title" class="flex flex-row items-center">
            Node List
        </span>
        <span slot="icon">
            <ComputerDesktopIcon/>
        </span>
        <span slot="end">
            <Button className="rounded-xl w-72 px-4" on:click={displayCreateModal}>
                <span class="flex flex-row gap-2 items-center">
                    <PlusIcon/>
                    Create New
                </span>
            </Button>
        </span>
        <div slot="content" class="pb-4 text-white w-full">
            <table class="border-collapse">
                <thead>
                    <tr class="text-xs">
                        <th scope="col">Health status</th>
                        <th scope="col">Name</th>
                        <th scope="col">Location</th>
                        <th scope="col">Memory</th>
                        <th scope="col">Disk</th>
                        <th scope="col">Servers</th>
                        <th scope="col">TLS</th>
                    </tr>
                </thead>
                <tbody>
                    <NodeRow status="online" name="test" location="test-loc" ram={4096} disk={16384} servers={2} tls={true}/>
                    <NodeRow status="offline" name="test" location="test-loc" ram={4096} disk={16384} servers={2} tls={false}/>
                </tbody>
            </table>
        </div>
    </Card>

    <NodeCreateModal open={isShowingModal} on:close={() => {isShowingModal = false;}}/>
<!--    <NodeCardList/>-->
</main>

<style>
    table {
        table-layout: fixed;
        width: 100%;
        border-collapse: collapse;
    }

    thead {
        text-align: left;
    }

    thead th {
        width: 7.5%;
    }

    thead th:nth-child(1) {
        width: 5%;
    }

    thead th:nth-child(2) {
        width: 15%;
    }

    thead th:nth-child(3), th:nth-child(4), th:nth-child(5) {
        width: 17.5%;
    }

    th {
        padding: 0.5rem 1.25rem 0.5em;
    }

</style>