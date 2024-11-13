import { getCollections } from '@/logic'

// 用于初始化加载数据
export async function loadData() {
  const _themes = ['light', 'dark']
  const { switchTheme } = useTheme(_themes)
  const store = useBasicStore()
  let { collectionsDatas, collections, navs, theme } = storeToRefs(store)

  switchTheme(theme.value)

  // 加载导航信息
  const data = await getCollections()

  collectionsDatas.value.length = 0
  navs.value.length = 0
  collections.value = new Map()

  collectionsDatas.value = [...collectionsDatas.value, ...data.datas]
  collectionsDatas.value.forEach((ele) => {
    // 填充导航菜单
    navs.value.push(ele.category)
    // 设置第一组group可见
    ele.groups[0].show = true

    ele.groups.forEach((group) => {
      group.collections.forEach((item) => {
        collections.value.set(item.cid, item)
      })
    })
  })
}
