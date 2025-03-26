package models

type Country struct {
	Idx  int    `gorm:"column:idx;type:int(11);primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;type:varchar(255);NOT NULL"`
}

func (m *Country) TableName() string {
	return "country"
}
