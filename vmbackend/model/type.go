package model

import "time"

type Good struct {
	GoodID string `json:"goodsId" gorm:"column:货物号;primaryKey"`
	Name   string `json:"name" gorm:"column:货物名"`
	Spec   string `json:"spec" gorm:"column:规格"`
	Model  string `json:"model" gorm:"column:型号"`
	Desc   string `json:"desc" gorm:"column:说明"`
}

func (Good) TableName() string {
	return "货物信息表"
}

//入库
type Inbound struct {
	InboundTime time.Time `json:"inboundTime" gorm:"column:入库时间;type:datetime"`
	GoodID      string    `json:"goodsId" gorm:"column:货物号;type:char(4)"`
	Quantity    float64   `json:"quantity" gorm:"column:入库数量;type:decimal(9,2)"`
	Operator    string    `json:"operator" gorm:"column:经办人;type:char(10)"`
}

func (Inbound) TableName() string {
	return "入库表"
}

type InboundGood struct {
	Inbound
	Name  string `json:"name" gorm:"column:货物名;type:char(20)"`
	Spec  string `json:"spec" gorm:"column:规格;type:char(20)"`
	Model string `json:"model" gorm:"column:型号;type:char(20)"`
}

func (InboundGood) TableName() string {
	return "入库信息"
}

//出库
type Outbound struct {
	OutboundTime time.Time `json:"outboundTime" gorm:"column:出库时间;type:datetime"`
	GoodID       string    `json:"goodsId" gorm:"column:货物号;type:char(4)"`
	Quantity     float64   `json:"quantity" gorm:"column:出库数量;type:decimal(9,2)"`
	Operator     string    `json:"operator" gorm:"column:经办人;type:char(10)"`
}

func (Outbound) TableName() string {
	return "出库表"
}

type OutboundGood struct {
	Outbound
	Name  string `json:"name" gorm:"column:货物名;type:char(20)"`
	Spec  string `json:"spec" gorm:"column:规格;type:char(20)"`
	Model string `json:"model" gorm:"column:型号;type:char(20)"`
}

func (OutboundGood) TableName() string {
	return "出库信息"
}

//库存
type Inventory struct {
	GoodID     string    `json:"goodsId" gorm:"column:货物号;type:char(4)"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:更新日期;type:datetime"`
	Quantity   float64   `json:"quantity" gorm:"column:库存量;type:decimal(9,2)"`
}

func (Inventory) TableName() string {
	return "库存表"
}

type InventoryGood struct {
	Inventory
	Name  string `json:"name" gorm:"column:货物名;type:char(20)"`
	Spec  string `json:"spec" gorm:"column:规格;type:char(20)"`
	Model string `json:"model" gorm:"column:型号;type:char(20)"`
}

func (InventoryGood) TableName() string {
	return "库存信息"
}
