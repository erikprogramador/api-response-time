<?php

use App\Models\Client;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/v1/clients', function () {
    return Client::all();
});

Route::get('/v1/clients/{client}', function (Client $client) {
    return $client;
});

Route::post('/v1/clients', function (Request $request) {
    $client = Client::create([
        'name' => $request->name,
        'email' => $request->email,
        'phone' => $request->phon,
        'domain' => $request->domain,
        'description' => $request->description,
        'tags' => $request->tags,
    ]);

    return $client;
});

Route::put('/v1/clients/{client}', function (Request $request, Client $client) {
    $client->update([
        'name' => $request->name,
        'email' => $request->email,
        'phone' => $request->phon,
        'domain' => $request->domain,
        'description' => $request->description,
        'tags' => $request->tags,
    ]);

    return $client->fresh();
});

Route::delete('/v1/clients/{client}', function (Client $client) {
    $client->delete();

    return [];
});
