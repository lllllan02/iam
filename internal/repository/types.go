package repository

import "gorm.io/gorm"

type Option func(*gorm.DB) *gorm.DB
