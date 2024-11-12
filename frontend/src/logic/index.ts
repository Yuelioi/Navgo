import * as api from '@/api'
import { db } from '@/db/db'

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
  return await query('collections', api.collections)
}
