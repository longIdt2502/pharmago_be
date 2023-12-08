CREATE TABLE "variant_media" (
     "id" bigserial PRIMARY KEY,
     "variant" bigserial NOT NULL,
     "media" bigserial NOT NULL
);

CREATE INDEX ON "variant_media" ("variant");

CREATE INDEX ON "variant_media" ("media");

CREATE UNIQUE INDEX ON "variant_media" ("variant", "media");

CREATE INDEX ON "product_media" ("product");

CREATE INDEX ON "product_media" ("media");

CREATE UNIQUE INDEX ON "product_media" ("product", "media");

ALTER TABLE "variant_media" ADD FOREIGN KEY ("variant") REFERENCES "variants" ("id");

ALTER TABLE "variant_media" ADD FOREIGN KEY ("media") REFERENCES "medias" ("id");