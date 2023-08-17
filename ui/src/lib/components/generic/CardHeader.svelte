<script lang="ts">
    export let size: "sm" | "md" | "lg" = "sm";
    export let className = "";

    const sizes = new Map<string, string>([
        ["sm", "text-2xl"],
        ["md", "text-3xl"],
        ["lg", "text-4xl"],
    ]);
</script>

<h2 class="font-semibold text-white {sizes.get(size)} {className}">
  <slot/>
</h2>
