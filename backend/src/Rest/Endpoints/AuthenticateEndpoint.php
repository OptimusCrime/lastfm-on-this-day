<?php
namespace Optimuscrime\Lastfm\Rest\Endpoints;

use Exception;
use Monolog\Logger;
use Optimuscrime\Lastfm\Services\AuthService;
use Psr\Container\ContainerExceptionInterface;
use Psr\Container\ContainerInterface;
use Psr\Container\NotFoundExceptionInterface;
use Slim\Http\Request;
use Slim\Http\Response;

class AuthenticateEndpoint extends BaseRestEndpoint
{
    private Logger $logger;
    private AuthService $authService;

    /**
     * @throws ContainerExceptionInterface
     * @throws NotFoundExceptionInterface
     */
    public function __construct(ContainerInterface $container)
    {
        $this->logger = $container->get(Logger::class);
        $this->authService = $container->get(AuthService::class);
    }

    public function authenticate(Request $request, Response $response): Response
    {
        $token = $request->getParam('token');

        if (!is_string($token) || mb_strlen($token) === 0) {
            return $this->returnBadRequest($response);
        }

        try {
            $requestToken = $this->authService->getSessionKey($token);

            // TODO: Create bearer token and store in database
            return $this->outputJson($response, [
                'data' => [
                    'token' => 'lorem-ipsum-dolor-sit-amet'
                ]
            ]);
        }
        catch (Exception $ex) {
            $this->logger->error('Failed to create session key from authentication token', [
                'error' => $ex
            ]);

            return $this->returnInternalServerError($response);
        }
    }
}
