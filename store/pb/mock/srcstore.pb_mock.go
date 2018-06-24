// generated by gen-mocks; DO NOT EDIT

package mock

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sourcegraph.com/sourcegraph/srclib/store/pb"
	"sourcegraph.com/sqs/pbtypes"
)

type MultiRepoImporterClient struct {
	Import_        func(ctx context.Context, in *pb.ImportOp) (*pbtypes.Void, error)
	CreateVersion_ func(ctx context.Context, in *pb.CreateVersionOp) (*pbtypes.Void, error)
	Index_         func(ctx context.Context, in *pb.IndexOp) (*pbtypes.Void, error)
}

func (s *MultiRepoImporterClient) Import(ctx context.Context, in *pb.ImportOp, opts ...grpc.CallOption) (*pbtypes.Void, error) {
	return s.Import_(ctx, in)
}

func (s *MultiRepoImporterClient) CreateVersion(ctx context.Context, in *pb.CreateVersionOp, opts ...grpc.CallOption) (*pbtypes.Void, error) {
	return s.CreateVersion_(ctx, in)
}

func (s *MultiRepoImporterClient) Index(ctx context.Context, in *pb.IndexOp, opts ...grpc.CallOption) (*pbtypes.Void, error) {
	return s.Index_(ctx, in)
}

var _ pb.MultiRepoImporterClient = (*MultiRepoImporterClient)(nil)

type MultiRepoImporterServer struct {
	Import_        func(v0 context.Context, v1 *pb.ImportOp) (*pbtypes.Void, error)
	CreateVersion_ func(v0 context.Context, v1 *pb.CreateVersionOp) (*pbtypes.Void, error)
	Index_         func(v0 context.Context, v1 *pb.IndexOp) (*pbtypes.Void, error)
}

func (s *MultiRepoImporterServer) Import(v0 context.Context, v1 *pb.ImportOp) (*pbtypes.Void, error) {
	return s.Import_(v0, v1)
}

func (s *MultiRepoImporterServer) CreateVersion(v0 context.Context, v1 *pb.CreateVersionOp) (*pbtypes.Void, error) {
	return s.CreateVersion_(v0, v1)
}

func (s *MultiRepoImporterServer) Index(v0 context.Context, v1 *pb.IndexOp) (*pbtypes.Void, error) {
	return s.Index_(v0, v1)
}

var _ pb.MultiRepoImporterServer = (*MultiRepoImporterServer)(nil)
