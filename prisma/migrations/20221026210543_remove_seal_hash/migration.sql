/*
  Warnings:

  - You are about to drop the column `hash` on the `Seal` table. All the data in the column will be lost.

*/
-- RedefineTables
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_Seal" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "slug" TEXT NOT NULL
);
INSERT INTO "new_Seal" ("id", "slug") SELECT "id", "slug" FROM "Seal";
DROP TABLE "Seal";
ALTER TABLE "new_Seal" RENAME TO "Seal";
CREATE UNIQUE INDEX "Seal_slug_key" ON "Seal"("slug");
PRAGMA foreign_key_check;
PRAGMA foreign_keys=ON;
