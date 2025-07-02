package info

import "Leech-ru/internal/domain/types"

type Info struct {
	Heading     string          `json:"heading"`
	Description string          `json:"description"`
	Fluid       bool            `json:"fluid"`
	Schedule    []ScheduleEntry `json:"schedule"`
	Links       []Link          `json:"links"`
}

type ScheduleEntry struct {
	Weekday types.Weekday `json:"weekday"`
	Hours   Hours         `json:"hours"`
}

type Hours struct {
	Open  string `json:"open"`
	Close string `json:"close"`
}

type Link struct {
	Label string `json:"label"`
	Href  string `json:"href"`
}
