CREATE TABLE IF NOT EXISTS "Teams" (
   id serial PRIMARY KEY,
   name VARCHAR(30) NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "Teams_name_idx" ON "Teams" ("name");

CREATE INDEX "Teams_CreatedAt_idx" ON "Teams" ("CreatedAt");
CREATE INDEX "Teams_UpdatedAt_idx" ON "Teams" ("UpdatedAt");
CREATE INDEX "Teams_DeletedAt_idx" ON "Teams" ("DeletedAt");