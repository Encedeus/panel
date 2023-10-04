<script lang="ts">
    import TabEnvironment from "$lib/components/internal/tabs/TabEnvironment.svelte";
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
    import SideBarTabLabel from "$lib/components/internal/nav/SideBarTabLabel.svelte";
    import CloudIcon from "$lib/components/heroicons/CloudIcon.svelte";
    import { page } from "$app/stores";
    import FolderIcon from "$lib/components/heroicons/FolderIcon.svelte";
    import DatabaseIcon from "$lib/components/heroicons/DatabaseIcon.svelte";
    import CalendarIcon from "$lib/components/heroicons/CalendarIcon.svelte";
    import DefaultUserIcon from "$lib/components/heroicons/DefaultUserIcon.svelte";
    import GlobeIcon from "$lib/components/heroicons/GlobeIcon.svelte";
    import StartupIcon from "$lib/components/heroicons/StartupIcon.svelte";
    import UserIcon from "$lib/components/heroicons/UserIcon.svelte";
    import ConsoleIcon from "$lib/components/heroicons/ConsoleIcon.svelte";

    $: serverId = $page.params.id;
</script>

<TabEnvironment>
    <div slot="tabs">
        <SideBarTab link="/dashboard/servers/{serverId}/console">
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
        </SideBarTab>
    </div>
    <div slot="content">
        <slot/>
    </div>
</TabEnvironment>