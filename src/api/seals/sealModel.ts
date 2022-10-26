import { Tag as Tag } from "../tags/tagModel";

export interface Seal {
  id: number;
  slug: string;
  tags: Tag[];
}

export interface Seals {
  count: number;
  offset: number;
  limit: number;
  seals: Seal[];
}
