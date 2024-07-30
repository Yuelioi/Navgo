export class Category {
    name: string = ''
    displayName: string = ''
    icon: string = ''
}

export class Nav {
    id: string = '' // 唯一标识符
    title: string = '' // 导航标题
    officialURL: string[] = [] // 官网地址
    docsURL: string[] = [] // 文档地址
    downloadURL: string[] = [] // 下载地址
    categories: Category[] = [] // 分类
    tags: string[] = [] // 标签
    links: string[] = [] // 相关链接
    icons: string[] = [] // 图标
    description: string = '' // 导航描述
    createdAt: Date = new Date() // 创建时间
    updatedAt: Date = new Date() // 更新时间
}
