from flask import url_for
from server import Server

from forum_blueprint import ForumBlueprint

app = Server(__name__)
app.debug = True

forum_blueprint = ForumBlueprint( app.auth )
app.register_blueprint( forum_blueprint )

@app.route('/')
def index():

    if app.auth.is_user_logged_in():

        api_token = ""

        try:
            api_token = app.auth.get_api_access_token()
        except Exception as e:
            print( e )


        auth_data = app.auth.get_auth_data_in_session()

        auth_data_dict = auth_data.to_dict()

        message = f'''
        <p>You are logged in as <strong>{auth_data.nameid}</strong>.
        The IdP sent back the following attributes:<p>
        '''

        attrs = '<dl>{}</dl>'.format(''.join(
            f'<dt>{attr}</dt><dd>{value}</dd>'
            for attr, value in auth_data.attributes.items()))

        logout_url = url_for('flask_saml2_sp.logout')
        logout = f'\n<form action="{logout_url}" method="POST"><input type="submit" value="Log out"></form>'

        token_preview = f'API TOKEN : {api_token}'

        return message + attrs + logout + token_preview
    else:

        message = '<p>You are logged out.</p>'

        login_url = url_for('flask_saml2_sp.login')
        link = f'<p><a href="{login_url}">Log in to continue</a></p>'

        return message + link

if __name__ == '__main__':
    app.run()
