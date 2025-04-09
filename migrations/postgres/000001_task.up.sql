CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
    title TEXT NOT NULL,               
    description TEXT,                 
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new' ,
	FOREIGN KEY (user_id) REFERENCES users
);