<script>
    import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
    import TabEnvironment from "$lib/components/tabs/TabEnvironment.svelte";
    import SideBarTab from "$lib/components/internal/nav/SideBarTab.svelte";
    import CloudIcon from "$lib/components/heroicons/CloudIcon.svelte";
</script>

<TabEnvironment>
    <div slot="tabs">
        <SideBarTab link="/dashboard/account/settings">
            <SettingsIcon slot="icon"/>
            <span class="text-white font-monti font-normal" slot="label">
                Settings
            </span>
        </SideBarTab>
        <SideBarTab link="/dashboard/account/api">
            <CloudIcon slot="icon"/>
            <span class="text-white font-monti font-normal" slot="label">
                API Credentials
            </span>
        </SideBarTab>
    </div>
    <div slot="content">
        <slot/>
    </div>
</TabEnvironment>

