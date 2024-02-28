<script lang="ts">
    import CardModal from "$lib/components/generic/CardModal.svelte";
    import DatabaseIcon from "$lib/components/heroicons/DatabaseIcon.svelte";
    import Input from "$lib/components/generic/Input.svelte";

    export let className = "";
    export let open: boolean;
</script>

<CardModal className="{className}" open={open} on:close>
    <DatabaseIcon slot="icon"/>
    <span slot="title">Create database</span>

    <div slot="content" id="content" class="">
        <Input label="name" placeholder="name" className="w-full"></Input>
        <Input label="location" placeholder="location" className="w-full"></Input>

        <div id="uriInput">
            <Input size=" " className="w-[375px]" label="address" placeholder="database address"></Input>
            <Input size=" " className="w-[95px]" label="port" placeholder="port"></Input>
        </div>

        <Input label="user" placeholder="database user" className="w-full"></Input>
        <Input label="user password" type="password" placeholder="database user password" className="w-full"></Input>
        <Input label="database name" placeholder="database name" className="w-full"></Input>
    </div>

</CardModal>

<style>
    #uriInput {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        gap: 10px 10px;
    }

    #content {
        margin: 10px 20px 20px;
        display: flex;
        flex-direction: column;
        justify-content: stretch;
        gap: 15px 10px;
    }
</style>