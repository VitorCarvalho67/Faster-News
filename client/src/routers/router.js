import Home from '../pages/Home.svelte';
import About from '../pages/About.svelte';
import Cadastro from '../pages/Cadastro.svelte';
import Login from '../pages/Login.svelte';
import Noticia from '../pages/Noticia.svelte';

export default {
    '/': Home,
    '/about': About,
    '/cadastro': Cadastro,
    '/login': Login,
    '/noticia': Noticia
};