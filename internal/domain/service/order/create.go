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
	totalPrice, err := s.priceService.CalculateTotal(
		req.OrderDetails.LeechSize1,
		req.OrderDetails.LeechSize2,
		req.OrderDetails.LeechSize3,
		req.OrderDetails.PackageType,
	)
	if err != nil {
		return nil, err
	}
	emailBody, err := s.mailBuilder.
		SetFIO(req.CustomerInfo.FIO).
		SetPhoneNumber(req.CustomerInfo.PhoneNumber).
		SetEmail(req.CustomerInfo.Email).
		SetAddress(req.CustomerInfo.Address).
		SetComment(req.CustomerInfo.Comment).
		SetLeechSize1(req.OrderDetails.LeechSize1).
		SetLeechSize2(req.OrderDetails.LeechSize2).
		SetLeechSize3(req.OrderDetails.LeechSize3).
		SetTotalCount(req.OrderDetails.LeechSize1 + req.OrderDetails.LeechSize2 + req.OrderDetails.LeechSize3).
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
