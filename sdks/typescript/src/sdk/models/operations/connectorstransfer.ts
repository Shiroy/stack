/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class ConnectorsTransferRequest extends SpeakeasyBase {
  @SpeakeasyMetadata({ data: "request, media_type=application/json" })
  transferRequest: shared.TransferRequest;

  /**
   * The name of the connector.
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=connector",
  })
  connector: shared.Connector;
}

export class ConnectorsTransferResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;

  /**
   * OK
   */
  @SpeakeasyMetadata()
  transferResponse?: shared.TransferResponse;
}
