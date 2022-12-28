# MDS - Activité 2

## Build all images

```bash
docker-compose build
```

## exercice 1

```bash
docker run --rm --name activity-2-ex-1 decima/mds-activity-2:ex-1
```

## exercice 2

```bash
docker run --rm --name activity-2-ex-2 decima/mds-activity-2:ex-2 Henri
```


## exercice 3

```bash
docker run --rm --name activity-2-ex-3 -v $PWD/motd.txt:/data/motd.txt decima/mds-activity-2:ex-3
```

## exercice 4

```bash
docker run --rm --name activity-2-ex-4 -p 8000:9000 decima/mds-activity-2:ex-4
```

## exercice 5

```bash
docker run --rm --name activity-2-ex-5 -p 8000:10000 -e HOST_AND_PORT=0.0.0.0:10000 decima/mds-activity-2:ex-5
```

## Exercice 6 - networking
```bash
docker network create chat
docker run --rm --name my-chat-server --network chat decima/mds-activity-2:server
docker run --rm -e SERVER=my-chat-server:9000  --network chat decima/mds-activity-2:client


```

## Exercice 7 - wordpress
```bash
docker network create wp
docker volume create wp
docker volume create bdd
docker run -d --rm --name bdd -e MYSQL_DATABASE=myblog -e MYSQL_ROOT_PASSWORD=somePassword -v bdd:/var/lib/mysql  --network wp mysql:8
docker run -d --rm --name wp -v wp:/var/www/html/wp-content -e WORDPRESS_DB_HOST=bdd:3306 -e WORDPRESS_DB_USER=root -e WORDPRESS_DB_PASSWORD=somePassword -e WORDPRESS_DB_NAME=myblog -p 12000:80 --network wp wordpress:6

```