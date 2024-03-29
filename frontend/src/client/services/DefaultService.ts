/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { EndpointsAuthOutput } from '../models/EndpointsAuthOutput';
import type { EndpointsCliInput } from '../models/EndpointsCliInput';
import type { EndpointsCliOutput } from '../models/EndpointsCliOutput';
import type { EndpointsGetFuzzBugsOutput } from '../models/EndpointsGetFuzzBugsOutput';
import type { EndpointsGetFuzzesOutput } from '../models/EndpointsGetFuzzesOutput';
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
   * Get Fuzz Bugs
   * Returns the list of all the Fuzz Bugs
   * @returns EndpointsGetFuzzBugsOutput OK
   * @throws ApiError
   */
  public endpointsGetFuzzesBug(): CancelablePromise<EndpointsGetFuzzBugsOutput> {
    return this.httpRequest.request({
      method: 'GET',
      url: '/fuzzBugs',
      errors: {
        400: `Bad Request`,
      },
    });
  }

  /**
   * Get Fuzzes
   * Returns the list of all the times the fuzzer ran
   * @returns EndpointsGetFuzzesOutput OK
   * @throws ApiError
   */
  public endpointsGetFuzzes(): CancelablePromise<EndpointsGetFuzzesOutput> {
    return this.httpRequest.request({
      method: 'GET',
      url: '/fuzzes',
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

  /**
   * Open Api Fuzz
   * Recieves the open-api file from client and adds to s3
   * @param requestBody
   * @returns EndpointsCliOutput OK
   * @throws ApiError
   */
  public endpointsFuzz(
    requestBody?: EndpointsCliInput,
  ): CancelablePromise<EndpointsCliOutput> {
    return this.httpRequest.request({
      method: 'POST',
      url: '/openapi',
      body: requestBody,
      mediaType: 'application/json',
      errors: {
        400: `Bad Request`,
      },
    });
  }

}
