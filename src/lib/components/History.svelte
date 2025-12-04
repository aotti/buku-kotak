<script lang="ts">
    import { PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public";
    import { sharedStates } from "$lib/helper/misc-helper.svelte";

    const historyList = [
        {name: 'belajar huruf lele', img: 'https://cdn.discordapp.com/attachments/826777163327602698/1443140862987997274/s25voucher.png?ex=692fe629&is=692e94a9&hm=4f4b48563079cc23aabcfd6ea25ba1cae47cb57d1c558fc7d5fe88adbb0a3ecf&'},
        {name: 'belajar huruf bapak', img: 'https://cdn.discordapp.com/attachments/818629322965385236/1177256050634797186/Thanksgiving.png?ex=692fa1e2&is=692e5062&hm=141a36a1e28d2a360c762aa5e9604ff23d04922a36ddad6693c43492c7674b37&'},
        {name: 'belajar huruf lele', img: 'https://cdn.discordapp.com/attachments/826777163327602698/1443140862987997274/s25voucher.png?ex=692fe629&is=692e94a9&hm=4f4b48563079cc23aabcfd6ea25ba1cae47cb57d1c558fc7d5fe88adbb0a3ecf&'},
        {name: 'belajar huruf bapak', img: 'https://cdn.discordapp.com/attachments/818629322965385236/1177256050634797186/Thanksgiving.png?ex=692fa1e2&is=692e5062&hm=141a36a1e28d2a360c762aa5e9604ff23d04922a36ddad6693c43492c7674b37&'},
        {name: 'belajar huruf lele', img: 'https://cdn.discordapp.com/attachments/826777163327602698/1443140862987997274/s25voucher.png?ex=692fe629&is=692e94a9&hm=4f4b48563079cc23aabcfd6ea25ba1cae47cb57d1c558fc7d5fe88adbb0a3ecf&'},
        {name: 'belajar huruf bapak', img: 'https://cdn.discordapp.com/attachments/818629322965385236/1177256050634797186/Thanksgiving.png?ex=692fa1e2&is=692e5062&hm=141a36a1e28d2a360c762aa5e9604ff23d04922a36ddad6693c43492c7674b37&'},
        {name: 'belajar huruf lele', img: 'https://cdn.discordapp.com/attachments/826777163327602698/1443140862987997274/s25voucher.png?ex=692fe629&is=692e94a9&hm=4f4b48563079cc23aabcfd6ea25ba1cae47cb57d1c558fc7d5fe88adbb0a3ecf&'},
        {name: 'belajar huruf bapak', img: 'https://cdn.discordapp.com/attachments/818629322965385236/1177256050634797186/Thanksgiving.png?ex=692fa1e2&is=692e5062&hm=141a36a1e28d2a360c762aa5e9604ff23d04922a36ddad6693c43492c7674b37&'},
    ]
    
    $effect(() => {
        if(sharedStates.currentPage == 'history') {
            fetch(`${PUBLIC_BUKU_KOTAK_API_URL}/history`, {method: 'GET', cache: 'no-store'})
            .then(data => data.json())
            .then(response => {
                console.log(response)
            })
            .catch(err => console.log(err))
        }
    })
</script>

<div class="col-span-10 flex flex-wrap gap-4">
    {#each historyList as list}
    <p class="flex flex-col items-center">
        <img src={list.img} alt={list.name} width="200" height="200">
        <span> {list.name} </span>
    </p>
    {/each}
</div>