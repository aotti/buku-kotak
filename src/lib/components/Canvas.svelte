<script lang="ts">
	import { onMount } from 'svelte'

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
	
	$effect(() => {
        if(context) {
			context.strokeStyle = color
	    }
    })
	
	const handleStart = ({ offsetX: x, offsetY: y }) => { 
		if(color === background) {
			context.clearRect(0, 0, width, height)
		} else {
			isDrawing = true
			start = { x, y }
		}
	}
	
	const handleEnd = () => { isDrawing = false }
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

<canvas class="border-2"
    {width}
    {height}
    style:background
    bind:this={canvas} 
    onmousedown={handleStart}	
    ontouchstart={e => {
        const { clientX, clientY } = e.touches[0]
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
    
<button class="border" style:background onclick={ev => {
    context.clearRect(0, 0, width, height)
}}> clear </button>