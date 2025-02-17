"""Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT."""

import requests as requests_http
from . import utils
from sdk.models import operations, shared
from typing import Optional

class Webhooks:
    _client: requests_http.Session
    _security_client: requests_http.Session
    _server_url: str
    _language: str
    _sdk_version: str
    _gen_version: str

    def __init__(self, client: requests_http.Session, security_client: requests_http.Session, server_url: str, language: str, sdk_version: str, gen_version: str) -> None:
        self._client = client
        self._security_client = security_client
        self._server_url = server_url
        self._language = language
        self._sdk_version = sdk_version
        self._gen_version = gen_version
        
    
    def activate_config(self, request: operations.ActivateConfigRequest) -> operations.ActivateConfigResponse:
        r"""Activate one config
        Activate a webhooks config by ID, to start receiving webhooks to its endpoint.
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.ActivateConfigRequest, base_url, '/api/webhooks/configs/{id}/activate', request)
        headers = {}
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('PUT', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ActivateConfigResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ConfigResponse])
                res.config_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def change_config_secret(self, request: operations.ChangeConfigSecretRequest) -> operations.ChangeConfigSecretResponse:
        r"""Change the signing secret of a config
        Change the signing secret of the endpoint of a webhooks config.
        
        If not passed or empty, a secret is automatically generated.
        The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.ChangeConfigSecretRequest, base_url, '/api/webhooks/configs/{id}/secret/change', request)
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "config_change_secret", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('PUT', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.ChangeConfigSecretResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ConfigResponse])
                res.config_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def deactivate_config(self, request: operations.DeactivateConfigRequest) -> operations.DeactivateConfigResponse:
        r"""Deactivate one config
        Deactivate a webhooks config by ID, to stop receiving webhooks to its endpoint.
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.DeactivateConfigRequest, base_url, '/api/webhooks/configs/{id}/deactivate', request)
        headers = {}
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('PUT', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.DeactivateConfigResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ConfigResponse])
                res.config_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def delete_config(self, request: operations.DeleteConfigRequest) -> operations.DeleteConfigResponse:
        r"""Delete one config
        Delete a webhooks config by ID.
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.DeleteConfigRequest, base_url, '/api/webhooks/configs/{id}', request)
        headers = {}
        headers['Accept'] = 'application/json'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('DELETE', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.DeleteConfigResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            pass
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def get_many_configs(self, request: operations.GetManyConfigsRequest) -> operations.GetManyConfigsResponse:
        r"""Get many configs
        Sorted by updated date descending
        """
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/webhooks/configs'
        headers = {}
        query_params = utils.get_query_params(operations.GetManyConfigsRequest, request)
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('GET', url, params=query_params, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.GetManyConfigsResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ConfigsResponse])
                res.configs_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def insert_config(self, request: shared.ConfigUser) -> operations.InsertConfigResponse:
        r"""Insert a new config
        Insert a new webhooks config.
        
        The endpoint should be a valid https URL and be unique.
        
        The secret is the endpoint's verification secret.
        If not passed or empty, a secret is automatically generated.
        The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)
        
        All eventTypes are converted to lower-case when inserted.
        """
        base_url = self._server_url
        
        url = base_url.removesuffix('/') + '/api/webhooks/configs'
        headers = {}
        req_content_type, data, form = utils.serialize_request_body(request, "request", 'json')
        if req_content_type not in ('multipart/form-data', 'multipart/mixed'):
            headers['content-type'] = req_content_type
        if data is None and form is None:
            raise Exception('request body is required')
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('POST', url, data=data, files=form, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.InsertConfigResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ConfigResponse])
                res.config_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    
    def test_config(self, request: operations.TestConfigRequest) -> operations.TestConfigResponse:
        r"""Test one config
        Test a config by sending a webhook to its endpoint.
        """
        base_url = self._server_url
        
        url = utils.generate_url(operations.TestConfigRequest, base_url, '/api/webhooks/configs/{id}/test', request)
        headers = {}
        headers['Accept'] = 'application/json;q=1, application/json;q=0'
        headers['user-agent'] = f'speakeasy-sdk/{self._language} {self._sdk_version} {self._gen_version}'
        
        client = self._security_client
        
        http_res = client.request('GET', url, headers=headers)
        content_type = http_res.headers.get('Content-Type')

        res = operations.TestConfigResponse(status_code=http_res.status_code, content_type=content_type, raw_response=http_res)
        
        if http_res.status_code == 200:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.AttemptResponse])
                res.attempt_response = out
        else:
            if utils.match_content_type(content_type, 'application/json'):
                out = utils.unmarshal_json(http_res.text, Optional[shared.ErrorResponse])
                res.error_response = out

        return res

    