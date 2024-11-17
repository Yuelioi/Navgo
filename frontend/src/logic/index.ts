import * as api from '@/api'
import { db } from '@/db/db'

// 高级查询 如果数据不存在，执行添加操作, 基于日期查询
async function query(tabName: string, apiFunc: any) {
  const now = new Date()

  const year = now.getFullYear()
  const month = now.getMonth() + 1
  const day = now.getDate()
  const today = `${year}${month}${day}`

  const collections = await db.queryAllData(tabName)

  var collection: any = null

  if (collections.length !== 0) {
    collections.forEach(async (element: any) => {
      if (element.id !== today) {
        await db.deleteData(tabName, element.id)
      } else {
        collection = element
      }
    })
  }

  // 只获取消息体 (data层)
  if (collection === null) {
    const resp = await apiFunc()
    const data = resp.data.data

    await db.addData(tabName, {
      id: today,
      data: data
    })
    return data
  } else {
    return collection['data']
  }
}

export async function getCollections() {
  // 先清除历史记录 再添加

  // await db.clearData('collections')
  return await query('collections', api.collections)
}

export async function getLikes() {
  const collections = await db.queryAllData('likes')
  if (collections) {
    return collections
  }
  return []
}

export async function getAnnounces() {
  // 先清除历史记录 再添加
  // await db.clearData('announces')
  return await query('announces', api.announces)
}

export async function net(id: string) {
  const resp = await api.net(id)
  return resp.data['data']
}
