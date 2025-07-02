package info

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

func (s infoService) Get(_ context.Context) (*dto.GetInfoResponse, error) {
	info, err := s.jsonInfoRepository.Read()
	if err != nil {
		return nil, err
	}

	response := &dto.GetInfoResponse{
		Heading:     info.Heading,
		Description: info.Description,
		Fluid:       &info.Fluid,
	}

	if len(info.Schedule) > 0 {
		var schedule []dto.ScheduleEntry
		for _, s := range info.Schedule {
			schedule = append(schedule, dto.ScheduleEntry{
				Weekday: s.Weekday,
				Hours: dto.Hours{
					Open:  s.Hours.Open,
					Close: s.Hours.Close,
				},
			})
		}
		response.Schedule = schedule
	}

	if len(info.Links) > 0 {
		var links []dto.InfoLinks
		for _, l := range info.Links {
			links = append(links, dto.InfoLinks{
				Label: l.Label,
				Href:  l.Href,
			})
		}
		response.Links = links
	}

	return response, nil
}
