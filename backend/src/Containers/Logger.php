<?php
namespace Optimuscrime\Lastfm\Containers;

use Monolog\Formatter\LineFormatter;
use Monolog\Handler\StreamHandler;
use Monolog\Logger as MonologLogger;
use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;
use Psr\Container\ContainerInterface;

class Logger
{
    public static function load(ContainerInterface $container): void
    {
        $container[MonologLogger::class] = function (): MonologLogger {
            $configuration = Configuration::getInstance();

            $logger = new MonologLogger($configuration->getLoggerName());

            $formatter = new LineFormatter(LineFormatter::SIMPLE_FORMAT, LineFormatter::SIMPLE_DATE);
            $formatter->includeStacktraces(true);

            $stream = new StreamHandler(
                'php://stdout',
                MonologLogger::DEBUG
            );

            $stream->setFormatter($formatter);

            $logger->pushHandler($stream);

            return $logger;
        };
    }
}
