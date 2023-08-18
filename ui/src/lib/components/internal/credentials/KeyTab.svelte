<script lang="ts">
    import KeyIcon from "$lib/components/heroicons/KeyIcon.svelte";
  import TrashCanIcon from "$lib/components/heroicons/TrashCanIcon.svelte";

    export let name = "";
    export let key = "";
    export let lastUsed = "";
    export let className = "";
</script>

<div class="flex items-center justify-between py-4 px-6 bg-indigo-900 rounded-xl {className}">
   <div class="flex items-center justify-center gap-3">
        <KeyIcon width={34} height={34}/>
        <div class="flex flex-col items-start justify-center">
            <span class="text-white text-lg font-semibold">{name}</span>
            <span class="text-white text-opacity-25 text-[9px] font-semibold -mt-0.5">{`Last Used: ${lastUsed.toUpperCase()}`}</span>
        </div>
   </div>
   <div class="flex items-center gap-6">
        <span class="rounded-xl bg-indigo-950 text-white text-sm py-1.5 px-7">{key}</span>
        <span class="hover:cursor-pointer">
            <TrashCanIcon width={34} height={34}/>
        </span>
   </div>
</div>
