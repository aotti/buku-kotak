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
                sharedStates.historyListRow1 = []
                sharedStates.historyListRow2 = []
                sharedStates.historyListRow3 = []
                // fill history list
                for(let i=0; i<tempHistoryList.length; i++) {
                    if(i % 3 == 0) sharedStates.historyListRow1.push(tempHistoryList[i])
                    else if(i % 3 == 1) sharedStates.historyListRow2.push(tempHistoryList[i])
                    else if(i % 3 == 2) sharedStates.historyListRow3.push(tempHistoryList[i])
                }
            })
            .catch((err) => qS('#history_error').textContent = err.message)
        }
    })

    function handleXScroll(ev: MouseEvent & {currentTarget: EventTarget & HTMLButtonElement}) {
        const historyContainer = qS('#history')
        // set scroll data
        const scrollDirection = ev.currentTarget.title.match('right') ? 'right' : 'left'
        const isUserMobile = screen.orientation.angle === 0 ? 'desktop' : 'mobile'
        const currentScroll = historyContainer.scrollLeft
        const scrollPower = isUserMobile == 'desktop' ? 800 : 400
        // scroll event
        historyContainer.scrollTo({
            behavior: 'smooth', 
            left: scrollDirection == 'right' ? currentScroll + scrollPower : currentScroll - scrollPower
        })
    }
</script>

<div id="history" class="relative col-span-10 flex flex-col gap-4 overflow-x-scroll">
    <!-- scroll horizontal button -->
    <div class="fixed top-1/3 w-[80vw] mx-3">
        <div class="flex justify-between">
            <button title="left-arrow" class="bg-darkbrown-2/50 rounded-full cursor-pointer opacity-25 hover:opacity-100" onclick={handleXScroll}>
                <img src="https://img.icons8.com/?size=100&id=122595&format=png&color=000000" alt="left-arrow">
            </button>
            <button title="right-arrow" class="bg-darkbrown-2/50 rounded-full cursor-pointer opacity-25 hover:opacity-100" onclick={handleXScroll}>
                <img src="https://img.icons8.com/?size=100&id=111410&format=png&color=000000" alt="right-arrow">
            </button>
        </div>
    </div>
    <!-- history fetch error -->
    <p id="history_error"></p>
    <!-- history row 1 -->
    <div class="flex gap-4">
        <!-- filtered history -->
        {#if sharedStates.historyFilteredRow1.length > 0}
            {#each sharedStates.historyFilteredRow1 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        <!-- raw history -->
        {:else}
            {#each sharedStates.historyListRow1 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        {/if}
    </div>

    <!-- history row 2 -->
    <div class="flex gap-4">
        <!-- filtered history -->
        {#if sharedStates.historyFilteredRow2.length > 0}
            {#each sharedStates.historyFilteredRow2 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        <!-- raw history -->
        {:else}
            {#each sharedStates.historyListRow2 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        {/if}
    </div>

    <!-- history row 3 -->
    <div class="flex gap-4">
        <!-- filtered history -->
        {#if sharedStates.historyFilteredRow3.length > 0}
            {#each sharedStates.historyFilteredRow3 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        <!-- raw history -->
        {:else}
            {#each sharedStates.historyListRow3 as history}
            {@const imageUrl = history}
            {@const imageTitle = history.match(/(?<=buku-kotak\/).*(?=.png)/)[0]}
                <p class="flex flex-col items-center border cursor-pointer shrink-0 basis-96">
                    <img src={imageUrl} alt={imageTitle} width="480" height="480">
                    <span> {imageTitle.replaceAll('%20', ' ')} </span>
                </p>
            {/each}
        {/if}
    </div>
</div>