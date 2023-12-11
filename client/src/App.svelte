<script>
    import { Router, Route } from 'svelte-routing';
    import routes from './routers/router.js';
  </script>
  
  <Router>
    {#each Object.entries(routes) as [path, Component]}
      <Route path={path}>
        <Component/>
      </Route>
    {/each}
  </Router>
  