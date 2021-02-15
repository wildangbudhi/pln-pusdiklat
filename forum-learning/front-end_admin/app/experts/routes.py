from flask import Blueprint, render_template

experts = Blueprint('experts', __name__, template_folder='templates', url_prefix='/experts')

# default top timeline
@experts.route('/')
@experts.route('/management')
def get_management():
    return render_template("experts/management.html")