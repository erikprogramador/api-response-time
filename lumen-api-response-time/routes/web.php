<?php

/** @var \Laravel\Lumen\Routing\Router $router */

use App\Models\Client;

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/v1/clients', function () use ($router) {
    return Client::all();
});
