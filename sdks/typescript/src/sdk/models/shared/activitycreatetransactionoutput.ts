/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Transaction } from "./transaction";
import { Expose, Type } from "class-transformer";

export class ActivityCreateTransactionOutput extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "data" })
  @Type(() => Transaction)
  data: Transaction;
}
