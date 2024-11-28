import axios, { type AxiosRequestConfig } from 'axios'
import * as components from './types'
export * from './types'

const formHeader = {
  headers: {
    'Content-Type': 'multipart/form-data'
  }
}

// const apiUrl = 'http://localhost:9200'
const apiUrl = '/api'

/**
 * @description "获取通知"
 */
export function announces() {
  return axios.get<{ data: components.AnnouncesData }>(apiUrl + `/v1/announces`)
}

/**
 * @description "增加页面"
 * @param req
 * @param headers
 */
export function addCollection(req: FormData) {
  return axios.post<{ code: number; msg: string }>(apiUrl + `/v1/collection`, req, formHeader)
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
    code: number
    data: components.IDResponse
    msg: string
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
export function comments() {
  return axios.get<{ data: components.CommentsResponse; code: number; msg: string }>(
    apiUrl + `/v1/comments`
  )
}

/**
 * @description "增加页面浏览量"
 * @param req
 */
export function statistics() {
  return axios.get<{
    data: components.SiteStats
  }>(apiUrl + `/v1/statistics`)
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

/**
 * @description "验证token信息"
 * @param req
 */
export function checkToken(req: components.IDRequest) {
  return axios.get<{ code: number; data: components.AuthResponse }>(apiUrl + `/v1/auth`, req)
}

/**
 * @description "获取验证信息"
 * @param req
 * @param headers
 */
export function auth(req: components.User) {
  return axios.post<{ data: components.AuthResponse }>(apiUrl + `/v1/auth`, req, formHeader)
}

/**
 * @description
 */
export function wallpaper() {
  return axios.get<{ data: components.IDResponse }>(apiUrl + `/v1/wallpaper`)
}
