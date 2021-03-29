from flask import url_for
from flask_saml2.sp import ServiceProvider
from jwt import encode
from typing import Tuple
from datetime import datetime, timedelta

class Auth(ServiceProvider):

    def __init__(self, api_secret_key: str):
        self.api_secret_key = api_secret_key

    def get_logout_return_url( self ):
        return url_for( 'homepage.index', _external=True )

    def get_default_login_return_url( self ):
        return url_for( 'homepage.index', _external=True )
    
    def get_user_auth_data( self ) -> dict:
        auth_data = self.get_auth_data_in_session()
        return auth_data.to_dict()["data"]["attributes"]

    def get_api_access_token( self ) -> str:

        if( not self.is_user_logged_in() ):
            raise Exception( "You Don't Have Permission" )

        auth_data = self.get_auth_data_in_session()
        payload = auth_data.to_dict()["data"]["attributes"]

        payload[ 'exp' ] = datetime.utcnow() + timedelta( days=1 )

        self.api_access_token = encode( payload=payload, key=self.api_secret_key, algorithm="HS512" )

        return self.api_access_token


