CREATE TABLE IF NOT EXISTS "TeamMemberInvitations" (
   id serial PRIMARY KEY,
   "teamId" INTEGER NOT NULL REFERENCES "Teams"(id),
   email VARCHAR(60) NOT NULL,
   "roleId" INTEGER NOT NULL REFERENCES "Roles"(id),
   "invitationKey" VARCHAR(32) NOT NULL,
   "CreatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "UpdatedAt" TIMESTAMP WITH TIME ZONE NOT NULL,
   "DeletedAt" TIMESTAMP WITH TIME ZONE NULL
);

CREATE INDEX "TeamMemberInvitations_teamId_idx" ON "TeamMemberInvitations" ("teamId");
CREATE INDEX "TeamMemberInvitations_email_idx" ON "TeamMemberInvitations" ("email");
CREATE INDEX "TeamMemberInvitations_roleId_idx" ON "TeamMemberInvitations" ("roleId");
CREATE INDEX "TeamMemberInvitations_invitationKey_idx" ON "TeamMemberInvitations" ("invitationKey");

CREATE INDEX "TeamMemberInvitations_CreatedAt_idx" ON "TeamMemberInvitations" ("CreatedAt");
CREATE INDEX "TeamMemberInvitations_UpdatedAt_idx" ON "TeamMemberInvitations" ("UpdatedAt");
CREATE INDEX "TeamMemberInvitations_DeletedAt_idx" ON "TeamMemberInvitations" ("DeletedAt");
