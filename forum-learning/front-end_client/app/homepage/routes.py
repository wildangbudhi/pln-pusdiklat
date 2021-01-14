from flask import Blueprint, render_template

homepage = Blueprint('homepage', __name__, template_folder='templates')

@homepage.route('/')
@homepage.route('/signin')
def signin():
    return render_template("homepage/signin.html")