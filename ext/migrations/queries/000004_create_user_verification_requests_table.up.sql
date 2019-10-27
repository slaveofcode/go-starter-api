CREATE TABLE IF NOT EXISTS "UserVerificationRequests" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   type VARCHAR(30) NOT NULL,
   "verificationKey" VARCHAR(100) NOT NULL,
   "verifiedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "UserVerificationRequests_userId_idx" ON "UserVerificationRequests" ("userId");
CREATE INDEX "UserVerificationRequests_type_idx" ON "UserVerificationRequests" (type);
CREATE INDEX "UserVerificationRequests_verificationKey_idx" ON "UserVerificationRequests" ("verificationKey");
CREATE INDEX "UserVerificationRequests_verifiedAt_idx" ON "UserVerificationRequests" ("verifiedAt");

CREATE INDEX "UserVerificationRequests_CreatedAt_idx" ON "UserVerificationRequests" ("CreatedAt");
CREATE INDEX "UserVerificationRequests_UpdatedAt_idx" ON "UserVerificationRequests" ("UpdatedAt");
CREATE INDEX "UserVerificationRequests_DeletedAt_idx" ON "UserVerificationRequests" ("DeletedAt");