from flask import Blueprint, render_template

qna = Blueprint('qna', __name__, template_folder='templates', url_prefix='/qna')

# default top timeline
@qna.route('/')
@qna.route('/management')
def get_management():
    return render_template("qna/management.html")

@qna.route('/replies_management')
def get_replies_management():
    return render_template('/qna/user_replies.html')