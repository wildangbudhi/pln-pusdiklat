from flask import Flask, render_template
from homepage.routes import homepage
from qna.routes import qna
from discussions.routes import discussions
from posts.routes import posts

app = Flask(__name__)

if __name__ == "__main__":
    app.register_blueprint(homepage)
    app.register_blueprint(qna)
    app.register_blueprint(discussions)
    app.register_blueprint(posts)
    app.run(host="0.0.0.0",port=8000, debug=True)