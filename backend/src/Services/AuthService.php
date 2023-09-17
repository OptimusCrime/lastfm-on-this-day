<?php
namespace Optimuscrime\Lastfm\Services;

class AuthService
{
    private LastFmService $lastFmService;

    public function __construct(LastFmService $lastFmService)
    {
        $this->lastFmService = $lastFmService;
    }

    public function getSessionKey(string $token): string
    {
        // TODO: Store in database and stuff
        return $this->lastFmService->getWebServiceSessionToken($token);
    }
}
