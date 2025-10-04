<?php

namespace App\Filament\Resources\Games\Schemas;

use Filament\Forms\Components\Textarea;
use Filament\Schemas\Schema;

class GameForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                Textarea::make('name')
                    ->required()
                    ->columnSpanFull(),
            ]);
    }
}
