---------------- user_detail table ----------------
CREATE TABLE IF NOT EXISTS user_detail(
  id                           SERIAL PRIMARY KEY NOT NULL,
  user_name                    VARCHAR(255),
  password                     VARCHAR(255),
  location                     VARCHAR(255),
  age                          INTEGER,
  created_at                   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at                   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()    
);