CREATE TABLE IF NOT EXISTS "TeamMembers" (
   id serial PRIMARY KEY,
   "teamId" INTEGER NOT NULL REFERENCES "Teams"(id),
   "userId" INTEGER NOT NULL REFERENCES "Users"(id),
   "roleId" INTEGER NOT NULL REFERENCES "Roles"(id),
   "isFrozen" BOOLEAN NOT NULL DEFAULT FALSE,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "TeamMembers_teamId_idx" ON "TeamMembers" ("teamId");
CREATE INDEX "TeamMembers_userId_idx" ON "TeamMembers" ("userId");
CREATE INDEX "TeamMembers_roleId_idx" ON "TeamMembers" ("roleId");
CREATE INDEX "TeamMembers_isFrozen_idx" ON "TeamMembers" ("isFrozen");

CREATE INDEX "TeamMembers_CreatedAt_idx" ON "TeamMembers" ("CreatedAt");
CREATE INDEX "TeamMembers_UpdatedAt_idx" ON "TeamMembers" ("UpdatedAt");
CREATE INDEX "TeamMembers_DeletedAt_idx" ON "TeamMembers" ("DeletedAt");