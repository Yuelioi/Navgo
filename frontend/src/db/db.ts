class IndexDB {
  dbName: string
  version = 2

  constructor(name: string) {
    this.dbName = name
  }

  // 初始化数据库, 需要填写要添加的表
  async init(storeNames: string[]): Promise<IDBDatabase> {
    let that = this

    return new Promise((resolve, reject) => {
      const request = indexedDB.open(this.dbName, that.version)
      request.onupgradeneeded = (event: IDBVersionChangeEvent) => {
        const db = (event.target as IDBOpenDBRequest).result as IDBDatabase
        storeNames.forEach((storeName) => {
          if (!db.objectStoreNames.contains(storeName)) {
            db.createObjectStore(storeName, { keyPath: 'id' })
          }
        })
      }
      request.onsuccess = (event) => {
        const target = event.target as IDBRequest
        resolve(target.result)
      }
      request.onerror = (event) => {
        const target = event.target as IDBRequest
        reject(target.error)
      }
    })
  }

  async launch(): Promise<IDBDatabase> {
    let that = this

    return new Promise((resolve, reject) => {
      const request = indexedDB.open(this.dbName, that.version)

      request.onsuccess = (event) => {
        const target = event.target as IDBRequest
        resolve(target.result)
      }

      request.onerror = (event) => {
        const target = event.target as IDBRequest
        reject(target.error)
      }
    })
  }

  // 添加, 有则更新
  async addData(tabName: string, data: any): Promise<void> {
    const db = await this.launch()
    return new Promise((resolve, reject) => {
      const transaction = db.transaction([tabName], 'readwrite')
      const store = transaction.objectStore(tabName)

      const get_request = store.get(data.id)

      get_request.onsuccess = (event) => {
        const existingData = (event.target as IDBRequest).result
        if (existingData) {
          // 如果数据存在，执行更新操作
          const update_request = store.put(data)
          update_request.onsuccess = () => {
            resolve()
          }
          update_request.onerror = () => {
            reject(update_request.error)
          }
        } else {
          // 如果数据不存在，执行添加操作
          const add_request = store.add(data)
          add_request.onsuccess = () => {
            resolve()
          }
          add_request.onerror = () => {
            reject(add_request.error)
          }
        }
      }

      get_request.onerror = () => {
        reject(get_request.error)
      }
    })
  }

  async deleteData(tabName: string, id: string): Promise<void> {
    const db = await this.launch()

    return new Promise((resolve, reject) => {
      const transaction = db.transaction([tabName], 'readwrite')
      const store = transaction.objectStore(tabName)
      const request = store.delete(id)

      request.onsuccess = () => {
        resolve()
      }

      request.onerror = () => {
        reject(request.error)
      }
    })
  }

  async clearData(tabName: string): Promise<void> {
    const db = await this.launch()

    return new Promise((resolve, reject) => {
      const transaction = db.transaction([tabName], 'readwrite')
      const store = transaction.objectStore(tabName)
      const request = store.clear()

      request.onsuccess = () => {
        resolve()
      }

      request.onerror = () => {
        reject(request.error)
      }
    })
  }

  async queryData(tabName: string, id: string): Promise<any> {
    const db = await this.launch()

    return new Promise((resolve, reject) => {
      const transaction = db.transaction([tabName], 'readwrite')
      const store = transaction.objectStore(tabName)
      const request = store.get(id)

      request.onsuccess = () => {
        resolve(request.result)
      }

      request.onerror = () => {
        reject(request.error)
      }
    })
  }
  async queryAllData(tabName: string): Promise<any> {
    const db = await this.launch()

    return new Promise((resolve, reject) => {
      const transaction = db.transaction([tabName], 'readwrite')
      const store = transaction.objectStore(tabName)
      const request = store.getAll()

      request.onsuccess = () => {
        resolve(request.result)
      }

      request.onerror = () => {
        reject(request.error)
      }
    })
  }
}

const db = new IndexDB('core')

db.init(['collections', 'likes', 'wallpaper'])

export { db }
