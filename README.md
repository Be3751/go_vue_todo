# go_vue_todo
Gin, Vue.js, MySQLによるToDoアプリ  
Web APIを介してマイクロサービス間を疎結合に連携させ、セッションを用いた認証機能やXSS・CSRF対策を実装  

 - XSS対策：Set-CookieヘッダにHttpOnly属性を付与し、JavaScriptからのクッキーへのアクセスを禁止  
 - CSRF対策：Set-CookieヘッダにSameSite属性のStrictを付与し、異なるドメインにあるホストとのクッキーの送受信を禁止  

# 使用技術
Docker環境上で以下3つに対応したコンテナを用意し、docker-composeでスタック
- フロントエンド：Vuetify（Vue.js）
- バックエンド：Gin（Go）
- データベース：MySQL

# デザイン
・タスク一覧画面
<img width="1440" alt="Screen Shot 2022-04-03 at 20 47 32" src="https://user-images.githubusercontent.com/49334354/161426472-f91dfbd0-cde3-4ab6-9ba7-7112a818423c.png">

・タスク作成画面
<img width="1440" alt="Screen Shot 2022-04-03 at 20 47 45" src="https://user-images.githubusercontent.com/49334354/161426487-ce338cc1-dbdb-4455-8bda-6fea3530cc24.png">

# 使用方法
起動：
`docker-compose up -d --build`  
停止：
`docker-compose down`  
