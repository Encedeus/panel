<script>
    import NodeCardList from "$lib/components/internal/nodes/NodeCardList.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import ComputerDesktopIcon from "$lib/components/heroicons/ComputerDesktopIcon.svelte";
    import HeartIcon from "$lib/components/heroicons/HeartIcon.svelte";
    import LockIcon from "$lib/components/heroicons/LockIcon.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import PlusIcon from "$lib/components/heroicons/PlusIcon.svelte";
    import NodeRow from "$lib/components/internal/nodes/NodeRow.svelte";
    import DatabaseIcon from "$lib/components/heroicons/DatabaseIcon.svelte";
    import DatabaseRow from "$lib/components/internal/databases/DatabaseRow.svelte";
</script>

<!--<div class="absolute bg-slate-900 w-screen h-screen top-0 -z-10"></div>-->

<main class="flex flex-col p-8 gap-3">
    <CardHeader size="lg">
        Databases
    </CardHeader>
    <Card height="lg" fixedHeight={true} className="overflow-x-hidden overflow-y-auto">
        <span slot="title" class="flex flex-row items-center">
            Database List
        </span>
        <span slot="icon">
            <DatabaseIcon/>
        </span>
        <span slot="end">
            <Button className="rounded-xl w-72 px-4">
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
                        <th scope="col">Size on disk</th>
                        <th scope="col">IOPS</th>
                        <th scope="col">TLS</th>
                    </tr>
                </thead>
                <tbody>
                    <DatabaseRow status="online" name="test" location="test-loc" sizeOnDisk={16384} iops={26} tls={true}/>
                    <DatabaseRow status="offline" name="test" location="test-loc" sizeOnDisk={16384} iops={20} tls={false}/>
                </tbody>
            </table>
        </div>
    </Card>
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