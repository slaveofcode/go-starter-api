CREATE TABLE IF NOT EXISTS "ReferralUsers" (
   id serial PRIMARY KEY,
   "referralCodeId" INTEGER NOT NULL REFERENCES "ReferralCodes"(id),
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "ReferralUsers_referralCodeId_idx" ON "ReferralUsers" ("referralCodeId");
CREATE INDEX "ReferralUsers_userId_idx" ON "ReferralUsers" ("userId");

CREATE INDEX "ReferralUsers_CreatedAt_idx" ON "ReferralUsers" ("CreatedAt");
CREATE INDEX "ReferralUsers_UpdatedAt_idx" ON "ReferralUsers" ("UpdatedAt");
CREATE INDEX "ReferralUsers_DeletedAt_idx" ON "ReferralUsers" ("DeletedAt");