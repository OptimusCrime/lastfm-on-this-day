<?php
namespace Optimuscrime\Lastfm\Helpers\Configuration;

use Exception;

class Configuration
{
    const DEV = 'dev';
    const SSL = 'ssl';

    const LOGGER_NAME = 'logger.name';

    const LASTFM_API_KEY = 'lastfm.api.key';
    const LASTFM_SHARED_SECRET = 'lastfm.shared.secret';

    private static ?Configuration $instance = null;
    private array $configuration;

    private function __construct()
    {
        $this->configuration = [];
    }

    public function isDev(): bool
    {
        return $this->lookup(static::DEV, '1') === '1';
    }

    public function isSSL(): bool
    {
        return $this->lookup(static::SSL, '0') === '1';
    }

    public function getLoggerName(): string
    {
        return $this->lookup(static::LOGGER_NAME, 'lastfm');
    }

    public function getLastFmApiKey(): string
    {
        return $this->lookup(static::LASTFM_API_KEY, '');
    }

    public function getLastFmSharedSecret(): string {
        return $this->lookup(static::LASTFM_SHARED_SECRET, '');
    }

    private function lookup(string $key, string $default): string
    {
        if (isset($this->configuration[$key])) {
            return $this->configuration[$key];
        }

        // Fetch from env (if it exists) and store for later lookups
        $envKey = str_replace('.', '_', strtoupper($key));
        $envValue = getenv($envKey);

        // No env variable. Store the default value for later lookups
        if ($envValue === false) {
            $this->configuration[$key] = $default;
            return $default;
        }

        $this->configuration[$key] = $envValue;
        return $envValue;
    }

    public static function getInstance(): Configuration
    {
        if (self::$instance === null) {
            self::$instance = new Configuration();
        }

        return self::$instance;
    }
}
