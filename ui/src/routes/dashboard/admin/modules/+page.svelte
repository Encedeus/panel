<script lang="ts">

    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import ModuleIcon from "$lib/components/heroicons/ModuleIcon.svelte";
    import {Plugin, PluginSearchByNameRequest} from "@encedeus/registry-js-api";
    import {} from "@encedeus/registry-js-api";
    import ModuleRow from "$lib/components/internal/modules/ModuleRow.svelte";
    import {getApi} from "$lib/api/api";
    import ModuleInfo from "$lib/components/internal/modules/ModuleInfo.svelte";

    let plugins: Plugin[] | undefined = [];

    let selectedModule: Plugin | undefined = undefined;

    const api = getApi();

    async function loadData() {
        const resp = await api.PluginService.SearchPlugins({} as PluginSearchByNameRequest);

        if (!resp.error) {
            plugins = resp.response?.plugins;
        }
    }

    function onInstallClick(e) {
        console.log("insatll", e.detail.plugin.name);

    }

    function onModuleClick(e) {
        selectedModule = e.detail.plugin;
        console.log("click", e.detail.plugin.name);
    }

    loadData();
</script>

<main class="flex flex-col p-8 gap-3">
    <CardHeader size="lg">
        Modules
    </CardHeader>

    <Card height="lg" fixedHeight={true} className="overflow-x-hidden overflow-y-auto">
        <span slot="title" class="flex flex-row items-center">
            Module List
        </span>
        <ModuleIcon slot="icon"/>

        <div slot="content" class="flex flex-row justify-center h-full ">
            <div id="pluginList" class="w-3/6 border-r-2 border-black pb-9 h-full min-h-[36rem]">
                {#if plugins}
                    {#each plugins as plugin}
                        <ModuleRow on:onInstall={onInstallClick} on:moduleClick={onModuleClick} plugin={plugin}/>
                    {/each}
                {/if}
            </div>

            <ModuleInfo plugin={selectedModule} on:onInstall={onInstallClick}/>
        </div>
    </Card>
</main>