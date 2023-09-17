<?php
namespace Optimuscrime\Lastfm\Containers;

use Exception;
use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;
use Psr\Container\ContainerInterface;
use Slim\Http\Request;
use Slim\Http\Response;

class InternalServerError implements ContainersInterface
{
    public static function load(ContainerInterface $container): void
    {
        $container['errorHandler'] = function () {
            return function (Request $request, Response $response, Exception $exception): Response {
                $configuration = Configuration::getInstance();
                if ($configuration->isDev()) {
                    var_dump(get_class($exception));
                    var_dump($exception->getMessage());
                    var_dump($exception->getTraceAsString());
                    die();
                }

                return $response->withStatus(500);
            };
        };
    }
}
