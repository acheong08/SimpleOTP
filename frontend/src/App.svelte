<script lang="ts">
  import Login from "./Login.svelte";
  import Dashboard from "./Dashboard.svelte";
  import AddEntry from "./AddEntry.svelte";

  let page_state = "login";
</script>

<main class="min-h-screen bg-base-100">
  {#if page_state == "login"}
    <Login bind:page_state />
  {:else if page_state == "dashboard"}
    <Dashboard bind:page_state />
  {:else if page_state == "addEntry"}
    <AddEntry bind:page_state />
  {/if}
</main>
