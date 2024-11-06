### 1. "单tag"

1. route definition

- Url: /v1/tag/:id
- Method: GET
- Request: `IDRequest`
- Response: `CollectionsResponse`

2. request definition



```golang
type IDRequest struct {
	ID int `json:"string"`
}
```


3. response definition



```golang
type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}
```

### 2. "tags页面集合"

1. route definition

- Url: /v1/tags
- Method: GET
- Request: `AnyRequest`
- Response: `TagsResponse`

2. request definition



```golang
type AnyRequest struct {
}
```


3. response definition



```golang
type TagsResponse struct {
	Tags []string `json:"tags"`
}
```

