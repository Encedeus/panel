<script>
    export let items = [];
    export let activeTabValue = 1;
    export let className = "";

    const handleClick = tabValue => () => (activeTabValue = tabValue);
</script>

<ul class="{className}">
    {#if items}
        {#each items as item}
            <li class={activeTabValue === item.value ? 'active' : ''}>
                <span on:click={handleClick(item.value)}>{item.label}</span>
            </li>
        {/each}
    {/if}
</ul>

<style lang="postcss">
    ul {
        display: flex;
        flex-wrap: wrap;
        justify-content: left;
        list-style: none;
        border-color: black;
    }

    span {
        border-bottom: 2px solid transparent;
        border-top: 2px transparent;
        /*display: block;*/
        padding: .35rem 1rem;
        padding-inline: 15px;
        cursor: pointer;
        border-color: black;

    }

    span:hover {
        background-color: rgb(49 46 129);
    }

    li.active > span {
        border-color: #4F46E5;
    }
</style>