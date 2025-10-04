<?php

namespace App\Filament\Resources\Contestants\Schemas;

use Filament\Forms\Components\TextInput;
use Filament\Forms\Components\Toggle;
use Filament\Schemas\Schema;

class ContestantForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                TextInput::make('game_id')
                    ->required()
                    ->numeric(),
                TextInput::make('name')
                    ->required(),
                TextInput::make('score')
                    ->required()
                    ->numeric()
                    ->default(0),
                Toggle::make('is_team')
                    ->required(),
            ]);
    }
}
