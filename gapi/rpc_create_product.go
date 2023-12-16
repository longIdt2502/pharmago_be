package gapi

import (
	"context"
	"fmt"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"path/filepath"
)

func (server *ServerGRPC) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, req.FileName)
	err := os.WriteFile(tempFilePath, req.Chunk, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %v", err)
	}

	reader, _ := os.Open(tempFilePath)
	metadata := make(map[string]string)
	_, err = server.b2Bucket.UploadFile(req.FileName, metadata, reader)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to upload image to b2", err)
	}

	return nil, status.Errorf(codes.Internal, "failed to upload image to b3")
}
