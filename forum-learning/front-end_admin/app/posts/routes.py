from flask import Blueprint, render_template

posts = Blueprint('posts', __name__, template_folder='templates', url_prefix='/posts')

# default top timeline
@posts.route('/')
@posts.route('/management')
def get_management():
    return render_template("posts/management.html")