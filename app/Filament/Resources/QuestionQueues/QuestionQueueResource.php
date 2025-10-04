<?php

namespace App\Filament\Resources\QuestionQueues;

use App\Filament\Resources\QuestionQueues\Pages\CreateQuestionQueue;
use App\Filament\Resources\QuestionQueues\Pages\EditQuestionQueue;
use App\Filament\Resources\QuestionQueues\Pages\ListQuestionQueues;
use App\Filament\Resources\QuestionQueues\Schemas\QuestionQueueForm;
use App\Filament\Resources\QuestionQueues\Tables\QuestionQueuesTable;
use App\Models\QuestionQueue;
use BackedEnum;
use Filament\Resources\Resource;
use Filament\Schemas\Schema;
use Filament\Support\Icons\Heroicon;
use Filament\Tables\Table;

class QuestionQueueResource extends Resource
{
    protected static ?string $model = QuestionQueue::class;

    protected static string|BackedEnum|null $navigationIcon = Heroicon::OutlinedRectangleStack;

    protected static ?string $recordTitleAttribute = 'QuestionQueue';

    public static function form(Schema $schema): Schema
    {
        return QuestionQueueForm::configure($schema);
    }

    public static function table(Table $table): Table
    {
        return QuestionQueuesTable::configure($table);
    }

    public static function getRelations(): array
    {
        return [
            //
        ];
    }

    public static function getPages(): array
    {
        return [
            'index' => ListQuestionQueues::route('/'),
            'create' => CreateQuestionQueue::route('/create'),
            'edit' => EditQuestionQueue::route('/{record}/edit'),
        ];
    }
}
