<?php

namespace App\Filament\Resources\QuestionQueues\Pages;

use App\Filament\Resources\QuestionQueues\QuestionQueueResource;
use Filament\Actions\DeleteAction;
use Filament\Resources\Pages\EditRecord;

class EditQuestionQueue extends EditRecord
{
    protected static string $resource = QuestionQueueResource::class;

    protected function getHeaderActions(): array
    {
        return [
            DeleteAction::make(),
        ];
    }
}
