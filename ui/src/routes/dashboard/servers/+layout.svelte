<script lang="ts">
    import TabEnvironment from "$lib/components/internal/tabs/TabEnvironment.svelte";
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import SideBarTabLabel from "$lib/components/internal/nav/SideBarTabLabel.svelte";
    import { page } from "$app/stores";
    import type { LayoutServerData } from "./$types";
    import ConsoleIcon from "$lib/components/heroicons/ConsoleIcon.svelte";

    export let data: LayoutServerData;
    $: serverId = $page.params.id;
</script>

<TabEnvironment>
    <div slot="tabs">
<!--        <SideBarTab link="/dashboard/servers/{serverId}/console">
            <ConsoleIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Console
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/files">
            <FolderIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                File Manager
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/databases">
            <DatabaseIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Databases
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/schedules">
            <CalendarIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Schedules
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/users">
            <UserIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Users
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/backup">
            <CloudIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Backup
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/network">
            <GlobeIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Network
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/startup">
            <StartupIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Startup
            </SideBarTabLabel>
        </SideBarTab>
        <SideBarTab link="/dashboard/servers/{serverId}/settings">
            <SettingsIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Settings
            </SideBarTabLabel>
        </SideBarTab>-->
        {#if data.modules}
            {#each data.modules as m}
                <SideBarTab link="/dashboard/servers/{serverId}/modules/{m.manifest?.name}">
                    <img src="http://localhost:{m?.frontendServer?.port?.value}/favicon.ico" height="34" width="34" alt="{m?.manifest?.name} icon" slot="icon">
                    <SideBarTabLabel slot="label">
                        {m.manifest?.frontend?.tabName}
                    </SideBarTabLabel>
                </SideBarTab>
            {/each}
        {/if}
        <SideBarTab link="/dashboard/servers/{serverId}/modules/dev_module_console">
            <ConsoleIcon slot="icon"/>
            <SideBarTabLabel slot="label">
                Console
            </SideBarTabLabel>
        </SideBarTab>
    </div>
    <div slot="content" class="w-full h-full">
        <slot/>
    </div>
</TabEnvironment>