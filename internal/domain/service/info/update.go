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
				Hours:   jsonInfo.Hours(s.Hours),
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

	if err := s.jsonInfoRepository.Write(current); err != nil {
		return nil, err
	}

	resp := &dto.UpdateInfoResponse{
		Heading:     current.Heading,
		Description: current.Description,
		Fluid:       &current.Fluid,
	}

	if len(current.Schedule) > 0 {
		for _, s := range current.Schedule {
			resp.Schedule = append(resp.Schedule, dto.ScheduleEntry{
				Weekday: s.Weekday,
				Hours: dto.Hours{
					Open:  s.Hours.Open,
					Close: s.Hours.Close,
				},
			})
		}
	}

	// Конвертация ссылок
	if len(current.Links) > 0 {
		for _, l := range current.Links {
			resp.Links = append(resp.Links, dto.InfoLinks{
				Label: l.Label,
				Href:  l.Href,
			})
		}
	}

	return resp, nil
}
