Описание: небольшое Create-Read приложение, реализованное на fiber + pgx
В нём реализуются следующие слои:

1) Репозиторий cmd содержит точку входа main, где парсится конфиг, инициализируются сервис, логгер и репозиторий.
2) Репозиторий internal содержит в себе api, config, dto, logger, repo, service
3) Репозиторий api содержит middleware с авторизацией, а также модель нашего роутера и его маршрутов
4) Репозиторий config содержит модели конфига, который делится на две составляющие: модель REST и PostgreSQL
5) Репозиторий dto содержит модели ответов и ошибок, а также их реализации
6) Репозиторий logger содержит реализацию создания сущности нашего логгера
7) Репозиторий repo содержит репозиторий с моками mocks, модель наших задач, а также файл repo.go, реализующий создание сущности репозитория, а также методы взаимодействия с базой данных (создание и чтение)
8) Репозиторий migrations содержит файлы для миграций: создание, удаление
9) Репозиторий viladator содержит реализацию авторизацию и тесты к ней

.env файл с переменными окружения


