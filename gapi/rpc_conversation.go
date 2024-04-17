package gapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataConversation struct {
	Items []pb.Conversation `json:"items"`
}

type DataMessage struct {
	Items []pb.Message `json:"items"`
}

type ResponseConversation struct {
	Data DataConversation `json:"data"`
}

type ResponseMessage struct {
	Data DataMessage `json:"data"`
}

func (server *ServerGRPC) ConversationList(ctx context.Context, req *pb.ListConversationRequest) (*pb.ListConversationResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	client := &http.Client{}

	url := fmt.Sprintf("%s/v1/conversations/%d?search=%s&page=%d&per_page=%d", server.config.WezoloServerAdress, req.GetOaId(), req.GetSerach(), req.GetPage(), req.GetPerPage())

	reqHttp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, nil
	}

	// Gửi request và nhận response
	resp, err := client.Do(reqHttp)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return nil, nil
	}

	// Đảm bảo rằng response body sẽ được đóng sau khi hoàn tất
	defer resp.Body.Close()

	// Đọc response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, nil
	}

	var responseData ResponseConversation
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Lỗi giải mã JSON:", err)
		return nil, nil
	}

	// Lấy các phần tử từ mảng "items" và in ra
	items := responseData.Data.Items

	var conversationsPb []*pb.Conversation
	for i := range items {
		conversationsPb = append(conversationsPb, &items[i])
	}

	return &pb.ListConversationResponse{
		Code:    200,
		Message: "success",
		Details: conversationsPb,
	}, nil
}

func (server *ServerGRPC) MessageList(ctx context.Context, req *pb.ListMessageRequest) (*pb.ListMessageResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	client := &http.Client{}

	url := fmt.Sprintf("%s/v1/conversations/%d/%s?page=%d&per_page=%d", server.config.WezoloServerAdress, req.GetOaId(), req.GetUserId(), req.GetPage(), req.GetPerPage())

	reqHttp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error creating HTTP request:", err))
	}

	// Gửi request và nhận response
	resp, err := client.Do(reqHttp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error sending HTTP request:", err))
	}

	// Đảm bảo rằng response body sẽ được đóng sau khi hoàn tất
	defer resp.Body.Close()

	// Đọc response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error reading response body:", err))
	}

	var responseData ResponseMessage
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Lỗi giải mã JSON:", err))
	}

	// Lấy các phần tử từ mảng "items" và in ra
	items := responseData.Data.Items

	var messagesPb []*pb.Message
	for i := range items {
		messagesPb = append(messagesPb, &items[i])
	}

	return &pb.ListMessageResponse{
		Code:    200,
		Message: "success",
		Details: messagesPb,
	}, nil
}
