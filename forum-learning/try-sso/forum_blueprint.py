from auth import Auth
from flask import Blueprint

class ForumBlueprint( Blueprint ):

    def __init__( self, auth: Auth ):
        super(ForumBlueprint, self).__init__( 'forum', __name__, url_prefix="/forum/" )
        self.auth = auth
        
        self.add_url_rule('/', view_func=self.index)
    
    def index( self ):
        if self.auth.is_user_logged_in():
            return "<h1>Welcome to Forum Learning</h1>"
        else:
            return "<h1>You don't have permission</h1>"