<script lang="ts">
    import type { LayoutData, PageData } from "./$types";

    export let layoutData: LayoutData;
    export let pageData: PageData;

    const module = layoutData.modules.find(m => m.manifest.name === pageData.module);
</script>

<iframe class="border-none overflow-x-hidden" src="http://localhost:{module.frontend_server.port.value}" title={module.manifest.name}></iframe>