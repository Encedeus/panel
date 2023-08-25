<script lang="ts">
    import CardModal from "$lib/components/generic/CardModal.svelte";
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import { createEventDispatcher } from "svelte";

    export let open = false;
    export let subject: string;

    const dispatch = createEventDispatcher();
    function onCancel() {
        dispatch("cancel");
    }
</script>

<CardModal on:close={onCancel} className="flex items-center" fixedHeight={true} headerHeight="md" height="md" {open} width="lg">
    <span class="text-2xl font-bold" slot="title">Change {subject}</span>
    <div class="flex flex-col w-full h-full" slot="content">
        <div class="flex flex-col items-center p-8 gap-6">
            <Input label="Old {subject}" placeholder="Enter Old {subject}" size="xl"/>
            <Input label="New {subject}" placeholder="Enter New {subject}" size="xl"/>
            <Input label="Confirm New {subject}" placeholder="Enter New {subject}" size="xl"/>
        </div>
        <div class="flex mt-6 mb-12 justify-center">
            <Button size="md">Save</Button>
        </div>
    </div>
</CardModal>