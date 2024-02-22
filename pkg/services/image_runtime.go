package services

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type RuntimeGrpc struct {
	client v1.RuntimeServiceClient
}

func NewRuntimeGrpc(client v1.RuntimeServiceClient) *RuntimeGrpc {
	return &RuntimeGrpc{client}
}

func (s *RuntimeGrpc) Version(ctx context.Context, req *v1.VersionRequest) (*v1.VersionResponse, error) {
	log.Println("Version()", req)

	return s.client.Version(ctx, req)
}

func (s *RuntimeGrpc) RunPodSandbox(ctx context.Context, req *v1.RunPodSandboxRequest) (*v1.RunPodSandboxResponse, error) {
	log.Println("RunPodSandbox()", req)

	return s.client.RunPodSandbox(ctx, req)
}

func (s *RuntimeGrpc) StopPodSandbox(ctx context.Context, req *v1.StopPodSandboxRequest) (*v1.StopPodSandboxResponse, error) {
	log.Println("StopPodSandbox()", req)

	return s.client.StopPodSandbox(ctx, req)
}

func (s *RuntimeGrpc) RemovePodSandbox(ctx context.Context, req *v1.RemovePodSandboxRequest) (*v1.RemovePodSandboxResponse, error) {
	log.Println("RemovePodSandbox()", req)

	return s.client.RemovePodSandbox(ctx, req)
}

func (s *RuntimeGrpc) PodSandboxStatus(ctx context.Context, req *v1.PodSandboxStatusRequest) (*v1.PodSandboxStatusResponse, error) {
	log.Println("PodSandboxStatus()", req)

	return s.client.PodSandboxStatus(ctx, req)
}

func (s *RuntimeGrpc) ListPodSandbox(ctx context.Context, req *v1.ListPodSandboxRequest) (*v1.ListPodSandboxResponse, error) {
	log.Println("ListPodSandbox()", req)

	return s.client.ListPodSandbox(ctx, req)
}

func (s *RuntimeGrpc) CreateContainer(ctx context.Context, req *v1.CreateContainerRequest) (*v1.CreateContainerResponse, error) {
	log.Println("CreateContainer()", req)

	return s.client.CreateContainer(ctx, req)
}

func (s *RuntimeGrpc) StartContainer(ctx context.Context, req *v1.StartContainerRequest) (*v1.StartContainerResponse, error) {
	log.Println("StartContainer()", req)

	return s.client.StartContainer(ctx, req)
}

func (s *RuntimeGrpc) StopContainer(ctx context.Context, req *v1.StopContainerRequest) (*v1.StopContainerResponse, error) {
	log.Println("StopContainer()", req)

	return s.client.StopContainer(ctx, req)
}

func (s *RuntimeGrpc) RemoveContainer(ctx context.Context, req *v1.RemoveContainerRequest) (*v1.RemoveContainerResponse, error) {
	log.Println("RemoveContainer()", req)

	return s.client.RemoveContainer(ctx, req)
}

func (s *RuntimeGrpc) ListContainers(ctx context.Context, req *v1.ListContainersRequest) (*v1.ListContainersResponse, error) {
	log.Println("ListContainers()", req)

	return s.client.ListContainers(ctx, req)
}

func (s *RuntimeGrpc) ContainerStatus(ctx context.Context, req *v1.ContainerStatusRequest) (*v1.ContainerStatusResponse, error) {
	log.Println("ContainerStatus()", req)

	return s.client.ContainerStatus(ctx, req)
}

func (s *RuntimeGrpc) UpdateContainerResources(ctx context.Context, req *v1.UpdateContainerResourcesRequest) (*v1.UpdateContainerResourcesResponse, error) {
	log.Println("UpdateContainerResources()", req)

	return s.client.UpdateContainerResources(ctx, req)
}

func (s *RuntimeGrpc) ReopenContainerLog(ctx context.Context, req *v1.ReopenContainerLogRequest) (*v1.ReopenContainerLogResponse, error) {
	log.Println("ReopenContainerLog()", req)

	return s.client.ReopenContainerLog(ctx, req)
}

