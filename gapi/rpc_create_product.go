package gapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	unitData := req.GetUnit()
	variantData := req.GetVariants()
	productData := req.GetProduct()

	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}
	account, err := server.store.GetAccountByUseName(ctx, tokenPayload.Username)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	unit, err := server.store.CreateUnit(ctx, db.CreateUnitParams{
		Name:        unitData.Name,
		SellPrice:   unitData.SellPrice,
		ImportPrice: unitData.ImportPrice,
		Weight: sql.NullFloat64{
			Float64: unitData.GetWeight(),
			Valid:   unitData.Weight != nil,
		},
		WeightUnit: sql.NullString{
			String: unitData.GetWeightUnit(),
			Valid:  unitData.WeightUnit != nil,
		},
		UserCreated: account.ID,
		UserUpdated: account.ID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record unit data: %e", err))
	}

	for _, value := range req.GetUnitChanges() {
		_, err := server.store.CreateUnitChange(ctx, db.CreateUnitChangeParams{
			Name:      value.GetName(),
			Value:     value.GetValue(),
			SellPrice: float64(value.GetSellPrice()),
			Unit: sql.NullInt32{
				Int32: unit.ID,
				Valid: true,
			},
			UserCreated: tokenPayload.UserID,
			UserUpdated: tokenPayload.UserID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record unit change: %e", err))
		}
	}

	codeProduct := fmt.Sprintf("PRODUCT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if productData.Code != nil {
		codeProduct = productData.GetCode()
	}
	product, err := server.store.CreateProduct(ctx, db.CreateProductParams{
		Name: productData.GetName(),
		Code: codeProduct,
		ProductCategory: sql.NullInt32{
			Int32: productData.GetCategory(),
			Valid: productData.Category != nil,
		},
		Type: sql.NullInt32{
			Int32: productData.GetType(),
			Valid: productData.Type != nil,
		},
		Brand: sql.NullInt32{
			Int32: productData.GetBrand(),
			Valid: productData.Type != nil,
		},
		Unit: int32(unit.ID),
		Taduoc: sql.NullString{
			String: productData.GetTaDuoc(),
			Valid:  productData.TaDuoc != nil,
		},
		Nongdo: sql.NullString{
			String: productData.GetNongDo(),
			Valid:  productData.NongDo != nil,
		},
		Lieudung: productData.GetLieuDung(),
		Chidinh:  productData.GetChiDinh(),
		Chongchidinh: sql.NullString{
			String: productData.GetChongChiDinh(),
			Valid:  productData.ChongChiDinh != nil,
		},
		Congdung:   productData.GetCongDung(),
		Tacdungphu: productData.GetTacDungPhu(),
		Thantrong:  productData.GetThanTrong(),
		Tuongtac: sql.NullString{
			String: productData.GetTuongTac(),
			Valid:  productData.TuongTac != nil,
		},
		Baoquan:     productData.GetBaoQuan(),
		Donggoi:     productData.GetDongGoi(),
		Congtysx:    productData.GetCongTySx(),
		Congtydk:    productData.GetCongTyDk(),
		Company:     productData.GetCompany(),
		UserCreated: account.ID,
		UserUpdated: account.ID,
		Phanloai:    productData.GetPhanLoai(),
		Dangbaoche:  productData.GetDangBaoChe(),
		Tieuchuansx: productData.GetTieuChuanSx(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record product data: %e", err))
	}

	for _, item := range req.Product.GetImage() {
		file, _ := utils.NewFileFromImage(item)
		_, err := server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to upload image to b2: %e", err))
		}

		url, err := server.b2Bucket.FileURL(file.Name)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "failed to get url by file name")
		}
		println("hay", url)

		media, err := server.store.CreateMedia(ctx, url)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record media: %e", err)
		}

		_, err = server.store.CreateProductMedia(ctx, db.CreateProductMediaParams{
			Product: product.ID,
			Media:   media.ID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record media product: %e", err)
		}
	}

	for _, value := range req.GetIngredients() {
		_, err := server.store.CreateIngredient(ctx, db.CreateIngredientParams{
			Name:   value.GetName(),
			Weight: float64(value.GetWeight()),
			Unit:   value.GetUnit(),
			Product: sql.NullInt32{
				Int32: int32(product.ID),
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record ingredient: %e", err)
		}
	}

	for _, value := range variantData {
		codeVariant := fmt.Sprintf("VARIANT-%s-%s-%d", product.Code, utils.RandomString(6), utils.RandomInt(100, 999))
		variant, err := server.store.CreateVariant(ctx, db.CreateVariantParams{
			Name:    value.GetName(),
			Code:    codeVariant,
			Barcode: value.GetCode(),
			Vat: sql.NullFloat64{
				Float64: 0,
				Valid:   true,
			},
			DecisionNumber: value.GetDecisionNumber(),
			RegisterNumber: value.GetRegisterNumber(),
			Longevity:      value.GetLongevity(),
			Product:        product.ID,
			UserCreated:    account.ID,
			UserUpdated:    account.ID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record variant data: %e", err)
		}

		if value.Image != nil {
			fileVariant, _ := utils.NewFileFromImage(value.Image)
			_, err = server.b2Bucket.UploadFile(fileVariant.Name, fileVariant.Meta, fileVariant.File)
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to upload image to b2: %e", err))
			}
			urlVariant, err := server.b2Bucket.FileURL(fileVariant.Name)
			if err != nil {
				return nil, status.Errorf(codes.NotFound, fmt.Sprintf("failed to record media: %e", err))
			}

			media, err := server.store.CreateMedia(ctx, urlVariant)
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record media: %e", err))
			}

			_, err = server.store.CreateVariantMedia(ctx, db.CreateVariantMediaParams{
				Variant: variant.ID,
				Media:   media.ID,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record media product: %e", err))
			}
		}

		_, err = server.store.CreateProductPriceList(ctx, db.CreateProductPriceListParams{
			VariantCode: variant.Code,
			VariantName: variant.Name,
			Unit:        unit.ID,
			PriceImport: unit.ImportPrice,
			PriceSell:   unit.SellPrice,
			UserCreated: tokenPayload.UserID,
			UserUpdated: sql.NullInt32{
				Int32: tokenPayload.UserID,
				Valid: true,
			},
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to record price list: %e", err))
		}
	}

	return &pb.CreateProductResponse{
		Message: fmt.Sprintf("created %s", product.Name),
		Code:    200,
		Details: product.ID,
	}, nil
}
