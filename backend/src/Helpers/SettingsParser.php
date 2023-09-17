<?php
namespace Optimuscrime\Lastfm\Helpers;

use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;

class SettingsParser
{
    public static function getSlimConfig(): array
    {
        return [
            'settings' => [
                'displayErrorDetails' => Configuration::getInstance()->isDev(),
                'addContentLengthHeader' => false,
            ]
        ];
    }

    public static function getPhinxConfig(): array
    {
        $configuration = Configuration::getInstance();
        return [
            'paths' => [
                'migrations' => '%%PHINX_CONFIG_DIR%%/../phinx/migrations',
                'seeds' => '%%PHINX_CONFIG_DIR%%/../phinx/seeds',
            ],
            'environments' => [
                'default_migration_table' => 'phinxlog',
                'default_database' => 'production',
                'production' => [
                    'adapter' => 'sqlite',
                    'name' => './tests/files/db.sqlite3',
                ],
                'test' => [
                    'adapter' => 'sqlite',
                    'name' => './tests/files/db.sqlite3',
                ]
            ]
        ];
    }
}
