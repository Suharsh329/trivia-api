<?php

namespace App\Services;

use App\Models\Question;
use App\Models\QuestionQueue;
use App\Models\Game;
use App\Models\Contestant;
use Illuminate\Support\Collection;
use InvalidArgumentException;

class GameService
{

    public function createSelectedGame(array $data): array
    {
        // TODO: Implement selected game creation logic
        return [];
    }

    public function createRandomGame(array $data): array
    {
        $this->validateRandomGameData($data);

        // Create the game first
        $game = $this->createGame();

        $this->createContestants($game->id, $data['contestants'], $data['is_team']);

        // Get random questions based on difficulty percentages
        $questions = $this->getRandomQuestions($data);

        // Add questions to the queue for this game
        $this->addQuestionsToQueue($game->id, $questions);

        return [
            'game' => $game,
            'questions' => $questions
        ];
    }

    public function generateRandomGameName(): string
    {
        $adjectives = config('app.game_name_adjectives');
        $nouns = config('app.game_name_nouns');

        $randomAdjective = $adjectives[array_rand($adjectives)];
        $randomNoun = $nouns[array_rand($nouns)];
        $randomNumber = rand(1, 1000);

        return "$randomAdjective-$randomNoun-$randomNumber";
    }

    private function createGame(): Game
    {
        $gameName = $this->generateRandomGameName();

        return Game::query()->create([
            'name' => $gameName
        ]);
    }

    private function getRandomQuestions(array $data): Collection
    {
        $percentages = $data['percentages'];
        $numberOfQuestions = $data['number_of_questions'];
        $originalNumberOfQuestions = $numberOfQuestions;
        $numberOfQuestions += 10; // Add buffer to ensure enough questions
        $questions = collect(); // Initialize as collection
        $totalCollected = 0;

        foreach ($percentages as $difficulty => $percentage) {
            $limit = (int)ceil($numberOfQuestions * ($percentage / 100));

            if ($limit <= 0) {
                continue;
            }

            // Ensure we don't exceed the total number of questions requested
            $remainingQuestions = $numberOfQuestions - $totalCollected;
            if ($limit > $remainingQuestions) {
                $limit = $remainingQuestions;
            }

            $difficultyQuestions = Question::query()
                ->with('subCategory') // Load subcategory relationship
                ->where('difficulty_level', $difficulty)
                ->inRandomOrder()
                ->limit($limit)
                ->get();

            $questions = $questions->merge($difficultyQuestions);
            $totalCollected += $difficultyQuestions->count();

            // Break if we've collected enough questions
            if ($totalCollected >= $numberOfQuestions) {
                break;
            }
        }

        // Shuffle the collection to mix difficulties
        $questions = $questions->shuffle();

        // Take only the exact number requested (without buffer)
        return $questions->take($originalNumberOfQuestions);
    }

    private function addQuestionsToQueue(int $gameId, $questions): void
    {
        $queueData = [];

        foreach ($questions as $question) {
            $queueData[] = [
                'game_id' => $gameId,
                'question_id' => $question->id,
                'created_at' => now(),
                'updated_at' => now(),
            ];
        }

        // Bulk insert for better performance
        QuestionQueue::query()->insert($queueData);
    }

    private function validateRandomGameData(array $data): void
    {
        if (!isset($data['percentages']) || !is_array($data['percentages'])) {
            throw new InvalidArgumentException('Percentages must be provided as an array');
        }

        if (!isset($data['number_of_questions']) || !is_numeric($data['number_of_questions'])) {
            throw new InvalidArgumentException('Number of questions must be provided as a number');
        }

        if ($data['number_of_questions'] <= 0) {
            throw new InvalidArgumentException('Number of questions must be greater than 0');
        }

        $totalPercentage = array_sum($data['percentages']);
        if ($totalPercentage != 100) {
            throw new InvalidArgumentException('Percentages must sum to 100');
        }
    }

    public function createContestants(int $gameId, array $contestants, bool $isTeam): array
    {
        $createdContestants = [];

        foreach ($contestants as $contestant) {
            $contestant = Contestant::query()->create([
                'game_id' => $gameId,
                'name' => $contestant,
                'is_team' => $isTeam,
            ]);

            $createdContestants[] = $contestant;
        }

        return $createdContestants;
    }
}
