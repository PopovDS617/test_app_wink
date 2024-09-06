package lb

import (
	"context"
	"errors"
	"log"
	"testappwink/pkg/customerr"
	gen "testappwink/pkg/lb_v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetVideoURL(ctx context.Context, req *gen.GetVideoURLRequest) (*gen.GetVideoURLResponse, error) {

	videoURL, err := i.lbService.GetURL(ctx, req.Video)

	if err != nil {

		if errors.Is(err, customerr.ErrInvalidURL) {
			return nil, status.Error(codes.Code(400), "not valid url")
		}
		return nil, status.Error(codes.InvalidArgument, "invalid input")

	}

	log.Print(videoURL)

	return &gen.GetVideoURLResponse{
		Video: videoURL,
	}, nil
}
