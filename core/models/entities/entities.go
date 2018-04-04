package entities

// ItemInfo store project information
type ItemInfo struct {
	ItemID int `xorm:"autoincr pk 'item_id'"`
	Label  string
	IsFinished bool
	CreateAt string
}

func (u ItemInfo) TableName() string {
	return "items"
}
