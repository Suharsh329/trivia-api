<?php

namespace App\Filament\Resources\QuestionQueues\Schemas;

use Filament\Forms\Components\TextInput;
use Filament\Schemas\Schema;

class QuestionQueueForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextInput::make('game_id')
                    ->required()
                    ->numeric(),
                TextInput::make('question_id')
                    ->required()
                    ->numeric(),
            ]);
    }
}
