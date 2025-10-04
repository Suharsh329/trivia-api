<?php

namespace App\Filament\Resources\QuestionQueues\Pages;

use App\Filament\Resources\QuestionQueues\QuestionQueueResource;
use Filament\Actions\CreateAction;
use Filament\Resources\Pages\ListRecords;

class ListQuestionQueues extends ListRecords
{
    protected static string $resource = QuestionQueueResource::class;

    protected function getHeaderActions(): array
    {
        return [
            CreateAction::make(),
        ];
    }
}
