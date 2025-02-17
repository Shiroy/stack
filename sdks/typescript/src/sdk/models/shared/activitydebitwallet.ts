/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { DebitWalletRequest } from "./debitwalletrequest";
import { Expose, Type } from "class-transformer";

export class ActivityDebitWallet extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "data" })
  @Type(() => DebitWalletRequest)
  data?: DebitWalletRequest;

  @SpeakeasyMetadata()
  @Expose({ name: "id" })
  id?: string;
}
