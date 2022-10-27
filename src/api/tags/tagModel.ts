import { Seal } from "../seals/sealApiModel";

export interface Tag {
  id: number;
  name: string;
  seals: Seal[];
}

export interface Tags {
  count: number;
  offset: number;
  limit: number;
  tags: Tag[];
}
