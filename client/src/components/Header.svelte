<script>
    import { Link } from 'svelte-routing';
    import Logo from '../assets/Logo.png';
</script>
  
<header>
    <div>
        <img src={Logo} alt="logo" />
    </div>
    <nav>
        <ul>
            <li>
                <a class="btn" href="/login">Login / Cadastro</a>
            </li>
            <li>
                <a class="a" href="/about">Sobre</a>
            </li>
            <li>
                <a class="a" href="/noticia">Notícias</a>
            </li>
        </ul>
    </nav>
</header>

<style>
    @import url(../app.css);

    header{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 70px;
        padding: 0 300px;
        width: 100vw;
        background-color: var(--bgcolor);
    }

    header div{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 70%;
    }

    header div img{
        height: 100%;
    }

    header nav{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 100%;
    }

    header nav ul{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 100%;
    }

    header nav ul li{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 100%;
        margin-left: 20px;
    }

    header nav ul li .a{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 100%;
        color: var(--textcolor);
        text-decoration: none;
        font-size: 20px;
        font-weight: 500;
    }

    header nav ul li .a:hover{
        color: var(--primarycolor);
    }

    header nav ul li .a:active{
        color: var(--primarycolor);
    }

    header nav ul li .btn{
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        padding: 10px 20px;
        border-radius: 50px;
        color: #cdd6f4;
        text-decoration: none;
        font-size: 20px;
        font-weight: 500;
        background-color: var(--primarycolor);
        border: none;
        cursor: pointer;
        border: solid transparent 1px;
    }

    header nav ul li .btn:hover{
        background-color: transparent;
        color: var(--primarycolor);
        border: solid var(--primarycolor) 1px;
    }
</style>
