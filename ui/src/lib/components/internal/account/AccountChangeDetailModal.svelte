<script lang="ts">
    import CardModal from "$lib/components/generic/CardModal.svelte";
    import Input from "$lib/components/generic/Input.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import { createEventDispatcher } from "svelte";
    import type { AccountChangeDetails } from "$lib/services/change_details_service";
    import { subjectAsUppercase } from "$lib/services/change_details_service";

    export let open = false;
    export let subjectDetails: AccountChangeDetails;

    let subject: string;
    $: { subject = subjectAsUppercase(subjectDetails?.subject); }

    export let oldSubjectError = false;
    export let newSubjectError = false;
    export let confirmNewSubjectError = false;

    const dispatch = createEventDispatcher();
    function onCancel() {
        dispatch("cancel");
    }

    function onSave() {
        dispatch("save");
    }
</script>

<CardModal on:close={onCancel} className="flex items-center" fixedHeight={true} headerHeight="md" height="md" {open} width="lg">
    <span class="text-2xl font-bold" slot="title">Change {subject}</span>
    <div class="flex flex-col w-full h-full" slot="content">
        <div class="flex flex-col items-center p-8 gap-6">
            <Input bind:error={oldSubjectError} on:input label="Old {subject}" bind:value={subjectDetails.oldSubject} placeholder="Enter Old {subject}" size="xl"/>
            <Input bind:error={newSubjectError} on:input label="New {subject}" bind:value={subjectDetails.newSubject} placeholder="Enter New {subject}" size="xl"/>
            <Input bind:error={confirmNewSubjectError} on:input label="Confirm New {subject}" bind:value={subjectDetails.confirmNewSubject} placeholder="Enter New {subject}" size="xl"/>
        </div>
        <div class="flex mt-6 mb-12 justify-center">
            <Button size="md" on:click={onSave}>Save</Button>
        </div>
    </div>
</CardModal>