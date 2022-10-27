import { Seal } from "../seals/sealApiModel";

/**
 * Model describing a Tag
 */
export interface Tag {
  /**
   * ID of the Tag
   */
  id: number;
  /**
   * Name of the Tag
   */
  name: string;
  /**
   * Seals associated with this Tag
   */
  seals: Seal[];
}

/**
 * Model containing multiple Tags
 */
export interface ManyTags {
  /**
   * Amount of Tags returned
   */
  /**
   * Offset started from
   */
  offset: number;
  /**
   * The limit that was set on the request
   */
  limit: number;
  /**
   * List of Tag objects
   */
  tags: Tag[];
}
