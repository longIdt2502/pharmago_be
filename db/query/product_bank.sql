-- name: CreateProductBank :one
INSERT INTO products_bank (
    name, code, "taDuoc", "nongDo", "lieuDung", "chiDinh", "chongChiDinh", "congDung", "tacDungPhu", "thanTrong", "tuongTac", "baoQuan",
    "dongGoi", "phanLoai", "dangBaoche", "tieuChuanSx", "congTySx", "congTyDk"
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
) RETURNING *;