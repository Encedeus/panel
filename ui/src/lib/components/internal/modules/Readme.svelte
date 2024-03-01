<script lang="ts">
    import {Plugin} from "@encedeus/registry-js-api";
    import {getApi} from "$lib/api/api";
    import {PluginGetReadmeRequest} from "@encedeus/registry-js-api/src/proto/plugin_api";
    import {marked} from 'marked';

    export let plugin: Plugin;
    export let prevPlugin: Plugin | undefined = undefined;
    export let className: string;
    let readme: string;

    async function loadData() {

        if (plugin == prevPlugin && (prevPlugin || plugin)) {
            return;
        }

        prevPlugin = plugin;

        const response = await getApi().PluginService.GetReadme({pluginId: plugin.id} as PluginGetReadmeRequest);

        if (response.error) {
            readme = response.error.message;
            return;
        }

        // translate markdown to html
        if (!response.response?.readme) {
            readme = "";
            return;
        }

        readme = <string>marked(response.response.readme);
        console.log(readme);
    }

    $: plugin && loadData();
</script>

<!-- todo: skeletal loading -->

{#if readme}
    <div class="readme break-words {className}">
        {@html readme}
    </div>
{:else}
    <div>
        <div class=""></div>
    </div>
{/if}