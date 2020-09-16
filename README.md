ローカル環境の構築とアプリケーションの実行手順：
 
0. golang docker git をinstallしてください

0. このリポジトリをローカルにクローンします
0. otenkiディレクトリへ移動
0. docker-compose pull
0. docker-compose up --build 
* 以降の作業は別の窓を立ち上げて行なってください

0. migrationの実行（スキーマ定義とseedのマイグレーション）
    0. `docker-compose exec migration /bin/sh`
    0. `goose -env local up` でマイグレーションを実行
    0. `exit` でコンテナからでる
    0. 下記のコマンドでseedingを実行（順番をこの通りやること。PWはpostgres）
    0. `psql -p 5432 -h localhost -d otenki -U postgres -c '\copy hourly_weather_infos from db/otenki.csv with csv header'`
    0. `psql -p 5432 -h localhost -d otenki -U postgres -c ' \copy weathers from db/otenki-weather.csv with csv header'`

0. fetcher-appの起動

    0. docker-compose exec fetcher-app /bin/sh
    
    0. fresh -c fresh.conf（ホットリロード環境が立ち上がります）
    
    0. localhost:8000へアクセス
    
0. servicer-appの起動

    0. docker-compose exec servicer-app /bin/sh
    
    0. fresh -c fresh.conf（ホットリロード環境が立ち上がります）
    
    0. localhost:8001へアクセス
    
0. バッチの手動実行

    0. localhost:8000/batch/hourly_weather をブラウザで叩くと直近1時間の天気が取得できます。
    
0. 気象情報の取得

    0. localhost:8000/batch/hourly_weather?from=hoge&to=fuga、unixtimeでfromからtoまでの時間の気象情報がCSV形式で吐き出されます。
        `例： localhost:8000/batch/hourly_weather?from=1599976800&to=1600182000`

* ローカルDBへの接続:

    `psql -p 5432 -h localhost -d otenki -U postgres`
    
    PW: postgres
