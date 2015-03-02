CREATE TABLE "quotes" (
  "id" SERIAL PRIMARY KEY,
  "key" VARCHAR(64) NOT NULL UNIQUE,
  "url" TEXT NOT NULL,
  "text" TEXT  NOT NULL,
  "created_at" TIMESTAMP NOT NULL
);

