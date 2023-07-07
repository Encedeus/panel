<script lang="ts">
  export let size: "sm" | "md" | "lg" | "xl" = "sm";
  export let value = "";
  export let placeholder = "";
  export let className = "";
  export let label = "";
  export let type: "password" | "text" = "text";

  const sizes = new Map<string, string>([
    ["sm", "pr-32"],
    ["md", "pr-56"],
    ["lg", "pr-72"],
    ["xl", "pr-96"],
  ]);

</script>

<div class="flex flex-col items-start gap-0.5">
  {#if label}
    <label class="text-white font-bold text-[1rem] ml-4">{label}</label>
  {/if}
  {#if type = "password"}
    <input bind:value={value} {placeholder} type="password" class="focus:outline-indigo-500 -outline-offset-0 focus:outline outline-none border-none {sizes.get(size)} pl-7 py-3 font-inter font-bold text-xs rounded-full placeholder-white placeholder-opacity-25 text-white bg-indigo-900 {className}">
  {:else}
    <input bind:value={value} {placeholder} type="text" class="focus:outline-indigo-500 -outline-offset-0 focus:outline outline-none border-none {sizes.get(size)} pl-7 py-3 font-inter font-bold text-xs rounded-full placeholder-white placeholder-opacity-25 text-white bg-indigo-900 {className}">
  {/if}
</div>

<style>
</style>