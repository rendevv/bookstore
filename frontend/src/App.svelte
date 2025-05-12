<svelte:head>
  <title>Bookstore CRUD</title>
  <meta name="description" content="A simple bookstore CRUD application." />
</svelte:head>

<script lang="ts">
  import axios from 'axios';
  import { onMount } from 'svelte';

  interface Book {
      id: number;
      name: string;
      author: string;
      publication: string;
  }

  let books: Book[] = [];
  let book: Partial<Book> = {};
  let isEditMode = false;

  const apiUrl = 'http://localhost:8080/book/';

  const fetchBooks = async () => {
      try {
          const response = await axios.get(apiUrl);
          books = response.data;
      } catch (error) {
          console.error('Error fetching books:', error);
      }
  };

  const saveBook = async () => {
      try {
          if (isEditMode && book.id) {
              await axios.put(`${apiUrl}/${book.id}`, book);
          } else {
              await axios.post(apiUrl, book);
          }
          book = {};
          isEditMode = false;
          fetchBooks();
      } catch (error) {
          console.error('Error saving book:', error);
      }
  };

  const editBook = (b: Book) => {
      book = { ...b };
      isEditMode = true;
  };

  const deleteBook = async (id: number) => {
      try {
          await axios.delete(`${apiUrl}/${id}`);
          fetchBooks();
      } catch (error) {
          console.error('Error deleting book:', error);
      }
  };

  onMount(fetchBooks);
</script>

<div class="max-w-4xl mx-auto p-6">
  <h2 class="text-3xl font-semibold text-center mb-8 text-gray-800">Bookstore CRUD</h2>
  <div class="bg-white shadow-lg rounded-lg p-6 mb-8">
      <input bind:value={book.name} placeholder="Book Name" class="border border-gray-300 rounded-lg p-4 w-full mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      <input bind:value={book.author} placeholder="Author" class="border border-gray-300 rounded-lg p-4 w-full mb-4 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      <input bind:value={book.publication} placeholder="Publication" class="border border-gray-300 rounded-lg p-4 w-full mb-6 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      <button on:click={saveBook} class="w-full bg-blue-600 text-white py-3 rounded-lg text-lg hover:bg-blue-700 focus:outline-none transition duration-200">
          {isEditMode ? 'Update Book' : 'Save Book'}
      </button>
  </div> 

  <div class="overflow-x-auto bg-white shadow-lg rounded-lg">
    <table class="w-full text-left border-collapse table-auto">
        <thead class="bg-blue-600 text-white">
            <tr>
                <th class="border-b p-4">Name</th>
                <th class="border-b p-4">Author</th>
                <th class="border-b p-4">Publication</th>
                <th class="border-b p-4">Actions</th>
            </tr>
        </thead>
        <tbody>
            {#each books as book}
                <tr class="hover:bg-gray-100 transition duration-150">
                    <td class="border-b p-4">{book.name}</td>
                    <td class="border-b p-4">{book.author}</td>
                    <td class="border-b p-4">{book.publication}</td>
                    <td class="border-b p-4 space-x-2">
                        <button on:click={() => editBook(book)} class="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600 transition duration-200">Edit</button>
                        <button on:click={() => deleteBook(book.id)} class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition duration-200">Delete</button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
  </div>
</div>

<style>
  input:focus {
    border-color: #2563eb;
  }
  button:focus {
    outline: none;
  }
  table th {
    text-transform: uppercase;
    letter-spacing: 1px;
  }
</style>
