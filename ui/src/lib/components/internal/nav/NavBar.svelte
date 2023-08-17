<script>
    import Search from "$lib/components/internal/search/Search.svelte";
</script>

<nav class="w-screen sticky flex gap-0 h-20">
    <div class="bg-indigo-900 basis-[20vw] xl:basis-[13.5vw] flex items-center justify-center">
        <slot name="logo"/>
    </div>
    <div class="bg-indigo-800 flex-grow flex items-center justify-between">
        <Search className="ml-8" size="lg"/>
        <slot name="links"/>
    </div>
</nav>
