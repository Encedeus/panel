<script lang="ts">
    export let size: "sm" | "md" | "lg" | "xl" = "sm";
    export let value = "";
    export let placeholder = "";
    export let className = "";
    export let label = "";
    export let error = false;

    const sizes: { [size: string] : { rows: number; cols: number; } } = {
        lg: {
            rows: 10,
            cols: 78,
        },
        md: {
            rows: 7,
            cols: 22,
        }
    };
</script>

<div class="flex flex-col items-start gap-0.5">
  {#if label}
    <label class="text-white font-bold text-[1rem] ml-4">{label}</label>
  {/if}
  <textarea on:input bind:value={value} {placeholder} id="" cols={sizes[size].cols} rows={sizes[size].rows} class=" {!error ? 'focus:outline-indigo-500' : 'outline-red-500'} outline-offset-0 outline-none rounded-3xl focus:outline outline-2 border-none pl-7 py-3 font-inter font-bold text-xs placeholder-white placeholder-opacity-25 text-white bg-indigo-900 {className}"/>
</div>
