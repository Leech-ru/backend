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

const minLeechSumRuleName = "min-leech-sum"

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

// Register custom validators
func RegisterCustomValidators() {
	gvalid.RegisterRule(minLeechSumRuleName, minLeechSumRule)
}
