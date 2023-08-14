<script lang="ts">
  export let className = "";
  export let height: "sm" | "md" | "lg" = "sm";
  export let headerHeight: "sm" | "md" = "sm";
  export let width: "sm" | "md" | "lg" = "sm";
  export let headerTextSize: "sm" | "md" | "lg" = "sm";

  const headerHeights = new Map<string, string>([
      ["sm", "h-14"],
      ["md", "h-20"],
      ["lg", ""],
  ]);

  const headerTextSizes = new Map<string, string>([
      ["sm", "text-base"],
  ]);

  const widths = new Map<string, string>([
      ["sm", ""],
      ["md", "w-[40rem]"],
      ["lg", "w-32"],
      ["screen", "w-[97.5vw]"]
  ]);

  const heights = new Map<string, string>([
      ["sm", "min-h-[20rem]"],
      ["md", "min-h-[28rem]"]
  ]);
</script>

<div class="drop-shadow-xl flex flex-col gap-0 {className}">
    <div class="bg-indigo-900 {headerHeights.get(headerHeight)} {widths.get(width)} rounded-t-lg flex gap-3 justify-start items-center">
    <span class="ml-7">
      <slot name="icon"/>
    </span>
        <span class="text-white {headerTextSizes.get(headerTextSize)} font-medium">
      <slot name="title"/>
    </span>
    </div>
    <div class="bg-indigo-950 rounded-b-lg {heights.get(height)} {widths.get(width)} w-full h-full">
        <div class="mt-5 w-full h-full">
            <slot name="content"/>
        </div>
    </div>
</div>