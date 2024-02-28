<script lang="ts">
    import CardModal from "$lib/components/generic/CardModal.svelte";
    import DatabaseIcon from "$lib/components/heroicons/DatabaseIcon.svelte";
    import Input from "$lib/components/generic/Input.svelte";
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
    import TagIcon from "$lib/components/heroicons/TagIcon.svelte";
    import TextArea from "$lib/components/generic/TextArea.svelte";
    import Button from "$lib/components/generic/Button.svelte";

    export let className = "";
    export let open: boolean;
</script>

<CardModal className="{className} h-min" open={open} on:close>
    <TagIcon slot="icon"/>
    <span slot="title">Create role</span>

    <div slot="content" id="content" class="">
        <Input label="Name" placeholder="role name" className="w-full"></Input>
        <Input label="Description" placeholder="description" className="w-full"></Input>
        <TextArea size="md" placeholder="role permissions" label="Role permissions" className="w-full"/>
        <Button className="m-auto">Create</Button>
    </div>

</CardModal>

<style>

    #content {
        margin: 10px 20px 20px;
        display: flex;
        flex-direction: column;
        justify-content: stretch;
        gap: 20px 10px;
    }
</style>