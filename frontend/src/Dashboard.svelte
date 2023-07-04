<script>
  let searchQuery = "";
  let cards = [
    {
      name: "Card 1",
      url: "https://example.com/card1",
      description: "This is the description for Card 1",
    },
    {
      name: "Card 2",
      url: "https://example.com/card2",
      description: "This is the description for Card 2",
    },
    // Add more cards here...
  ];
  let showSettings = false;

  function search() {
    // Perform search logic here
    // You can update the 'cards' array based on the search results
    console.log("Searching for:", searchQuery);
  }

  function toggleSettings() {
    showSettings = !showSettings;
  }
</script>

<div class="p-4 m-auto">
  <div class="flex justify-between items-center">
    <input
      type="text"
      placeholder="Search..."
      class="w-64 px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
      bind:value={searchQuery}
    />
    <button
      class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
      on:click={search}
    >
      Search
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

<div class="p-4 flex justify-between items-center">
  <button
    class="py-2 px-2 focus:outline-none bg-neutral rounded-md hover:bg-primary"
    on:click={toggleSettings}
  >
    {#if showSettings}
      Hide Settings
    {:else}
      Show Settings
    {/if}
  </button>
  <button
    class="px-4 py-2 focus:outline-none hover:bg-primary bg-neutral"
    on:click={() => alert("Home button clicked!")}
  >
    Home
  </button>
</div>

{#if showSettings}
  <div class="p-4">
    <!-- Add your settings content here -->
  </div>
{/if}
