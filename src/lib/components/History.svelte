<script lang="ts">
    import { PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public";
    import { qS, sharedStates } from "$lib/helper/misc-helper.svelte";
    
    $effect(() => {
        if(sharedStates.currentPage == 'history') {
            fetch(`${PUBLIC_BUKU_KOTAK_API_URL}/history`, {method: 'GET', cache: 'no-store'})
            .then(data => data.json())
            .then(response => sharedStates.historyList = response.data)
            .catch((err) => qS('#history_error').textContent = err.message)
        }
    })
</script>

<div class="col-span-10 flex flex-wrap gap-4">
    <p id="history_error"></p>
    <!-- filtered history -->
    {#if sharedStates.historyFiltered.length > 0}
        {#each sharedStates.historyFiltered as history}
        {@const imageUrl = history}
        {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
        <p class="flex flex-col items-center cursor-pointer">
            <img src={imageUrl} alt={imageTitle} width="400" height="400">
            <span> {imageTitle.replaceAll('%20', ' ')} </span>
        </p>
        {/each}
    <!-- raw history -->
    {:else}
        {#each sharedStates.historyList as history}
        {@const imageUrl = history}
        {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
        <p class="flex flex-col items-center cursor-pointer">
            <img src={imageUrl} alt={imageTitle} width="400" height="400">
            <span> {imageTitle.replaceAll('%20', ' ')} </span>
        </p>
        {/each}
    {/if}
</div>