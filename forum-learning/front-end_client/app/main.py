from flask import Flask, render_template
from homepage.routes import homepage
from qna.routes import qna
from discussions.routes import discussions
from posts.routes import posts

app = Flask(__name__)
app.config["TEMPLATES_AUTO_RELOAD"] = True

app.register_blueprint(homepage)
app.register_blueprint(qna)
app.register_blueprint(discussions)
app.register_blueprint(posts)
