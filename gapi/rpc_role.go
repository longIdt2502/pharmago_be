package gapi

import (
	"context"
	"database/sql"
	"fmt"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (server *ServerGRPC) RoleCreate(ctx context.Context, req *pb.RoleCreateRequest) (*pb.RoleCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	params := db.CreateRoleTxParams{
		CreateRoleParams: db.CreateRoleParams{
			Code:  req.GetCode(),
			Title: req.GetTitle(),
			Note: sql.NullString{
				String: req.GetNote(),
				Valid:  req.Note != nil,
			},
			Company: sql.NullInt32{
				Int32: req.GetCompany(),
				Valid: true,
			},
			UserCreated: tokenPayload.UserID,
			UserUpdated: sql.NullInt32{
				Int32: tokenPayload.UserID,
				Valid: true,
			},
		},
		Items: req.GetItems(),
	}

	role, err := server.store.CreateRoleTx(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record: %v", err))
	}

	return &pb.RoleCreateResponse{
		Code:    200,
		Message: "success",
		Details: role.Id,
	}, nil
}

func (server *ServerGRPC) RoleList(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListResponse, error) {
	roles, err := server.store.ListRole(ctx, db.ListRoleParams{
		Company: sql.NullInt32{
			Int32: req.GetCompany(),
			Valid: true,
		},
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		CreatedStart: sql.NullTime{
			Time:  time.Unix(req.GetCreatedAtStart().GetSeconds(), 0),
			Valid: req.CreatedAtStart != nil,
		},
		CreatedEnd: sql.NullTime{
			Time:  time.Unix(req.GetCreatedAtEnd().GetSeconds(), 0),
			Valid: req.CreatedAtStart != nil,
		},
		UpdatedStart: sql.NullTime{
			Time:  time.Unix(req.GetUpdatedAtStart().GetSeconds(), 0),
			Valid: req.CreatedAtStart != nil,
		},
		UpdatedEnd: sql.NullTime{
			Time:  time.Unix(req.GetUpdatedAtEnd().GetSeconds(), 0),
			Valid: req.CreatedAtStart != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: true,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get record roles: %v", err))
	}

	var rolesPb []*pb.Role
	for _, item := range roles {
		var note *string
		if item.Note.Valid {
			note = &(item.Note.String)
		}

		data := &pb.Role{
			Id:              item.ID,
			Code:            item.Code,
			Title:           item.Title,
			Note:            note,
			Company:         item.Company.Int32,
			UserCreatedName: item.CreatedName,
			UserUpdatedName: item.UpdatedName,
			CreatedAt:       timestamppb.New(item.CreatedAt),
			UpdatedAt:       timestamppb.New(item.UpdatedAt.Time),
		}
		rolesPb = append(rolesPb, data)
	}

	return &pb.RoleListResponse{
		Code:    200,
		Message: "success",
		Details: rolesPb,
	}, nil
}

func (server *ServerGRPC) AppList(ctx context.Context, req *pb.AppListRequest) (*pb.AppListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	apps, err := server.store.ListApp(ctx, db.ListAppParams{
		Level: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get record app: %v", err))
	}

	var appsPb []*pb.App
	for _, item := range apps {

		subApps, err := server.store.ListApp(ctx, db.ListAppParams{
			Parent: sql.NullString{
				String: item.Code,
				Valid:  true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get record app: %v", err))
		}

		var subAppPb []*pb.App
		for _, app := range subApps {
			subData := pb.App{
				Id:     app.ID,
				Title:  app.Title,
				Code:   app.Code,
				SubApp: nil,
				Level:  app.Level.Int32,
			}

			subAppPb = append(subAppPb, &subData)
		}

		data := pb.App{
			Id:     item.ID,
			Title:  item.Title,
			Code:   item.Code,
			SubApp: subAppPb,
			Level:  item.Level.Int32,
		}

		appsPb = append(appsPb, &data)
	}

	return &pb.AppListResponse{
		Code:    200,
		Message: "success",
		Details: appsPb,
	}, nil
}
