<?php
namespace Optimuscrime\Lastfm\Containers;

use Illuminate\Container\Container;
use Illuminate\Database\Capsule\Manager as DB;
use Illuminate\Database\ConnectionResolver;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Events\Dispatcher;
use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;
use Psr\Container\ContainerInterface;

class Database implements ContainersInterface
{
    public static function load(ContainerInterface $container): void
    {
        $configuration = Configuration::getInstance();

        $connection = [
            'driver' => 'sqlite',
            'charset'   => 'utf8',
            'collation' => 'utf8_unicode_ci',
            'prefix'    => '',
        ];

        $capsule = new DB;
        $capsule->addConnection($connection);

        $capsule->setEventDispatcher(new Dispatcher(new Container()));

        // Make it possible to use $app->get('db') -> whatever
        $capsule->setAsGlobal();
        $capsule->bootEloquent();

        // Make it possible to use Model :: whatever
        $resolver = new ConnectionResolver();
        $resolver->addConnection('default', $capsule->getConnection());
        $resolver->setDefaultConnection('default');
        Model::setConnectionResolver($resolver);

        if ($configuration->isDev()) {
            DB::connection()->enableQueryLog();
        }

        $container['db'] = function () use ($capsule): DB {
            return $capsule;
        };
    }
}
