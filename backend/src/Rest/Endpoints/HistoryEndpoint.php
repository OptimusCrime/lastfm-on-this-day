<?php

namespace Optimuscrime\Lastfm\Rest\Endpoints;

use DateTimeImmutable;
use Exception;

use Monolog\Logger;
use Optimuscrime\Lastfm\Services\LastFmService;
use Psr\Container\ContainerExceptionInterface;
use Psr\Container\ContainerInterface;
use Psr\Container\NotFoundExceptionInterface;
use Slim\Http\Request;
use Slim\Http\Response;

class HistoryEndpoint extends BaseRestEndpoint
{
    private Logger $logger;
    private LastFmService $lastFmService;

    /**
     * @throws ContainerExceptionInterface
     * @throws NotFoundExceptionInterface
     */
    public function __construct(ContainerInterface $container)
    {
        $this->logger = $container->get(Logger::class);
        $this->lastFmService = $container->get(LastFmService::class);
    }

    public function get(Request $request, Response $response): Response
    {
        // TODO: Read from headers
        $token = $request->getParam('sk');

        try {
            // TODO: Read from request
            $from = DateTimeImmutable::createFromFormat("Y-m-d H:i:s", "2022-09-15 00:00:00", new \DateTimeZone('Europe/Oslo'));
            $to = DateTimeImmutable::createFromFormat("Y-m-d H:i:s", "2022-09-15 23:59:59", new \DateTimeZone('Europe/Oslo'));

            $recentTracksResponse = $this->lastFmService->getRecentTracks($token, $from, $to);

            // Last.fm API has a weird tendency to break. Let's just return an error here and let the frontend
            // try to do some "sane" retry stuff instead. Or something.
            if ($recentTracksResponse === null) {
                return $this->returnServiceUnavailable($response);
            }

            // TODO: Maybe we should do some of this inside of the service, and throw an error or something
            $recentTracksResponseDecoded = json_decode($recentTracksResponse, true);
            if (!is_array($recentTracksResponseDecoded)) {
                return $this->returnServiceUnavailable($response);
            }

            $recentTracks = static::formatRecentTracksResponse($recentTracksResponseDecoded['recenttracks']['track']);

            return $this->outputJson($response, [
                'data' => $recentTracks
            ]);
        } catch (Exception $ex) {
            $this->logger->error('Failed to create session key from authentication token', [
                'error' => $ex
            ]);

            return $this->returnInternalServerError($response);
        }
    }

    private static function formatRecentTracksResponse(array $recentTracks): array
    {
        $lookup = [];
        foreach ($recentTracks as $recentTrack) {
            // Skip the song we are currently listening to (which is, for some weird reason, always returned)
            if (isset($recentTrack['@attr']['nowplaying']) && $recentTrack['@attr']['nowplaying']) {
                continue;
            }

            // Whatever
            $checksum = md5($recentTrack['artist']['mbid'] . $recentTrack['album']['mbid'] . $recentTrack['name']);

            if (!isset($lookup[$checksum])) {
                $lookup[$checksum] = [
                    'artist' => $recentTrack['artist']['#text'],
                    'album' => $recentTrack['album']['#text'],
                    'song' => $recentTrack['name'],
                    'playCount' => 1,
                    'playedAt' => [
                        $recentTrack['date']['uts']
                    ]
                ];

                continue;
            }

            $lookup[$checksum]['playCount'] += 1;
            $lookup[$checksum]['playedAt'][] = $recentTrack['date']['uts'];
        }

        $output = [];
        foreach ($lookup as $v) {
            $output[] = $v;
        }

        usort($output, fn($a, $b) => $b['playCount'] - $a['playCount']);

        return $output;
    }
}
