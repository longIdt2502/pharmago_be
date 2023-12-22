package gapi

import (
	"context"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/tealeg/xlsx/v3"
)

func (server *ServerGRPC) ImportProductMasterData(ctx context.Context, req *pb.ImportProductMasterDataRequest) (*pb.ImportProductMasterDataResponse, error) {
	excelFileName := "./utils/import/drugbank_full_v2.xlsx"
	xlFile, _ := xlsx.OpenFile(excelFileName)
	for _, sheet := range xlFile.Sheets {
		for i := 1; i < sheet.MaxRow; i++ {
			row, _ := sheet.Row(i)

			classifyName := row.GetCell(9)
			if classifyName != nil {
				classifyCode := utils.ExtractFirstLetters(classifyName.String())

				_, _ = server.store.CreateClassify(ctx, db.CreateClassifyParams{
					Code: classifyCode,
					Name: classifyName.String(),
				})
			}

			productionStandardName := row.GetCell(14)
			if productionStandardName != nil {

				_, _ = server.store.CreateProductionStandard(ctx, db.CreateProductionStandardParams{
					Code: productionStandardName.String(),
					Name: productionStandardName.String(),
				})
			}

			preparationTypeName := row.GetCell(12)
			if preparationTypeName != nil {
				preparationTypeCode := utils.ExtractFirstLetters(preparationTypeName.String())

				_, _ = server.store.CreatePreparationType(ctx, db.CreatePreparationTypeParams{
					Code: preparationTypeCode,
					Name: preparationTypeName.String(),
				})
			}

		}
	}

	return &pb.ImportProductMasterDataResponse{
		Code:    200,
		Message: "success",
	}, nil
}
