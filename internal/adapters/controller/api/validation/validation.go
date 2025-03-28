package validation

import (
	"LutiLeech/internal/domain/dto"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/gvalid"
	"strconv"
	"strings"
)

// minLeechSumRule Custom validation function to ensure sum of leeches is at least 50
func minLeechSumRule(ctx context.Context, in gvalid.RuleFuncInput) error {
	orderDetails := dto.OrderDetails{}
	if err := in.Value.Struct(&orderDetails); err != nil {
		return errors.New(fmt.Sprintf("invalid order details format: %v", err))
	}

	minSum := 50
	if sum, err := strconv.Atoi(strings.Split(in.Rule, ":")[1]); err == nil {
		minSum = sum
	}

	totalLeeches := orderDetails.LeechSize1 + orderDetails.LeechSize2 + orderDetails.LeechSize3
	if totalLeeches < minSum {
		return errors.New(in.Message)
	}

	return nil
}

// maxLeechSumRule Custom validation function to ensure sum of leeches is no more than 1000
func maxLeechSumRule(ctx context.Context, in gvalid.RuleFuncInput) error {
	orderDetails := dto.OrderDetails{}
	if err := in.Value.Struct(&orderDetails); err != nil {
		return errors.New(fmt.Sprintf("invalid order details format: %v", err))
	}

	maxSum := 1000
	if sum, err := strconv.Atoi(strings.Split(in.Rule, ":")[1]); err == nil {
		maxSum = sum
	}

	totalLeeches := orderDetails.LeechSize1 + orderDetails.LeechSize2 + orderDetails.LeechSize3
	if totalLeeches > maxSum {
		return errors.New(in.Message)
	}

	return nil
}

// Register custom validators
func RegisterCustomValidators() {
	gvalid.RegisterRule("min-leech-sum", minLeechSumRule)
	gvalid.RegisterRule("max-leech-sum", maxLeechSumRule)
}
