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

<div class="container mx-auto flex justify-center items-center h-screen">
  <div class="card shadow-xl place-self-center w-2/3 h-1/2 bg-neutral">
    <div class="container m-5">
      <h1 class="text-5xl font-bold my-5">Login</h1>
      <div class="pr-8">
        <PasswordInput
          labelText="Password"
          placeholder="Enter password..."
          bind:value={password}
        />
      </div>
      <button class="btn btn-primary my-2" on:click={login}>Submit</button>
    </div>
  </div>
</div>
