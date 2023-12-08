-- Xoá ràng buộc trên bảng "variant_media"
ALTER TABLE IF EXISTS "variant_media" DROP CONSTRAINT IF EXISTS "variant_media_variants_fkey";
ALTER TABLE IF EXISTS "variant_media" DROP CONSTRAINT IF EXISTS "variant_media_medias_fkey";

DROP TABLE IF EXISTS variant_media;