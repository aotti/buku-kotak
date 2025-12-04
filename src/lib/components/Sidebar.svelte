<script lang="ts">
    import { PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public";
    import { qS, sharedStates } from "$lib/helper/misc-helper.svelte";
    import html2canvas from "html2canvas";

    const sidebarButtons = [
        {id: 'paper', text: 'paper'},
        {id: 'history', text: 'history'},
    ]

    function categoryHandler(id: string) {
        sharedStates.currentPage = id
    }

    function screenshotThenUpload() {
        const imageTitle = prompt('Please enter image title')
        if(!imageTitle) return alert('image title cannot be empty')
        // set element to screenshot
        html2canvas(qS('#paper'))
        // generate element to canvas
        .then(async canvas => {
            const imageData = {
                public_id: imageTitle,
                // then convert to base64
                img_base64: canvas.toDataURL()
            }
            const fetchOptions: RequestInit = {
                method: 'POST',
                headers: {'content-type': 'application/json'},
                body: JSON.stringify(imageData)
            }
            const imageUpload = await fetch(`${PUBLIC_BUKU_KOTAK_API_URL}/upload`, fetchOptions)
            if(imageUpload.status < 400) return alert('image upload success!')
        })
    }
</script>

<!-- desktop ver -->
<div class="hidden col-span-2 bg-darkbrown-2 p-1 md:flex lg:flex md:flex-col lg:flex-col md:gap-1 lg:gap-1">
    <h1> sidebar desktop </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
    <a href={`?section=save_upload`} class="border p-1" onclick={() => screenshotThenUpload()}> save & upload </a>
</div>

<!-- mobile ver -->
<div class="col-span-2 flex flex-col gap-1 bg-darkbrown-2 p-1 md:hidden lg:hidden">
    <h1> sidebar mobile </h1>
    {#each sidebarButtons as button}
    <a href={`?section=${button.id}`} class="border p-1" onclick={() => categoryHandler(button.id)}> {button.text} </a>
    {/each}
    <a href={`?section=save_upload`} class="border p-1" onclick={() => screenshotThenUpload()}> save & upload </a>
</div>