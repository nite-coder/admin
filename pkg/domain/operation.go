package domain

import (
	"context"
	"time"
)

type OperationGroup struct {
	ID   uint32 `gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name string `gorm:"column:name;type:string;size:32;not null"`
}

func (o *OperationGroup) TableName() string {
	return "operation_groups"
}

type Operation struct {
	ID          uint32    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	GroupID     uint32    `gorm:"column:group_id;type:int;not null"`
	Name        string    `gorm:"column:name;type:string;size:32;not null"`
	Description string    `gorm:"column:description;type:string;size:256;not null"`
	Code        string    `gorm:"column:code;type:string;size:32;not null"`
	Version     uint32    `gorm:"column:version;type:int;not null"`
	CreatorID   uint64    `gorm:"column:creator_id;type:bigint;not null"`
	CreatorName string    `gorm:"column:creator_name;type:string;size:32;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:'1970-01-01 00:00:00';not null"`
	UpdaterID   uint64    `gorm:"column:updater_id;type:bigint;not null"`
	UpdaterName string    `gorm:"column:updater_name;type:string;size:32;not null"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:'1970-01-01 00:00:00';not null"`
}

func (o *Operation) TableName() string {
	return "operations"
}

type OperationUsecase interface {
	Operation(ctx context.Context, id uint32) (*Operation, error)
	Operations(ctx context.Context) ([]Operation, error)
	CreatOperation(ctx context.Context, operation *Operation) error
	UpdateOperation(ctx context.Context, operation *Operation) error
	DeleteOperation(ctx context.Context, id uint32) error
}

type OperationRepository interface {
	Operation(ctx context.Context, id uint32) (*Operation, error)
	Operations(ctx context.Context) ([]Operation, error)
	CreatOperation(ctx context.Context, operation *Operation) error
	UpdateOperation(ctx context.Context, operation *Operation) error
	DeleteOperation(ctx context.Context, id uint32) error
}
