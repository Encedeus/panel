<script lang="ts">
    import Button from "$lib/components/generic/Button.svelte";
    import { createEventDispatcher } from "svelte";

    export let label = "";
    export let value = "";

    const dispatch = createEventDispatcher();
</script>

<div class="flex items-center justify-between py-3 px-5 bg-indigo-900 rounded-xl self-stretch">
    <div class="flex items-center gap-3">
      <span>
          <slot name="icon"/>
      </span>
      <div class="flex flex-col">
          <span class="text-white">{label}</span>
          <span class="text-white text-opacity-25 text-xs">{value}</span>
      </div>
    </div>
    <div>
        <Button color="indigo" size="sm" on:click={() => dispatch("change")}>Change</Button>
    </div>
</div>
