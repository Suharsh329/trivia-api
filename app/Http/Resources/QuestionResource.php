<?php

namespace App\Http\Resources;

use Illuminate\Http\Request;
use Illuminate\Http\Resources\Json\JsonResource;

class QuestionResource extends JsonResource
{
    /**
     * Transform the resource into an array.
     *
     * @return array<string, mixed>
     */
    public function toArray(Request $request): array
    {
        return [
            'id' => $this->id,
            'sub_category_name' => $this->subCategory?->name,
            'question_text' => $this->question_text,
            'correct_answer' => $this->correct_answer,
            'accepted_answers' => $this->accepted_answers,
            'difficulty_level' => $this->difficulty_level,
            'image_url' => $this->image_url,
            'created_at' => $this->created_at,
            'updated_at' => $this->updated_at,
        ];
    }
}
