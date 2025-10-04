<?php

use App\Http\Controllers\GameController;
use Illuminate\Support\Facades\Route;

Route::post('/games/random', [GameController::class, 'createRandomGame']);

