<script lang="ts">
  export let mode: "error" | "ok" | "warning" | "info" = "ok";
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
  ]);

  const icons = new Map<string, string>([
      ["error", "    <svg width=\"14\" height=\"14\" viewBox=\"0 0 14 14\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n" +
    "      <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M0.227806 0.227806C0.531547 -0.0759353 1.02401 -0.0759353 1.32775 0.227806L7 5.90006L12.6722 0.227807C12.976 -0.0759344 13.4685 -0.0759344 13.7722 0.227807C14.0759 0.531548 14.0759 1.02401 13.7722 1.32775L8.09994 7L13.7722 12.6722C14.0759 12.976 14.0759 13.4685 13.7722 13.7722C13.4685 14.0759 12.976 14.0759 12.6722 13.7722L7 8.09994L1.32775 13.7722C1.02401 14.0759 0.531547 14.0759 0.227806 13.7722C-0.0759352 13.4685 -0.0759352 12.976 0.227806 12.6722L5.90006 7L0.227806 1.32775C-0.0759353 1.02401 -0.0759353 0.531547 0.227806 0.227806Z\" fill=\"white\"/>\n" +
    "    </svg>"]
  ]);
</script>

<div class="{className} {colors.get(mode)} h-10 w-80 rounded-lg flex flex-row items-center text-white">
  <span class="ml-3.5">
    {@html icons.get(mode)}
  </span>
  <span class="ml-3.5 font-monti font-semibold text-sm">
    <slot name="label"/>
  </span>
</div>