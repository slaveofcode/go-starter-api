CREATE TABLE IF NOT EXISTS "ApiConfigs" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "callLimits" INTEGER NOT NULL,
   "refreshHourCycle" INTEGER NOT NULL,
   "maxApiKeys" INTEGER NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "ApiConfigs_userId_idx" ON "ApiConfigs" ("userId");
CREATE INDEX "ApiConfigs_callLimits_idx" ON "ApiConfigs" ("callLimits");
CREATE INDEX "ApiConfigs_refreshHourCycle_idx" ON "ApiConfigs" ("refreshHourCycle");
CREATE INDEX "ApiConfigs_maxApiKeys_idx" ON "ApiConfigs" ("maxApiKeys");

CREATE INDEX "ApiConfigs_CreatedAt_idx" ON "ApiConfigs" ("CreatedAt");
CREATE INDEX "ApiConfigs_UpdatedAt_idx" ON "ApiConfigs" ("UpdatedAt");
CREATE INDEX "ApiConfigs_DeletedAt_idx" ON "ApiConfigs" ("DeletedAt");
