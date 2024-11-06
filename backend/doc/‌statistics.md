### 1. ""

1. route definition

- Url: /v1/statistics/:id
- Method: POST
- Request: `IDRequest`
- Response: `IDResponse`

2. request definition



```golang
type IDRequest struct {
	ID int `json:"string"`
}
```


3. response definition



```golang
type IDResponse struct {
	ID int `json:"string"`
}
```

