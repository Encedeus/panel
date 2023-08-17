<script>
    import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
    import TabEnvironment from "$lib/components/internal/tabs/TabEnvironment.svelte";
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import CloudIcon from "$lib/components/heroicons/CloudIcon.svelte";
    import SideBarTabLabel from "$lib/components/internal/nav/SideBarTabLabel.svelte";
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
    </div>
    <div slot="content">
        <slot/>
    </div>
</TabEnvironment>

