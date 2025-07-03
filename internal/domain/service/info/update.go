package info

import (
	jsonInfo "Leech-ru/internal/adapters/repository/json/info"
	"Leech-ru/internal/domain/dto"
	"context"
)

func (s infoService) Update(_ context.Context, req *dto.UpdateInfoRequest) (*dto.UpdateInfoResponse, error) {
	current, err := s.jsonInfoRepository.Read()
	if err != nil {
		return nil, err
	}

	// Обновление полей, если они заданы
	if req.Heading != nil {
		current.Heading = *req.Heading
	}
	if req.Description != nil {
		current.Description = *req.Description
	}
	if req.Fluid != nil {
		current.Fluid = *req.Fluid
	}
	if req.Schedule != nil {
		var newSchedule []jsonInfo.ScheduleEntry
		for _, s := range *req.Schedule {
			newSchedule = append(newSchedule, jsonInfo.ScheduleEntry{
				Weekday: s.Weekday,
				Hours: jsonInfo.Hours{
					Open:  s.Hours.Open,
					Close: s.Hours.Close,
				},
			})
		}
		current.Schedule = newSchedule
	}
	if req.Links != nil {
		var newLinks []jsonInfo.Link
		for _, l := range *req.Links {
			newLinks = append(newLinks, jsonInfo.Link{
				Label: l.Label,
				Href:  l.Href,
			})
		}
		current.Links = newLinks
	}

	updated, err := s.jsonInfoRepository.Write(current)
	if err != nil {
		return nil, err
	}

	resp := &dto.UpdateInfoResponse{
		Heading:     updated.Heading,
		Description: updated.Description,
		Fluid:       &updated.Fluid,
	}

	if len(updated.Schedule) > 0 {
		for _, s := range updated.Schedule {
			resp.Schedule = append(resp.Schedule, dto.ScheduleEntry{
				Weekday: s.Weekday,
				Hours: dto.Hours{
					Open:  s.Hours.Open,
					Close: s.Hours.Close,
				},
			})
		}
	}

	if len(updated.Links) > 0 {
		for _, l := range updated.Links {
			resp.Links = append(resp.Links, dto.InfoLinks{
				Label: l.Label,
				Href:  l.Href,
			})
		}
	}

	return resp, nil
}
