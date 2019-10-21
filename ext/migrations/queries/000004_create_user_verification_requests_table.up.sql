CREATE TABLE IF NOT EXISTS "UserVerificationsRequests" (
   id serial PRIMARY KEY,
   userId INTEGER NOT NULL REFERENCES "Users"(id),
   type VARCHAR(30) NOT NULL,
   "verificationKey" VARCHAR(100) NOT NULL,
   "requestedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "userId_idx" ON "UserVerificationsRequests" ("userId");
CREATE INDEX type_idx ON "UserVerificationsRequests" (type);
CREATE INDEX "verificationKey" ON "UserVerificationsRequests" ("verificationKey");
CREATE INDEX "requestedAt" ON "UserVerificationsRequests" ("requestedAt");

CREATE INDEX "CreatedAt_idx" ON "Credentials" ("CreatedAt");
CREATE INDEX "UpdatedAt_idx" ON "Credentials" ("UpdatedAt");
CREATE INDEX "DeletedAt_idx" ON "Credentials" ("DeletedAt");