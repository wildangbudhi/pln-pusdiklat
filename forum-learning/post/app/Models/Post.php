<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Post extends Model
{
    use HasFactory;
    protected $connection = 'mysql';
    protected $table = 'posts';
    protected $fillable = [ 'title', 'discussion_id', 'author_user_id', 'category_id', 'articles', 'thumbnail'];
}
