CREATE TABLE IF NOT EXISTS "Subscriptions" (
   id serial PRIMARY KEY,
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "planType" VARCHAR(15) NOT NULL,
   "endOfPlanAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "isRecurring" BOOLEAN NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "Subscriptions_userId_idx" ON "Subscriptions" ("userId");
CREATE INDEX "Subscriptions_planType_idx" ON "Subscriptions" ("planType");
CREATE INDEX "Subscriptions_endOfPlanAt_idx" ON "Subscriptions" ("endOfPlanAt");
CREATE INDEX "Subscriptions_isRecurring_idx" ON "Subscriptions" ("isRecurring");

CREATE INDEX "Subscriptions_CreatedAt_idx" ON "Subscriptions" ("CreatedAt");
CREATE INDEX "Subscriptions_UpdatedAt_idx" ON "Subscriptions" ("UpdatedAt");
CREATE INDEX "Subscriptions_DeletedAt_idx" ON "Subscriptions" ("DeletedAt");
