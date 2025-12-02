<script lang="ts">
    import { sharedStates } from "$lib/helper/misc-helper.svelte";

    const sidebarButtons = [
        {id: 'paper', text: 'paper'},
        {id: 'history', text: 'history'},
        {id: 'save_upload', text: 'save & upload'},
    ]

    function categoryHandler(id: string) {
        sharedStates.currentPage = id
    }
</script>

<!-- desktop ver -->
<div class="hidden col-span-2 bg-darkbrown-2 p-1 md:flex lg:flex md:flex-col lg:flex-col md:gap-1 lg:gap-1">
    <h1> sidebar desktop </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
</div>

<!-- mobile ver -->
<div class="col-span-2 flex flex-col gap-1 bg-darkbrown-2 p-1 md:hidden lg:hidden">
    <h1> sidebar mobile </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
</div>