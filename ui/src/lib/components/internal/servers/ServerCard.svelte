<script lang="ts">
    import ServerStackIcon from "$lib/components/heroicons/ServerStackIcon.svelte";
    import CpuIcon from "$lib/components/heroicons/CpuIcon.svelte";
    import BoxIcon from "$lib/components/heroicons/BoxIcon.svelte";
    import DriveIcon from "$lib/components/heroicons/DriveIcon.svelte";
    import type {Server} from "@encedeus/js-api";
    export let className = "";
    export let server: Server;
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
            <span class="font-lato font-semibold text-sm">{server.ipv4}</span>
        </div>
        <div class="flex gap-3 items-center">
            <CpuIcon/>
            <!-- todo: implement realtime data -->
            <span class="font-lato font-semibold text-sm">69.420%</span>
        </div>
        <div class="flex gap-3 items-center">
            <BoxIcon/>
            <span class="flex flex-col gap-0">
                <span class="font-lato font-semibold text-sm">
                    0 kB
                </span>
                <span class="font-lato font-bold text-white text-opacity-25 text-[10px]">
                    Of {server.ram}
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
                    Of {server.storage}
                </span>
            </span>
        </div>
    </div>
</div>
