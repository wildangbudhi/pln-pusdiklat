from setup.auth import Auth
from flask import Blueprint, render_template, redirect, url_for
import urllib.parse

# temp data that should be fetched
fetched_category = {
  "category": [
    {
      "id": 1,
      "category_name": "Pembangkitan"
    },
    {
      "id": 2,
      "category_name": "Transmisi"
    },
    {
      "id": 3,
      "category_name": "Distribusi"
    },
    {
      "id": 4,
      "category_name": "Niaga dan Manajemen Pelanggan"
    },
    {
      "id": 5,
      "category_name": "Produksi Peralatan Ketenagalistrikan"
    },
    {
      "id": 6,
      "category_name": "K2, K3, Keamanan dan Lingkungan"
    },
    {
      "id": 7,
      "category_name": "Manajemen Proyek, Enjiniring (Engineering) dan Konstruksi"
    },
    {
      "id": 8,
      "category_name": "Penelitian dan Pengembangan"
    },
    {
      "id": 9,
      "category_name": "Pembelajaran"
    },
    {
      "id": 10,
      "category_name": "Sertifikasi"
    },
    {
      "id": 11,
      "category_name": "Supply Chain Management"
    },
    {
      "id": 12,
      "category_name": "Regulatory and Compliance"
    },
    {
      "id": 13,
      "category_name": "Teknologi Informasi"
    },
    {
      "id": 14,
      "category_name": "SDM"
    },
    {
      "id": 15,
      "category_name": "Keuangan"
    },
    {
      "id": 16,
      "category_name": "Komunikasi, CSR dan Pengelolaan Kantor"
    },
    {
      "id": 17,
      "category_name": "Manajemen Perusahaan"
    },
    {
      "id": 18,
      "category_name": "Bebas"
    }
  ]
}
fetched_qna = {
  "forum": [
    {
      "id": "41c2a806-a2d0-4d6f-9b04-3dfd98dbd441",
      "title": "Bagaimana Cara Membuka Kaleng Sarden",
      "question": "Saya sudah beli Sarden Kalengan, Tapi saya tidak bisa membukan nya",
      "author_user_id": 2,
      "author_full_name": "Wildan G Budhi",
      "author_username": "0511174000184",
      "status": "TERBUKA",
      "category_id": 4,
      "category_name": "Niaga dan Manajemen Pelanggan",
      "up_vote": 0,
      "down_vote": 0,
      "is_up_voted": False,
      "is_down_voted": False,
      "replies_count": 0
    },
    {
      "id": "a8948c1a-b47f-4bda-872b-bfd449145379",
      "title": "Bagaimana Cara Memasang Baut",
      "question": "Saya mengalami masalah dalam memperbaiki sepeda saya, ketika ingin memasang baut pada lampu sepeda bautnya tidak bisa di pasang padahal saya sudah memutarnya berlawanan arah jarum jam.",
      "author_user_id": 1,
      "author_full_name": "Rangga Kusuma Dinata",
      "author_username": "0511174000120",
      "status": "TERBUKA",
      "category_id": 18,
      "category_name": "Bebas",
      "up_vote": 2,
      "down_vote": 1,
      "is_up_voted": False,
      "is_down_voted": True,
      "replies_count": 3
    }
  ]
}


class QNABlueprint( Blueprint ):

	def __init__( self, auth: Auth ):
		super( QNABlueprint, self ).__init__( 'qna', __name__, template_folder='templates', url_prefix='/qna' )
		self.auth = auth

		self.add_url_rule( "/", view_func=self.index )
	
	def index( self ):

		if( not self.auth.is_user_logged_in() ):
			login_url = url_for('flask_saml2_sp.login')
			return redirect( login_url )
		
		user_auth_data = self.auth.get_user_auth_data()

		category = [c["category_name"] for c in fetched_category["category"]]
		url_category = [urllib.parse.quote(c) for c in category]
		category_dict = {c:uc for c, uc in zip(category, url_category)}
		
		qna_dict = { 
			urllib.parse.quote("/".join((q['author_username'], q['id']))):q for q in fetched_qna["forum"]
		}

		render_data = {
			"user_auth_data" 	: user_auth_data,
			"category_dict" 	: category_dict,
			"qna_dict"			: qna_dict
		}

		return render_template( "qna/timeline.html", **render_data )

# @qna.route('/<string:username>/<string:qna_id>')
# def user_question(username, qna_id):
#     return render_template('qna/user_question.html', username=username, qna_id=qna_id)