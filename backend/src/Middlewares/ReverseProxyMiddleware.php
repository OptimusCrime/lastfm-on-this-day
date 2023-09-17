<?php
namespace Optimuscrime\Lastfm\Middlewares;

use Exception;
use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;
use Slim\Http\Request;
use Slim\Http\Response;

class ReverseProxyMiddleware
{
    public function __invoke(Request $request, Response $response, callable $next): Response
    {
        $scheme = Configuration::getInstance()->isSSL() ? 'https' : 'http';
        $uri = $request->getUri()->withScheme($scheme);
        $request = $request->withUri($uri);
        return $next($request, $response);
    }
}
