<?php
namespace Optimuscrime\Lastfm;

use Exception;
use Monolog\Logger as MonologLogger;
use Optimuscrime\Lastfm\Containers\Database;
use Optimuscrime\Lastfm\Containers\InternalServerError;
use Optimuscrime\Lastfm\Containers\Logger;
use Optimuscrime\Lastfm\Containers\PageNotFound;
use Optimuscrime\Lastfm\Containers\Services;
use Optimuscrime\Lastfm\Middlewares\ReverseProxyMiddleware;
use Optimuscrime\Lastfm\Rest\Endpoints\AuthenticateEndpoint;
use Optimuscrime\Lastfm\Rest\Endpoints\HistoryEndpoint;
use Psr\Container\ContainerExceptionInterface;
use Psr\Container\NotFoundExceptionInterface;
use Slim\App as Slim;
use Slim\Exception\MethodNotAllowedException;
use Slim\Exception\NotFoundException;

class App
{
    private Slim $app;

    public function __construct(array $settings)
    {
        $this->app = new Slim($settings);
    }

    /**
     * @throws NotFoundExceptionInterface
     * @throws MethodNotAllowedException
     * @throws NotFoundException
     * @throws ContainerExceptionInterface
     */
    public function run(): void
    {
        $this->setup();

        try {
            $this->app->run();
        } catch (Exception $ex) {
            /** @var MonologLogger $logger */
            $logger = $this->app->getContainer()->get(Logger::class);

            $logger->error($ex);

            // Rethrow exception to the outer exception handler
            throw $ex;
        }
    }

    private function setup(): void
    {
        $this->dependencies();
        $this->routes();
    }

    private function routes(): void
    {
        $app = $this->app;
        $app->add(new ReverseProxyMiddleware());

        $app->group('/api', function () use ($app) {
            $app->get('/authenticate', AuthenticateEndpoint::class . ':authenticate');
            $app->get('/history', HistoryEndpoint::class . ':get');
        });
    }

    private function dependencies(): void
    {
        $containers = [
            Database::class,
            InternalServerError::class,
            PageNotFound::class,
            Services::class,
            Logger::class,
        ];

        foreach ($containers as $container) {
            call_user_func([$container, 'load'], $this->app->getContainer());
        }
    }
}
