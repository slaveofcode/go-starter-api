CREATE TABLE IF NOT EXISTS "Roles" (
   id serial PRIMARY KEY,
   name VARCHAR(30) UNIQUE NOT NULL,
   scopes VARCHAR[] NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "Roles_name_idx" ON "Roles" ("name");
CREATE INDEX "Roles_scopes_idx" ON "Roles" ("scopes");

CREATE INDEX "Roles_CreatedAt_idx" ON "Roles" ("CreatedAt");
CREATE INDEX "Roles_UpdatedAt_idx" ON "Roles" ("UpdatedAt");
CREATE INDEX "Roles_DeletedAt_idx" ON "Roles" ("DeletedAt");