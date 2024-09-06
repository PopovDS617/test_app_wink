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

```
ghz --insecure --call lb_v1.LB_v1.GetVideoURL -d '{"video": "http://s{{randomInt 1 6}}.origin-cluster/video/{{randomInt 1 11}}/pl{{randomInt 0 1000}}.m3u8"}' -n 10000 :9000
```
