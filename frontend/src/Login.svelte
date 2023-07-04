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

<div class="card place-self-center bg-neutral h-screen w-screen">
  <div class="container m-5">
    <h1 class="text-5xl font-bold my-1">Login</h1>
    <div class="pr-9">
      <PasswordInput
        labelText="Password"
        placeholder="Enter password..."
        bind:value={password}
      />
    </div>
    <button class="btn btn-primary my-2" on:click={login}>Submit</button>
  </div>
</div>
