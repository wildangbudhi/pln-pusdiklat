from flask import Flask, render_template
from setup.server import Server
from connector import depedency_injection

app = Server( __name__ )
app.config["TEMPLATES_AUTO_RELOAD"] = True
depedency_injection( app )

