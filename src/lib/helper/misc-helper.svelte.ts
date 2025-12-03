export const sharedStates = $state({
    currentPage: 'paper'
})

export function qS<T = HTMLElement & HTMLInputElement>(selector: string) {
    return document.querySelector(selector) as T
}