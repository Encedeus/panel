<script lang="ts">
  export let size: "sm" | "md" | "lg" | "xl" = "sm";
  export let value = "";
  export let placeholder = "";
  export let className = "";
  export let label = "";
  export let type: "password" | "text" | "search" = "text";
  export let error = false;

  const sizes = new Map<string, string>([
      ["sm", "pr-32"],
      ["md", "pr-56"],
      ["lg", "pr-72"],
      ["xl", "w-[43rem]"],
  ]);

</script>

<div class="flex flex-col items-start gap-0.5">
  {#if label}
    <label class="text-white font-bold text-[1rem] ml-4">{label}</label>
  {/if}
  {#if type === "password"}
    <input on:input bind:value={value} {placeholder} type="password"
           class=" {!error ? 'focus:outline-indigo-500' : 'outline-red-500'} outline-offset-0 outline-none focus:outline outline-2 border-none {sizes.get(size)} pl-7 py-3 font-inter font-bold text-xs rounded-full placeholder-white placeholder-opacity-25 text-white bg-indigo-900 {className}">
  {:else}
    <input on:input bind:value={value} {placeholder} type="text"
           class="{!error ? 'focus:outline-indigo-500' : 'outline-red-500'} outline-offset-0 outline-none focus:outline outline-2 border-none {sizes.get(size)} pl-7 py-3 font-inter font-bold text-xs rounded-full placeholder-white placeholder-opacity-25 text-white bg-indigo-900 {className}">
  {/if}
</div>

