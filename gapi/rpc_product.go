package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/woker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}
	account, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	payloadTx := db.CreateProductTxParams{
		CreateProductRequest: req,
		Account:              account,
		TokenPayload:         tokenPayload,
		UploadImageVariant: func(idVariant int32, image []byte) {
			payload := woker.PayloadUploadImageVariant{
				Image: image,
				Id:    idVariant,
			}

			opts := []asynq.Option{
				asynq.MaxRetry(0),
				asynq.ProcessIn(1 * time.Second),
				asynq.Queue(woker.QueueCritical),
			}

			_ = server.taskDistributor.DistributorUploadImageVariant(ctx, &payload, opts...)
		},
	}

	productId, err := server.store.CreateProductTx(ctx, payloadTx)
	if err != nil {
		return nil, err
	}

	for _, item := range req.Product.GetImage() {
		payload := woker.PayloadUploadImageProduct{
			Image: item,
			Id:    productId,
		}

		opts := []asynq.Option{
			asynq.MaxRetry(0),
			asynq.ProcessIn(1 * time.Second),
			asynq.Queue(woker.QueueCritical),
		}

		_ = server.taskDistributor.DistributorUploadImageProduct(ctx, &payload, opts...)
	}

	return &pb.CreateProductResponse{
		Message: "success",
		Code:    200,
		Details: productId,
	}, nil
}

func (server *ServerGRPC) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	products, err := server.store.GetProducts(ctx, db.GetProductsParams{
		Company:         sql.NullInt32{Int32: req.GetCompany(), Valid: req.Company != nil},
		Search:          sql.NullString{String: req.GetSearch(), Valid: req.Search != nil},
		Brand:           sql.NullInt32{},
		ProductCategory: sql.NullInt32{},
		Active:          sql.NullBool{Bool: req.GetActive(), Valid: req.Active != nil},
		Page:            sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:           sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get product: %e", err))
	}

	var productsPb []*pb.Product
	for _, value := range products {
		productsPb = append(productsPb, mapper.ProductMapper(ctx, server.store, value))
	}

	counts, _ := server.store.CountProduct(ctx, req.GetCompany())

	var countsPb []*pb.SimpleData
	for _, item := range counts {
		value := int32(item.Total)
		switch item.Active {
		case true:
			countsPb = append(countsPb, &pb.SimpleData{
				Name:  "Đang bán",
				Code:  "TRUE",
				Value: &value,
			})
		case false:
			countsPb = append(countsPb, &pb.SimpleData{
				Name:  "Đang bán",
				Code:  "FALSE",
				Value: &value,
			})
		}

	}

	return &pb.ListProductResponse{
		Code:    200,
		Message: "success",
		Details: productsPb,
		Counts:  countsPb,
	}, nil
}

func (server *ServerGRPC) DetailProduct(ctx context.Context, req *pb.DetailProductRequest) (*pb.DetailProductResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	productDb, err := server.store.DetailProduct(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "product not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get product")
	}

	imagesDb, err := server.store.GetProductMedia(ctx, productDb.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get image")
	}

	var images []string
	for _, item := range imagesDb {
		images = append(images, item.MediaUrl)
	}

	productPb := &pb.Product{
		Id:           productDb.ID,
		Name:         productDb.Name,
		Code:         productDb.Code,
		Category:     &productDb.ProductCategory.Int32,
		Type:         &productDb.Type.Int32,
		TaDuoc:       &productDb.TaDuoc.String,
		NongDo:       &productDb.NongDo.String,
		LieuDung:     productDb.LieuDung.String,
		ChiDinh:      productDb.ChiDinh.String,
		ChongChiDinh: &productDb.ChongChiDinh.String,
		CongDung:     productDb.CongDung.String,
		TacDungPhu:   productDb.TacDungPhu.String,
		ThanTrong:    productDb.ThanTrong.String,
		TuongTac:     &productDb.TuongTac.String,
		BaoQuan:      productDb.BaoQuan.String,
		DongGoi:      productDb.DongGoi.String,
		NoiSx:        productDb.Name_9.String,
		CongTyDk:     productDb.CongTyDk.Int32,
		CongTySx:     productDb.CongTySx.Int32,
		Image:        images,
	}

	var unitsPb []*pb.Unit
	unitsPb = append(unitsPb, &pb.Unit{
		Id:          productDb.ID_4.Int32,
		Name:        productDb.Name_4.String,
		Value:       1,
		SellPrice:   float32(productDb.SellPrice.Float64),
		ImportPrice: float32(productDb.ImportPrice.Float64),
		Default:     true,
		Weight:      float32(productDb.Weight.Float64),
		WeightUnit:  productDb.WeightUnit.String,
	})

	units, err := server.store.GetListUnitChange(ctx, productDb.Unit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get unit")
	}

	for _, item := range units {
		unitsPb = append(unitsPb, &pb.Unit{
			Id:          item.ID,
			Name:        item.Name,
			Value:       int32(item.Value),
			SellPrice:   float32(item.SellPrice),
			ImportPrice: 0,
			Default:     false,
		})
	}

	ingredients, err := server.store.ListIngredient(ctx, productDb.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get ingredient")
	}

	var ingredientsPb []*pb.Ingredient
	for _, item := range ingredients {
		ingredientsPb = append(ingredientsPb, &pb.Ingredient{
			Name:   item.Name,
			Weight: float32(item.Weight),
			Unit:   item.Unit,
		})
	}

	variants, err := server.store.GetVariants(ctx, db.GetVariantsParams{
		Product: productDb.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get variants: %e", err)
	}

	var variantsPb []*pb.Variant
	for _, item := range variants {
		variantsPb = append(variantsPb, mapper.VariantMapper(ctx, server.store, item))
	}

	return &pb.DetailProductResponse{
		Code:    200,
		Message: "success",
		Details: &pb.ProductDetail{
			Product:     productPb,
			Variants:    variantsPb,
			Unit:        unitsPb,
			Ingredients: ingredientsPb,
		},
	}, nil
}
