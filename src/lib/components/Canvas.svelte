<script lang="ts">
    import { canvasCopy, canvasPaste, disableScrolling, enableScrolling, qS, sharedStates } from '$lib/helper/misc-helper.svelte';
	import { onMount } from 'svelte'

	const {row, col} = $props()
    let width = 100,
        height = 100,
        color = '#000',
        background = '#fff';

	let canvas: HTMLCanvasElement
	let context: CanvasRenderingContext2D
	let isDrawing: boolean
	let start: {x: number, y: number}
	
	let t: number, l: number
	
	onMount(() => {
		context = canvas.getContext('2d')
		context.lineWidth = 3
		handleSize()
	})
	
	const handleStart = ({ offsetX: x, offsetY: y }) => { 
		if(color === background) {
			context.clearRect(0, 0, width, height)
		} else {
			disableScrolling()
			isDrawing = true
			start = { x, y }
		}
	}
	
	const handleEnd = () => { 
		enableScrolling()
		isDrawing = false 
	}
	const handleMove = (({ offsetX: x1, offsetY: y1 }) => {
		if(!isDrawing) return
		
		const { x, y } = start
		context.beginPath()
		context.moveTo(x, y)
		context.lineTo(x1, y1)
		context.closePath()
		context.stroke()
		
		start = { x: x1, y: y1 }
	})
	
	const handleSize = () => {
		const { top, left } = canvas.getBoundingClientRect()
		t = top
		l = left
	}
</script>

<svelte:window on:resize={handleSize} />

<div class="col-span-2 flex flex-col">
	<input type="text" class="border bg-darkbrown-1 text-center h-8" placeholder="~">
	<canvas id={`canvas_${row}${col}`} class="bg-darkbrown-4"
		{width}
		{height}
		bind:this={canvas} 
		onmousedown={handleStart}	
		ontouchstart={e => {
			const { clientX, clientY } = e.touches[0]
			// for mobile, handle size every touch start
			// because the viewport position changed on scroll
			handleSize()
			handleStart({
				offsetX: clientX - l,
				offsetY: clientY - t
			})
		}}	
		onmouseup={handleEnd}				
		ontouchend={handleEnd}				
		onmouseleave={handleEnd}
		onmousemove={handleMove}
		ontouchmove={e => {
			const { clientX, clientY } = e.touches[0]
			handleMove({
				offsetX: clientX - l,
				offsetY: clientY - t
			})
		}}></canvas>
	<!-- canvas action -->
	<div class="flex">
		<!-- clear button -->
		<button class="grow bg-darkbrown-2 border" onclick={ev => context.clearRect(0, 0, width, height)}> 
			{`cls - ${row}`}
		</button>
		<!-- copy & paste button -->
		{#if sharedStates.canvasCopy != null}
			<button class="grow bg-darkbrown-2 border" onclick={() => canvasPaste(`#canvas_${row}${col}`)}>
				{`pste - ${col}`}
			</button>
		{:else}
			<button class="grow bg-darkbrown-2 border" onclick={() => canvasCopy(`#canvas_${row}${col}`)}>
				{`cpy - ${col}`}
			</button>
		{/if}
	</div>
</div>