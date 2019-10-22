CREATE TABLE IF NOT EXISTS "ReferralCodes" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   code INTEGER UNIQUE NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "ReferralCodes_userId_idx" ON "ReferralCodes" ("userId");
CREATE INDEX "ReferralCodes_code_idx" ON "ReferralCodes" ("code");

CREATE INDEX "ReferralCodes_CreatedAt_idx" ON "ReferralCodes" ("CreatedAt");
CREATE INDEX "ReferralCodes_UpdatedAt_idx" ON "ReferralCodes" ("UpdatedAt");
CREATE INDEX "ReferralCodes_DeletedAt_idx" ON "ReferralCodes" ("DeletedAt");