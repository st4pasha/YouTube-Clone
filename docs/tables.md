# Таблицы 

1) ```sql
    Table videos {
    id int
    title string
    description string
    views int
    likes int
    } 
    ```

2) ``` sql
    Table channels {
    id int
    channel_name string
    followers int
    } 
    ```

3) ``` sql
    Table channel_videos {
        channel_id int
        video_id int
    }
    ```
