CREATE TABLE IF NOT EXISTS "ResetCredentials" (
   id serial PRIMARY KEY,
   "credentialId" INTEGER NOT NULL REFERENCES "ResetCredentials"(id),
   "resetToken" VARCHAR(128) NOT NULL,
   "validUntil" TIMESTAMP WITH TIME ZONE NOT NULL,
   "validatedAt" TIMESTAMP WITH TIME ZONE NULL,
   "isExpired" BOOLEAN NOT NULL DEFAULT FALSE,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "ResetCredentials_credentialId_idx" ON "ResetCredentials" ("credentialId");
CREATE INDEX "ResetCredentials_resetToken_idx" ON "ResetCredentials" ("resetToken");
CREATE INDEX "ResetCredentials_validUntil_idx" ON "ResetCredentials" ("validUntil");
CREATE INDEX "ResetCredentials_validatedAt_idx" ON "ResetCredentials" ("validatedAt");
CREATE INDEX "ResetCredentials_isExpired_idx" ON "ResetCredentials" ("isExpired");

CREATE INDEX "ResetCredentials_CreatedAt_idx" ON "ResetCredentials" ("CreatedAt");
CREATE INDEX "ResetCredentials_UpdatedAt_idx" ON "ResetCredentials" ("UpdatedAt");
CREATE INDEX "ResetCredentials_DeletedAt_idx" ON "ResetCredentials" ("DeletedAt");