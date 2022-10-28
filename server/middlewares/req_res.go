package middlewares

// ReqRes is a middleware that logs the request and response.
type ReqRes struct {
	URI          string                 `json:"URI"`
	Method       string                 `json:"Method"`
	RequestBody  map[string]interface{} `json:"RequestBody,omitempty"`
	ResponseCode int                    `json:"ResponseCode"`
	ResponseBody map[string]interface{} `json:"ResponseBody,omitempty"`
}
