<script lang="ts">
    import Button from "$lib/components/generic/Button.svelte";
    import Modal from "$lib/components/generic/AlertModal.svelte";
    import { createEventDispatcher } from "svelte";

    export let open = false;
    export let description = "";
    
    const dispatch = createEventDispatcher();
    function onProceed() {
        dispatch("proceed");
    }

    function onCancel() {
        dispatch("cancel");
    }
</script>

<Modal on:close={onCancel} {open}>
    <span slot="title">Are you sure?</span>
    <span slot="description">
        {description}
    </span>
    <div slot="content" class="flex gap-5">
        <Button on:click={onCancel}>
            Cancel
        </Button>
        <Button color="red" on:click={onProceed}>
            Delete
        </Button>
    </div>
</Modal>
