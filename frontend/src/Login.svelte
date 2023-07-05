<script lang="ts">
  import { Login } from "../wailsjs/go/main/App.js";
  import { PasswordInput } from "carbon-components-svelte";

  let password: string;

  function login(): void {
    Login(password).then((result: string) => {
      console.log(result);
      if (result == "success") {
        page_state = "dashboard";
      } else {
        alert("Incorrect password");
      }
    });
  }
  export let page_state = "login";
</script>

<div class="login-container">
  <div class="max-h-md">
    <h1 class="text-5xl font-bold mb-6">Login</h1>
    <div class="pr-9">
      <PasswordInput
        labelText="Password"
        placeholder="Enter password..."
        bind:value={password}
        on:change={login}
      />
    </div>
    <button class="btn btn-primary mt-4" on:click={login}>Submit</button>
  </div>
</div>

<style>
  .login-container {
    max-width: 30rem;
    margin: auto;
    margin-top: 4rem; /* Adjust the value to increase or decrease the top margin */
  }
</style>
