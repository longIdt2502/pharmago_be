package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateProductTxParams struct {
	*pb.CreateProductRequest
	Account            Account
	TokenPayload       *token.Payload
	UploadImageVariant func(idVariant int32, image []byte)
}

func (store *Store) CreateProductTx(ctx context.Context, req CreateProductTxParams) (int32, error) {
	var id int32

	opts := &sql.TxOptions{
		Isolation: 1,
		ReadOnly:  false,
	}

	err := store.execTx(ctx, opts, func(q *Queries) error {
		unitData := req.GetUnit()
		variantData := req.GetVariants()
		productData := req.GetProduct()

		unit, err := q.CreateUnit(ctx, CreateUnitParams{
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
			UserCreated: req.Account.ID,
			UserUpdated: req.Account.ID,
		})
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("failed to record unit data: %e", err))
		}

		for _, value := range req.GetUnitChanges() {
			_, err := q.CreateUnitChange(ctx, CreateUnitChangeParams{
				Name:      value.GetName(),
				Value:     value.GetValue(),
				SellPrice: float64(value.GetSellPrice()),
				Unit: sql.NullInt32{
					Int32: unit.ID,
					Valid: true,
				},
				UserCreated: req.TokenPayload.UserID,
				UserUpdated: req.TokenPayload.UserID,
			})
			if err != nil {
				return status.Errorf(codes.Internal, fmt.Sprintf("failed to record unit change: %e", err))
			}
		}

		codeProduct := fmt.Sprintf("PRODUCT-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
		if productData.Code != nil {
			codeProduct = productData.GetCode()
		}
		product, err := q.CreateProduct(ctx, CreateProductParams{
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
			Lieudung: sql.NullString{
				String: productData.GetLieuDung(),
				Valid:  productData.LieuDung != nil,
			},
			Chidinh: sql.NullString{
				String: productData.GetChiDinh(),
				Valid:  productData.ChiDinh != nil,
			},
			Chongchidinh: sql.NullString{
				String: productData.GetChongChiDinh(),
				Valid:  productData.ChongChiDinh != nil,
			},
			Congdung: sql.NullString{
				String: productData.GetCongDung(),
				Valid:  productData.CongDung != nil,
			},
			Tacdungphu: sql.NullString{
				String: productData.GetTacDungPhu(),
				Valid:  productData.TacDungPhu != nil,
			},
			Thantrong: sql.NullString{
				String: productData.GetThanTrong(),
				Valid:  productData.ThanTrong != nil,
			},
			Tuongtac: sql.NullString{
				String: productData.GetTuongTac(),
				Valid:  productData.TuongTac != nil,
			},
			Baoquan: sql.NullString{
				String: productData.GetBaoQuan(),
				Valid:  productData.BaoQuan != nil,
			},
			Donggoi: sql.NullString{
				String: productData.GetDongGoi(),
				Valid:  productData.DongGoi != nil,
			},
			Congtysx: sql.NullInt32{
				Int32: productData.GetCongTySx(),
				Valid: productData.CongTySx != nil,
			},
			Congtydk: sql.NullInt32{
				Int32: productData.GetCongTyDk(),
				Valid: productData.CongTyDk != nil,
			},
			Company:     productData.GetCompany(),
			UserCreated: req.Account.ID,
			UserUpdated: req.Account.ID,
			Phanloai: sql.NullString{
				String: productData.GetPhanLoai(),
				Valid:  productData.DongGoi != nil,
			},
			Dangbaoche: sql.NullString{
				String: productData.GetDangBaoChe(),
				Valid:  productData.DangBaoChe != nil,
			},
			Tieuchuansx: sql.NullString{
				String: productData.GetTieuChuanSx(),
				Valid:  productData.TieuChuanSx != nil,
			},
		})
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("failed to record product data: %e", err))
		}
		id = product.ID

		// for _, item := range req.Product.GetImage() {
		// 	file, _ := utils.NewFileFromImage(item)
		// 	_, err := server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
		// 	if err != nil {
		// 		return status.Errorf(codes.Internal, fmt.Sprintf("failed to upload image to b2: %e", err))
		// 	}

		// 	url, err := server.b2Bucket.FileURL(file.Name)
		// 	if err != nil {
		// 		return status.Errorf(codes.NotFound, "failed to get url by file name")
		// 	}
		// 	println("hay", url)

		// 	media, err := server.store.CreateMedia(ctx, url)
		// 	if err != nil {
		// 		return status.Errorf(codes.Internal, "failed to record media: %e", err)
		// 	}

		// 	_, err = server.store.CreateProductMedia(ctx, db.CreateProductMediaParams{
		// 		Product: product.ID,
		// 		Media:   media.ID,
		// 	})
		// 	if err != nil {
		// 		return status.Errorf(codes.Internal, "failed to record media product: %e", err)
		// 	}
		// }

		for _, value := range req.GetIngredients() {
			_, err := q.CreateIngredient(ctx, CreateIngredientParams{
				Name:   value.GetName(),
				Weight: float64(value.GetWeight()),
				Unit:   value.GetUnit(),
				Product: sql.NullInt32{
					Int32: int32(product.ID),
					Valid: true,
				},
			})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to record ingredient: %e", err)
			}
		}

		for _, value := range variantData {
			codeVariant := fmt.Sprintf("VARIANT-%s-%s-%d", product.Code, utils.RandomString(6), utils.RandomInt(100, 999))
			variant, err := q.CreateVariant(ctx, CreateVariantParams{
				Name: value.GetName(),
				Code: codeVariant,
				Barcode: sql.NullString{
					String: value.GetBarcode(),
					Valid:  value.Barcode != nil,
				},
				Vat: sql.NullFloat64{
					Float64: 0,
					Valid:   true,
				},
				DecisionNumber: sql.NullString{
					String: value.GetDecisionNumber(),
					Valid:  value.DecisionNumber != nil,
				},
				RegisterNumber: sql.NullString{
					String: value.GetRegisterNumber(),
					Valid:  value.RegisterNumber != nil,
				},
				Longevity: sql.NullString{
					String: value.GetLongevity(),
					Valid:  value.Longevity != nil,
				},
				Product:          product.ID,
				UserCreated:      req.Account.ID,
				UserUpdated:      req.Account.ID,
				InitialInventory: value.InitialInventory,
				RealInventory:    value.InitialInventory,
			})
			if err != nil {
				return status.Errorf(codes.Internal, "failed to record variant data: %e", err)
			}

			if value.Image != nil {
				req.UploadImageVariant(variant.ID, value.Image)
				// 	fileVariant, _ := utils.NewFileFromImage(value.Image)
				// 	_, err = server.b2Bucket.UploadFile(fileVariant.Name, fileVariant.Meta, fileVariant.File)
				// 	if err != nil {
				// 		return status.Errorf(codes.Internal, fmt.Sprintf("failed to upload image to b2: %e", err))
				// 	}
				// 	urlVariant, err := server.b2Bucket.FileURL(fileVariant.Name)
				// 	if err != nil {
				// 		return status.Errorf(codes.NotFound, fmt.Sprintf("failed to record media: %e", err))
				// 	}

				// 	media, err := server.store.CreateMedia(ctx, urlVariant)
				// 	if err != nil {
				// 		return status.Errorf(codes.Internal, fmt.Sprintf("failed to record media: %e", err))
				// 	}

				// 	_, err = server.store.CreateVariantMedia(ctx, db.CreateVariantMediaParams{
				// 		Variant: variant.ID,
				// 		Media:   media.ID,
				// 	})
				// 	if err != nil {
				// 		return status.Errorf(codes.Internal, fmt.Sprintf("failed to record media product: %e", err))
				// 	}
			}

			_, err = q.CreateProductPriceList(ctx, CreateProductPriceListParams{
				VariantCode: variant.Code,
				VariantName: variant.Name,
				Unit:        unit.ID,
				PriceImport: unit.ImportPrice,
				PriceSell:   unit.SellPrice,
				UserCreated: req.TokenPayload.UserID,
				UserUpdated: sql.NullInt32{
					Int32: req.TokenPayload.UserID,
					Valid: true,
				},
			})
			if err != nil {
				return status.Errorf(codes.Internal, fmt.Sprintf("failed to record price list: %e", err))
			}
		}

		return err
	})

	return id, err
}
