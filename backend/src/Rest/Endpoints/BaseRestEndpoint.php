<?php
namespace Optimuscrime\Lastfm\Rest\Endpoints;

use Slim\Http\Response;

class BaseRestEndpoint
{
    protected function outputJson(Response $response, array $object): Response
    {
        return $response
            ->withHeader('Content-Type', 'application/json')
            ->withJson($object);
    }

    protected function returnBadRequest(Response $response): Response
    {
        return $response->withStatus(400);
    }

    protected function returnServiceUnavailable(Response $response): Response
    {
        return $response->withStatus(503);
    }

    protected function returnForbidden(Response $response): Response
    {
        return $response->withStatus(403);
    }

    protected function returnInternalServerError(Response $response): Response
    {
        return $response->withStatus(500);
    }
}
