package order

import (
	"LutiLeech/internal/domain/dto"
	"context"
)

const (
	corporateEmail = "corporate0mail@gmail.com"
	emailSubject   = "Новый заказ пиявок"
)

func (s *Service) Create(ctx context.Context, req *dto.CreateOrderRequest) (*dto.CreateOrderResponse, error) {
	leechSize1 := 0
	if req.OrderDetails.LeechSize1 != nil {
		leechSize1 = *req.OrderDetails.LeechSize1
	}

	leechSize2 := 0
	if req.OrderDetails.LeechSize2 != nil {
		leechSize2 = *req.OrderDetails.LeechSize2
	}

	leechSize3 := 0
	if req.OrderDetails.LeechSize3 != nil {
		leechSize3 = *req.OrderDetails.LeechSize3
	}

	totalPrice, err := s.priceService.CalculateTotal(
		leechSize1,
		leechSize2,
		leechSize3,
		req.OrderDetails.PackageType,
	)
	if err != nil {
		return nil, err
	}
	// TODO Если что-то там nil, автоматизировать
	emailBody, err := s.mailBuilder.
		SetFIO(req.CustomerInfo.FIO).
		SetPhoneNumber(req.CustomerInfo.PhoneNumber).
		SetEmail(req.CustomerInfo.Email).
		SetAddress(req.CustomerInfo.Address).
		SetComment(req.CustomerInfo.Comment).
		SetLeechSize1(req.OrderDetails.LeechSize1).
		SetLeechSize2(req.OrderDetails.LeechSize1).
		SetLeechSize3(req.OrderDetails.LeechSize1).
		SetTotalCount(leechSize1 + leechSize2 + leechSize3).
		SetPackageType(req.OrderDetails.PackageType).
		SetTotalPrice(totalPrice).
		Build()
	if err != nil {
		return nil, err
	}

	if err := sendOrderEmails(s.mailService, req.CustomerInfo.Email, emailBody); err != nil {
		return nil, err
	}

	return &dto.CreateOrderResponse{
		CustomerInfo: req.CustomerInfo,
		OrderDetails: req.OrderDetails,
		TotalPrice:   totalPrice,
	}, nil
}
