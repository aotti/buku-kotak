<script>
    import { page } from "$app/stores";
    import Canvas from "$lib/components/Canvas.svelte";
    import History from "$lib/components/History.svelte";
    import Sidebar from "$lib/components/Sidebar.svelte";
    import { sharedStates } from "$lib/helper/misc-helper.svelte";

	let currentPage = $state('paper')
	const [cols, rows] = [Array(6), Array(5)]
    // get page section from search param
    const getSection = $page.url.searchParams.get('section')
    if(getSection) sharedStates.currentPage = getSection
</script>

<main class="grid grid-cols-12 gap-2 p-1 bg-darkbrown-3 text-white">
	<Sidebar />

	{#if sharedStates.currentPage == 'history'}
		<!-- history section -->
		<History />
	{:else}
		<!-- paper section -->
		<div class="col-span-10 md:col-span-6 lg:col-span-6 grid grid-cols-12 gap-1">
		{#each rows as row, r}
			{#each cols as col, c}
			<!-- 6 cols 5 rows -->
			<Canvas row={r+1} />
			{/each}
		{/each}
		</div>
	{/if}
</main>