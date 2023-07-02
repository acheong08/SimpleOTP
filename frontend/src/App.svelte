<script lang="ts">
  import Login from "./Login.svelte";
  import Dashboard from "./Dashboard.svelte";

  let page_state = "login";
</script>

<main class="">
  {#if page_state == "login"}
    <Login bind:page_state />
  {:else if page_state == "dashboard"}
    <Dashboard />
  {/if}
</main>
