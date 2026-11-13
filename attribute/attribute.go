package attribute

import (
	"context"
	"encoding/xml"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【8. 商品属性情報取得】 =================

type GetItemAttributeRequest struct {
	ItemCode      string // 商品コード (〇)
	AttributeType string // 商品属性タイプ (〇) 1:属性① 2:属性②
	AttributeCode string // 商品属性コード
}

type GetItemAttributeResponse struct {
	XMLName   xml.Name `xml:"GetItemAttribute"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string          `xml:"TotalResult,attr"`
		ResultStatus util.Status     `xml:"ResultStatus"`
		Results      []AttributeItem `xml:"Result"`
	} `xml:"ResultSet"`
}

type AttributeItem struct {
	No             string `xml:"No,attr"`
	AttributeCode  string `xml:"attribute_code"`
	AttributeName  string `xml:"attribute_name"`
	SortNo         string `xml:"sort_no"`         // 並び順
}

func (c *Client) GetItemAttribute(ctx context.Context, req *GetItemAttributeRequest) (*GetItemAttributeResponse, error) {
	params := map[string]string{
		"item_code":      req.ItemCode,
		"attribute_type": req.AttributeType,
		"attribute_code": req.AttributeCode,
	}
	var res GetItemAttributeResponse
	if err := c.DoGet(ctx, "/webapi2/get_item_attribute", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【9. 商品属性情報更新】 =================

type UpdItemAttributeRequest struct {
	ItemCode      string // 商品コード (〇)
	AttributeType string // 商品属性タイプ (〇) 1:属性① 2:属性②
	AttributeCode string // 商品属性コード (〇)
	AttributeName string // 商品属性名称 (〇 新規/修正 时)
	SortNo        string // 並び順
	UpdateMode    string // 更新モード (〇) u:更新 d:削除
}

type UpdItemAttributeResponse struct {
	XMLName   xml.Name `xml:"UpdItemAttribute"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

func (c *Client) UpdItemAttribute(ctx context.Context, req *UpdItemAttributeRequest) (*UpdItemAttributeResponse, error) {
	params := map[string]string{
		"item_code":      req.ItemCode,
		"attribute_type": req.AttributeType,
		"attribute_code": req.AttributeCode,
		"attribute_name": req.AttributeName,
		"sort_no":        req.SortNo,
		"update_mode":    req.UpdateMode,
	}
	var res UpdItemAttributeResponse
	if err := c.DoGet(ctx, "/webapi2/upd_item_attribute", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}