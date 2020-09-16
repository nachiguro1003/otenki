ローカル環境の構築とアプリケーションの実行手順：
 
0. golang / docker / git をinstallしてください

1. このリポジトリをローカルにクローンします
1. otenkiディレクトリへ移動
1. `./src/fetcher/env.yaml`のapi_keyに`b253c507227e58ee1cc5e3a440e5c6b4`を追加
1. docker-compose pull
1. docker-compose up --build 
* 以降の作業は別の窓を立ち上げて行なってください

1. migrationの実行（スキーマ定義とseedのマイグレーション）
    1. `docker-compose exec migration /bin/sh`
    1. `goose -env local up` でマイグレーションを実行
    1. `exit` でコンテナからでる
    1. 下記のコマンドでseedingを実行（順番をこの通りやること。PWはpostgres）
    1. `psql -p 5432 -h localhost -d otenki -U postgres -c '\copy hourly_weather_infos from db/otenki.csv with csv header'`
    1. `psql -p 5432 -h localhost -d otenki -U postgres -c ' \copy weathers from db/otenki-weather.csv with csv header'`

1. fetcher-appの起動

    1. docker-compose exec fetcher-app /bin/sh
    
    1. fresh -c fresh.conf（ホットリロード環境が立ち上がります）
    
    1. localhost:8000へアクセス
    
1. servicer-appの起動

    1. docker-compose exec servicer-app /bin/sh
    
    1. fresh -c fresh.conf（ホットリロード環境が立ち上がります）
    
    1. localhost:8001へアクセス
    
1. バッチの手動実行

    1. localhost:8000/batch/hourly_weather をブラウザで叩くと直近1時間の天気が取得できます。
    
1. 気象情報の取得

    1. localhost:8000/batch/hourly_weather?from=hoge&to=fuga、unixtimeでfromからtoまでの時間の気象情報がCSV形式で吐き出されます。
        `例： localhost:8000/batch/hourly_weather?from=1599976800&to=1600182000`

* ローカルDBへの接続:

    `psql -p 5432 -h localhost -d otenki -U postgres`
    
    PW: postgres
