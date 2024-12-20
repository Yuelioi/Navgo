export function open(url: string) {
  window.open(url, '_blank')
}

export function imageLoadError(event: Event) {
  const target = event.target as HTMLImageElement
  if (target) {
    target.style.display = 'none'
  }
}

export function isMobileDevice() {
  // return /iPhone|iPad|iPod|Android/i.test(navigator.userAgent)
  return screen.width < 768
}

export function getToken() {}
