package order

import (
	"context"
	"encoding/xml"
	"strings"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【10. 注文伝票情報取得】 =================

type GetOrderRequest struct {
	OrderNumber            string   // 管理番号 (△: 0:= 时必填)
	OrderCode              string   // 注文番号
	OrderDateFr            string   // 注文日(Fr)
	OrderDateTo            string   // 注文日(To)
	CreatedAtFr            string   // 作成日時(Fr)
	CreatedAtTo            string   // 作成日時(To)
	UpdatedAtFr            string   // 更新日時(Fr)
	UpdatedAtTo            string   // 更新日時(To)
	PhaseName              string   // 処理フェーズ
	CancelFlag             string   // キャンセルフラグ 0:通常 1:保留 2:キャンセル
	ShopCode               string   // 店舗コード
	DeliveryCode           string   // 配送便コード
	PaymentCode            string   // 支払方法コード
	CheckMark1             string   // チェックマーク1 0:OFF 1:ON
	CheckMark2             string   // チェックマーク2 0:OFF 1:ON
	CheckMark3             string   // チェックマーク3 0:OFF 1:ON
	Condition              string   // 抽出方法 0:=, 1:>, 2:>= (省略时 1:>)
	CreatingSourceType     []string // 作成方法 (支持多选，如 []string{"3", "4"})
	DeliveryDate           string   // 発送日
	OrderOption1           string   // 取引オプション1~10
	SenderType             string   // 出荷元 0:通常, 1:AmazonFBA, 2:FBAマルチチャネル...
	SenderTypeCondition    string   // 出荷元抽出方法 0:ALL, 1:EXISTS, 2:NOT EXISTS
	SocialGiftReplaceOrder string   // ソーシャルギフト注文者情報置換 0:置換しない, 1:置換する
}

type GetOrderResponse struct {
	XMLName   xml.Name `xml:"GetOrder"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string      `xml:"TotalResult,attr"`
		ResultStatus util.Status `xml:"ResultStatus"`
		Results      []OrderData `xml:"Result"`
	} `xml:"ResultSet"`
}

type OrderData struct {
	No             string `xml:"No,attr"`
	OrderNumber    string `xml:"order_number"`     // 管理番号
	OrderDate      string `xml:"order_date"`       // 注文日時
	ShopCode       string `xml:"shop_code"`        // 店舗コード
	ShopName       string `xml:"shop_name"`        // 店舗名
	OrderCode      string `xml:"order_code"`       // 注文番号
	ClientName     string `xml:"client_name"`      // 注文者氏名
	// ... (此处省略了文档中的几十个字段，实际开发请对照 PDF 27-29页补齐)
	PhaseName      string `xml:"phase_name"`       // 処理フェーズ名
	CancelFlag     string `xml:"cancel_flag"`      // キャンセルフラグ
	TotalPrice     string `xml:"total_price"`      // 合計
	SocialGiftFlag string `xml:"social_gift_flag"` // ソーシャルギフトフラグ
	// ...
}

func (c *Client) GetOrder(ctx context.Context, req *GetOrderRequest) (*GetOrderResponse, error) {
	params := map[string]string{
		"order_number":             req.OrderNumber,
		"order_code":               req.OrderCode,
		"order_date_fr":            req.OrderDateFr,
		"order_date_to":            req.OrderDateTo,
		"created_at_fr":            req.CreatedAtFr,
		"created_at_to":            req.CreatedAtTo,
		"updated_at_fr":            req.UpdatedAtFr,
		"updated_at_to":            req.UpdatedAtTo,
		"phase_name":               req.PhaseName,
		"cancel_flag":              req.CancelFlag,
		"shop_code":                req.ShopCode,
		"delivery_code":            req.DeliveryCode,
		"payment_code":             req.PaymentCode,
		"check_mark1":              req.CheckMark1,
		"check_mark2":              req.CheckMark2,
		"check_mark3":              req.CheckMark3,
		"condition":                req.Condition,
		"delivery_date":            req.DeliveryDate,
		"order_option1":            req.OrderOption1,
		"sender_type":              req.SenderType,
		"sender_type_condition":    req.SenderTypeCondition,
		"social_gift_replace_order": req.SocialGiftReplaceOrder,
	}

	// 特殊处理数组参数 creating_source_type[]
	if len(req.CreatingSourceType) > 0 {
		params["creating_source_type[]"] = strings.Join(req.CreatingSourceType, ",")
	}

	var res GetOrderResponse
	if err := c.DoGet(ctx, "/webapi2/get_order", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【11. 注文詳細情報取得】 =================

type GetOrderDetailRequest struct {
	OrderNumber string // 管理番号 (〇)
}

type GetOrderDetailResponse struct {
	XMLName   xml.Name `xml:"GetOrderDetail"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string            `xml:"TotalResult,attr"`
		ResultStatus util.Status       `xml:"ResultStatus"`
		Results      []OrderDetailData `xml:"Result"`
	} `xml:"ResultSet"`
}

