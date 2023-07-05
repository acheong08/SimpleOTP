<script lang="ts">
  import { AddEntry } from "../wailsjs/go/main/App.js";
  import { database } from "../wailsjs/go/models.js";

  export let page_state = "addEntry";

  let entry: database.Entry = new database.Entry(); // Create an instance of the Entry class

  function submitEntry() {
    // Call the AddEntry function with the entry data
    AddEntry(entry);
    // Reset the entry form
    entry = new database.Entry();
  }
</script>

<div class="max-w-md mx-auto">
  <form class="my-5">
    <h2 class="text-5xl font-bold mb-6">Add OTP Details</h2>
    <div class="mb-4">
      <label class="block text-sm font-bold mb-2" for="name">Name:</label>
      <input
        class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        type="text"
        id="name"
        placeholder="Enter the name"
        bind:value={entry.name}
      />
    </div>

    <div class="mb-4">
      <label class="block text-sm font-bold mb-2" for="description"
        >Description:</label
      >
      <textarea
        class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        id="description"
        placeholder="Enter the description"
        bind:value={entry.description}
      />
    </div>

    <div class="mb-4">
      <label class="block text-sm font-bold mb-2" for="url">URL:</label>
      <input
        class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        type="text"
        id="url"
        placeholder="Enter the URL"
        bind:value={entry.url}
      />
    </div>

    <div class="mb-4">
      <label class="block text-sm font-bold mb-2" for="secret">Secret:</label>
      <input
        class="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        type="password"
        id="secret"
        placeholder="Enter the secret"
        bind:value={entry.secret}
      />
    </div>

    <div class="flex justify-between">
      <button
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        type="button"
        on:click={submitEntry}
      >
        Submit
      </button>

      <button
        class="bg-warning hover:bg-error text-warning-content font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        type="button"
        on:click={() => (page_state = "dashboard")}
      >
        Cancel
      </button>
    </div>
  </form>
</div>
