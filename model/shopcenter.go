package model

type Goods struct {
	ID          uint   `gorm:"primarykey" db:"id"`
	SalesVolume uint   `db:"sales_volume" form:"sales_volume" json:"sales_volume" `
	GoodsName   string `db:"goods_name" form:"goods_name" json:"goods_name" `
	Cost        uint   `db:"cost" form:"cost" json:"cost"`
	Describe    string `db:"describe" form:"describe" json:"describe" `
}
type AddGoods struct {
	GoodsName string `db:"goods_name" form:"goods_name" json:"goods_name" binding:"required"`
	Cost      uint   `db:"cost" form:"cost" json:"cost" binding:"required"`
	Describe  string `db:"describe" form:"describe" json:"describe" binding:"required"`
	Secret    string `db:"secret" form:"secret" json:"secret" binding:"required"`
}
