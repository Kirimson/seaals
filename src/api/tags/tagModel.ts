import { Seal } from "../seals/sealApiModel";

export interface Tag {
  id: number;
  name: string;
  seals: Seal[];
}
