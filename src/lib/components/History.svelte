<script lang="ts">
    import { PUBLIC_BUKU_KOTAK_API_DEMO, PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public";
    import { qS, sharedStates } from "$lib/helper/misc-helper.svelte";
    
    $effect(() => {
        if(sharedStates.currentPage == 'history') {
            fetch(`${PUBLIC_BUKU_KOTAK_API_URL}/history`, {method: 'GET', cache: 'no-store'})
            .then(data => data.json())
            .then(response => {
                const tempHistoryList = response.data as string[]
                // set history list to empty everytime fetching history
                sharedStates.historyList = []
                // fill history list
                for(let i=tempHistoryList.length-1; i>=0; i--) {
                    const filename = tempHistoryList[i].split('/buku-kotak/')[1].replaceAll('%20', ' ')
                    sharedStates.historyList.push({
                        filename, link: tempHistoryList[i]
                    })
                }
            })
            .catch((err) => qS('#history_error').textContent = err.message)
        }
    })

    function handlePage(ev: MouseEvent & {currentTarget: EventTarget & HTMLButtonElement}) {
        const pageDirection = ev.currentTarget.title.match('right') ? 'right' : 'left'
        const [first, last] = sharedStates.historyPage
        sharedStates.historyPage = pageDirection == 'left' 
                                // dont move page if out of bounds
                                // move left
                                ? first-8 < 0 
                                    ? [0, 8] 
                                    : [first-8, last-8] 
                                // move right
                                : last+8 > sharedStates.historyList.length-1
                                    ? [sharedStates.historyList.length-9, sharedStates.historyList.length-1]
                                    : [first+8, last+8]
    }
</script>

<div id="history" class="relative col-span-10 flex flex-col gap-4 overflow-x-scroll">
    <!-- scroll horizontal button -->
    <div class="fixed top-1/3 w-[80vw] mx-3">
        <div class="flex justify-between">
            <button title="left-arrow" class="bg-darkbrown-2/50 rounded-full cursor-pointer opacity-25 hover:opacity-100" onclick={handlePage}>
                <img src="https://img.icons8.com/?size=100&id=122595&format=png&color=000000" alt="left-arrow">
            </button>
            <button title="right-arrow" class="bg-darkbrown-2/50 rounded-full cursor-pointer opacity-25 hover:opacity-100" onclick={handlePage}>
                <img src="https://img.icons8.com/?size=100&id=111410&format=png&color=000000" alt="right-arrow">
            </button>
        </div>
    </div>
    <!-- history fetch error -->
    <p id="history_error"></p>
    <!-- history row 1 -->
    <div class="grid grid-cols-3 gap-4">
        <!-- filtered history -->
        {#if sharedStates.historyFiltered.length > 0}
            {#each sharedStates.historyFiltered as history,i}
            {@const [first, last] = sharedStates.historyPage}
            {#if i >= first && i <= last}
                <p class="flex flex-col items-center cursor-pointer shrink-0 basis-96">
                    <img src={history.link} alt={history.filename} width="480" height="480">
                    <span> {history.filename.replaceAll('%20', ' ')} </span>
                </p>
            {/if}
            {/each}
        <!-- raw history -->
        {:else}
            {#each sharedStates.historyList as history,i}
            {@const [first, last] = sharedStates.historyPage}
            {#if i >= first && i <= last}
                <p class="flex flex-col items-center cursor-pointer shrink-0 basis-96">
                    <img src={history.link} alt={history.filename} width="480" height="480">
                    <span> {history.filename.replaceAll('%20', ' ')} </span>
                </p>
            {/if}
            {/each}
        {/if}
    </div>
</div>