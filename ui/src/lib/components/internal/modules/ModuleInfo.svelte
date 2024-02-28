<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import Button from "$lib/components/generic/Button.svelte";
    import {createEventDispatcher} from "svelte";
    import TabSelector from "$lib/components/generic/TabSelector.svelte";

    export let plugin: Plugin | undefined;

    const dispatch = createEventDispatcher();
    let isButtonDisabled = false;
    let currentTab = 1;

    function onInstallClick(e) {
        e.plugin = plugin;
        dispatch("onInstall", e);
    }

    function loadData() {
        isButtonDisabled = !plugin?.releases;
    }

    $:  plugin && loadData();

</script>

<div class="w-3/6 text-white">
    {#if plugin}
        <div class="h-[300px] flex flex-col justify-start mt-2.5 gap-2.5">
            <h1 class="text-3xl ml-7">{plugin.name}</h1>
            <div class="flex flex-row justify-start gap-2.5 font-light text-indigo-700 ml-7">
                <p>{plugin.ownerName}</p>
                <!-- todo: add link to plugin page -->
                <a href="https://github.com">Plugin homepage</a>
            </div>
            <div class="flex flex-row justify-start gap-5 ml-7">
                <Button on:click={onInstallClick} isDisabled={isButtonDisabled} size="sm">Install</Button>
                {#if plugin.releases}
                    <p>{plugin.releases[0].name}</p>
                {:else }
                    <p>no releases</p>
                {/if}
            </div>

            <TabSelector bind:activeTabValue={currentTab} className={"border-b-[2px] pb-[2px] pl-7"} items={[
                {label: "Description", value: 1},
                {label: "What's New", value: 2},
                {label: "Reviews", value: 3}]
                }/>
            <div class="ml-7">
            {#if currentTab === 1}
                not implemented1
            {:else if currentTab === 2 }
                not implemented2
            {:else if currentTab === 3 }
                not implemented3
            {/if}
            </div>
        </div>
    {/if}
</div>
