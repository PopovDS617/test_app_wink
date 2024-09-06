## Как запустить:

### Env

```
Переименовать .env-example в .env и заполнить актуальными данными
```

### Локальный запуск

```
1. Make install-deps
2. Make get-deps
3. Make start
4. Make restart
5. Make stop
```

### Запуск в контейнере

```
1. Make db
2. Make du
3. Make dd
```

### Отправка запросов на GRP сервер

```
grpcurl -plaintext -d '{"video": "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' \
:9000 lb_v1.LB_v1.GetVideoURL
```

### Нагрузочное тестирование

1. Простая отправка 100к запросов

```
ghz --insecure --call lb_v1.LB_v1.GetVideoURL -d  '{"video": "http://s{{randomInt 1 10}}.origin-cluster/video/{{randomInt 1 300}}/{{randomString 10}}{{randomInt 0 50}}.m3u8"}' -n 100000 :9000
```

2. Отправка 100к запросов с rps 15000

```
ghz --insecure --call lb_v1.LB_v1.GetVideoURL -d  '{"video": "http://s{{randomInt 1 10}}.origin-cluster/video/{{randomInt 1 300}}/{{randomString 10}}{{randomInt 0 50}}.m3u8"}' -n 100000 --rps 15000 :9000
```

3. Работа в течение 30 секунд под максимальной нагрузкой

```
ghz --insecure --call lb_v1.LB_v1.GetVideoURL -d  '{"video": "http://s{{randomInt 1 10}}.origin-cluster/video/{{randomInt 1 300}}/{{randomString 10}}{{randomInt 0 50}}.m3u8"}' -n 10000000 --duration=30s :9000
```
