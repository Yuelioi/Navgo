export interface WallpaperSetting {
  show: boolean
  useLocal: boolean
  opacity: number
  link: string
}

export interface Search {
  name: string
  url: string
  active?: boolean
  placeholder: string
}

export interface ShowSetting {
  siderBar: boolean
  likes: boolean
  footer: boolean
  collections: boolean
  announce: boolean
}
