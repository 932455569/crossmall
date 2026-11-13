package context

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/932455569/crossmall/config"
)

// Context 请求上下文，封装底层的 HTTP GET、签名生成与 XML 解析
type Context struct {
	*config.Config
}

// DoGet 发送 HTTP GET 请求，自动生成签名并解析 XML 响应
// path: 接口路径 (如 "/webapi2/get_stock")
// params: 请求参数 map
// res: 用于接收解析结果的指针
func (c *Context) DoGet(ctx context.Context, path string, params map[string]string, res interface{}) error {
	// 1. 处理 BaseURL
	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = "https://crossmall.jp"
	}
	u, err := url.Parse(baseURL + path)
	if err != nil {
		return fmt.Errorf("invalid base url: %v", err)
	}

	// 2. 构建 Query 参数
	q := u.Query()
	// 强制加入 account (会社コード)
	q.Set("account", c.Account)
	for k, v := range params {
		// 忽略空字符串参数
		if v != "" {
			q.Set(k, v)
		}
	}

	// 3. 生成签名 (signing)
	// 文档 2. 署名：リクエスト引数の末尾に認証鍵を結合した文字列をｍｄ５によりハッシュ化
	// 注意：转换为MD5前，参数必须经过 URL Encode
	// Go 的 q.Encode() 会自动进行 URL Encode 并按 Key 升序排序
	q.Del("signing")
	signStr := q.Encode() + c.AuthKey
	hash := md5.Sum([]byte(signStr))
	signature := hex.EncodeToString(hash[:])

	// 将签名加入参数
	q.Set("signing", signature)
	u.RawQuery = q.Encode()

	// 4. 发起 HTTP GET 请求
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	fmt.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %v", err)
	}
	defer resp.Body.Close()

	// 5. 读取响应内容
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return fmt.Errorf("read response body failed: %v", err)
	}

	// 6. 解析 XML
	if err := xml.Unmarshal(body, res); err != nil {
		return fmt.Errorf("xml unmarshal error: %v, raw response: %s", err, string(body))
	}

	return nil
}

// BuildArrayParam 处理类似于 creating_source_type[] 的数组参数
// 根据文档 10. 注文伝票情報取得 (get_order) 的备注，多值参数格式为 xxx[]=a&xxx[]=b
func BuildArrayParam(key string, values []string) (string, string) {
	if len(values) == 0 {
		return "", ""
	}
	// 返回 key[] 和 拼接好的字符串，由底层 DoGet 处理
	// 注意：由于 url.Values 原生不支持直接传同 key 的数组，这里做特殊的处理
	// 实际上我们可以传给 DoGet 一个特殊格式，为了简化，我们在具体业务模块中用分号拼接，然后在 DoGet 里展开
	// 最简单的方式是在业务层直接把参数设成 map[string]string{"creating_source_type[]": "3,4"}
	// 根据测试，CROSS MALL 实际支持 `creating_source_type[]=3&creating_source_type[]=4` 或直接传逗号。
	// 我们这里推荐在请求层特殊处理，不污染 DoGet。
	return key, strings.Join(values, ",")
}