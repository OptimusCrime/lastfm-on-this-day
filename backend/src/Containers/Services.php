<?php
namespace Optimuscrime\Lastfm\Containers;

use Optimuscrime\Lastfm\Services\AuthService;
use Optimuscrime\Lastfm\Services\LastFmService;
use Psr\Container\ContainerInterface;

class Services implements ContainersInterface
{
    public static function load(ContainerInterface $container): void
    {
        $container[LastFmService::class] = function (ContainerInterface $container): LastFmService {
            return new LastFmService();
        };

        $container[AuthService::class] = function (ContainerInterface $container): AuthService {
            return new AuthService($container->get(LastFmService::class));
        };

    }
}
