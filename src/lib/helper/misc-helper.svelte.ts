import { PUBLIC_BUKU_KOTAK_API_URL } from "$env/static/public"
import html2canvas from "html2canvas"

interface ISharedStates {
    currentPage: string,
    historyList: string[],
    historyFiltered: string[],
    canvasCopy: HTMLCanvasElement,
}
export const sharedStates = $state<ISharedStates>({
    currentPage: 'paper',
    historyList: [],
    historyFiltered: [],
    canvasCopy: null,
})

export function qS<T = HTMLElement & HTMLInputElement>(selector: string) {
    return document.querySelector(selector) as T
}

export function qSA<T = NodeListOf<HTMLElement & HTMLInputElement>>(selector: string) {
    return document.querySelectorAll(selector) as T
}

export function categoryHandler(id: string) {
    sharedStates.currentPage = id
}

export function canvasCopy(element_id: string) {
    const canvasCopyElement = qS(element_id) as HTMLCanvasElement
    canvasCopyElement.classList.toggle('bg-green-400/50')
    setTimeout(() => canvasCopyElement.classList.toggle('bg-green-400/50'), 5000);
    // set canvas copy
    sharedStates.canvasCopy = canvasCopyElement
}

export function canvasPaste(element_id: string) {
    const canvasPasteElement = qS(element_id) as HTMLCanvasElement
    canvasPasteElement.classList.toggle('bg-green-400/50')
    setTimeout(() => canvasPasteElement.classList.toggle('bg-green-400/50'), 1000);
    // paste canvas
    const canvasPasteCtx = canvasPasteElement.getContext('2d')
    canvasPasteCtx.drawImage(sharedStates.canvasCopy, 0, 0)
    // reset canvas copy
    sharedStates.canvasCopy = null
}

export function screenshotThenUpload() {
    // hide action buttons before screenshot
    const actionButtons = qSA('.action-buttons')
    actionButtons.forEach(el => el.classList.add('hidden'))
    
    // set image title
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
        
        // set upload notif
        const imageUploadNotifContainer = qS('#upload_notif_container')
        const imageUploadNotif = qS('#upload_notif')
        imageUploadNotifContainer.classList.remove('hidden')
        imageUploadNotif.textContent = 'uploading image..'
        
        // uploading image
        const fetchOptions: RequestInit = {
            method: 'POST',
            headers: {'content-type': 'application/json'},
            body: JSON.stringify(imageData)
        }
        const imageUpload = await fetch(`${PUBLIC_BUKU_KOTAK_API_URL}/upload`, fetchOptions)
        // show action buttons
        actionButtons.forEach(el => el.classList.remove('hidden'))
        // upload image response
        if(imageUpload.status < 400) {
            screenshotThenUploadNotif('image upload success!')
            .then(() => alert('image upload success!'))
        }
        else {
            screenshotThenUploadNotif('fail to upload image..')
            .then(async () => alert(await imageUpload.json()))
        }
    })
}

function screenshotThenUploadNotif(message: string) {
    return new Promise(resolve => {
        const imageUploadNotifContainer = qS('#upload_notif_container')
        const imageUploadNotif = qS('#upload_notif')
        // display and set notif text
        imageUploadNotif.textContent = message
        setTimeout(() => resolve(imageUploadNotifContainer.classList.add('hidden')), 3000)
    })
}

export function filterHistory(ev: Event & {currentTarget: HTMLInputElement}) {
    const inputValue = ev.currentTarget.value
    if(inputValue.length > 0) 
        sharedStates.historyFiltered = sharedStates.historyList.filter(v => v.match(inputValue))
    else if(inputValue.length === 0)
        sharedStates.historyFiltered = []
}

function preventDefault(e) {e.preventDefault()}

export function disableScrolling(){
    let x = window.scrollX;
    let y = window.scrollY;
    window.onscroll=() => {window.scrollTo(x, y)};
    window.addEventListener('touchmove', preventDefault, {passive: false})
}

export function enableScrolling(){
    window.onscroll=() => {};
    window.removeEventListener('touchmove', preventDefault)
}