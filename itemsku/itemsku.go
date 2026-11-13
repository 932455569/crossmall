package itemsku

import (
	"context"
	"encoding/xml"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【17. 商品SKU情報取得】 =================

type GetItemSkuRequest struct {
	ItemCode       string // 商品コード (〇)
	Attribute1Code string // 属性1コード
	Attribute2Code string // 属性2コード
}

type GetItemSkuResponse struct {
	XMLName   xml.Name `xml:"GetItemSku"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string      `xml:"TotalResult,attr"`
		ResultStatus util.Status `xml:"ResultStatus"`
		Results      []ItemSkuData `xml:"Result"`
	} `xml:"ResultSet"`
}

type ItemSkuData struct {
	No             string `xml:"No,attr"`
	ItemCode       string `xml:"item_code"`
	Attribute1Code string `xml:"attribute1_code"`
	Attribute1Name string `xml:"attribute1_name"`
	Attribute2Code string `xml:"attribute2_code"`
	Attribute2Name string `xml:"attribute2_name"`
	ItemSkuCode    string `xml:"item_sku_code"`    // 商品SKUコード
	OrderType      string `xml:"order_type"`       // 発注管理区分 0:しない 1:通常 2:補充
	OrderPoint     string `xml:"order_point"`      // 発注点
	SafetyStock    string `xml:"safety_stock"`     // 安全在庫数量
	LeadTime       string `xml:"lead_time"`        // 調達期間
	OrderLot       string `xml:"order_lot"`        // 発注ロット
	KeepStock      string `xml:"keep_stock"`       // 確保在庫数量
	StockUnlimitFlag string `xml:"stock_unlimit_flag"` // 在庫無制限 0:OFF 1:ON
	JanCode        string `xml:"jan_code"`         // JANコード
}

func (c *Client) GetItemSku(ctx context.Context, req *GetItemSkuRequest) (*GetItemSkuResponse, error) {
	params := map[string]string{
		"item_code":       req.ItemCode,
		"attribute1_code": req.Attribute1Code,
		"attribute2_code": req.Attribute2Code,
	}
	var res GetItemSkuResponse
	if err := c.DoGet(ctx, "/webapi2/get_item_sku", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【18. 商品SKU情報更新】 =================

type UpdItemSkuRequest struct {
	ItemCode         string // 商品コード (〇)
	Attribute1Code   string // 属性1コード (〇)
	Attribute2Code   string // 属性2コード (〇)
	OrderType        string // 発注管理区分 0:しない 1:通常 2:補充
	OrderPoint       string // 発注点 -999~
	SafetyStock      string // 安全在庫数量 0~
	LeadTime         string // 調達期間 0~
	OrderLot         string // 発注ロット 0~
	KeepStock        string // 確保在庫数量 0~
	StockUnlimitFlag string // 在庫無制限 0:OFF 1:ON
	JanCode          string // JANコード
}

type UpdItemSkuResponse struct {
	XMLName   xml.Name `xml:"UpdItemSku"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

func (c *Client) UpdItemSku(ctx context.Context, req *UpdItemSkuRequest) (*UpdItemSkuResponse, error) {
	params := map[string]string{
		"item_code":          req.ItemCode,
		"attribute1_code":    req.Attribute1Code,
		"attribute2_code":    req.Attribute2Code,
		"order_type":         req.OrderType,
		"order_point":        req.OrderPoint,
		"safety_stock":       req.SafetyStock,
		"lead_time":          req.LeadTime,
		"order_lot":          req.OrderLot,
		"keep_stock":         req.KeepStock,
		"stock_unlimit_flag": req.StockUnlimitFlag,
		"jan_code":           req.JanCode,
	}
	var res UpdItemSkuResponse
	if err := c.DoGet(ctx, "/webapi2/upd_item_sku", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【21. 商品SKU情報差分取得】 =================

type GetDiffItemSkuRequest struct {
	UpdatedAtFr string // 更新日時(Fr) (〇)
	UpdatedAtTo string // 更新日時(To) (〇)
	ItemCode    string // 商品コード (△)
	Condition   string // 抽出方法 0:=, 1:>, 2:>= (省略时 1:>)
}

type GetDiffItemSkuResponse struct {
	XMLName   xml.Name `xml:"GetItemSku"` // 文档 69页 显示返回 XML 根节点也是 GetItemSku
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string        `xml:"TotalResult,attr"`
		ResultStatus util.Status   `xml:"ResultStatus"`
		Results      []ItemSkuData `xml:"Result"`
	} `xml:"ResultSet"`
}

func (c *Client) GetDiffItemSku(ctx context.Context, req *GetDiffItemSkuRequest) (*GetDiffItemSkuResponse, error) {
	params := map[string]string{
		"updated_at_fr": req.UpdatedAtFr,
		"updated_at_to": req.UpdatedAtTo,
		"item_code":     req.ItemCode,
		"condition":     req.Condition,
	}
	var res GetDiffItemSkuResponse
	if err := c.DoGet(ctx, "/webapi2/get_diff_item_sku", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}