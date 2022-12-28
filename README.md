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
docker run --rm --name activity-2-ex-5 -p 8000:10000 -e HOST_AND_PORT=0.0.0.0:10000 decima/mds-activity-2:ex-5 serve
```