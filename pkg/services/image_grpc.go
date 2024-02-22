package services

import (
	"context"
	"log"

	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ImageGrpc struct {
	client v1.ImageServiceClient
}

func NewImageGrpc(client v1.ImageServiceClient) *ImageGrpc {
	return &ImageGrpc{client}
}

func (s *ImageGrpc) ListImages(ctx context.Context, req *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	log.Println("ListImages()", req)

	return s.client.ListImages(ctx, req)
}

func (s *ImageGrpc) ImageStatus(ctx context.Context, req *v1.ImageStatusRequest) (*v1.ImageStatusResponse, error) {
	log.Println("ImageStatus()", req)

	return s.client.ImageStatus(ctx, req)
}

func (s *ImageGrpc) PullImage(ctx context.Context, req *v1.PullImageRequest) (*v1.PullImageResponse, error) {
	log.Println("PullImage()", req)

	return s.client.PullImage(ctx, req)
}

func (s *ImageGrpc) RemoveImage(ctx context.Context, req *v1.RemoveImageRequest) (*v1.RemoveImageResponse, error) {
	log.Println("RemoveImage()", req)

	return s.client.RemoveImage(ctx, req)
}

func (s *ImageGrpc) ImageFsInfo(ctx context.Context, req *v1.ImageFsInfoRequest) (*v1.ImageFsInfoResponse, error) {
	log.Println("ImageFsInfo()", req)

	return s.client.ImageFsInfo(ctx, req)
}
