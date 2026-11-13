package service

import (
	"context"
	"encoding/xml"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【15. 送受信サービス状況取得】 =================

type GetServiceStatusRequest struct {
	ServiceType string // サービス区分 (〇) 1:受信サービス 2:送信サービス
}

type GetServiceStatusResponse struct {
	XMLName   xml.Name `xml:"GetServiceStatus"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string        `xml:"TotalResult,attr"`
		ResultStatus util.Status   `xml:"ResultStatus"`
		Results      []ServiceData `xml:"Result"`
	} `xml:"ResultSet"`
}

type ServiceData struct {
	No            string `xml:"No,attr"`
	ServiceStatus string `xml:"service_status"` // サービス状況 0:停止 1:停止中 2:開始中 3:開始
}

func (c *Client) GetServiceStatus(ctx context.Context, req *GetServiceStatusRequest) (*GetServiceStatusResponse, error) {
	params := map[string]string{
		"service_type": req.ServiceType,
	}
	var res GetServiceStatusResponse
	if err := c.DoGet(ctx, "/webapi2/get_service_status", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}

// ================= 【16. 送受信サービス状況更新】 =================

type UpdServiceStatusRequest struct {
	ServiceType   string // サービス区分 (〇) 1:受信サービス 2:送信サービス
	ServiceStatus string // サービス状況 (〇) 0:停止 3:開始
}

type UpdServiceStatusResponse struct {
	XMLName   xml.Name `xml:"UpdServiceStatus"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		ResultStatus struct {
			Status  string `xml:"UpdStatus"`
			Message string `xml:"Message"`
		} `xml:"ResultStatus"`
	} `xml:"ResultSet"`
}

func (c *Client) UpdServiceStatus(ctx context.Context, req *UpdServiceStatusRequest) (*UpdServiceStatusResponse, error) {
	params := map[string]string{
		"service_type":   req.ServiceType,
		"service_status": req.ServiceStatus,
	}
	var res UpdServiceStatusResponse
	if err := c.DoGet(ctx, "/webapi2/upd_service_status", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}