// Code generated by goctl. DO NOT EDIT.
package types


type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}
