/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class ReadConnectorConfigRequest extends SpeakeasyBase {
  /**
   * The name of the connector.
   */
  @SpeakeasyMetadata({
    data: "pathParam, style=simple;explode=false;name=connector",
  })
  connector: shared.Connector;
}

export class ReadConnectorConfigResponse extends SpeakeasyBase {
  /**
   * OK
   */
  @SpeakeasyMetadata()
  connectorConfigResponse?: shared.ConnectorConfigResponse;

  @SpeakeasyMetadata()
  contentType: string;

  @SpeakeasyMetadata()
  statusCode: number;

  @SpeakeasyMetadata()
  rawResponse?: AxiosResponse;
}
