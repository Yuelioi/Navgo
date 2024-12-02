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

  // 切换主题
  switchTheme(theme.value)

  // 加载我的收藏
  const likeData = await getLikes()
  likeCollectionsList.value = likeData
  likeCollectionsList.value.forEach((el) => {
    el.like = true
  })

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
    // 设置第一组group可见, 防止首屏加载太多

    for (let i = 0; i < ele.groups.length; i++) {
      const group = ele.groups[i]
      group.show = true
      break
    }

    ele.groups.forEach((group) => {
      group.collections.forEach((item) => {
        // 添加like
        if (
          likeCollectionsList.value.some((ele) => {
            return ele.cid === item.cid
          })
        ) {
          item.like = true
        }
        collectionsMap.value.set(item.cid, item)
        collectionsList.value.push(item)
      })
    })
  })


  try {
    // 加载网站数据
    const siteData = await statistics()
    if (siteData) {
      siteStats.value = siteData.data.data
    }
  } catch (e) {
    console.log(e)
  }
}
