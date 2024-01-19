<script lang="ts">
    import CloudIcon from "$lib/components/heroicons/CloudIcon.svelte";
    import SideBarTabLabel from "$lib/components/internal/nav/SideBarTabLabel.svelte";
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import TabEnvironment from "$lib/components/internal/tabs/TabEnvironment.svelte";
    import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
    import type { LayoutServerData } from "./$types";

    export let data: LayoutServerData;
</script>

<TabEnvironment>
    <div slot="tabs">
        <SideBarTab link="/dashboard/account/settings">
            <SettingsIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Settings
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/account/api">
            <CloudIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                API Credentials
            </SideBarTabLabel>
        </SideBarTab>
        {#each data.modules as m}
            <SideBarTab link="/dashboard/account/modules/{m.manifest.name}">
                <img src="http://localhost:{m?.frontend_server?.port?.value}/favicon.ico" height="34" width="34" alt="{m?.manifest?.name} icon" slot="icon">
                <SideBarTabLabel slot="label">
                    {m.manifest.frontend.tab_name}
                </SideBarTabLabel>
            </SideBarTab>
        {/each}
    </div>
    <div class="w-full h-full" slot="content">
        <slot/>
    </div>
</TabEnvironment>