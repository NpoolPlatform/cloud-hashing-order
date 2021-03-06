package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-order"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func assertOrder(t *testing.T, actual, expected *npool.Order) {
	assert.Equal(t, actual.GoodID, expected.GoodID)
	assert.Equal(t, actual.AppID, expected.AppID)
	assert.Equal(t, actual.UserID, expected.UserID)
	assert.Equal(t, actual.Units, expected.Units)
	assert.Equal(t, actual.DiscountCouponID, expected.DiscountCouponID)
	assert.Equal(t, actual.UserSpecialReductionID, expected.UserSpecialReductionID)
	assert.Equal(t, actual.Start, expected.Start)
	assert.Equal(t, actual.End, expected.End)
	assert.Equal(t, actual.CouponID, expected.CouponID)
	assert.Equal(t, actual.PromotionID, expected.PromotionID)
}

func TestCRUD(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	second := uint32(time.Now().Unix())

	order := npool.Order{
		GoodID:                 uuid.New().String(),
		AppID:                  uuid.New().String(),
		UserID:                 uuid.New().String(),
		Units:                  10,
		DiscountCouponID:       uuid.New().String(),
		UserSpecialReductionID: uuid.New().String(),
		Start:                  second,
		End:                    second + 20,
		CouponID:               uuid.New().String(),
		PromotionID:            uuid.New().String(),
	}

	resp, err := Create(context.Background(), &npool.CreateOrderRequest{
		Info: &order,
	})

	if assert.Nil(t, err) {
		assert.NotEqual(t, resp.Info.ID, uuid.UUID{})
		assertOrder(t, resp.Info, &order)
	}

	order.ID = resp.Info.ID
	resp1, err := Get(context.Background(), &npool.GetOrderRequest{
		ID: resp.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp1.Info.ID, resp.Info.ID)
		assertOrder(t, resp1.Info, &order)
	}

	resp2, err := GetByAppUser(context.Background(), &npool.GetOrdersByAppUserRequest{
		AppID:  resp.Info.AppID,
		UserID: resp.Info.UserID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp2.Infos), 1)
	}

	resp3, err := GetByApp(context.Background(), &npool.GetOrdersByAppRequest{
		AppID: resp.Info.AppID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp3.Infos), 1)
	}

	resp4, err := GetByGood(context.Background(), &npool.GetOrdersByGoodRequest{
		GoodID: resp.Info.GoodID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, len(resp4.Infos), 1)
	}
}
