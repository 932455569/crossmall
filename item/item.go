package item

import (
	"context"
	"encoding/xml"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【6. 商品情報取得】 =================

// GetItemRequest 商品情報取得 请求
type GetItemRequest struct {
	ItemCode    string // 商品コード (△: 抽出方法 0:= 时必填)
	Condition   string // 抽出方法 0:=, 1:>, 2:>= (省略时 0:=)
	UpdatedAtFr string // 更新日時(Fr) yyyy-mm-dd hh:mm:ss
	UpdatedAtTo string // 更新日時(To) yyyy-mm-dd hh:mm:ss
	SortType    string // 並び順 0:商品コード昇順, 1:更新日時、商品コード昇順
}

// GetItemResponse 商品情報取得 响应
type GetItemResponse struct {
	XMLName   xml.Name `xml:"GetItem"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string       `xml:"TotalResult,attr"`
		ResultStatus util.Status  `xml:"ResultStatus"`
		Results      []ItemDetail `xml:"Result"`
	} `xml:"ResultSet"`
}

// ItemDetail 商品详细信息
type ItemDetail struct {
	No                   string `xml:"No,attr"`
	ItemCode             string `xml:"item_code"`               // 商品コード
	ItemName             string `xml:"item_name"`               // 商品名
	Attribute1Name       string `xml:"attribute1_name"`         // 属性1グループ名
	Attribute2Name       string `xml:"attribute2_name"`         // 属性2グループ名
	FixedPrice           string `xml:"fixed_price"`             // 定価 (税区分により定価(税別)or定価(税込))
	UnitPrice            string `xml:"unit_price"`              // 売上単価 (同上)
	NotaxFixedPrice      string `xml:"notax_fixed_price"`       // 定価(税別)
	TaxinFixedPrice      string `xml:"taxin_fixed_price"`       // 定価(税込)
	NotaxUnitPrice       string `xml:"notax_unit_price"`        // 売上単価(税別)
	TaxinUnitPrice       string `xml:"taxin_unit_price"`        // 売上単価(税込)
	NotaxPurchasePrice   string `xml:"notax_purchase_price"`    // 仕入単価(税別)
	TaxinPurchasePrice   string `xml:"taxin_purchase_price"`    // 仕入単価(税込)
	CostPrice            string `xml:"cost_price"`              // 原価
	TaxType              string `xml:"tax_type"`                // 税区分 0:非課税 1:税別 2:税込
	ReducedTaxRateType   string `xml:"reduced_tax_rate_type"`   // 軽減税率区分 0:対象外 1:対象
	FreightType          string `xml:"freight_type"`            // 送料区分 1:送料別 2:送料込
	StockType            string `xml:"stock_type"`              // 在庫管理区分 0:しない 1:する
	StockBalancerCode    string `xml:"stock_balancer_code"`     // 在庫配分パターンコード
	StockBalancerName    string `xml:"stock_balancer_name"`     // 在庫配分パターン名
	SupplierCd           string `xml:"supplier_cd"`             // 仕入先コード
	ProductCode          string `xml:"product_code"`            // 製品コード
	ItemMemo             string `xml:"item_memo"`               // 商品備考
	ShelfCd              string `xml:"shelf_cd"`                // 棚番
	ItemGoods            string `xml:"item_goods"`              // 品目
	MainCategory1Name    string `xml:"main_category1_name"`     // メインカテゴリ1名
	MainCategory2Name    string `xml:"main_category2_name"`     // メインカテゴリ2名
	MainCategory3Name    string `xml:"main_category3_name"`     // メインカテゴリ3名
	SubCategory1Name     string `xml:"sub_category1_name"`      // サブカテゴリ1名
	SubCategory2Name     string `xml:"sub_category2_name"`      // サブカテゴリ2名
	SubCategory3Name     string `xml:"sub_category3_name"`      // サブカテゴリ3名
	UpdatedAt            string `xml:"updated_at"`              // 更新日時
}

// GetItem 获取商品信息
func (c *Client) GetItem(ctx context.Context, req *GetItemRequest) (*GetItemResponse, error) {
	params := map[string]string{
		"item_code":     req.ItemCode,
		"condition":     req.Condition,
		"updated_at_fr": req.UpdatedAtFr,
		"updated_at_to": req.UpdatedAtTo,
		"sort_type":     req.SortType,
	}
	var res GetItemResponse
	if err := c.DoGet(ctx, "/webapi2/get_item", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【7. 商品情報更新】 =================

// UpdItemRequest 商品情報更新 请求
type UpdItemRequest struct {
	ItemCode             string // 商品コード (〇)
	ItemName             string // 商品名 (〇)
	Attribute1Group      string // 属性グループ1名
	Attribute2Group      string // 属性グループ2名
	FixedPrice           string // 定価
	UnitPrice            string // 売上単価
	NotaxFixedPrice      string // 定価(税別)
	TaxinFixedPrice      string // 定価(税込)
	NotaxUnitPrice       string // 売上単価(税別)
	TaxinUnitPrice       string // 売上単価(税込)
	NotaxPurchasePrice   string // 仕入単価(税別)
	TaxinPurchasePrice   string // 仕入単価(税込)
	CostPrice            string // 原価
	TaxType              string // 税区分 0:非課税 1:税別 2:税込
	ReducedTaxRateType   string // 軽減税率区分 0:対象外 1:対象
	FreightType          string // 送料区分 1:送料別 2:送料込
	StockType            string // 在庫管理区分 0:しない 1:する
	StockBalancerCode    string // 在庫配分パターンコード
	SupplierCode         string // 仕入先コード
	ProductCode          string // 製品コード
	ItemMemo             string // 商品備考
	ShelfCd              string // 棚番
	ItemGoods            string // 品目
	MainCategory1Name    string // メインカテゴリ1名
	MainCategory2Name    string // メインカテゴリ2名
	MainCategory3Name    string // メインカテゴリ3名
	SubCategory1Name     string // サブカテゴリ1名
	SubCategory2Name     string // サブカテゴリ2名
	SubCategory3Name     string // サブカテゴリ3名
	CategoryAdd          string // カテゴリ追加 0:しない 1:する
	UpdateMode           string // 更新モード (〇) u:更新 d:削除
}

// UpdItemResponse 商品情報更新 响应
type UpdItemResponse struct {
	XMLName   xml.Name    `xml:"UpdItem"`
	Version   string      `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

// UpdItem 更新商品信息
func (c *Client) UpdItem(ctx context.Context, req *UpdItemRequest) (*UpdItemResponse, error) {
	params := map[string]string{
		"item_code":              req.ItemCode,
		"item_name":              req.ItemName,
		"attribute1_group":       req.Attribute1Group,
		"attribute2_group":       req.Attribute2Group,
		"fixed_price":            req.FixedPrice,
		"unit_price":             req.UnitPrice,
		"notax_fixed_price":      req.NotaxFixedPrice,
		"taxin_fixed_price":      req.TaxinFixedPrice,
		"notax_unit_price":       req.NotaxUnitPrice,
		"taxin_unit_price":       req.TaxinUnitPrice,
		"notax_purchase_price":   req.NotaxPurchasePrice,
		"taxin_purchase_price":   req.TaxinPurchasePrice,
		"cost_price":             req.CostPrice,
		"tax_type":               req.TaxType,
		"reduced_tax_rate_type":  req.ReducedTaxRateType,
		"freight_type":           req.FreightType,
		"stock_type":             req.StockType,
		"stock_balancer_code":    req.StockBalancerCode,
		"supplier_code":          req.SupplierCode,
		"product_code":           req.ProductCode,
		"item_memo":              req.ItemMemo,
		"shelf_cd":               req.ShelfCd,
		"item_goods":             req.ItemGoods,
		"main_category1_name":    req.MainCategory1Name,
		"main_category2_name":    req.MainCategory2Name,
		"main_category3_name":    req.MainCategory3Name,
		"sub_category1_name":     req.SubCategory1Name,
		"sub_category2_name":     req.SubCategory2Name,
		"sub_category3_name":     req.SubCategory3Name,
		"category_add":           req.CategoryAdd,
		"update_mode":            req.UpdateMode,
	}
	var res UpdItemResponse
	if err := c.DoGet(ctx, "/webapi2/upd_item", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}