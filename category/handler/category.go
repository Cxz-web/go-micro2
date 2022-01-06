package handler

import (
	"context"
	"github.com/Cxz-web/go-micro2/user/domain/catefory/common"
	"github.com/Cxz-web/go-micro2/user/domain/catefory/domain/model"
	"github.com/Cxz-web/go-micro2/user/domain/catefory/domain/service"
	pb "github.com/Cxz-web/go-micro2/user/domain/catefory/proto"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func (c *CategoryHandler) CreateCategory(context context.Context, request *pb.CategoryInfo, resp *pb.CreateCategoryResponse) error {

	newCategory := &model.Category{}

	err := common.SwapTo(request, newCategory)

	if err != nil {
		return err
	}

	resp.Id = newCategory.ID

	return nil
}

func (c *CategoryHandler) FindCategoryById(ctx context.Context, in *pb.FindCategoryByIdRequest, our *pb.CategoryInfo) error {
	category, err := c.service.FindCategoryById(in.Id)

	if err != nil {
		return err
	}

	our.CategoryName = category.CategoryName

	our.CategoryCode = category.CategoryCode

	return nil
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		&service.CategoryService{},
	}
}
