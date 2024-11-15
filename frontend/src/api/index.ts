import axios from 'axios'
import * as components from './types'
export * from './types'

const apiUrl = '/api'

/**
 * @description "增加页面"
 * @param req
 * @param headers
 */
export function addCollection(req: components.Collection) {
  return axios.post<components.Collection>(apiUrl + `/v1/collection`, req)
}

/**
 * @description "删除页面"
 * @param req
 */
export function deleteCollection(req: components.IDRequest) {
  return axios.delete<components.Collection>(apiUrl + `/v1/collection`)
}

/**
 * @description "更新页面"
 * @param req
 * @param headers
 */
export function updateCollection(req: components.Collection) {
  return axios.put<components.Collection>(apiUrl + `/v1/collection`, req)
}

/**
 * @description "单页面"
 */
export function collection(id: string) {
  return axios.get<components.Collection>(apiUrl + `/v1/collection/${id}`)
}

/**
 * @description "页面集合"
 */
export function collections() {
  return axios.get<{ data: components.CollectionsResponse }>(apiUrl + `/v1/collections`)
}

/**
 * @description "发布评论"
 * @param req
 */
export function addComment(req: components.CommentRequest) {
  return axios.post<{
    data: components.IDResponse
  }>(apiUrl + `/v1/comment`, req)
}

/**
 * @description "删除评论"
 * @param req
 */
export function deleteComment(req: components.IDRequest) {
  return axios.delete<components.IDResponse>(apiUrl + `/v1/comment`)
}

/**
 * @description "获取页面评论"
 * @param req
 */
export function comment(req: components.IDRequest) {
  return axios.get<components.CommentsResponse>(apiUrl + `/v1/comments`)
}

/**
 * @description "增加页面浏览量"
 * @param req
 */
export function statistics(req: components.IDRequest) {
  return axios.post<components.IDResponse>(apiUrl + `/v1/statistics/${req.id}`)
}

/**
 * @description "单tag"
 * @param req
 */
export function tag(id: string) {
  return axios.get<components.CollectionsResponse>(apiUrl + `/v1/tag/${id}`)
}

/**
 * @description "tags页面集合"
 */
export function tags() {
  return axios.get<components.TagsResponse>(apiUrl + `/v1/tags`)
}

/**
 * @description "获取页面信息"
 */
export function net(id: string) {
  return axios.post<components.CollectionResponse>(
    apiUrl + `/v1/net`,
    {
      id: id
    },
    {
      withCredentials: false // 如果不需要发送凭证，可以设置为 false
    }
  )
}
