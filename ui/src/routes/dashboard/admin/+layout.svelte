<script lang="ts">
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import SideBarTabLabel from "$lib/components/internal/nav/SideBarTabLabel.svelte";
    import TabEnvironment from "$lib/components/internal/tabs/TabEnvironment.svelte";
    import type { LayoutServerData } from "./$types";
    import { page } from "$app/stores";
    import ComputerDesktop from "$lib/components/heroicons/ComputerDesktopIcon.svelte";
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
</script>


<TabEnvironment>
    <div slot="tabs">
        <SideBarTab link="/dashboard/admin/nodes">
            <ComputerDesktop width="32" height="32" slot="icon"/>
            <SideBarTabLabel slot="label">
                Nodes
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/admin/servers">
            <ServerStackIcon width="30" height="28" slot="icon"/>
            <SideBarTabLabel slot="label">
                Servers
            </SideBarTabLabel>
        </SideBarTab>
    </div>
    <div slot="content" class="w-full h-full">
        <slot/>
    </div>
</TabEnvironment>
