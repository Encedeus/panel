<script lang="ts">
  import DefaultUserIcon from "$lib/components/heroicons/DefaultUserIcon.svelte";
  import DoorExitIcon from "$lib/components/heroicons/DoorExitIcon.svelte";
  import ServerIcon from "$lib/components/heroicons/ServerIcon.svelte";
  import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
  import NavBar from "$lib/components/internal/nav/NavBar.svelte";
  import NavItem from "$lib/components/internal/nav/NavItem.svelte";
  import ServersCardList from "$lib/components/internal/servers/ServersCardList.svelte";
  import { signOut } from "$lib/services/auth_service";
</script>

<div class="bg-slate-900 absolute w-screen h-screen"></div>
<!---->
<!-- <NavBar> -->
<!--   <div slot="logo" class="text-white text-3xl font-bold font-lato">Encedeus</div> -->
<!--   <div class="flex items-center justify-between gap-3 mr-5" slot="links"> -->
<!--     <NavItem link="/dashboard/account/settings"> -->
<!--       <SettingsIcon/> -->
<!--     </NavItem> -->
<!--     <NavItem link="/dashboard/servers"> -->
<!--       <ServerIcon/> -->
<!--     </NavItem> -->
<!--     <NavItem on:click={signOut}> -->
<!--       <DoorExitIcon/> -->
<!--     </NavItem> -->
<!--     <NavItem> -->
<!--       <DefaultUserIcon/> -->
<!--     </NavItem> -->
<!--   </div> -->
<!-- </NavBar> -->

<main class="flex justify-center mt-14">
  <ServersCardList/>
</main>
