package gapi

import (
	"context"
	"database/sql"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/rs/zerolog/log"
	"github.com/tealeg/xlsx/v3"
)

func (server *ServerGRPC) ImportProduct(ctx context.Context, req *pb.ImportProductRequest) (*pb.ImportProductResponse, error) {
	excelFileName := "./utils/import/drugbank_full_v2.xlsx"
	xlFile, _ := xlsx.OpenFile(excelFileName)
	for _, sheet := range xlFile.Sheets {
		for i := 1; i < sheet.MaxRow; i++ {
			row, _ := sheet.Row(i)

			congTySx, _ := server.store.GetCompanyPharmaByName(ctx, row.GetCell(16).String())
			congTyDk, _ := server.store.GetCompanyPharmaByName(ctx, row.GetCell(20).String())

			_, err := server.store.CreateProductBank(ctx, db.CreateProductBankParams{
				Name: row.GetCell(2).String(),
				Code: row.GetCell(0).String(),
				TaDuoc: sql.NullString{
					String: row.GetCell(11).String(),
					Valid:  row.GetCell(11) != nil,
				},
				NongDo: sql.NullString{
					String: row.GetCell(10).String(),
					Valid:  row.GetCell(10) != nil,
				},
				LieuDung:     "",
				ChiDinh:      "",
				ChongChiDinh: sql.NullString{},
				CongDung:     "",
				TacDungPhu:   "",
				ThanTrong:    "",
				TuongTac:     sql.NullString{},
				BaoQuan:      "",
				DongGoi:      row.GetCell(13).String(),
				PhanLoai: sql.NullString{
					String: utils.ExtractFirstLetters(row.GetCell(9).String()),
					Valid:  row.GetCell(9) != nil,
				},
				DangBaoChe:  utils.ExtractFirstLetters(row.GetCell(12).String()),
				TieuChuanSx: row.GetCell(14).String(),
				CongTySx:    congTySx.ID,
				CongTyDk:    congTyDk.ID,
			})
			if err != nil {
				log.Printf("=====: %e", err)
			}
		}
	}
	return &pb.ImportProductResponse{
		Code:    200,
		Message: "success",
	}, nil
}
