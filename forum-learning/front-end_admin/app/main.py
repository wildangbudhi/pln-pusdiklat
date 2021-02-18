from flask import Flask, render_template
from homepage.routes import homepage
from qna.routes import qna
from discussions.routes import discussions
from posts.routes import posts
from experts.routes import experts

app = Flask(__name__)

app.register_blueprint(homepage)
app.register_blueprint(qna)
app.register_blueprint(discussions)
app.register_blueprint(posts)
app.register_blueprint(experts)
