<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import ModuleIcon from "$lib/components/heroicons/ModuleIcon.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import {createEventDispatcher} from "svelte";

    export let plugin: Plugin;


    const dispatch = createEventDispatcher();

    function onInstallClick(e) {
        e.plugin = plugin;
        dispatch("onInstall", e);
    }
    function onModuleClick(e) {
        e.plugin = plugin;
        dispatch("moduleClick", e);
    }

</script>

<div class="text-red-50 flex flex-row justify-between hover:bg-indigo-900 cursor-pointer" on:click={onModuleClick}>
    <div class="flex flex-row justify-start">
        <ModuleIcon height="75" width="75" className="m-2.5"/>
        <div class="flex flex-col justify-center">
            <p>{plugin.name}</p>
            <p class="text-sm font-thin">{plugin.ownerName}</p>
        </div>

    </div>
    <Button size="sm" className="mt-auto mb-auto mr-2.5" on:click={onInstallClick} isDisabled={!plugin.releases}>Install</Button>
</div>