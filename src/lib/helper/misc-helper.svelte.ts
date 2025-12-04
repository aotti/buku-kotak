import { PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public"
import html2canvas from "html2canvas"

interface ISharedStates {
    currentPage: string,
    historyList: string[],
    historyFiltered: string[],
}
export const sharedStates = $state<ISharedStates>({
    currentPage: 'paper',
    historyList: [],
    historyFiltered: [],
})

export function qS<T = HTMLElement & HTMLInputElement>(selector: string) {
    return document.querySelector(selector) as T
}

export function categoryHandler(id: string) {
    sharedStates.currentPage = id
}

export function screenshotThenUpload() {
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

export function filterHistory(ev: Event & {currentTarget: HTMLInputElement}) {
    const inputValue = ev.currentTarget.value
    if(inputValue.length > 0) 
        sharedStates.historyFiltered = sharedStates.historyList.filter(v => v.match(inputValue))
    else if(inputValue.length === 0)
        sharedStates.historyFiltered = []
}