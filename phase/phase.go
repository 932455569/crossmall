package phase

import (
	"context"
	"encoding/xml"

	cmctx "github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/util"
)

type Client struct{ *cmctx.Context }

func NewClient(ctx *cmctx.Context) *Client { return &Client{Context: ctx} }

// ================= 【19. 処理フェーズ情報取得】 =================

type GetPhaseResponse struct {
	XMLName   xml.Name `xml:"GetPhase"`
	Version   string   `xml:"version,attr"`
	ResultSet struct {
		TotalResult  string      `xml:"TotalResult,attr"`
		ResultStatus util.Status `xml:"ResultStatus"`
		Results      []PhaseData `xml:"Result"`
	} `xml:"ResultSet"`
}

type PhaseData struct {
	No        string `xml:"No,attr"`
	PhaseName string `xml:"phase_name"` // 処理フェーズ名
}

func (c *Client) GetPhase(ctx context.Context) (*GetPhaseResponse, error) {
	params := map[string]string{}
	var res GetPhaseResponse
	if err := c.DoGet(ctx, "/webapi2/get_phase", params, &res); err != nil {
		return nil, err
	}
	if res.ResultSet.ResultStatus.Status == "error" {
		return &res, &util.APIError{Message: res.ResultSet.ResultStatus.Message}
	}
	return &res, nil
}