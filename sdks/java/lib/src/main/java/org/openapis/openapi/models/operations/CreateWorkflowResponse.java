/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package org.openapis.openapi.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.net.http.HttpResponse;

public class CreateWorkflowResponse {
    
    public String contentType;

    public CreateWorkflowResponse withContentType(String contentType) {
        this.contentType = contentType;
        return this;
    }
    
    /**
     * Created workflow
     */
    
    public org.openapis.openapi.models.shared.CreateWorkflowResponse createWorkflowResponse;

    public CreateWorkflowResponse withCreateWorkflowResponse(org.openapis.openapi.models.shared.CreateWorkflowResponse createWorkflowResponse) {
        this.createWorkflowResponse = createWorkflowResponse;
        return this;
    }
    
    /**
     * General error
     */
    
    public org.openapis.openapi.models.shared.Error error;

    public CreateWorkflowResponse withError(org.openapis.openapi.models.shared.Error error) {
        this.error = error;
        return this;
    }
    
    
    public Integer statusCode;

    public CreateWorkflowResponse withStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
        return this;
    }
    
    
    public HttpResponse<byte[]> rawResponse;

    public CreateWorkflowResponse withRawResponse(HttpResponse<byte[]> rawResponse) {
        this.rawResponse = rawResponse;
        return this;
    }
    
    public CreateWorkflowResponse(@JsonProperty("ContentType") String contentType, @JsonProperty("StatusCode") Integer statusCode) {
        this.contentType = contentType;
        this.statusCode = statusCode;
  }
}
