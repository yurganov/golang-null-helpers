<?php

namespace App\Http\Resources;

use Illuminate\Http\Resources\Json\JsonResource;

class ClientResource extends JsonResource
{
    public function toArray($client)
    {
        return [
            'id' => $this->ID,
            'name' => $this->Client,
            'group' => [
                'name' => $this->group->name,
                'id' => $this->group->id,
            ],
            'program' => [
                'name' => $this->program->name,
                'id' => $this->program->id,
            ],

        ];
    }
}
