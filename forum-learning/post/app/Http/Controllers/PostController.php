<?php

namespace App\Http\Controllers;

use App\Models\Post;
use Illuminate\Http\Request;

class PostController extends Controller
{
    public function index()
    {
        $post = Post::all();
        return $post;
    }

    public function create(Request $request)
    {
        Post::create($request->all());

        return response()->json([
            'status' => 'OK',
            'message' => 'Data Berhasil Disimpan'
        ], 200);
    }

    public function show($id)
    {
        //
    }

    public function edit($id)
    {
        //
    }

    public function update(Request $request, $id)
    {
        Post::find($id)->update($request->all());

        return response()->json([
            'status' => 'OK',
            'message' => 'Data Berhasil Update'
        ], 200);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        Post::destroy($id);

        return response()->json([
            'status' => 'OK',
            'message' => 'Data Berhasil Hapus'
        ], 200);
    }
}
