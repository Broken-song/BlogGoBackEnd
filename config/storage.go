package config

import (
	"github.com/jassue/go-storage/local"
	"github.com/jassue/go-storage/storage"
)

type Storage struct {
	Default storage.DiskName `mapstructure:"default" json:"default" yaml:"default"` // local本地
	Disks   Disks            `mapstructure:"disks" json:"disks" yaml:"disks"`
}

type Disks struct {
	Local local.Config `mapstructure:"local" json:"local" yaml:"local"`
}
