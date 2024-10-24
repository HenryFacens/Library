-- Use ou crie o keyspace (banco de dados) chamado 'library'
CREATE KEYSPACE IF NOT EXISTS library WITH replication = { 'class': 'SimpleStrategy',
'replication_factor': 1 };
-- Use o keyspace criado
USE library;
-- Criar a tabela 'users' para armazenar usuários
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email text,
    name text
);
-- Criar a tabela 'books' para armazenar livros
CREATE TABLE IF NOT EXISTS books (
    id uuid PRIMARY KEY,
    author text,
    genre text,
    publish_year int,
    title text
);
-- Criar a tabela 'borrowed_books' para registrar empréstimos
CREATE TABLE IF NOT EXISTS borrowed_books (
    book_id uuid,
    user_id uuid,
    borrow_date date,
    PRIMARY KEY (book_id, user_id)
);
-- Inserindo um usuário na tabela 'users'
INSERT INTO users (id, email, name)
VALUES (
        48d32eb7-700e-43a7-b4db-9a8737d4c946,
        'john.doe@example.com',
        'John Doe'
    );
-- Inserindo um livro na tabela 'books'
INSERT INTO books (id, author, genre, publish_year, title)
VALUES (
        3d2f0b2f-63e3-4a2f-9e12-c1a0efb1fa5e,
        'George Orwell',
        'Dystopian',
        1949,
        '1984'
    );
-- Registrando que o usuário pegou um livro emprestado
INSERT INTO borrowed_books (book_id, user_id, borrow_date)
VALUES (
        3d2f0b2f-63e3-4a2f-9e12-c1a0efb1fa5e,
        48d32eb7-700e-43a7-b4db-9a8737d4c946,
        '2024-10-24'
    );