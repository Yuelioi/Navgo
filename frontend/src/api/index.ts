import axios from 'axios'
import * as components from './types'
export * from './types'

class Header {
  private data: Map<string, string>

  constructor() {
    // 初始化一个空的Map来存储header数据
    this.data = new Map<string, string>()
  }

  withForm() {
    this.data.set('Content-Type', 'multipart/form-data')
    return this
  }

  withAuth() {
    const store = useBasicStore()
    const { token } = storeToRefs(store)
    this.data.set('Authorization', token.value)
    return this
  }

  build(): { [key: string]: any } {
    let obj: { [key: string]: string } = {}
    this.data.forEach((value, key) => {
      obj[key] = value
    })
    return { headers: obj }
  }
}

const formHeader = {
  headers: {
    'Content-Type': 'multipart/form-data'
  }
}

type ApiResponse<T> = {
  code: number
  data: T
  msg: string
}

// const apiUrl = 'http://localhost:9200'
const apiUrl = '/api'

/**
 * @description "获取通知"
 */
export function announces() {
  return axios.get<ApiResponse<components.AnnouncesData>>(apiUrl + `/v1/announces`)
}

/**
 * @description "菜单集合"
 */
export function navs() {
  return axios.get<ApiResponse<components.NavsResponse>>(apiUrl + `/v1/navs`)
}

/**
 * @description "增加页面"
 * @param req
 * @param headers
 */
export function addCollection(req: components.CollectionUpdateParams) {
  return axios.post<ApiResponse<null>>(
    apiUrl + `/v1/collection`,
    req,
    new Header().withForm().withAuth().build()
  )
}

/**
 * @description "删除页面"
 * @param req
 */
export function deleteCollection(req: components.IDRequest) {
  return axios.delete<ApiResponse<components.Collection>>(apiUrl + `/v1/collection`, {
    params: req,
    ...new Header().withAuth().build()
  })
}

/**
 * @description "更新页面"
 * @param req
 * @param headers
 */
export function updateCollection(req: components.CollectionUpdateParams) {
  return axios.put<ApiResponse<components.Collection>>(
    apiUrl + `/v1/collection`,
    req,
    new Header().withForm().withAuth().build()
  )
}

/**
 * @description "单页面"
 */
export function collection(id: string) {
  return axios.get<ApiResponse<components.Collection>>(apiUrl + `/v1/collection/${id}`)
}

/**
 * @description "页面集合"
 */
export function collections() {
  return axios.get<ApiResponse<components.CollectionsDataResponse>>(apiUrl + `/v1/collections`)
}

/**
 * @description "分页集合"
 * @param req
 */
export function filteredCollections(req: components.CollectionFilter) {
  return axios.post<ApiResponse<components.CollectionsResponse>>(
    apiUrl + `/v1/filteredCollections`,
    req
  )
}

/**
 * @description "发布评论"
 * @param req
 */
export function addComment(req: components.CommentRequest) {
  return axios.post<ApiResponse<components.IDResponse>>(apiUrl + `/v1/comment`, req)
}

/**
 * @description "删除评论"
 * @param req
 */
export function deleteComment(req: components.IDRequest) {
  return axios.delete<ApiResponse<components.IDResponse>>(apiUrl + `/v1/comment`)
}

/**
 * @description "获取页面评论"
 * @param req
 */
export function comments() {
  return axios.get<ApiResponse<components.CommentsResponse>>(apiUrl + `/v1/comments`)
}

/**
 * @description "增加页面浏览量"
 * @param req
 */
export function statistics() {
  return axios.get<ApiResponse<components.SiteStats>>(apiUrl + `/v1/statistics`)
}

/**
 * @description "单tag"
 * @param req
 */
export function tag(id: string) {
  return axios.get<ApiResponse<components.CollectionsResponse>>(apiUrl + `/v1/tag/${id}`)
}

/**
 * @description "tags页面集合"
 */
export function tags() {
  return axios.get<ApiResponse<components.TagsResponse>>(apiUrl + `/v1/tags`)
}

/**
 * @description "获取页面信息"
 */
export function net(id: string) {
  return axios.post<components.CollectionResponse>(
    apiUrl + `/v1/net`,
    {},
    {
      params: {
        id: id
      },
      withCredentials: false // 如果不需要发送凭证，可以设置为 false
    }
  )
}

/**
 * @description "验证token信息"
 * @param req
 */
export function checkToken(token: string) {
  return axios.get<ApiResponse<components.AuthResponse>>(
    apiUrl + `/v1/auth`,
    new Header().withAuth().build()
  )
}

/**
 * @description "获取验证信息"
 * @param req
 * @param headers
 */
export function auth(req: components.User) {
  return axios.post<ApiResponse<components.AuthResponse>>(
    apiUrl + `/v1/auth`,
    req,
    new Header().withForm().withAuth().build()
  )
}

/**
 * @description
 */
export function wallpaper() {
  return axios.get<ApiResponse<components.IDResponse>>(apiUrl + `/v1/wallpaper`)
}
