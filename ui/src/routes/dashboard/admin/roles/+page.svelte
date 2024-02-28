<script lang="ts">
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import ComputerDesktopIcon from "$lib/components/heroicons/ComputerDesktopIcon.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import PlusIcon from "$lib/components/heroicons/PlusIcon.svelte";
    import NodeRow from "$lib/components/internal/nodes/NodeRow.svelte";
    import NodeCreateModal from "$lib/components/internal/nodes/NodeCreateModal.svelte";
    import TagIcon from "$lib/components/heroicons/TagIcon.svelte";
    import RoleRow from "$lib/components/internal/roles/RoleRow.svelte";
    import RoleCreateModal from "$lib/components/internal/roles/RoleCreateModal.svelte";

    let isShowingModal = false;

    function displayCreateModal() {
        isShowingModal = true;
    }
</script>

<!--<div class="absolute bg-slate-900 w-screen h-screen top-0 -z-10"></div>-->

<main class="flex flex-col p-8 gap-3">
    <CardHeader size="lg">
        Roles
    </CardHeader>
    <Card height="lg" fixedHeight={true} className="overflow-x-hidden overflow-y-auto">
        <span slot="title" class="flex flex-row items-center">
            Role List
        </span>
        <span slot="icon">
            <TagIcon/>
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
                    <th scope="col">Name</th>
                    <th scope="col">Description</th>
                    <th scope="col">Permissions</th>
                    <th scope="col">User List</th>
                    <th scope="col"></th>
                </tr>
                </thead>
                <tbody>
                <RoleRow name="role1" description="lorem ipsum1" permissions={["1", "2", "3"]} userList={["testus", "mestus"]}/>
                <RoleRow name="role2" description="lorem ipsum2" permissions={["1", "2", "3", "4"]} userList={["testus", "mestus", "pestus"]}/>
                <RoleRow name="role1" description="lorem ipsum3" permissions={["1", "2", ]} userList={["testus", "mestus", "pestus", "destus"]}/>
                </tbody>
            </table>
        </div>
    </Card>

    <RoleCreateModal open={isShowingModal} on:close={() => {isShowingModal = false;}}/>
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

    th {
        padding: 0.5rem 1.25rem 0.5em;
    }

</style>