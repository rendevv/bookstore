import { Book } from "../types/booktype";

const BASE_URL = import.meta.env.BOOK_URL;

export const getBooks = async () => {
  const res = await fetch(`${BASE_URL}/`);
  return res.json();
};

export const getBookById = async (id: string) => {
  const res = await fetch(`${BASE_URL}/${id}`);
  return res.json();
};

export const createBook = async (book: Book) => {
  const res = await fetch(`${BASE_URL}/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(book),
  });
  return res.json();
};

export const updateBook = async (id: string, book: Book) => {
  const res = await fetch(`${BASE_URL}/${id}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(book),
  });
  return res.json();
};

export const deleteBook = async (id: string) => {
  const res = await fetch(`${BASE_URL}/${id}`, { method: "DELETE" });
  return res.json();
};
