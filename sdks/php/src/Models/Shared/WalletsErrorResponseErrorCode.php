<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


enum WalletsErrorResponseErrorCode: string
{
    case VALIDATION = 'VALIDATION';
    case INTERNAL_ERROR = 'INTERNAL_ERROR';
    case INSUFFICIENT_FUND = 'INSUFFICIENT_FUND';
    case HOLD_CLOSED = 'HOLD_CLOSED';
}
