// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"github.com/donknap/dpanel/common/accessor"
)

const TableNameBackup = "ims_backup"

// Backup mapped from table <ims_backup>
type Backup struct {
	ID          int32                         `gorm:"column:id;primaryKey" json:"id"`
	ContainerID string                        `gorm:"column:container_id" json:"containerId"`
	Setting     *accessor.BackupSettingOption `gorm:"column:setting;serializer:json" json:"setting"`
	CreatedAt   time.Time                     `gorm:"column:created_at" json:"createdAt"`
}

// TableName Backup's table name
func (*Backup) TableName() string {
	return TableNameBackup
}
