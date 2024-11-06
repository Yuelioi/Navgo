### 1. "增加页面"

1. route definition

- Url: /v1/collection
- Method: POST
- Request: `CollectionRequest`
- Response: `Collection`

2. request definition



```golang
type CollectionRequest struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
}
```


3. response definition



```golang
type Collection struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
	View int `json:"view"`
}
```

### 2. "删除页面"

1. route definition

- Url: /v1/collection
- Method: DELETE
- Request: `IDRequest`
- Response: `Collection`

2. request definition



```golang
type IDRequest struct {
	ID int `json:"string"`
}
```


3. response definition



```golang
type Collection struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
	View int `json:"view"`
}
```

### 3. "更新页面"

1. route definition

- Url: /v1/collection
- Method: PUT
- Request: `CollectionRequest`
- Response: `Collection`

2. request definition



```golang
type CollectionRequest struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
}
```


3. response definition



```golang
type Collection struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
	View int `json:"view"`
}
```

### 4. "单页面"

1. route definition

- Url: /v1/collection/:id
- Method: GET
- Request: `IDRequest`
- Response: `Collection`

2. request definition



```golang
type IDRequest struct {
	ID int `json:"string"`
}
```


3. response definition



```golang
type Collection struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Category string `json:"category"`
	Link string `json:"link"`
	Description string `json:"description"`
	Country string `json:"country"`
	Thumbnail string `json:"thumbnail"`
	Tags []string `json:"tags"`
	View int `json:"view"`
}
```

### 5. "页面集合"

1. route definition

- Url: /v1/collections
- Method: GET
- Request: `AnyRequest`
- Response: `CollectionsResponse`

2. request definition



```golang
type AnyRequest struct {
}
```


3. response definition



```golang
type CollectionsResponse struct {
	Collections []Collection `json:"collections"`
}
```

