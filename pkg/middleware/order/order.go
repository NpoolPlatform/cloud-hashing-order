package order

import (
	"context"

	"github.com/NpoolPlatform/cloud-hashing-order/message/npool"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/compensate"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/gas-paying"  //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/good-paying" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/order"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/out-of-gas" //nolint
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"    //nolint

	"golang.org/x/xerrors"
)

func constructOrderDetail(
	info *npool.Order,
	goodPaying *npool.GoodPaying,
	gasPayings []*npool.GasPaying,
	compensates []*npool.Compensate,
	outOfGases []*npool.OutOfGas,
	goodPayment *npool.Payment) *npool.OrderDetail {
	return &npool.OrderDetail{
		ID:                     info.ID,
		GoodID:                 info.GoodID,
		AppID:                  info.AppID,
		UserID:                 info.UserID,
		Units:                  info.Units,
		DiscountCouponID:       info.DiscountCouponID,
		UserSpecialReductionID: info.UserSpecialReductionID,
		GoodPaying:             goodPaying,
		GasPayings:             gasPayings,
		Compensates:            compensates,
		OutOfGases:             outOfGases,
		Payment:                goodPayment,
		Start:                  info.Start,
		End:                    info.End,
		CouponID:               info.CouponID,
	}
}

func getOrderDetail(ctx context.Context, info *npool.Order) (*npool.OrderDetail, error) {
	goodPayment, err := payment.GetByOrder(ctx, &npool.GetPaymentByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("cannot find payment for order: %v", err)
	}

	goodPaying, err := goodpaying.GetByOrder(ctx, &npool.GetGoodPayingByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("cannot find good paying for order: %v", err)
	}

	gasPayings, err := gaspaying.GetByOrder(ctx, &npool.GetGasPayingsByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("cannot find gas paying for order: %v", err)
	}

	compensates, err := compensate.GetByOrder(ctx, &npool.GetCompensatesByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to get compensates for order: %v", err)
	}

	outOfGases, err := outofgas.GetByOrder(ctx, &npool.GetOutOfGasesByOrderRequest{
		OrderID: info.ID,
	})
	if err != nil {
		return nil, xerrors.Errorf("fail to get out of gas for order: %v", err)
	}

	return constructOrderDetail(
		info,
		goodPaying.Info,
		gasPayings.Infos,
		compensates.Infos,
		outOfGases.Infos,
		goodPayment.Info), nil
}

func Get(ctx context.Context, in *npool.GetOrderDetailRequest) (*npool.GetOrderDetailResponse, error) {
	info, err := order.Get(ctx, &npool.GetOrderRequest{
		ID: in.GetID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get order: %v", err)
	}

	detail, err := getOrderDetail(ctx, info.Info)
	if err != nil {
		return nil, xerrors.Errorf("fail get order detail: %v", err)
	}

	return &npool.GetOrderDetailResponse{
		Detail: detail,
	}, nil
}

func getOrdersDetail(ctx context.Context, infos []*npool.Order) ([]*npool.OrderDetail, error) {
	details := []*npool.OrderDetail{}
	for _, info := range infos {
		detail, err := getOrderDetail(ctx, info)
		if err != nil {
			return nil, xerrors.Errorf("fail get order detail: %v", err)
		}
		details = append(details, detail)
	}
	return details, nil
}

func GetByAppUser(ctx context.Context, in *npool.GetOrdersDetailByAppUserRequest) (*npool.GetOrdersDetailByAppUserResponse, error) {
	resp, err := order.GetByAppUser(ctx, &npool.GetOrdersByAppUserRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app user: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppUserResponse{
		Details: details,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetOrdersDetailByAppRequest) (*npool.GetOrdersDetailByAppResponse, error) {
	resp, err := order.GetByApp(ctx, &npool.GetOrdersByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByAppResponse{
		Details: details,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetOrdersDetailByGoodRequest) (*npool.GetOrdersDetailByGoodResponse, error) {
	resp, err := order.GetByGood(ctx, &npool.GetOrdersByGoodRequest{
		GoodID: in.GetGoodID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get orders by app: %v", err)
	}

	details, err := getOrdersDetail(ctx, resp.Infos)
	if err != nil {
		return nil, xerrors.Errorf("fail get orders detail: %v", err)
	}

	return &npool.GetOrdersDetailByGoodResponse{
		Details: details,
	}, nil
}
