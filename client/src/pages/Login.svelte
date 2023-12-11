<script>
    import { onMount } from 'svelte';
    import { GetNews } from '../services/api';
  
    let dados = null;
  
    onMount(async () =>{
      try {
        dados = await GetNews();
      } catch (error) {
        console.log(error);
      }
    });
  </script>
  
  <main>
    <h1>Login</h1>
      {#if dados}
          {#each dados as dado}
          <div>
              <h2>{dado.title}</h2>
              <p>{dado.body}</p>
          </div>
          {/each}
      {/if}
  </main>
  
  <style>
  
  </style>
  