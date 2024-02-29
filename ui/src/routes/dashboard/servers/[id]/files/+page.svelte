<script lang="ts">
    import { onMount } from "svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import Card from "$lib/components/generic/Card.svelte";
    import File from "$lib/components/internal/file_manager/File.svelte";
    import Breadcrumbs from "$lib/components/internal/file_manager/Breadcrumbs.svelte";
    import RightArrowIcon from "$lib/components/heroicons/RightArrowIcon.svelte";
    import RefreshIcon from "$lib/components/heroicons/RefreshIcon.svelte";
    import LeftArrowIcon from "$lib/components/heroicons/LeftArrowIcon.svelte";

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

    let urlParams;
    let serverId;
    onMount(() => {
        urlParams = new URLSearchParams(window.location.search);
        serverId = urlParams.get("serverId");
    })
/*    const urlParams = new URLSearchParams(window.location.search);
    const serverId = urlParams.get("serverId");*/
</script>

<main class="p-10 flex flex-col h-full">
    <CardHeader className="self-start mb-5 text-white font-inter" size="lg">
        File Manager
    </CardHeader>
    <Card headerHeight="md" width="full" height="xl">
            <span slot="title" class="gap-10 flex flex-row items-center">
                <LeftArrowIcon width={34} height={34}/>
                <RightArrowIcon width={34} height={34}/>
                <RefreshIcon width={34} height={34}/>
                <Breadcrumbs crumbs={[{ name: "home", path: "" }, { name: "container", path: "" }]}></Breadcrumbs>
            </span>
        <div bind:this={ref} slot="content" class="h-full w-full">
            <File name="testFile.txt" lastEdited={new Date(Date.now())}/>
            <File name="apps" isFolder={true} lastEdited={new Date(Date.now())}/>
        </div>
    </Card>
</main>

<style>
</style>
