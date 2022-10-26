/*
  Warnings:

  - Added the required column `hash` to the `Seal` table without a default value. This is not possible if the table is not empty.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_Seal" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "slug" TEXT NOT NULL,
    "hash" TEXT NOT NULL
);
INSERT INTO "new_Seal" ("id", "slug") SELECT "id", "slug" FROM "Seal";
DROP TABLE "Seal";
ALTER TABLE "new_Seal" RENAME TO "Seal";
CREATE UNIQUE INDEX "Seal_slug_key" ON "Seal"("slug");
CREATE UNIQUE INDEX "Seal_hash_key" ON "Seal"("hash");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
