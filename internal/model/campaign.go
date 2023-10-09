package model

import (
	"fmt"
	"strings"
	"time"
)

type Campaign struct {
	ID        int        `json:"id"`
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Amount    float64    `json:"amount"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	Quota     int        `json:"quota"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (c *Campaign) GenerateCode() {
	c.Code = fmt.Sprintf("%s%d%d", strings.ToLower(c.Name[:5]), c.Quota, c.EndDate.Unix())
}
