<script lang="ts">

import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
import CpuIcon from "$lib/components/heroicons/CpuIcon.svelte";
import BoxIcon from "$lib/components/heroicons/BoxIcon.svelte";
import DriveIcon from "$lib/components/heroicons/DriveIcon.svelte";
</script>

<div class="bg-indigo-900 h-14 w-[95vw] rounded-lg flex items-center justify-between text-white px-4">
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
            <span class="font-lato font-semibold text-sm">69.42%</span>
        </div>
        <div class="flex gap-3 items-center">
            <DriveIcon/>
            <span class="font-lato font-semibold text-sm">69.42%</span>
        </div>
    </div>
</div>