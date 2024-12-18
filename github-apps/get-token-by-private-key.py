# pip install PyJWT requests cryptography

import jwt
import time
import requests
import os

# 環境変数から設定を取得
app_id = os.getenv('GITHUB_APP_ID')
installation_id = os.getenv('GITHUB_APP_INSTALLATION_ID')
private_key_path = os.getenv('GITHUB_APP_PRIVATE_KEY')

# プライベートキーをファイルから読み込む
with open(private_key_path, 'r') as key_file:
    private_key = key_file.read()

# JWTの生成
def generate_jwt(app_id, private_key):
    payload = {
        'iat': int(time.time()),
        'exp': int(time.time()) + (10 * 60),
        'iss': app_id
    }
    token = jwt.encode(payload, private_key, algorithm='RS256')
    return token

# アクセストークンの取得
def get_access_token(jwt_token, installation_id):
    headers = {
        'Authorization': f'Bearer {jwt_token}',
        'Accept': 'application/vnd.github.v3+json'
    }
    url = f'https://api.github.com/app/installations/{installation_id}/access_tokens'
    response = requests.post(url, headers=headers)
    response.raise_for_status()
    return response.json()['token']

# JWTの生成
jwt_token = generate_jwt(app_id, private_key)

# アクセストークンの取得
access_token = get_access_token(jwt_token, installation_id)

print(f'{access_token}')
