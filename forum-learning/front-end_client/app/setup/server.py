from flask import Flask, url_for
from flask_saml2.utils import certificate_from_file, private_key_from_file
from .auth import Auth
from os import getenv

class Server( Flask ):

    def __init__( self, import_name ):
        super(Server, self).__init__(import_name)
        self.read_config()
        self.setup_auth()

    def read_config( self ):
        self.config['SECRET_KEY'] = getenv("SECRET_KEY")
        self.config['API_SECRET_KEY'] = getenv("API_SECRET_KEY")
        # self.config['SERVER_NAME'] = getenv("SERVER_NAME")
        self.config['API'] = getenv("API")

    def setup_auth( self ):
        
        IDP_CERTIFICATE = certificate_from_file( "./sso-config/idp.crt" )
        SP_CERTIFICATE = certificate_from_file( "./sso-config/sp.crt" )
        SP_PRIVATE_KEY = private_key_from_file( "./sso-config/sp.pem" )

        self.auth = Auth( api_secret_key=self.config['API_SECRET_KEY'] )
        self.register_blueprint( self.auth.create_blueprint(), url_prefix='/saml/' )

        self.config['SAML2_SP'] = {
            'certificate': SP_CERTIFICATE,
            'private_key': SP_PRIVATE_KEY,
            'entity_id' : 'forumlearning.pln.co.id'
        }

        self.config['SAML2_IDENTITY_PROVIDERS'] = [
            {
                'CLASS': 'flask_saml2.sp.idphandler.IdPHandler',
                'OPTIONS': {
                    'display_name': "pln-pusdiklat",
                    'entity_id': 'https://idp.pln-pusdiklat.co.id',
                    'sso_url': 'https://elearning.pln-pusdiklat.co.id/sso/saml2/idp/SSOService.php',
                    'slo_url': 'https://elearning.pln-pusdiklat.co.id/sso/saml2/idp/SingleLogoutService.php',
                    'certificate': IDP_CERTIFICATE,
                },
            },
        ]