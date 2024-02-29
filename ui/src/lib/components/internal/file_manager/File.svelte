<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";
    import TrashCanIcon from "../../heroicons/TrashCanIcon.svelte";
    import TextDocumentIcon from "$lib/components/heroicons/TextDocumentIcon.svelte";
    import DownloadIcon from "$lib/components/heroicons/DownloadIcon.svelte";
    import PencilIcon from "$lib/components/heroicons/PencilIcon.svelte";
    import ThreeDotsIcon from "$lib/components/heroicons/ThreeDotsIcon.svelte";
    import Checkbox from "$lib/components/generic/Checkbox.svelte";
    import FolderIcon from "$lib/components/heroicons/FolderIcon.svelte";
    import FileIcon from "$lib/components/heroicons/FileIcon.svelte";

    export let isFolder = false;
    export let name: string;
    export let lastEdited: Date;

    let isMenuOpen = false;

    onMount(() => {
        document.addEventListener("hideMenus", () => {
            isMenuOpen = false;
        });
    });

    let ref;
    function toggleMenu() {
        console.log("test");
        if (isMenuOpen === false) {
            const event  = new CustomEvent("hideMenus", {
                bubbles: true
            });
            ref.dispatchEvent(event);
        }
        isMenuOpen = !isMenuOpen;
    }
</script>

<section class="flex flex-row items-center justify-between py-8 px-12 border-b-indigo-900 border-b-2 text-white">
    <div class="flex flex-row items-center gap-3.5">
        <Checkbox/>
        {#if isFolder}
            <FolderIcon width="34" height="34"/>
        {:else}
            <FileIcon width="34" height="34"/>
        {/if}
        <span class="font-bold">{name}</span>
    </div>
    <div class="flex flex-row items-center gap-36 relative file-action-menu">
        <!--        <span>less than a minute ago</span>-->
        <span>{lastEdited.toLocaleString()}</span>
        <div class="relative z-10">
            <div class="{isMenuOpen ? '' : 'hidden'} absolute bg-indigo-900 w-64 -translate-x-20 translate-y-4 rounded-xl p-3 flex flex-col gap-2">
                <div on:click={() => { }} class="cursor-pointer hover:bg-indigo-700 rounded-xl py-2 px-3 flex flex-row items-center gap-3 text-center">
                    <DownloadIcon width={27} height={27}/>
                    <span class="font-semibold">Download</span>
                </div>
                <div on:click={() => { }} class="cursor-pointer hover:bg-indigo-700 rounded-xl py-2 px-3 flex flex-row items-center gap-3 text-center">
                    <TextDocumentIcon width={27} height={27}/>
                    <span class="font-semibold">Compress</span>
                </div>
                <div on:click={() => { }} class="cursor-pointer hover:bg-indigo-700 rounded-xl py-2 px-3 flex flex-row items-center gap-3 text-center">
                    <PencilIcon width={26} height={26}/>
                    <span class="font-semibold">Rename</span>
                </div>
                <div on:click={() => { }} class="cursor-pointer hover:bg-indigo-700 rounded-xl py-2 px-3 flex flex-row items-center gap-3 text-center">
                    <TrashCanIcon strokeWidth={2} width={26} height={26}/>
                    <span class="font-semibold">Delete</span>
                </div>
            </div>
        </div>
        <span bind:this={ref} role="button" tabindex="0" on:keydown={toggleMenu} on:click={toggleMenu}>
            <ThreeDotsIcon width="24" height="6"/>
        </span>
    </div>
</section>

<style>
</style>