package gapi

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/tealeg/xlsx/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *ServerGRPC) ImportCompany(ctx context.Context, req *pb.ImportCompanyRequest) (*pb.ImportCompanyResponse, error) {
	excelFileName := "./utils/import/drugbank_full_v2.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	for _, sheet := range xlFile.Sheets {
		for i := 1; i < sheet.MaxRow; i++ {
			row, _ := sheet.Row(i)

			_, _ = server.store.CreateCompanyPharma(ctx, db.CreateCompanyPharmaParams{
				Name: row.GetCell(16).String(),
				Code: sql.NullString{
					String: row.GetCell(17).String(),
					Valid:  true,
				},
				Country: sql.NullString{
					String: row.GetCell(18).String(),
					Valid:  true,
				},
				Address: sql.NullString{
					String: row.GetCell(19).String(),
					Valid:  true,
				},
				CompanyPharmaType: "PRODUCTION",
			})

			_, _ = server.store.CreateCompanyPharma(ctx, db.CreateCompanyPharmaParams{
				Name: row.GetCell(20).String(),
				Code: sql.NullString{},
				Country: sql.NullString{
					String: row.GetCell(21).String(),
					Valid:  true,
				},
				Address: sql.NullString{
					String: row.GetCell(22).String(),
					Valid:  true,
				},
				CompanyPharmaType: "REGISTERED",
			})
		}
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to read file: %e", err))
	}

	return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to read file: %e", err))
}
