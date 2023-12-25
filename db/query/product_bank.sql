-- name: CreateProductBank :one
INSERT INTO products_bank (
    name, code, ta_duoc, nong_do, lieu_dung, chi_dinh, chong_chi_dinh, cong_dung, tac_dung_phu, than_trong, tuong_tac, bao_quan,
    dong_goi, phan_loai, dang_bao_che, tieu_chuan_sx, cong_ty_sx, cong_ty_dk
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
) RETURNING *;