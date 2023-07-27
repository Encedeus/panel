<nav class="w-screen sticky flex gap-0 h-20">
    <div class="bg-indigo-900 basis-2/12 flex items-center justify-center">
        <slot name="logo"/>
    </div>
    <div class="bg-indigo-800 flex-grow flex justify-between">
        <div></div>
        <slot name="links"/>
    </div>
</nav>