package resp

type List struct {
	Count int64       `json:"count"`
	List  interface{} `json:"list"`
}
