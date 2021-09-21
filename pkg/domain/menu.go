package domain

import (
	"context"
	"time"
)

type Menu struct {
	ID          uint32    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	ParentID    uint32    `gorm:"column:parent_id;type:int;not null"`
	Name        string    `gorm:"column:name;type:string;size:32;not null"`
	Depth       uint32    `gorm:"column:depth;type:int;not null"`
	Code        string    `gorm:"column:code;type:string;size:32;not null"`
	Enabled     bool      `gorm:"column:enabled;type:bool;not null"`
	Sort        uint32    `gorm:"column:sort;type:int;not null"`
	Version     uint32    `gorm:"column:version;type:int;not null"`
	CreatorID   uint64    `gorm:"column:creator_id;type:bigint;not null"`
	CreatorName string    `gorm:"column:creator_name;type:string;size:32;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:'1970-01-01 00:00:00';not null"`
	UpdaterID   uint64    `gorm:"column:updater_id;type:bigint;not null"`
	UpdaterName string    `gorm:"column:updater_name;type:string;size:32;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:'1970-01-01 00:00:00';not null"`
}

func (m *Menu) TableName() string {
	return "menus"
}

type MenuUsecase interface {
	Menu(ctx context.Context, id uint32) (*Menu, error)
	Menus(ctx context.Context) ([]Menu, error)
	CreatMenu(ctx context.Context, menu *Menu) error
	UpdateMenu(ctx context.Context, menu *Menu) error
	DeleteMenu(ctx context.Context, id uint32) error
}

type MenuRepository interface {
	Menu(ctx context.Context, id uint32) (*Menu, error)
	Menus(ctx context.Context) ([]Menu, error)
	CreatMenu(ctx context.Context, menu *Menu) error
	UpdateMenu(ctx context.Context, menu *Menu) error
	DeleteMenu(ctx context.Context, id uint32) error
}
