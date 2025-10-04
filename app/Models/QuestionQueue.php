<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class QuestionQueue extends Model
{
    use HasFactory;

    protected $fillable = [
        'game_id',
        'question_id',
    ];
}
