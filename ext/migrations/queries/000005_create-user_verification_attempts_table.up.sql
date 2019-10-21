CREATE TABLE IF NOT EXISTS "UserVerificationAttempts" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "userVerificationRequestId" INTEGER NOT NULL REFERENCES "UserVerificationRequests"(id),
   "userAgent" VARCHAR(500) NULL,
   "ipAddr" VARCHAR(15) NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "UserVerificationAttempts_userId_idx" ON "UserVerificationAttempts" ("userId");
CREATE INDEX "UserVerificationAttempts_userVerificationRequestId_idx" ON "UserVerificationAttempts" ("userVerificationRequestId");

CREATE INDEX "UserVerificationAttempts_CreatedAt_idx" ON "UserVerificationAttempts" ("CreatedAt");
CREATE INDEX "UserVerificationAttempts_UpdatedAt_idx" ON "UserVerificationAttempts" ("UpdatedAt");
CREATE INDEX "UserVerificationAttempts_DeletedAt_idx" ON "UserVerificationAttempts" ("DeletedAt");