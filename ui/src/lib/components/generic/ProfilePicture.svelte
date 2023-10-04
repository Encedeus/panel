<script lang="ts">
    import DefaultIcon from "$lib/assets/generic-user.png";
    import EditIcon from "$lib/components/heroicons/EditIcon.svelte";
    import { createEventDispatcher } from "svelte";

    export let width = 32;
    export let height = 32;
    export let src = DefaultIcon;
    export let user = "";
    export let className = "";
    export let changeable = false;

    let fileInput: HTMLInputElement;

    const dispatch = createEventDispatcher();

    function onClick() {
        fileInput.click();

        let file: File;
        fileInput.onchange = () => {
            file = fileInput.files.item(0);
        };

        dispatch("change", {
            file
        });
    }
</script>

{#if changeable}
    <span role="button" tabindex="0" on:keydown={onClick} on:click={onClick} class="rounded-full cursor-pointer transition-all hover:bg-indigo-700 active:bg-indigo-800 bg-indigo-600 p-2.5 absolute right-32 top-[82%]">
        <EditIcon/>
    </span>
    <input bind:this={fileInput} type="file" class="hidden">
{/if}
<img class="rounded-full border-2 border-white {className}" {width} {height} {src} alt="{user}'s profile picture">
