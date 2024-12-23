package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
	"strings"
)

type CreateBonsaiStorage interface {
	CreateItem(ctx context.Context, data *model.BonsaiItemCreate) error
}

type createBonsaiBiz struct {
	storage CreateBonsaiStorage
}

func NewCreateBonsaiBiz(storage CreateBonsaiStorage) *createBonsaiBiz {
	return &createBonsaiBiz{storage: storage}
}

func (biz *createBonsaiBiz) CreateNewBonsai(ctx context.Context, data *model.BonsaiItemCreate) error {

	title := strings.TrimSpace(data.TenSP)

	if title != "" {
		return model.ErrTitleIsBlank
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}