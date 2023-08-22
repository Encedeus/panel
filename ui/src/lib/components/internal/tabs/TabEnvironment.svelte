<script lang="ts">
    import SideBar from "$lib/components/internal/nav/SideBar.svelte";
    import NavBar from "$lib/components/internal/nav/NavBar.svelte";
    import NavItem from "$lib/components/internal/nav/NavItem.svelte";
    import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
    import ServerIcon from "$lib/components/heroicons/ServerIcon.svelte";
    import DoorExitIcon from "$lib/components/heroicons/DoorExitIcon.svelte";
    import { signOut } from "$lib/services/auth_service";
    import { page } from "$app/stores";
    import ProfilePicture from "$lib/components/generic/ProfilePicture.svelte";
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
            <NavItem link="/dashboard/account/settings">
                <ProfilePicture/>
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
            <div class="bg-slate-900 basis-full scroll">
                <slot name="content"/>
            </div>
        </div>
    {:else}
        <div class="scroll">
            <slot name="content"/>
        </div>
    {/if}
</div>

<style>
    :global(body) {
        overflow: hidden;
    }

    .scroll {
        height: 100vh;
        width: 100vw;
        overflow-y: auto !important;
        padding-bottom: 4.5rem;
    }
</style>
