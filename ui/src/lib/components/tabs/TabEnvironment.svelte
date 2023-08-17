<script lang="ts">
    import SideBar from "../internal/nav/SideBar.svelte";
    import NavBar from "../internal/nav/NavBar.svelte";
    import NavItem from "../internal/nav/NavItem.svelte";
    import SettingsIcon from "../heroicons/SettingsIcon.svelte";
    import ServerIcon from "../heroicons/ServerIcon.svelte";
    import DoorExitIcon from "../heroicons/DoorExitIcon.svelte";
    import DefaultUserIcon from "../heroicons/DefaultUserIcon.svelte";
    import { signOut } from "$lib/services/auth_service";
    import { page } from "$app/stores";
</script>

<div class="flex flex-col grow-0">
    <NavBar>
        <div slot="logo" class="text-white text-3xl font-bold font-lato">Encedeus</div>
        <div class="flex items-center justify-between gap-3 mr-5" slot="links">
            <NavItem link="/dashboard/account">
                <SettingsIcon/>
            </NavItem>
            <NavItem link="/dashboard/servers">
                <ServerIcon/>
            </NavItem>
            <NavItem on:click={signOut}>
                <DoorExitIcon/>
            </NavItem>
            <NavItem>
                <DefaultUserIcon/>
            </NavItem>
        </div>
    </NavBar>
    {#if $page.route.id !== "/dashboard/servers"}
        <div class="flex flex-row">
            <SideBar>
                <div>
                    <slot name="tabs"/>  
                </div>
            </SideBar>
            <div class="bg-slate-900 basis-full">
                <slot name="content"/>
            </div>
        </div>
    {:else}
        <slot name="content"/>
    {/if}
</div>



<style>
    :global(body) {
        overflow: hidden;
    }
</style>
