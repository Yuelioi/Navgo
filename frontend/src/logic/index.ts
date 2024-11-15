import * as api from '@/api'
import { db } from '@/db/db'

// 高级查询 如果数据不存在，执行添加操作
async function query(tabName: string, apiFunc: any) {
  const now = new Date()

  const year = now.getFullYear()
  const month = now.getMonth() + 1
  const day = now.getDate()
  const today = `${year}${month}${day}`

  const collections = await db.queryData(tabName, today)

  if (collections == null) {
    const resp = await apiFunc()
    const data = resp.data.data

    await db.addData(tabName, {
      id: today,
      data: data
    })
    return data
  } else {
    return collections['data']
  }
}

export async function getCollections() {
  // 先清除历史记录 再添加
  await db.clearData('collections')
  return await query('collections', api.collections)
}

export async function net(id: string) {
  const resp = await api.net(id)
  return resp.data['data']
}
