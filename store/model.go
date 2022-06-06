package store

type License struct {
	License string `gorm:"license"`
}

func (l *License) TableName() string {
	return "license"
}
