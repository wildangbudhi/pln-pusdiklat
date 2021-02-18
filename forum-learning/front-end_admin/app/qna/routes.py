from flask import Blueprint, render_template
import urllib.parse

qna = Blueprint('qna', __name__, template_folder='templates', \
    url_prefix='/qna')

# temp data that should be fetched
category = [
    "Generation",
    "Transmission",
    "Distribution",
    "Commerce & Customer Management",
    "Electricity Equipment Production",
    "Electric Safety, OHS, Security & Environment",
    "Project Management, Engineering & Construction",
    "Research & Development",
    "Learning",
    "Certification",
    "Supply Chain Management",
    "Regulatory & Compliance",
    "Information Technology",
    "HR",
    "Finance",
    "Communication, CSR & Office Management",
    "Company Management",
    "Miscellaneous",
]

# all category management, default category = Generation
@qna.route('/')
@qna.route('/management')
def get_management():
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    return render_template("qna/management.html", category_dict=category_dict)

@qna.route('/management/<url_filter>')
def get_management_by_filter(url_filter):
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    filter = urllib.parse.unquote(url_filter)
    return render_template("qna/management_by_filter.html", filter=filter, \
        category_dict=category_dict)

@qna.route('/replies/<string:qna_id>')
def get_replies_management(qna_id):
    return render_template('/qna/replies.html', qna_id=qna_id)