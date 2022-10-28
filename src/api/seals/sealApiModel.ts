import { Tag as Tag } from "../tags/tagModel";

/**
 * Model describing a Seal
 */
export interface Seal {
  /**
   * ID of the Seal
   */
  id: number;
  /**
   * Slug for the Seal's filename
   */
  slug: string;
  /**
   * Tags associated with this Seal
   */
  tags: Tag[];
}

/**
 * Model containing multiple Seals
 */
export interface ManySeals {
  /**
   * Amount of Seals returned
   */
  count: number;
  /**
   * Offset started from
   */
  offset: number;
  /**
   * The limit that was set on the request
   */
  limit: number;
  /**
   * List of Seal objects
   */
  seals: Seal[];
}

export interface SealCreationParams {
  file: Buffer;
  filename: string;
  tags: string[];
}

export interface SealEdit {
  id: number;
  tags: string[];
}

/**
 * Response when something goes wrong
 */
export interface SealError {
  /**
   * Message detailing the error, if the error is known
   */
  message: string;
  /**
   * Error code for the error. Usually a Prisma error
   */
  error: string;
}

/**
 * Generic response message
 */
export interface SealResponse {
  /**
   * Message providing information about the response
   */
  message: string;
}
