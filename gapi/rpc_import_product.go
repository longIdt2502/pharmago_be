package gapi

import (
	"context"
	"database/sql"
	"log"
	"math"
	"strconv"

	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/tealeg/xlsx/v3"
)

func (server *ServerGRPC) ImportProduct(ctx context.Context, req *pb.ImportProductRequest) (*pb.ImportProductResponse, error) {
	account, err := server.store.GetAccountByPhone(ctx, "0359012720")
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.ImportProductResponse{
			Code:    int32(errApp.StatusCode),
			Message: "tài khoản không tồn tại",
		}, nil
	}

	company, err := server.store.GetCompanyByPhone(ctx, sql.NullString{
		String: "0359012720",
		Valid:  true,
	})
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.ImportProductResponse{
			Code:    int32(errApp.StatusCode),
			Message: "công ty không tồn tại",
		}, nil
	}

	excelFileName := "./utils/import/vanquan_drugs.xlsx"
	xlFile, _ := xlsx.OpenFile(excelFileName)
	for _, sheet := range xlFile.Sheets {
		for i := 1; i < sheet.MaxRow; i++ {
			row, _ := sheet.Row(i)
			if row.GetCell(21).String() != "" {
				product, err := server.store.GetProductsByCode(ctx, row.GetCell(21).String())
				if err != nil {
					continue
					// errApp := common.ErrDB(err)
					// return &pb.ImportProductResponse{
					// 	Code:    int32(errApp.StatusCode),
					// 	Message: "lỗi lấy dữ liệu sản phẩm",
					// }, nil
				}
				variant, err := server.store.GetVariantByProduct(ctx, product.ID)
				if err != nil {
					continue
					// errApp := common.ErrDB(err)
					// return &pb.ImportProductResponse{
					// 	Code:    int32(errApp.StatusCode),
					// 	Message: "lỗi lấy dữ liệu variant",
					// }, nil
				}
				sellPrice, _ := strconv.ParseFloat(row.GetCell(13).String(), 64)
				value, _ := strconv.Atoi(row.GetCell(22).String())
				_, err = server.store.CreateUnitChange(ctx, db.CreateUnitChangeParams{
					Name:      row.GetCell(20).String(),
					Value:     int32(value),
					SellPrice: sellPrice,
					Unit: sql.NullInt32{
						Int32: product.Unit,
						Valid: true,
					},
					UserCreated: account.ID,
					UserUpdated: account.ID,
				})
				if err != nil {
					errApp := common.ErrDB(err)
					return &pb.ImportProductResponse{
						Code:    int32(errApp.StatusCode),
						Message: "lỗi lấy tạo đơn vị quy đổi",
					}, nil
				}

				inventoryFloat64, _ := strconv.ParseFloat(row.GetCell(37).String(), 64)
				inventoryChange := inventoryFloat64 * float64(value)
				_, err = server.store.UpdateVariant(ctx, db.UpdateVariantParams{
					ID: variant.ID,
					RealInventory: sql.NullInt32{
						Int32: int32(math.Round(inventoryChange)) + variant.RealInventory,
					},
				})
				if err != nil {
					errApp := common.ErrDB(err)
					return &pb.ImportProductResponse{
						Code:    int32(errApp.StatusCode),
						Message: "lỗi cập nhật tồn kho",
					}, nil
				}

			} else {
				sellPrice, _ := strconv.ParseFloat(row.GetCell(13).String(), 64)
				importPrice, _ := strconv.ParseFloat(row.GetCell(14).String(), 64)
				unit, err := server.store.CreateUnit(ctx, db.CreateUnitParams{
					Name:        row.GetCell(20).String(),
					SellPrice:   sellPrice,
					ImportPrice: importPrice,
					Weight:      sql.NullFloat64{},
					WeightUnit:  sql.NullString{},
					UserCreated: account.ID,
					UserUpdated: account.ID,
				})
				if err != nil {
					errApp := common.ErrDB(err)
					return &pb.ImportProductResponse{
						Code:    int32(errApp.StatusCode),
						Message: "lỗi đơn vị sản phẩm",
					}, nil
				}

				product, err := server.store.CreateProduct(ctx, db.CreateProductParams{
					Name:            row.GetCell(3).String(),
					Code:            row.GetCell(2).String(),
					ProductCategory: sql.NullInt32{},
					Type:            sql.NullInt32{},
					Brand:           sql.NullInt32{},
					Unit:            unit.ID,
					Taduoc: sql.NullString{
						String: row.GetCell(7).String(),
						Valid:  row.GetCell(7) != nil,
					},
					Nongdo: sql.NullString{
						String: row.GetCell(8).String(),
						Valid:  row.GetCell(8) != nil,
					},
					Lieudung: sql.NullString{},
					Chidinh: sql.NullString{
						String: row.GetCell(12).String(),
						Valid:  row.GetCell(12) != nil,
					},
					Chongchidinh: sql.NullString{},
					Congdung:     sql.NullString{},
					Tacdungphu:   sql.NullString{},
					Thantrong:    sql.NullString{},
					Tuongtac:     sql.NullString{},
					Baoquan:      sql.NullString{},
					Donggoi: sql.NullString{
						String: row.GetCell(11).String(),
						Valid:  row.GetCell(11) != nil,
					},
					Congtysx:    sql.NullInt32{},
					Congtydk:    sql.NullInt32{},
					Company:     company.ID,
					UserCreated: account.ID,
					UserUpdated: account.ID,
					Phanloai:    sql.NullString{},
					Dangbaoche:  sql.NullString{},
					Tieuchuansx: sql.NullString{},
				})
				if err != nil {
					errApp := common.ErrDB(err)
					log.Printf("===== product: %s", errApp.Log)
				}

				inventoryFloat64, _ := strconv.ParseFloat(row.GetCell(37).String(), 64)
				inventoryInt := int(math.Round(inventoryFloat64))
				_, err = server.store.CreateVariant(ctx, db.CreateVariantParams{
					Name:           row.GetCell(3).String(),
					Code:           row.GetCell(2).String(),
					Barcode:        sql.NullString{},
					Vat:            sql.NullFloat64{},
					DecisionNumber: sql.NullString{},
					RegisterNumber: sql.NullString{
						String: row.GetCell(5).String(),
						Valid:  row.GetCell(5) != nil,
					},
					Longevity:        sql.NullString{},
					Product:          product.ID,
					UserCreated:      account.ID,
					UserUpdated:      account.ID,
					InitialInventory: int32(inventoryInt),
					RealInventory:    int32(inventoryInt),
				})
				if err != nil {
					errApp := common.ErrDB(err)
					log.Printf("===== variant: %s", errApp.Log)
				}
			}

		}
	}
	return &pb.ImportProductResponse{
		Code:    200,
		Message: "success",
	}, nil
}
