<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Client extends Model
{
    protected $table = 'clients';

    public function program()
    {
        return $this->hasOne(Developer::class, 'id', 'developer_id');
    }
    public function group()
    {
        return $this->hasOne(Developer::class, 'id', 'developer_id');
    }

}
