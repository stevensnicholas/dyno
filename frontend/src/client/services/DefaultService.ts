/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { EndpointsAuthOutput } from '../models/EndpointsAuthOutput';
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
  public endpointsPostEcho(
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

  /**
   * Authentication
   * Return token
   * @param code
   * @returns EndpointsAuthOutput OK
   * @throws ApiError
   */
  public endpointsAuthentication(
    code?: string,
  ): CancelablePromise<EndpointsAuthOutput> {
    return this.httpRequest.request({
      method: 'GET',
      url: '/login',
      query: {
        'code': code,
      },
      errors: {
        400: `Bad Request`,
      },
    });
  }

}
