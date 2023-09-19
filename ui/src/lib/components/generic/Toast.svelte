<script context="module" lang="ts">
  export type NotificationMode = "error" | "ok" | "warning" | "info";
</script>

<script lang="ts">
    import CheckmarkIcon from "$lib/components/heroicons/CheckmarkIcon.svelte";
    import ErrorIcon from "$lib/components/heroicons/ErrorIcon.svelte";

    export let mode: NotificationMode = "ok";
    export let label = "";
    export let size: "sm" | "md" | "lg" = "md";
    export let className = "";

    const sizes = new Map<string, string>([
        ["sm", ""],
        ["md", ""],
        ["lg", ""],
    ]);

    const colors = new Map<string, string>([
        ["error", "bg-red-500"],
        ["ok", "bg-blue-600"],
    ]);

    const icons = new Map<string, string>([
        ["error", ""],
        ["ok", "<CheckmarkIcon width={28} height={28}/>"],
    ]);
</script>

<div class="{className} {colors.get(mode)} h-10 w-80 rounded-lg flex flex-row items-center text-white">
  <span class="ml-3.5 flex justify-center items-center">
    {#if mode === "ok"}
      <CheckmarkIcon width={28} height={28}/>
    {:else if mode === "error"}
      <ErrorIcon width={30} height={30}/>
    {/if}
  </span>
    <span class="ml-3.5 font-monti font-semibold text-sm">
    <slot/>
  </span>
</div>