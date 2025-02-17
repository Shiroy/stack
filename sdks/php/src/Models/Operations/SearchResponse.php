<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class SearchResponse
{
	
    public string $contentType;
    
    /**
     * Success
     * 
     * @var ?\formance\stack\Models\Shared\Response $response
     */
	
    public ?\formance\stack\Models\Shared\Response $response = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->response = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}
