### 1. "获取页面评论"

1. route definition

- Url: /v1/comments
- Method: GET
- Request: `IDRequest`
- Response: `CommentsResponse`

2. request definition



```golang
type IDRequest struct {
	ID int `json:"string"`
}
```


3. response definition



```golang
type CommentsResponse struct {
	Comments []string `json:"comments"`
}
```

