package crossmall

import (
	"github.com/932455569/crossmall/config"
	"github.com/932455569/crossmall/context"
	"github.com/932455569/crossmall/attribute"
	"github.com/932455569/crossmall/item"
	"github.com/932455569/crossmall/itemsku"
	"github.com/932455569/crossmall/order"
	"github.com/932455569/crossmall/phase"
	"github.com/932455569/crossmall/service"
	"github.com/932455569/crossmall/stock"
)

// CrossMall 实例
type CrossMall struct {
	ctx *context.Context
}

// NewCrossMall 实例化 CROSS MALL SDK
func NewCrossMall(cfg *config.Config) *CrossMall {
	return &CrossMall{
		ctx: &context.Context{Config: cfg},
	}
}

// GetContext 获取底层 Context (可用于自定义扩展)
func (c *CrossMall) GetContext() *context.Context {
	return c.ctx
}

// GetStock 获取库存管理 API
func (c *CrossMall) GetStock() *stock.Client { return stock.NewClient(c.ctx) }

// GetItem 获取商品管理 API
func (c *CrossMall) GetItem() *item.Client { return item.NewClient(c.ctx) }

// GetAttribute 获取商品属性管理 API
func (c *CrossMall) GetAttribute() *attribute.Client { return attribute.NewClient(c.ctx) }

// GetOrder 获取订单管理 API
func (c *CrossMall) GetOrder() *order.Client { return order.NewClient(c.ctx) }

// GetItemSKU 获取商品 SKU 管理 API
func (c *CrossMall) GetItemSKU() *itemsku.Client { return itemsku.NewClient(c.ctx) }

// GetService 获取送受信服务状况 API
func (c *CrossMall) GetService() *service.Client { return service.NewClient(c.ctx) }

// GetPhase 获取处理阶段 API
func (c *CrossMall) GetPhase() *phase.Client { return phase.NewClient(c.ctx) }