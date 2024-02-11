<script lang="ts">
    import type { LayoutServerData } from "./$types";
    import { page } from "$app/stores";

    // export let data: LayoutServerData;

    // const module = data.modules.find(m => m.manifest.name == $page.params.module);
</script>

<!--<iframe class="border-none overflow-x-hidden w-full h-full" src="http://localhost:{module?.frontendServer?.port?.value}" title={module.manifest.name} loading="eager"></iframe>-->
<iframe class="border-none overflow-x-hidden w-full h-full" src="http://localhost:7500" title={"Console"} loading="eager"></iframe>