import { Tag as Tag } from "../tags/tagModel";

export interface Seal {
  id: number;
  slug: string;
  tags: Tag[];
}