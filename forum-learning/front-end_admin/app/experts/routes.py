from flask import Blueprint, render_template
import urllib.parse

experts = Blueprint('experts', __name__, template_folder='templates', \
    url_prefix='/experts')

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
@experts.route('/')
@experts.route('/management')
def get_management():
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    return render_template("experts/management.html", \
        category_dict=category_dict)

@experts.route('/management/<url_filter>')
def get_management_by_filter(url_filter):
    url_category = [urllib.parse.quote(c) for c in category]
    category_dict = {c:uc for c, uc in zip(category, url_category)}
    filter = urllib.parse.unquote(url_filter)
    return render_template("experts/management_by_filter.html", \
        filter=filter, category_dict=category_dict)