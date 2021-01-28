from flask import Blueprint, render_template

posts = Blueprint('posts', __name__, template_folder='templates', url_prefix='/posts')

# default top timeline
@posts.route('/')
@posts.route('/timeline')
def latest_timeline():
    return render_template("posts/timeline.html")

@posts.route('/<string:username>/<string:posts_id>')
def user_post(username, posts_id):
    return render_template('/posts/user_post.html', username=username, posts_id=posts_id)