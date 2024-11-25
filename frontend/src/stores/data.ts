import { statistics } from '@/api'
import { getCollections, getLikes } from '@/logic'

// 用于初始化加载数据
export async function loadData() {
  const _themes = ['light', 'dark']
  const { switchTheme } = useTheme(_themes)
  const store = useBasicStore()
  const {
    collectionsDatas,
    collectionsMap,
    navs,
    theme,
    collectionsList,
    likeCollectionsList,
    siteStats
  } = storeToRefs(store)

  switchTheme(theme.value)

  // 加载导航信息

  const guideData = await getCollections()

  collectionsDatas.value.length = 0
  navs.value.length = 0
  collectionsList.value = []
  collectionsMap.value = new Map()

  collectionsDatas.value = [...collectionsDatas.value, ...guideData.datas]
  collectionsDatas.value.forEach((ele) => {
    // 填充导航菜单
    navs.value.push(ele.category)
    // 设置第一组group可见
    ele.groups[0].show = true

    ele.groups.forEach((group) => {
      group.collections.forEach((item) => {
        collectionsMap.value.set(item.cid, item)
        collectionsList.value.push(item)
      })
    })
  })

  // 加载我的收藏
  const likeData = await getLikes()
  likeCollectionsList.value = likeData

  // 加载网站数据
  const siteData = await statistics()

  siteStats.value = siteData.data.data
}
