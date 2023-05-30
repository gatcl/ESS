package model

import (
	"ess/global"

	"gorm.io/gorm"
)

type Model struct {
	ID       uint64         `gorm:"primarykey;" json:"id"`
	CreateAt global.Time    `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"create_at"`
	UpdateAt global.Time    `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at"`
	DeleteAt gorm.DeletedAt `gorm:"index;" json:"delete_at"`
}
