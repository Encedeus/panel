<script lang="ts">
  export let className = "";
  export let height: "sm" | "md" | "lg" | "xl" | "full" = "sm";
  export let headerHeight: "sm" | "md" = "sm";
  export let width: "sm" | "md" | "lg" | "screen" | "full" = "sm";
  export let headerTextSize: "sm" | "md" | "lg" = "sm";
  export let fixedHeight = false;
  export let fixedWidth = false;

  const headerHeights = new Map<string, string>([
      ["sm", "h-14"],
      ["md", "h-20"],
      ["lg", ""],
  ]);

  const headerTextSizes = new Map<string, string>([
      ["sm", "text-base"],
  ]);

  const widths = new Map<string, string>([
      ["sm", `${!fixedWidth ? "min-w-[30rem]" : "w-[30rem]"}`],
      ["md", `${!fixedWidth ? "min-[40rem]" : "w-[40rem]"}`],
      ["lg", `${!fixedWidth ? "min-w-[48rem]" : "w-[48rem]"}`],
      ["screen", `${!fixedWidth ? "min-[97.5vw]" : "w-[97.5vw]"}`],
      ["full", "w-full"]
  ]);

  const heights = new Map<string, string>([
      ["sm", `${fixedHeight ? "min-h-[20rem] max-h-[20rem]" : "min-h-[20rem]"}`],
      ["md", `${fixedHeight ? "min-h-[27.5rem] max-h-[27.5rem]" : "min-h-[27.5rem]"}`],
      ["lg", `${fixedHeight ? "min-h-[36rem] max-h-[36rem]" : "min-h-[36rem]"}`],
      ["xl", `${fixedHeight ? "min-h-[52rem] max-h-[52rem]" : "min-h-[52rem]"}`],
      ["full", "h-full"]
  ]);
</script>

<div class="drop-shadow-xl flex flex-col gap-0 {className} w-full">
    <div class="bg-indigo-900 {headerHeights.get(headerHeight)} {widths.get(width)} rounded-t-lg flex gap-3 justify-start items-center drop-shadow-2xl">
        <span class="ml-7">
            <slot name="icon"/>
        </span>
        <span class="text-white {headerTextSizes.get(headerTextSize)} font-medium">
            <slot name="title"/>
        </span>
    </div>
    <div class="bg-indigo-950 rounded-b-lg {fixedHeight ? 'overflow-y-auto' : ''} {heights.get(height)} {widths.get(width)} w-full">
        <slot name="content"/>
    </div>
</div>
