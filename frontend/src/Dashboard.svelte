<script>
  import { List, Search } from "../wailsjs/go/main/App.js";
  import { onMount } from "svelte";
  let searchQuery = "";

  let cards = [];

  onMount(async () => {
    cards = await List();
  });

  async function search() {
    // Wait for the search query to be updated
    await new Promise((resolve) => setTimeout(resolve, 0));
    if (searchQuery == "") {
      cards = await List();
      return;
    }
    cards = await Search(searchQuery);
    console.log(cards);
  }

  function addEntryPage() {
    page_state = "addEntry";
  }

  export let page_state = "dashboard";
</script>

<div class="p-4 m-auto">
  <div class="flex justify-between items-center">
    <input
      type="text"
      placeholder="Search..."
      class="w-64 px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
      bind:value={searchQuery}
      on:input={search}
    />
  </div>
  <div class="flex justify-end items-center mt-4">
    <button
      class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500"
      on:click={addEntryPage}
    >
      +
    </button>
  </div>
</div>

<div class="p-4">
  {#each cards as card}
    <div class="mb-4 shadow-md p-4 card bg-neutral">
      <h3 class="text-lg font-bold mb-2">{card.name}</h3>
      <a
        href={card.url}
        target="_blank"
        class="text-blue-500 hover:underline mb-2">{card.url}</a
      >
      <p>{card.description}</p>
    </div>
  {/each}
</div>