/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { EndpointsPostEchoInput } from '../models/EndpointsPostEchoInput';
import type { EndpointsPostEchoOutput } from '../models/EndpointsPostEchoOutput';

import type { CancelablePromise } from '../core/CancelablePromise';
import type { BaseHttpRequest } from '../core/BaseHttpRequest';

export class DefaultService {

  constructor(public readonly httpRequest: BaseHttpRequest) {}

  /**
   * Echo
   * Returns the same string as provided
   * @param requestBody
   * @returns EndpointsPostEchoOutput OK
   * @throws ApiError
   */
  public cmdEndpointsPostEcho(
    requestBody?: EndpointsPostEchoInput,
  ): CancelablePromise<EndpointsPostEchoOutput> {
    return this.httpRequest.request({
      method: 'POST',
      url: '/echo',
      body: requestBody,
      mediaType: 'application/json',
      errors: {
        400: `Bad Request`,
      },
    });
  }

}
