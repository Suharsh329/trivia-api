<?php

namespace App\Http\Controllers;

use App\Services\GameService;
use App\Http\Resources\GameResource;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class GameController extends Controller
{
    public function __construct(private readonly GameService $gameService) {}

    public function createRandomGame(Request $request): JsonResponse
    {
        $data = $request->validate([
            'is_team' => 'required|boolean',
            'percentages' => 'required|array',
            'number_of_questions' => 'required|integer|min:1|max:50',
            'contestants' => 'required|array',
        ]);

        $game = $this->gameService->createRandomGame($data);

        return response()->json([
            'message' => 'Game created successfully',
            'data' => new GameResource($game)
        ], 201);
    }
}
