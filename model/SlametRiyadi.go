package model

type SlametRiyadi struct {
	No_transaksi   string `json:"no_transaksi" gorm:"primaryKey;column:no_transaksi"`
	No_bill        string `json:"no_bill" gorm:"column:no_bill"`
	Trans_date     string `json:"trans_date" gorm:"column:trans_date"`
	Jam            string `json:"jam" gorm:"column:jam"`
	Gross_sales    string `json:"gross_sales" gorm:"column:gross_sales"`
	Sales_discount string `json:"sales_discount" gorm:"column:sales_discount"`
	Item_discount  string `json:"item_discount" gorm:"column:item_discount"`
	Service        string `json:"service" gorm:"column:service"`
	Tax            string `json:"tax" gorm:"column:tax"`
	Sub_total      string `json:"sub_total" gorm:"column:sub_total"`
	Rounding       string `json:"rounding" gorm:"column:rounding"`
	Status         string `json:"status" gorm:"column:status"`
	Branch_id      string `json:"branch_id" grom:"column:branch_id"`
	Valid_status   string `json:"valid_status" gorm:"column:valid_status"`
	Tacharge       string `json:"tacharge" gorm:"column:tacharge"`
	Nmhhb          string `json:"nmhhb" gorm:"column:nmhhb"`
}

func (SlametRiyadi) TableName() string {
	return "api_pajak"
}
