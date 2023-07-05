<script lang="ts">
  import {
    List,
    Search,
    GenerateCode,
    DeleteEntry,
    SaveState,
  } from "../wailsjs/go/main/App.js";
  import { onMount } from "svelte";

  let searchQuery = "";
  let cards = [];

  onMount(async () => {
    cards = await List();
  });

  async function search() {
    // Wait for the search query to be updated
    await new Promise((resolve) => setTimeout(resolve, 0));
    if (searchQuery === "") {
      cards = await List();
      return;
    }
    cards = await Search(searchQuery);
    console.log(cards);
  }

  async function generateCode(name: string, event) {
    const div = event.target;
    const code = await GenerateCode(name);

    // Show the code in pre and hide the button
    div.nextElementSibling.textContent = code;
    div.nextElementSibling.style.display = "block";
    div.style.display = "none";

    setInterval(() => {
      div.nextElementSibling.style.display = "none";
      div.style.display = "block";
    }, 10000);
  }

  async function deleteEntry(name: string) {
    await DeleteEntry(name);
    cards = await List();
    SaveState();
  }

  function addEntryPage() {
    page_state = "addEntry";
  }

  export let page_state = "dashboard";
</script>

<div class="p-4 m-auto">
  <h2 class="text-5xl font-bold mb-6">SimpleOTP</h2>
  <div class="flex justify-between">
    <div class="flex items-center">
      <input
        type="text"
        placeholder="Search..."
        class="w-64 px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        bind:value={searchQuery}
        on:input={search}
      />
    </div>
    <div class="flex items-center mt-4">
      <button
        class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
        on:click={addEntryPage}
      >
        +
      </button>
    </div>
  </div>
</div>
<div class="p-4">
  {#each cards as card}
    <div class="mb-4 shadow-md p-4 card bg-neutral justify-between">
      <div class="flex">
        <h3 class="flex text-lg font-bold mb-2">{card.name}</h3>
        <div
          class="badge badge-error gap-1 flex justify ml-9"
          on:click={() => {
            deleteEntry(card.name);
          }}
          on:keypress={() => {}}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            class="inline-block w-4 h-4 stroke-current"
            ><path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            /></svg
          > delete
        </div>
      </div>
      <a
        href={card.url}
        target="_blank"
        class="text-blue-500 hover:underline mb-2">{card.url}</a
      >
      <p>{card.description}</p>

      <!-- A button to generate & show the code -->
      <div class="flex m-2 justify-center">
        <button
          class="flex bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 p-2"
          on:click={(event) => generateCode(card.name, event)}
        >
          Generate Code
        </button>
        <pre
          style="display:none;"
          class="flex bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 p-2 selection:bg-warning"
        />
      </div>
    </div>
  {/each}
</div>
