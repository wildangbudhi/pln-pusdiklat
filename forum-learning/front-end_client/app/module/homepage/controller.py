from setup.auth import Auth
from flask import Blueprint, render_template, url_for, redirect

class HomePageBlueprint( Blueprint ):

    def __init__( self, auth: Auth ):
        super( HomePageBlueprint, self ).__init__( 'homepage', __name__, template_folder='templates' )
        self.auth = auth

        self.add_url_rule( "/", view_func=self.index )
    
    def index( self ):

        if( self.auth.is_user_logged_in() ):
            return redirect( "/qna" )
        else :
            return render_template( "homepage/signin.html" )