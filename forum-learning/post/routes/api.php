<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('post', 'PostController@index')->name('view_post');
Route::post('post', 'PostController@create')->name('post_post');
Route::put('post/{id}/update', 'PostController@update')->name('update_post');
Route::delete('post/{id}/delete', 'PostController@destroy')->name('delete_post');
