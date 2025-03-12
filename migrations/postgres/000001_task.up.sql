CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,             -- Уникальный идентификатор задачи
    title TEXT NOT NULL,               -- Заголовок задачи
    description TEXT,                  -- Описание задачи (необязательное поле)
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new' -- Статус задачи
);