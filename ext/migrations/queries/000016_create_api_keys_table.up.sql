CREATE TABLE IF NOT EXISTS "ApiKeys" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   token VARCHAR(128) NOT NULL,
   "currentCalls" INTEGER NOT NULL,
   limits INTEGER NOT NULL,
   "lastRefreshedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "refreshHourCycle" INTEGER NOT NULL,
   "isFrozen" BOOLEAN NOT NULL DEFAULT FALSE,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "ApiKeys_userId_idx" ON "ApiKeys" ("userId");
CREATE INDEX "ApiKeys_token_idx" ON "ApiKeys" ("token");
CREATE INDEX "ApiKeys_currentCalls_idx" ON "ApiKeys" ("currentCalls");
CREATE INDEX "ApiKeys_limits_idx" ON "ApiKeys" ("limits");
CREATE INDEX "ApiKeys_lastRefreshedAt_idx" ON "ApiKeys" ("lastRefreshedAt");
CREATE INDEX "ApiKeys_refreshHourCycle_idx" ON "ApiKeys" ("refreshHourCycle");
CREATE INDEX "ApiKeys_isFrozen_idx" ON "ApiKeys" ("isFrozen");

CREATE INDEX "ApiKeys_CreatedAt_idx" ON "ApiKeys" ("CreatedAt");
CREATE INDEX "ApiKeys_UpdatedAt_idx" ON "ApiKeys" ("UpdatedAt");
CREATE INDEX "ApiKeys_DeletedAt_idx" ON "ApiKeys" ("DeletedAt");
