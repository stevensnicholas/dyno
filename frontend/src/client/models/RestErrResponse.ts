/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type RestErrResponse = {
  /**
   * Application-specific error code.
   */
  code: number;
  /**
   * Application context.
   */
  context: Record<string, any>;
  /**
   * Error message.
   */
  error: string;
  /**
   * Status text.
   */
  status: string;
};

