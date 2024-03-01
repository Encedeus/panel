<script lang="ts">
    import { onMount } from "svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import File from "$lib/components/internal/file_manager/File.svelte";
    import Breadcrumbs from "$lib/components/internal/file_manager/Breadcrumbs.svelte";
    import RightArrowIcon from "$lib/components/heroicons/RightArrowIcon.svelte";
    import RefreshIcon from "$lib/components/heroicons/RefreshIcon.svelte";
    import LeftArrowIcon from "$lib/components/heroicons/LeftArrowIcon.svelte";
    import type { PageServerData } from "./$types";
    import { fileManagerHistory, fileManagerPathIndex } from "$lib/store";
    import { page } from "$app/stores";
    import { goto, invalidate, invalidateAll } from "$app/navigation";

    let ref: HTMLDivElement;
    onMount(() => {
        document.body.addEventListener("click", (e) => {
            const els = document.getElementsByClassName("file-action-menu");
            for (const el of els) {
                if (el.contains(e.target)) {
                    return;
                }
            }

            const event  = new CustomEvent("hideMenus", {
                bubbles: true
            });
            ref?.dispatchEvent(event);
        });
    });


    onMount(() => {
        // $fileManagerHistory.push($page.params.path);
    });

    function goBack() {
        if ($fileManagerPathIndex === 0) {
            console.log("NO BACK")
            return;
        }

        $fileManagerPathIndex--;
        const url = location.href.split("files")[0];
        const target = `${url}files/${$fileManagerHistory[$fileManagerPathIndex]}`;
        // $fileManagerHistory.push($page.params.path);

        console.log(target);
        goto(target);
    }

    function goForward() {
        if ($fileManagerPathIndex === $fileManagerHistory.length - 1) {
            console.log("NO FORWARD")
            return;
        }

        $fileManagerPathIndex++;
        const url = location.href.split("files")[0];
        const target = `${url}files/${$fileManagerHistory[$fileManagerPathIndex]}`;

        console.log(target);
        goto(target);
    }

    function refresh() {
        invalidateAll();
        // goto(location.href);
    }

    export let data: PageServerData;
</script>

<main class="p-10 flex flex-col h-full">
    <CardHeader className="self-start mb-5 text-white font-inter" size="lg">
        File Manager
    </CardHeader>
    <Card headerHeight="md" width="full" height="xl">
            <span slot="title" class="gap-10 flex flex-row items-center">
                <span class="cursor-pointer" on:click={goBack}>
                    <LeftArrowIcon width={34} height={34}/>
                </span>
                <span class="cursor-pointer" on:click={goForward}>
                    <RightArrowIcon width={34} height={34}/>
                </span>
                <span class="cursor-pointer" on:click={refresh}>
                    <RefreshIcon width={34} height={34}/>
                </span>
                <Breadcrumbs crumbs={[{ name: "home", path: "" }, { name: "container", path: "" }]}></Breadcrumbs>
            </span>
        <div bind:this={ref} slot="content" class="h-full w-full">
            {#each data.files as f}
                <File name={f.name} isFolder={f.type === "d"} lastEdited={new Date(f.modifyTime)}/>
            {/each}
        </div>
    </Card>
</main>

<style>
</style>
