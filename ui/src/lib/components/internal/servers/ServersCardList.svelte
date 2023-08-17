<script lang="ts">
  import Card from "$lib/components/generic/Card.svelte";
  import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
  import ServerCard from "$lib/components/internal/servers/ServerCard.svelte";

  export let className = "";
</script>

<div>
  <Card headerHeight="sm" height="md" width="screen" className={className}>
    <span slot="title">Servers</span>
    <ServerStackIcon slot="icon"/>
    <div slot="content" class="flex items-center flex-col gap-2">
      <ServerCard/>
    </div>
  </Card>
</div>
