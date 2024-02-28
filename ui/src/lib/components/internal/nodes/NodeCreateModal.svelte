<script lang="ts">
    import CardModal from "$lib/components/generic/CardModal.svelte";
    import DatabaseIcon from "$lib/components/heroicons/DatabaseIcon.svelte";
    import Input from "$lib/components/generic/Input.svelte";
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";

    export let className = "";
    export let open: boolean;
</script>

<CardModal className="{className} h-min" open={open} on:close>
    <ServerStackIcon slot="icon"/>
    <span slot="title">Create node</span>

    <div slot="content" id="content" class="">
        <Input label="name" placeholder="name" className="w-full"></Input>
        <Input label="location" placeholder="location" className="w-full"></Input>
    </div>

</CardModal>

<style>

    #content {
        margin: 10px 20px 20px;
        display: flex;
        flex-direction: column;
        justify-content: stretch;
        gap: 15px 10px;
    }
</style>