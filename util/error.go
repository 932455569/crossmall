package util

import "fmt"

// APIError CROSS MALL API 业务错误
// 对应 XML 中 <Message> 返回的内容
type APIError struct {
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("crossmall api error: %s", e.Message)
}

// Status 获取系（Get）通用状态
// 对应 XML 中的 <ResultStatus><GetStatus>success/error</GetStatus><Message>...</Message></ResultStatus>
type Status struct {
	Status  string `xml:"GetStatus"` // success 或 error
	Message string `xml:"Message"`   // 失败时的错误内容
}