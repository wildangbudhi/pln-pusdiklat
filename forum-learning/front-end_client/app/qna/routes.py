from flask import Blueprint, render_template

qna = Blueprint('qna', __name__, template_folder='templates', url_prefix='/qna')

# default top timeline
@qna.route('/')
@qna.route('/timeline')
def latest_timeline():
    return render_template("qna/timeline.html")

@qna.route('/<string:username>/<string:qna_id>')
def user_question(username, qna_id):
    return render_template('/qna/user_question.html', username=username, qna_id=qna_id)