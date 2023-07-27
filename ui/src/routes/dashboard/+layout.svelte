<script>
  import NavBar from "$lib/components/internal/nav/NavBar.svelte";
  import NavItem from "$lib/components/internal/nav/NavItem.svelte";
  import SettingsIcon from "$lib/components/heroicons/SettingsIcon.svelte";
  import ServerIcon from "$lib/components/heroicons/ServerIcon.svelte";
  import DoorExitIcon from "../../lib/components/heroicons/DoorExitIcon.svelte";
  import DefaultUserIcon from "../../lib/components/heroicons/DefaultUserIcon.svelte";

</script>

<NavBar>
  <div slot="logo" class="text-white text-3xl xl:text-4xl font-bold">Encedeus</div>
  <div class="flex items-center justify-between gap-3 mr-5" slot="links">
    <NavItem link="/dashboard/account/settings">
      <SettingsIcon/>
    </NavItem>
    <NavItem link="/dashboard/servers">
      <ServerIcon/>
    </NavItem>
    <NavItem>
      <DoorExitIcon/>
    </NavItem>
    <NavItem>
      <DefaultUserIcon/>
    </NavItem>
  </div>
</NavBar>

<slot/>