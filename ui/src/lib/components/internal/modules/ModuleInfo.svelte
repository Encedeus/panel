<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import Button from "$lib/components/generic/Button.svelte";
    import {createEventDispatcher} from "svelte";
    import TabSelector from "$lib/components/generic/TabSelector.svelte";
    import Readme from "$lib/components/internal/modules/Readme.svelte";
    import ArrowRightUpIcon from "$lib/components/heroicons/ArrowRightUpIcon.svelte";

    export let plugin: Plugin;
    export let className;

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

<div class="w-3/6 text-white {className}">
    {#if plugin}
        <div class="h-[300px] flex flex-col justify-start mt-2.5 gap-2.5">
            <h1 class="text-3xl ml-7">{plugin.name}</h1>
            <div class="flex flex-row justify-start items-baseline gap-2.5 font-light text-[#4F46E5] ml-7">
                <p>{plugin.ownerName}</p>
                <!-- todo: add link to plugin page -->
                <a target="_blank" href={plugin.source.repoUri} class="flex flex-row gap-1 align-baseline"><p>Plugin homepage</p> <ArrowRightUpIcon className="m-auto" width="12" height="12"/></a>
            </div>
            <div class="flex flex-row justify-start gap-5 ml-7">
                <Button on:click={onInstallClick} isDisabled={isButtonDisabled} size="sm">Install</Button>
                {#if plugin.releases}
                    <p>{plugin.releases[0].name}</p>
                {:else }
                    <p>no releases</p>
                {/if}
            </div>

            <TabSelector bind:activeTabValue={currentTab} className={"border-b-[2px] pb-[4px] pl-7 mt-4"} items={[
                {label: "Description", value: 1},
                {label: "What's New", value: 2},
                {label: "Reviews", value: 3}]
                }/>
            <div class="ml-7">
                <div class:invisible={currentTab !== 1}>
                    <Readme plugin={plugin}/>
                </div>
                {#if currentTab === 2 }
                    No changes
                {:else if currentTab === 3 }
                    No reviews
                {/if}
            </div>
        </div>
    {/if}
</div>

<style>
    .invisible {
        display: none;
    }
</style>