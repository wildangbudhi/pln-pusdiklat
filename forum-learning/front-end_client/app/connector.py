from flask import Flask
from module.homepage.controller import HomePageBlueprint
from module.qna.controller import QNABlueprint
# from module.discussions.routes import discussions
# from module.posts.routes import posts

def depedency_injection( app: Flask ):

    homepage_blueprint = HomePageBlueprint( app.auth )
    qna_blueprint = QNABlueprint( app.auth )

    app.register_blueprint( homepage_blueprint )
    app.register_blueprint( qna_blueprint )