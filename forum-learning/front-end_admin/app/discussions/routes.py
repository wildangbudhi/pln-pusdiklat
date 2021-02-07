from flask import Blueprint, render_template

discussions = Blueprint('discussions', __name__, template_folder='templates', url_prefix='/discussions')

# default top timeline
@discussions.route('/')
@discussions.route('/management')
def get_management():
    return render_template("discussions/management.html")

@discussions.route('/request')
def get_request_management():
    return render_template('/discussions/request.html')