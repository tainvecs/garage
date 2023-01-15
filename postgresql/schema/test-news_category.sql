CREATE TABLE IF NOT EXISTS "news" (
   "id"           INT           PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   "uuid"         UUID          NOT NULL,
   "link"         VARCHAR(512)  NOT NULL,
   "title"        TEXT          NOT NULL,
   "description"  TEXT,
   "created_at"   DATE,
   "category"     VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS "authors" (
   "id"   INT          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   "name" VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS "news_authors" (
   "id"         INT  PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   "news_id"    INT,
   "authors_id" INT,
   CONSTRAINT "news_authors_news_id_fkey" FOREIGN KEY ("news_id") REFERENCES "news"("id") ON DELETE CASCADE,
   CONSTRAINT "news_authors_authors_id_fkey" FOREIGN KEY ("authors_id") REFERENCES "authors"("id") ON DELETE CASCADE
);
