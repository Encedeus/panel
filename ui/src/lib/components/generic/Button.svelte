<script lang="ts">
  export let size: "sm" | "md" | "lg" | "xl" = "md";
  export let color: "indigo" | "red" = "indigo";
  export let type: "button" | "submit" = "button";
  export let className = "";

  const sizes = new Map<string, string>([
      ["sm", "w-24 h-8"],
      ["md", "w-32 h-11"],
      ["lg", ""],
      ["xl", ""],
  ]);

  const colors = new Map<string, string>([
      ["indigo", "bg-indigo-600 hover:bg-indigo-700 active:bg-indigo-800 "],
      ["red", "bg-red-600 hover:bg-red-700 active:bg-red-800"],
  ]);
</script>

<button on:click {type} class="text-white font-bold text-sm rounded-full {colors.get(color)} {sizes.get(size)} hover:shadow-xl active:shadow-xl transition-all {className}">
  <slot/>
</button>
