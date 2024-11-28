import axios from "./gocliRequest"
import * as components from "./mainComponents"
export * from "./mainComponents"

/**
 * @description "获取通知"
 */
export function announces() {
	return axios.get<components.AnnouncesData>(`/v1/announces`)
}

/**
 * @description "获取验证信息"
 * @param req
 * @param headers
 */
export function auth(req: components.User, headers: components.UserHeaders) {
	return axios.post<components.AuthResponse>(`/v1/auth`, req, headers)
}

/**
 * @description "验证token信息"
 * @param req
 */
export function checkToken(req: components.IDRequest) {
	return axios.get<components.AuthResponse>(`/v1/auth`, req)
}

/**
 * @description "增加页面"
 * @param req
 * @param headers
 */
export function addCollection(req: components.Collection, headers: components.CollectionHeaders) {
	return axios.post<components.Collection>(`/v1/collection`, req, headers)
}

/**
 * @description "删除页面"
 * @param req
 */
export function deleteCollection(req: components.IDRequest) {
	return axios.delete<components.Collection>(`/v1/collection`, req)
}

/**
 * @description "更新页面"
 * @param req
 * @param headers
 */
export function updateCollection(req: components.Collection, headers: components.CollectionHeaders) {
	return axios.put<components.Collection>(`/v1/collection`, req, headers)
}

/**
 * @description "单页面"
 */
export function collection() {
	return axios.get<components.Collection>(`/v1/collection/${id}`)
}

/**
 * @description "页面集合"
 */
export function collections() {
	return axios.get<components.CollectionsResponse>(`/v1/collections`)
}

/**
 * @description "发布评论"
 * @param req
 * @param headers
 */
export function addComment(req: components.Comment, headers: components.CommentHeaders) {
	return axios.post<components.IDResponse>(`/v1/comment`, req, headers)
}

/**
 * @description "删除评论"
 * @param req
 */
export function deleteComment(req: components.IDRequest) {
	return axios.delete<components.IDResponse>(`/v1/comment`, req)
}

/**
 * @description "获取页面评论"
 */
export function comment() {
	return axios.get<components.CommentsResponse>(`/v1/comments`)
}

/**
 * @description "获取网站信息"
 * @param req
 */
export function net(req: components.IDRequest) {
	return axios.post<components.Collection>(`/v1/net`, req)
}

/**
 * @description 
 */
export function wallpaper() {
	return axios.get<components.IDResponse>(`/v1/wallpaper`)
}

/**
 * @description "获取网站数量"
 */
export function statistics() {
	return axios.get<components.SiteStats>(`/v1/statistics`)
}

/**
 * @description "单tag"
 * @param req
 */
export function tag(req: components.IDRequest) {
	return axios.get<components.CollectionsResponse>(`/v1/tag/${id}`, req)
}

/**
 * @description "tags页面集合"
 */
export function tags() {
	return axios.get<components.TagsResponse>(`/v1/tags`)
}

/**
 * @description "tags页面集合"
 */
export function users() {
	return axios.get<components.UserResponse>(`/v1/users`)
}
