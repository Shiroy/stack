/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Workflow } from "./workflow";
import { Expose, Type } from "class-transformer";

/**
 * Created workflow
 */
export class CreateWorkflowResponse extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "data" })
  @Type(() => Workflow)
  data: Workflow;
}
