<script>
    import Canvas from "$lib/components/Canvas.svelte";
    import Sidebar from "$lib/components/Sidebar.svelte";

	const [cols, rows] = [Array(6), Array(5)]
</script>

<main class="flex gap-2 p-1 bg-darkbrown-3 text-white">
	<Sidebar />
	<!-- 6 cols 5 rows -->
	<div class="grid grid-cols-12 gap-1">
	{#each rows as row, r}
		{#each cols as col, c}
		<Canvas row={r+1} />
		{/each}
	{/each}
	</div>
</main>