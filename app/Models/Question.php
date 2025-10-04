<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsTo;

class Question extends Model
{

    use HasFactory;

    protected $fillable = [
        'sub_category_id',
        'question_text',
        'correct_answer',
        'accepted_answers',
        'difficulty_level',
        'image_url',
    ];

    public function subCategory(): BelongsTo
    {
        return $this->belongsTo(SubCategory::class);
    }
}
