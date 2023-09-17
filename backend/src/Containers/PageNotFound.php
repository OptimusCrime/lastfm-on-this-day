<?php
namespace Optimuscrime\Lastfm\Containers;

use Closure;
use Psr\Container\ContainerInterface;
use Slim\Http\Request;
use Slim\Http\Response;

class PageNotFound implements ContainersInterface
{
    public static function load(ContainerInterface $container): void
    {
        $container['notFoundHandler'] = function (ContainerInterface $container): Closure {
            return function (Request $request, Response $response) use ($container) {
                return $response->withStatus(404);
            };
        };
    }
}
