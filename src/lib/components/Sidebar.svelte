<script lang="ts">
    import { categoryHandler, filterHistory, screenshotThenUpload, sharedStates } from "$lib/helper/misc-helper.svelte";

    const sidebarButtons = [
        {id: 'paper', text: 'paper'},
        {id: 'history', text: 'history'},
    ]
</script>

<!-- desktop ver -->
<div class="hidden col-span-2 bg-darkbrown-2 p-1 md:flex lg:flex md:flex-col lg:flex-col md:gap-1 lg:gap-1">
    <h1> sidebar desktop </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
    <a href={`?section=save_upload`} class="border p-1" onclick={() => screenshotThenUpload()}> save & upload </a>
    {#if sharedStates.currentPage == 'history'}
        <input type="text" class="bg-darkbrown-3/70 border p-1" placeholder="search history" oninput={filterHistory}>
    {/if}
</div>

<!-- mobile ver -->
<div class="col-span-2 flex flex-col gap-1 bg-darkbrown-2 p-1 md:hidden lg:hidden">
    <h1> sidebar mobile </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
    <a href={`?section=save_upload`} class="border p-1" onclick={() => screenshotThenUpload()}> save & upload </a>
    {#if sharedStates.currentPage == 'history'}
        <input type="text" class="bg-darkbrown-3/70 border p-1" placeholder="search history" oninput={filterHistory}>
    {/if}
</div>