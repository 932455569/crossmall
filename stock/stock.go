package stock

import (
	"context"
	"encoding/xml"
	"fmt"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct {
	*cmctx.Context
}

func NewClient(ctx *cmctx.Context) *Client {
	return &Client{Context: ctx}
}

// ================= 【4. 在庫情報取得】 =================

// GetStockRequest 在庫情報取得 请求参数
type GetStockRequest struct {
	SkuCode        string // SKUコード (△: item_code / jan_code 未指定时必填)
	ItemCode       string // 商品コード (△: sku_code / jan_code 未指定时必填)
	JanCode        string // JANコード (△: sku_code / item_code 未指定时必填)
	Attribute1Code string // 属性1コード (商品代码指定时有效)
	Attribute2Code string // 属性2コード
}

// GetStockResponse 在庫情報取得 响应
type GetStockResponse struct {
	XMLName   xml.Name `xml:"GetStock"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string `xml:"TotalResult,attr"` // 取得レコード数
		ResultStatus struct {
			Status  string `xml:"GetStatus"` // success 或 error
			Message string `xml:"Message"`   // 失败时的错误内容
		} `xml:"ResultStatus"`
		Results []StockItem `xml:"Result"`
	} `xml:"ResultSet"`
}

// StockItem 库存明细数据
type StockItem struct {
	No             string `xml:"No,attr"` // 序号 1から連番
	ItemCd         string `xml:"item_cd"`
	Attribute1Code string `xml:"attribute1_code"`
	Attribute1Name string `xml:"attribute1_name"`
	Attribute2Code string `xml:"attribute2_code"`
	Attribute2Name string `xml:"attribute2_name"`
	Stock          int    `xml:"stock"` // 在庫数量
}

// GetStock 获取库存信息
func (c *Client) GetStock(ctx context.Context, req *GetStockRequest) (*GetStockResponse, error) {
	params := map[string]string{
		"sku_code":        req.SkuCode,
		"item_code":       req.ItemCode,
		"jan_code":        req.JanCode,
		"attribute1_code": req.Attribute1Code,
		"attribute2_code": req.Attribute2Code,
	}
	var res GetStockResponse
	err := c.DoGet(ctx, "/webapi2/get_stock", params, &res)
	if err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【5. 在庫情報更新】 =================

// UpdStockRequest 在庫情報更新 请求参数
type UpdStockRequest struct {
	SkuCode                 string // SKUコード (△)
	ItemCode                string // 商品コード (△)
	JanCode                 string // JANコード (△)
	Attribute1Code          string // 属性1コード (△: item_code 指定时必填)
	Attribute2Code          string // 属性2コード (△)
	Stock                   int    // 在庫数量 (〇)
	StockCalculation        string // 在庫数量計算 0:OFF 1:ON
	StockCalculationReserve string // 在庫数量計算(仮注文) 0:OFF 1:ON
	StockCalculationErr     string // 在庫数量計算(注文エラー) 0:OFF 1:ON
	MarkUpdate              string // 商品マーク更新 空值:不更新, 1~5:MarkX ON
	SendStockMode           string // 在庫送信モード 0:標準 1:強制
	UpdateMode              string // 更新モード (〇) u:更新(上書) a:更新(加算) d:削除
}

// UpdStockResponse 在庫情報更新 响应
type UpdStockResponse struct {
	XMLName   xml.Name `xml:"UpdStock"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"` // success 或 error
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

// UpdStock 更新库存信息
func (c *Client) UpdStock(ctx context.Context, req *UpdStockRequest) (*UpdStockResponse, error) {
	params := map[string]string{
		"sku_code":                  req.SkuCode,
		"item_code":                 req.ItemCode,
		"jan_code":                  req.JanCode,
		"attribute1_code":           req.Attribute1Code,
		"attribute2_code":           req.Attribute2Code,
		"stock":                     fmt.Sprintf("%d", req.Stock),
		"stock_calculation":         req.StockCalculation,
		"stock_calculation_reserve": req.StockCalculationReserve,
		"stock_calculation_err":     req.StockCalculationErr,
		"mark_update":               req.MarkUpdate,
		"send_stock_mode":           req.SendStockMode,
		"update_mode":               req.UpdateMode,
	}
	var res UpdStockResponse
	err := c.DoGet(ctx, "/webapi2/upd_stock", params, &res)
	if err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【20. 在庫情報差分取得】 =================

// GetDiffStockRequest 在庫情報差分取得 请求参数
type GetDiffStockRequest struct {
	UpdatedAtFr string // 更新日時(Fr) (〇) yyyy-mm-dd hh:mm:ss
	UpdatedAtTo string // 更新日時(To) (〇) yyyy-mm-dd hh:mm:ss
	SkuCode     string // SKUコード (△)
	JanCode     string // JANコード (△)
	ItemCode    string // 商品コード (△)
	Condition   string // 抽出方法 0:=, 1:>, 2:>= (省略时 1:>)
}

// GetDiffStockResponse 在庫情報差分取得 响应 (与 GetStockResponse 完全一致)
type GetDiffStockResponse struct {
	GetStockResponse
}

// GetDiffStock 获取库存差分信息
func (c *Client) GetDiffStock(ctx context.Context, req *GetDiffStockRequest) (*GetDiffStockResponse, error) {
	params := map[string]string{
		"updated_at_fr": req.UpdatedAtFr,
		"updated_at_to": req.UpdatedAtTo,
		"sku_code":      req.SkuCode,
		"jan_code":      req.JanCode,
		"item_code":     req.ItemCode,
		"condition":     req.Condition,
	}
	var res GetDiffStockResponse
	err := c.DoGet(ctx, "/webapi2/get_diff_stock", params, &res)
	if err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}