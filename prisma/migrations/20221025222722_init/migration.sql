-- CreateTable
CREATE TABLE "Seal" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT
);

-- CreateTable
CREATE TABLE "Tag" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "name" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "_SealToTag" (
    "A" INTEGER NOT NULL,
    "B" INTEGER NOT NULL,
    CONSTRAINT "_SealToTag_A_fkey" FOREIGN KEY ("A") REFERENCES "Seal" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "_SealToTag_B_fkey" FOREIGN KEY ("B") REFERENCES "Tag" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

-- CreateIndex
CREATE UNIQUE INDEX "Tag_name_key" ON "Tag"("name");

-- CreateIndex
CREATE UNIQUE INDEX "_SealToTag_AB_unique" ON "_SealToTag"("A", "B");

-- CreateIndex
CREATE INDEX "_SealToTag_B_index" ON "_SealToTag"("B");
