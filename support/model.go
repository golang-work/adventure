package support

type Model struct {
	ID        uint  `json:"id" gorm:"primarykey"`
	CreatedAt *Time `json:"createdAt" gorm:"comment:创建时间;type:timestamp"`
	UpdatedAt *Time `json:"updatedAt" gorm:"comment:更新时间;type:timestamp"`
}
