package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
			Valid: req.Company != nil,
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

		data := pb.Role{
			Id:              item.ID,
			Code:            item.Code,
			Title:           item.Title,
			Note:            note,
			Company:         item.Company.Int32,
			TotalEmployee:   int32(item.Count),
			UserCreatedName: item.CreatedName,
			UserUpdatedName: item.UpdatedName.String,
			CreatedAt:       timestamppb.New(item.CreatedAt),
			UpdatedAt:       timestamppb.New(item.UpdatedAt.Time),
		}
		rolesPb = append(rolesPb, &data)
	}

	return &pb.RoleListResponse{
		Code:    200,
		Message: "success",
		Details: rolesPb,
	}, nil
}

func (server *ServerGRPC) AppList(ctx context.Context, _ *pb.AppListRequest) (*pb.AppListResponse, error) {
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

func (server *ServerGRPC) RoleDetail(ctx context.Context, req *pb.RoleDetailRequest) (*pb.RoleDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	role, err := server.store.RoleDetail(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "role not exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get role: %v", err))
	}

	roleItem, err := server.store.ListRoleItem(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get role item: %v", err))
	}

	// var roleItemPb []*pb.RoleItem
	// for _, item := range roleItem {
	// 	data := &pb.RoleItem{
	// 		Id: item.ID,
	// 		App: &pb.App{
	// 			Id:     item.ID_2,
	// 			Title:  item.Title,
	// 			Code:   item.Code,
	// 			SubApp: nil,
	// 			Level:  item.Level.Int32,
	// 		},
	// 		Value: item.Value.Bool,
	// 	}
	// 	roleItemPb = append(roleItemPb, data)
	// }

	var note *string
	if role.Note.Valid {
		note = &(role.Note.String)
	}

	apps, err := server.AppList(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get app: %v", err))
	}

	var appsRes []*pb.App
	for _, item := range apps.Details {
		for _, y := range item.SubApp {
			a := funk.Find(roleItem, func(x db.ListRoleItemRow) bool {
				return x.App == y.Code
			})
			if a != nil {
				if foundItem, ok := a.(db.ListRoleItemRow); ok {
					y.Value = &(foundItem.Value.Bool)
				} else {
					v := false
					y.Value = &v
				}
			} else {
				v := false
				y.Value = &v
			}
		}
		appsRes = append(appsRes, item)
	}

	return &pb.RoleDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.RoleDetailResponseDetail{
			Role: &pb.Role{
				Id:              role.ID,
				Code:            role.Code,
				Title:           role.Title,
				Note:            note,
				Company:         role.Company.Int32,
				UserCreatedName: role.CreatedName,
				UserUpdatedName: role.UpdatedName.String,
				CreatedAt:       timestamppb.New(role.CreatedAt),
				UpdatedAt:       timestamppb.New(role.UpdatedAt.Time),
			},
			Items: appsRes,
		},
	}, nil
}

func (server *ServerGRPC) RoleUpdate(ctx context.Context, req *pb.RoleUpdateRequest) (*pb.RoleUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	params := db.UpdateRoleParams{
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  req.Code != nil,
		},
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  true,
		},
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  req.Note != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		ID: req.GetId(),
	}

	_, err = server.store.UpdateRole(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "role not exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to updated record: %v", err))
	}

	for _, item := range req.GetItems() {

		_, err = server.store.UpdateRoleItem(ctx, db.UpdateRoleItemParams{
			Value: sql.NullBool{
				Bool:  item.GetChecked(),
				Valid: true,
			},
			Roles: req.GetId(),
			App:   item.GetAppCode(),
		})
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				_, err := server.store.CreateRoleItem(ctx, db.CreateRoleItemParams{
					Roles: req.GetId(),
					App:   item.GetAppCode(),
					Value: sql.NullBool{
						Bool:  item.GetChecked(),
						Valid: true,
					},
				})
				if err != nil {
					return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record role item: %v", err))
				}
			}
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to updated role item: %v", err))
		}
	}

	return &pb.RoleUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest) (*pb.RoleDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	users, _ := server.store.ListAccount(ctx, db.ListAccountParams{
		Role: sql.NullInt32{
			Int32: req.GetId(),
			Valid: true,
		},
	})
	if len(users) != 0 {
		return nil, status.Errorf(codes.Unavailable, "role has account implement")
	}

	_, err = server.store.DeleteRole(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "role not exists")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to delete record: %v", err))
	}

	return &pb.RoleDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