func (s *RuntimeGrpc) ExecSync(ctx context.Context, req *v1.ExecSyncRequest) (*v1.ExecSyncResponse, error) {
	log.Println("ExecSync()", req)

	return s.client.ExecSync(ctx, req)
}

func (s *RuntimeGrpc) Exec(ctx context.Context, req *v1.ExecRequest) (*v1.ExecResponse, error) {
	log.Println("Exec()", req)

	return s.client.Exec(ctx, req)
}

func (s *RuntimeGrpc) Attach(ctx context.Context, req *v1.AttachRequest) (*v1.AttachResponse, error) {
	log.Println("Attach()", req)

	return s.client.Attach(ctx, req)
}

func (s *RuntimeGrpc) PortForward(ctx context.Context, req *v1.PortForwardRequest) (*v1.PortForwardResponse, error) {
	log.Println("PortForward()", req)

	return s.client.PortForward(ctx, req)
}

func (s *RuntimeGrpc) ContainerStats(ctx context.Context, req *v1.ContainerStatsRequest) (*v1.ContainerStatsResponse, error) {
	log.Println("ContainerStats()", req)

	return s.client.ContainerStats(ctx, req)
}

func (s *RuntimeGrpc) ListContainerStats(ctx context.Context, req *v1.ListContainerStatsRequest) (*v1.ListContainerStatsResponse, error) {
	log.Println("ListContainerStats()", req)

	return s.client.ListContainerStats(ctx, req)
}

func (s *RuntimeGrpc) PodSandboxStats(ctx context.Context, req *v1.PodSandboxStatsRequest) (*v1.PodSandboxStatsResponse, error) {
	log.Println("PodSandboxStats()", req)

	return s.client.PodSandboxStats(ctx, req)
}

func (s *RuntimeGrpc) ListPodSandboxStats(ctx context.Context, req *v1.ListPodSandboxStatsRequest) (*v1.ListPodSandboxStatsResponse, error) {
	log.Println("ListPodSandboxStats()", req)

	return s.client.ListPodSandboxStats(ctx, req)
}

func (s *RuntimeGrpc) UpdateRuntimeConfig(ctx context.Context, req *v1.UpdateRuntimeConfigRequest) (*v1.UpdateRuntimeConfigResponse, error) {
	log.Println("UpdateRuntimeConfig()", req)

	return s.client.UpdateRuntimeConfig(ctx, req)
}

func (s *RuntimeGrpc) Status(ctx context.Context, req *v1.StatusRequest) (*v1.StatusResponse, error) {
	log.Println("Status()", req)

	return s.client.Status(ctx, req)
}

func (s *RuntimeGrpc) CheckpointContainer(ctx context.Context, req *v1.CheckpointContainerRequest) (*v1.CheckpointContainerResponse, error) {
	log.Println("CheckpointContainer()", req)

	return s.client.CheckpointContainer(ctx, req)
}

func (s *RuntimeGrpc) GetContainerEvents(req *v1.GetEventsRequest, srv v1.RuntimeService_GetContainerEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetContainerEvents not implemented")
}

func (s *RuntimeGrpc) ListMetricDescriptors(ctx context.Context, req *v1.ListMetricDescriptorsRequest) (*v1.ListMetricDescriptorsResponse, error) {
	log.Println("ListMetricDescriptors()", req)

	return s.client.ListMetricDescriptors(ctx, req)
}

func (s *RuntimeGrpc) ListPodSandboxMetrics(ctx context.Context, req *v1.ListPodSandboxMetricsRequest) (*v1.ListPodSandboxMetricsResponse, error) {
	log.Println("ListPodSandboxMetrics()", req)

	return s.client.ListPodSandboxMetrics(ctx, req)
}

func (s *RuntimeGrpc) RuntimeConfig(ctx context.Context, req *v1.RuntimeConfigRequest) (*v1.RuntimeConfigResponse, error) {
	log.Println("RuntimeConfig()", req)

	return s.client.RuntimeConfig(ctx, req)
}
