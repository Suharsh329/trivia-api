<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class GameResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @return array<string, mixed>
     */
    public function toArray(Request $request): array
    {
        return [
            'game' => [
                'id' => $this->resource['game']->id,
                'name' => $this->resource['game']->name,
                'created_at' => $this->resource['game']->created_at,
                'updated_at' => $this->resource['game']->updated_at,
            ],
            'questions' => QuestionResource::collection($this->resource['questions']),
        ];
    }
}
