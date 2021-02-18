from flask import Blueprint, render_template
import urllib.parse

discussions = Blueprint('discussions', __name__, template_folder='templates', \
    url_prefix='/discussions')

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
@discussions.route('/')
@discussions.route('/management')
def get_management():
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    return render_template("discussions/management.html", \
        category_dict=category_dict)

@discussions.route('/management/<url_filter>')
def get_management_by_filter(url_filter):
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    filter = urllib.parse.unquote(url_filter)
    return render_template("discussions/management_by_filter.html", \
        filter=filter, category_dict=category_dict)

@discussions.route('/request')
def get_request_management():
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    return render_template('/discussions/request.html', \
        category_dict=category_dict)

@discussions.route('/request/<url_filter>')
def get_request_management_by_filter(url_filter):
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    filter = urllib.parse.unquote(url_filter)
    return render_template("discussions/request_by_filter.html", \
        filter=filter, category_dict=category_dict)