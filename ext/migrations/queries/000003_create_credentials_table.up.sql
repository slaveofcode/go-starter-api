CREATE TABLE IF NOT EXISTS "Credentials" (
   id serial PRIMARY KEY,
   userId INTEGER NOT NULL REFERENCES "Users"(id),
   email VARCHAR(30) NOT NULL,
   password VARCHAR(100) NOT NULL,
   "isCurrentlyUsed" BOOLEAN NOT NULL DEFAULT false,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "userId_idx" ON "Credentials" ("userId");
CREATE INDEX email_idx ON "Credentials" (email);
CREATE INDEX password_idx ON "Credentials" (password);
CREATE INDEX "isCurrentlyUsed" ON "Credentials" ("isCurrentlyUsed");

CREATE INDEX "CreatedAt_idx" ON "Credentials" ("CreatedAt");
CREATE INDEX "UpdatedAt_idx" ON "Credentials" ("UpdatedAt");
CREATE INDEX "DeletedAt_idx" ON "Credentials" ("DeletedAt");