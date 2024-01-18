package gapi

import (
	"context"
	"database/sql"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

func (server *ServerGRPC) ListVariant(ctx context.Context, req *pb.ListVariantRequest) (*pb.ListVariantResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	variants, err := server.store.GetVariants(ctx, db.GetVariantsParams{
		Company: req.GetCompany(),
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get variant record: ", err)
	}

	var variantsPb []*pb.Variant
	for _, value := range variants {
		data := mapper.VariantMapper(ctx, server.store, value)
		variantsPb = append(variantsPb, data)
	}

	return &pb.ListVariantResponse{
		Code:    200,
		Message: "success",
		Details: variantsPb,
	}, nil
}

func (server *ServerGRPC) ScanVariant(srv pb.Pharmago_ScanVariantServer) error {
	//_, err := server.authorizeUser(ctx)
	//if err != nil {
	//	return nil, config.UnauthenticatedError(err)
	//}
	ctx := srv.Context()
	for {
		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Debug().Msgf("exit")
			return nil
		}
		if err != nil {
			log.Debug().Msgf("receive error %v", err)
			continue
		}

		code := req.GetCode()
		company := req.GetCompany()
		variants, _ := server.store.GetVariants(ctx, db.GetVariantsParams{
			Company: company,
			Search: sql.NullString{
				String: code,
				Valid:  true,
			},
		})

		var variantPb *pb.Variant
		if len(variants) != 0 {
			variantPb = mapper.VariantMapper(ctx, server.store, variants[0])
		}

		resp := pb.VariantScanResponse{
			Code:    200,
			Message: "success",
			Details: variantPb,
		}
		if err := srv.Send(&resp); err != nil {
			log.Debug().Msgf("send new response failed: ", err.Error())
		}
		log.Info().Msg("send new response success")
	}
}