type OrderDetailData struct {
	No                string `xml:"No,attr"`
	OrderNumber       string `xml:"order_number"` // 管理番号
	LineNo            string `xml:"line_no"`      // 行No
	ItemCode          string `xml:"item_code"`    // 商品コード
	ItemName          string `xml:"item_name"`    // 商品名
	Attribute1Code    string `xml:"attribute1_code"`
	Attribute1Name    string `xml:"attribute1_name"`
	Attribute2Code    string `xml:"attribute2_code"`
	Attribute2Name    string `xml:"attribute2_name"`
	Amount            string `xml:"amount"`       // 数量
	UnitPrice         string `xml:"unit_price"`   // 単価
	AmountPrice       string `xml:"amount_price"` // 金額
	TaxType           string `xml:"tax_type"`     // 税区分
	ReducedTaxRateType string `xml:"reduced_tax_rate_type"` // 軽減税率区分
	// ... (其他字段略，可对照文档 34-35页)
	SenderType        string `xml:"sender_type"`       // 出荷元
	TargetStockType   string `xml:"target_stock_type"` // 在庫引当元
}

func (c *Client) GetOrderDetail(ctx context.Context, req *GetOrderDetailRequest) (*GetOrderDetailResponse, error) {
	params := map[string]string{"order_number": req.OrderNumber}
	var res GetOrderDetailResponse
	if err := c.DoGet(ctx, "/webapi2/get_order_detail", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【12. 注文構成品情報取得】 =================

type GetOrderComponentRequest struct {
	OrderNumber string // 管理番号 (〇)
	LineNo      string // 行No (〇)
}

type GetOrderComponentResponse struct {
	XMLName   xml.Name `xml:"GetOrderComponent"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string               `xml:"TotalResult,attr"`
		ResultStatus util.Status          `xml:"ResultStatus"`
		Results      []OrderComponentData `xml:"Result"`
	} `xml:"ResultSet"`
}

type OrderComponentData struct {
	No             string `xml:"No,attr"`
	OrderNumber    string `xml:"order_number"`
	LineNo         string `xml:"line_no"`
	SetItemCode    string `xml:"set_item_code"`   // セット商品コード
	ItemCode       string `xml:"item_code"`       // 商品コード
	ItemName       string `xml:"item_name"`       // 商品名
	Attribute1Code string `xml:"attribute1_code"`
	Attribute1Name string `xml:"attribute1_name"`
	Attribute2Code string `xml:"attribute2_code"`
	Attribute2Name string `xml:"attribute2_name"`
	ComponentCount string `xml:"component_count"` // 員数
	JanCd          string `xml:"jan_cd"`          // JANコード
}

func (c *Client) GetOrderComponent(ctx context.Context, req *GetOrderComponentRequest) (*GetOrderComponentResponse, error) {
	params := map[string]string{
		"order_number": req.OrderNumber,
		"line_no":      req.LineNo,
	}
	var res GetOrderComponentResponse
	if err := c.DoGet(ctx, "/webapi2/get_order_component", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【13. 注文伝票フェーズ（発送情報）更新】 =================

type UpdOrderPhaseRequest struct {
	OrderNumber              string // 管理番号 (〇)
	BeforePhaseName          string // 更新前処理フェーズ
	AfterPhaseName           string // 更新後処理フェーズ (〇)
	DeliveryDate             string // 発送日 yyyy-mm-dd
	DeliveryDateUpdate       string // 発送日上書き 0:しない 1:する
	DeliveryNo               string // 配送番号
	DeliveryName             string // 配送便名
	DeliveryTimeConvertError string // 配送時間変換エラー 0:更新不可 1:未設定で更新
}

type UpdOrderPhaseResponse struct {
	XMLName   xml.Name `xml:"UpdOrderStatus"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

func (c *Client) UpdOrderPhase(ctx context.Context, req *UpdOrderPhaseRequest) (*UpdOrderPhaseResponse, error) {
	params := map[string]string{
		"order_number":                req.OrderNumber,
		"before_phase_name":           req.BeforePhaseName,
		"after_phase_name":            req.AfterPhaseName,
		"delivery_date":               req.DeliveryDate,
		"delivery_date_update":        req.DeliveryDateUpdate,
		"delivery_no":                 req.DeliveryNo,
		"delivery_name":               req.DeliveryName,
		"delivery_time_convert_error": req.DeliveryTimeConvertError,
	}
	var res UpdOrderPhaseResponse
	if err := c.DoGet(ctx, "/webapi2/upd_order_phase", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【14. 注文伝票Mark更新】 =================

type UpdOrderCheckMarkRequest struct {
	OrderNumber    string // 管理番号 (〇)
	CheckMarkType  string // チェックマーク区分 (〇) 1:Check_Mark1 2:Check_Mark2 3:Check_Mark3
	CheckMarkValue string // チェックマーク値 (〇) 0:OFF 1:ON
}

type UpdOrderCheckMarkResponse struct {
	XMLName   xml.Name `xml:"UpdOrderCheckMark"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

func (c *Client) UpdOrderCheckMark(ctx context.Context, req *UpdOrderCheckMarkRequest) (*UpdOrderCheckMarkResponse, error) {
	params := map[string]string{
		"order_number":     req.OrderNumber,
		"check_mark_type":  req.CheckMarkType,
		"check_mark_value": req.CheckMarkValue,
	}
	var res UpdOrderCheckMarkResponse
	if err := c.DoGet(ctx, "/webapi2/upd_order_check_mark", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}