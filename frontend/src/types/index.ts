class Model {
    // 这里可以添加一些通用的模型方法或属性
}

class Category extends Model {
    cid?: string
    depth: number = 0
    title: string = ''
    order?: number
    path: string = ''
}

class Collection extends Model {
    cid?: string = ''
    title: string = ''
    link: string = ''
    path?: string = ''
    categoryID: string = ''
    category: Category | null = null
    description?: string = ''
    country?: string = ''
    thumbnail?: string = ''
    tags: string[] = []
    view?: number = 0
}

class Group {
    category: Category
    collections: Collection[]

    constructor(category: Category, collections: Collection[]) {
        this.category = category
        this.collections = collections
    }
}

class CollectionsData {
    category: Category = new Category()
    groups: Group[] = []
}

class CollectionsResponse {
    datas: CollectionsData[] = []
}
