<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import Button from "$lib/components/generic/Button.svelte";
    import {createEventDispatcher} from "svelte";

    export let plugin: Plugin | undefined;

    const dispatch = createEventDispatcher();
    let isButtonDisabled = false;

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
        <div class="h-[300px] flex flex-col justify-start ml-2.5 mt-2.5 gap-2.5">
            <h1 class="text-3xl">{plugin.name}</h1>
            <div class="flex flex-row justify-start gap-2.5 font-light text-indigo-700">
                <p>{plugin.ownerName}</p>
                <!-- todo: add link to plugin page -->
                <a href="https://github.com">Plugin homepage</a>
            </div>
            <div class="flex flex-row justify-start gap-5">
                <Button on:click={onInstallClick} isDisabled={isButtonDisabled} size="sm">Install</Button>
                {#if plugin.releases}
                    <p>{plugin.releases[0].name}</p>
                {:else }
                    <p>no releases</p>
                {/if}
            </div>
        </div>


    {/if}
</div>
