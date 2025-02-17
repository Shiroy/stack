/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import org.openapis.openapi.utils.SpeakeasyMetadata;

public class CreditWalletRequest {
    @SpeakeasyMetadata("request:mediaType=application/json")
    public org.openapis.openapi.models.shared.CreditWalletRequest creditWalletRequest;

    public CreditWalletRequest withCreditWalletRequest(org.openapis.openapi.models.shared.CreditWalletRequest creditWalletRequest) {
        this.creditWalletRequest = creditWalletRequest;
        return this;
    }
    
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=id")
    public String id;

    public CreditWalletRequest withId(String id) {
        this.id = id;
        return this;
    }
    
    public CreditWalletRequest(@JsonProperty("id") String id) {
        this.id = id;
  }
}
