// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

import type { Kuuzuki } from '../client';

export abstract class APIResource {
  protected _client: Kuuzuki;

  constructor(client: Kuuzuki) {
    this._client = client;
  }
}
