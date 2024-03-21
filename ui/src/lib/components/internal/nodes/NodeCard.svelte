<script lang="ts">
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
    import CpuIcon from "$lib/components/heroicons/CpuIcon.svelte";
    import BoxIcon from "$lib/components/heroicons/BoxIcon.svelte";
    import DriveIcon from "$lib/components/heroicons/DriveIcon.svelte";
    export let className = "";
</script>

<div class="bg-indigo-900 h-14 rounded-lg flex items-center justify-between text-white px-4 {className}">
    <div class="flex gap-3">
            <span class="self-center">
                <ServerStackIcon/>
            </span>
        <span class="font-medium">Minecraft</span>
    </div>
    <div class="flex gap-11 items-center">
        <div class="flex gap-3 items-center">
            <ServerStackIcon/>
            <span class="font-lato font-semibold text-sm">192.168.0.1:25565</span>
        </div>
        <div class="flex gap-3 items-center">
            <CpuIcon/>
            <span class="font-lato font-semibold text-sm">69.42%</span>
        </div>
        <div class="flex gap-3 items-center">
            <BoxIcon/>
            <span class="flex flex-col gap-0">
                <span class="font-lato font-semibold text-sm">
                    0 kB
                </span>
                <span class="font-lato font-bold text-white text-opacity-25 text-[10px]">
                    Of 128 MB
                </span>
            </span>
        </div>
        <div class="flex gap-3 items-center">
            <DriveIcon/>
            <span class="flex flex-col gap-0">
                <span class="font-lato font-semibold text-sm">
                    2.7 MB
                </span>
                <span class="font-lato font-bold text-white text-opacity-25 text-[10px]">
                    Of 512 MB
                </span>
            </span>
        </div>
    </div>
</div>
