import type { WallpaperSetting } from '@/models'

export const defaultSearchList = [
  {
    name: '百度',
    url: 'https://www.baidu.com/s?wd=',
    active: true,
    placeholder: '百度一下'
  },
  {
    name: '必应',
    url: 'https://www.bing.com/search?q=',
    placeholder: '必应一下'
  },
  {
    name: '谷歌',
    url: 'https://www.google.com/search?q=',
    placeholder: '谷歌一下'
  },
  {
    name: '站内',
    url: '/search?q=',
    placeholder: '站内搜索'
  }
]

export const defaultShowSetting = {
  siderBar: true,
  likes: true,
  footer: true,
  collections: true,
  announce: true
}

export const defaultWallpaperSetting = {
  show: true,
  useLocal: false,
  opacity: 0.5,
  link: ''
}
