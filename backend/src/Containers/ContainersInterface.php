<?php
namespace Optimuscrime\Lastfm\Containers;

use Psr\Container\ContainerInterface;

interface ContainersInterface
{
    public static function load(ContainerInterface $container): void;
}
