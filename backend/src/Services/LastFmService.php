<?php

namespace Optimuscrime\Lastfm\Services;

use DateTimeImmutable;
use Optimuscrime\Lastfm\Helpers\Configuration\Configuration;

class LastFmService
{
    const CURL_CONNECTION_TIMEOUT_IN_SECONDS = 5;
    const CURL_TIMEOUT_IN_SECONDS = 8;

    const API_BASE_URL = 'https://ws.audioscrobbler.com/2.0/';

    public function getWebServiceSessionToken(string $token): string
    {
        return static::makeUnauthorizedRequest('auth.getSession', $token);
    }

    public function getRecentTracks(string $sk, DateTimeImmutable $from, DateTimeImmutable $to): string
    {
        return static::makeAuthorizedRequest(
            'user.getRecentTracks',
            [
                'from' => $from->getTimestamp(),
                'to' => $to->getTimestamp()
            ],
            $sk
        );
    }

    private function makeUnauthorizedRequest(string $operation, string $token): string
    {
        return static::makeRequest($operation, [
            'token' => $token
        ]);
    }

    private function makeAuthorizedRequest(string $operation, array $params, string $sessionKey): string
    {
        return static::makeRequest(
            $operation,
            array_merge(
                $params, [
                    'sk' => $sessionKey
                ]
            )
        );
    }

    private static function makeRequest(string $operation, array $params): string
    {
        $requestUrl = static::createRequestUrl($operation, $params);

        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $requestUrl);
        curl_setopt($ch, CURLOPT_HEADER, 0);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, true);
        curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, static::CURL_CONNECTION_TIMEOUT_IN_SECONDS);
        curl_setopt($ch, CURLOPT_TIMEOUT, static::CURL_TIMEOUT_IN_SECONDS);

        $output = curl_exec($ch);
        curl_close($ch);

        return $output;
    }

    private static function createRequestUrl(string $operation, array $params): string
    {
        $configuration = Configuration::getInstance();

        $params['api_key'] = $configuration->getLastFmApiKey();
        $params['method'] = $operation;

        $signedChecksum = static::createSignedChecksum($params);

        $params['api_sig'] = $signedChecksum;
        $params['format'] = 'json';

        return static::API_BASE_URL . '?' . static::buildRequestParams($params);
    }

    private static function createSignedChecksum(array $params): string
    {
        // Got to love mutating sort <3
        ksort($params);

        $keyValuePairs = [];
        foreach ($params as $key => $value) {
            $keyValuePairs[] = $key . $value;
        }

        $checksumString = join('', $keyValuePairs);

        $configuration = Configuration::getInstance();

        return md5($checksumString . $configuration->getLastFmSharedSecret());
    }

    private static function buildRequestParams(array $params): string
    {
        $outParams = [];
        foreach ($params as $key => $value) {
            $outParams[] = $key . '=' . $value;
        }

        return join('&', $outParams);
    }
}