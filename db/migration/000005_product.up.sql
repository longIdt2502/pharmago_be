ALTER TABLE products ADD COLUMN taDuoc varchar(255);
ALTER TABLE products ADD COLUMN nongDo varchar(255);
ALTER TABLE products ADD COLUMN lieuDung varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN chiDinh varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN chongChiDinh varchar(255);
ALTER TABLE products ADD COLUMN congDung varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN tacDungPhu varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN thanTrong varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN tuongTac varchar(255);
ALTER TABLE products ADD COLUMN baoQuan varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN dongGoi varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN noiSx varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN congTySx varchar(255) NOT NULL;
ALTER TABLE products ADD COLUMN congTyDk varchar(255) NOT NULL;

-- ALTER TABLE variants ALTER COLUMN decision_number TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN decision_number SET DATA TYPE varchar(255),
ALTER COLUMN decision_number SET NOT NULL,
ADD CONSTRAINT decision_number_unique UNIQUE (decision_number);

-- ALTER TABLE variants ALTER COLUMN register_number TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN register_number SET DATA TYPE varchar(255),
ALTER COLUMN register_number SET NOT NULL,
ADD CONSTRAINT register_number_unique UNIQUE (register_number);

ALTER TABLE variants ADD COLUMN longevity varchar(255) NOT NULL;
ALTER TABLE variants DROP COLUMN IF EXISTS "discount";
-- ALTER TABLE variants ALTER COLUMN code TYPE varchar(255) NOT NULL UNIQUE;
ALTER TABLE variants ALTER COLUMN code SET DATA TYPE varchar(255),
ALTER COLUMN code SET NOT NULL,
ADD CONSTRAINT code_unique UNIQUE (code);

