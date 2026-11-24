# HTTP-Контракты

1) GET /watch

* Watch отправляет запрос в Studio: GET /video

``` json
    Request: None

        Response: 
    [
    {
        "id": 2,
        "channel_id": 1,
        "title": "Самое популярное видео",
        "description": "Описание",
        "views": 1500000,
        "likes": 85000
    },
    {
        "id": 1,
        "channel_id": 1,
        "title": "Мой первый влог",
        "description": "Описание ролика",
        "views": 120,
        "likes": 15
    }
    ]
```

2) GET /watch/{channel_id}

* Watch отправляет запрос в Studio: GET /video/{channel_id}

``` json
    Request: None

    Response: [
    {
        "id": 1,
        "channel_id": 1,
        "title": "Мой первый влог",
        "description": "Описание ролика",
        "views": 120,
        "likes": 15
    }
    ]
```

3) GET /watch/{channel_id}/{video_id} 

* Watch отправляет запрос в Studio: GET /video/{channel_id}/{video_id}

``` json 
    Request: None

    Response: 
    {
    "id": 1,
    "channel_id": 1,
    "title": "Мой первый влог",
    "description": "Описание ролика",
    "views": 120,
    "likes": 15
    }
```

4) POST /watch/likes/{channel_id}/{video_id}

* Watch отправляет запрос в Studio: POST /video/like/{channel_id}/{video_id}

``` json
    Request: None

    Response: 
    {
        "success": true
    }  
``` 

5) POST /watch/subscribe/{channel_id}

* Watch отправляет запрос в Studio: POST /channel/subscribe/{channel_id}

``` json
    Request: None

    Response: 

    {
        "success": true
    }
```

