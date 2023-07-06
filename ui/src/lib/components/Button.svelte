<script lang="ts">
  export let size: "sm" | "md" | "lg" | "xl" = "md";
  export let color: "indigo" = "indigo";
  export let type: "button" | "submit" = "button";
  export let className = "";

  const sizes = new Map<string, string>([
    ["sm", ""],
    ["md", "w-32 h-11"],
    ["lg", ""],
    ["xl", ""],
  ]);

  const colors = new Map<string, string>([
    ["indigo", "bg-indigo-600"],
  ]);
</script>

<button {type} class="text-white font-bold text-sm rounded-full {colors.get(color)} {sizes.get(size)} {className}">
  <slot/>
</button>