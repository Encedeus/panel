<script lang="ts">
  export let className = "";
  export let height: "sm" | "md" | "lg" = "sm";
  export let headerHeight: "sm" | "md" = "sm";
  export let width: "sm" | "md" | "lg" | "screen" = "sm";
  export let headerTextSize: "sm" | "md" | "lg" = "sm";
  export let fixedHeight = false;

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
      ["lg", "w-[48rem]"],
      ["screen", "w-[97.5vw]"]
  ]);

  const heights = new Map<string, string>([
      ["sm", `${fixedHeight ? "max-h-[20rem]" : "min-h-[20rem]"}`],
      ["md", `${fixedHeight ? "max-h-[27.5rem]" : "min-h-[27.5rem]"}`]
  ]);
</script>

<div class="drop-shadow-xl flex flex-col gap-0 {className} w-full h-full">
    <div class="bg-indigo-900 {headerHeights.get(headerHeight)} {widths.get(width)} rounded-t-lg flex gap-3 justify-start items-center">
    <span class="ml-7">
      <slot name="icon"/>
    </span>
        <span class="text-white {headerTextSizes.get(headerTextSize)} font-medium">
      <slot name="title"/>
    </span>
    </div>
    <div class="bg-indigo-950 rounded-b-lg {fixedHeight ? 'overflow-y-auto' : ''} {heights.get(height)} {widths.get(width)}">
        <slot name="content"/>
        <!-- <slot/> -->
    </div>
</div>
