# Image gallery service
For use with S3 & TynyPnG

# Требования
## Загрузка файла
1. Создать endpoint для загрузки файла через multipart\form-data
2. Создать таблицы: 
 - user
 - folder
 - file
 - user_foler
 - user_file
3. Создать endpoint для получения списка загруженных пользователем файлов
4. Подключить S3 библиотеку и грузить файлы в S3
5. Подключить TinyPNG и обрабатывать файлы через него