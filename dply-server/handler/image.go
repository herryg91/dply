package handler

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	grst_errors "github.com/herryg91/cdd/grst/errors"
	image_usecase "github.com/herryg91/dply/dply-server/app/usecase/image"
	pbImage "github.com/herryg91/dply/dply-server/handler/grst/image"
	"github.com/herryg91/dply/dply-server/pkg/interceptor"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type handlerImage struct {
	image_uc image_usecase.UseCase
	pbImage.UnimplementedImageApiServer
}

func NewImageHandler(image_uc image_usecase.UseCase) pbImage.ImageApiServer {
	return &handlerImage{image_uc: image_uc}
}
func (h *handlerImage) Get(ctx context.Context, req *pbImage.GetReq) (*pbImage.Images, error) {
	if err := pbImage.ValidateRequest(req); err != nil {
		return nil, err
	}

	datas, err := h.image_uc.Get(req.Project, req.Repository, int(req.Page), int(req.Size))
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 10000, err.Error())
	}
	resp := &pbImage.Images{
		Images: []*pbImage.Image{},
	}
	for _, data := range datas {
		resp.Images = append(resp.Images, &pbImage.Image{
			Id:          int32(data.Id),
			Digest:      data.Digest,
			Image:       data.Image,
			Project:     data.Project,
			Repository:  data.Repository,
			Description: data.Description,
			CreatedBy:   int32(data.CreatedBy),
			CreatedAt:   timestamppb.New(*data.CreatedAt),
		})
	}

	return resp, nil
}
func (h *handlerImage) Add(ctx context.Context, req *pbImage.AddReq) (*empty.Empty, error) {
	if err := pbImage.ValidateRequest(req); err != nil {
		return nil, err
	}

	userCtx := interceptor.ExtractMustLoginContext(ctx)
	err := h.image_uc.Add(req.Project, req.Repository, req.Image, req.Description, userCtx.Id)
	if err != nil {
		return nil, grst_errors.New(http.StatusInternalServerError, codes.Internal, 11000, err.Error())
	}

	return &empty.Empty{}, nil
}
