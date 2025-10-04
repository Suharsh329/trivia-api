<?php

namespace App\Filament\Resources\Questions\Schemas;

use Filament\Forms\Components\TextInput;
use Filament\Forms\Components\Textarea;
use Filament\Schemas\Schema;

class QuestionForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextInput::make('sub_category_id')
                    ->required()
                    ->numeric(),
                Textarea::make('question_text')
                    ->required()
                    ->columnSpanFull(),
                Textarea::make('correct_answer')
                    ->required()
                    ->columnSpanFull(),
                Textarea::make('acceptable_answers')
                    ->columnSpanFull(),
                TextInput::make('difficulty_level')
                    ->required()
                    ->numeric()
                    ->default(1),
                Textarea::make('image_url')
                    ->columnSpanFull(),
            ]);
    }
}
