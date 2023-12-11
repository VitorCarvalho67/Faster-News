<script>
    import { onMount } from "svelte";
    import { GetNews } from "../services/api";
    import Header from "../components/Header.svelte";
    import logo from "../assets/Faster_News_Logo.png";
    let dados = null;

    onMount(async () => {
        try {
            dados = await GetNews();
        } catch (error) {
            console.log(error);
        }
    });

    const date = new Date().toLocaleDateString("pt-BR");
</script>

<main>
    <Header />
    <div class="main">
        <div class="box">
            <div>
                <h1>Top 10 Notícias</h1>
                <!-- <p>No dia {date}:</p> -->
                <h2>do dia {date}, de acordo com o Google News, são:</h2>
                <h2>
                    <a href="/news">Ver mais</a>
                </h2>
            </div>
        </div>
        <div class="box">
            <img src={logo} alt="logo" />
        </div>
    </div>
    <div class="TopNews">
        {#if dados}
            {#each dados as dado}
                <div class="News">
                    <img src="#" alt="">
                    <div>
                        <h2>{dado.title}</h2>
                        <p>{dado.body}</p>
                        <p>Data:</p>
                        <p>Fonte:</p>
                    </div>
                </div>
            {/each}
        {/if}
    </div>
</main>

<style>
    .main {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        height: 500px;
    }

    .main .box {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        height: 100%;
        width: 50%;
        padding: 0 100px;
    }

    .main .box div {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        justify-content: center;
        height: 100%;
        width: 60%;
    }

    main .box h1 {
        font-size: 50px;
        font-weight: 500;
        color: var(--textcolor);
        margin-bottom: 20px;
    }

    main .box h2 {
        font-size: 30px;
        font-weight: 500;
        color: var(--textcolor);
        margin-bottom: 20px;
    }

    .main .box div a {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: flex-start;
        padding: 10px 20px;
        border-radius: 50px;
        text-decoration: none;
        font-size: 20px;
        font-weight: 500;
        background-color: transparent;
        color: var(--primarycolor);
        border: solid var(--primarycolor) 1px;
        cursor: pointer;
    }

    .main .box div a:hover {
        color: var(--textcolor);
        background-color: var(--primarycolor);
        border: solid transparent 1px;
    }

    .main .box img {
        width: 60%;
    }

    .TopNews {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: flex-start;
        width: 100%;
        padding: 20px 100px;
        background-color: #0000000c;
    }

    .TopNews .News {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: flex-start;
        height: 100%;
        width: 100%;
        padding: 20px;
        margin: 20px 0;
        border-radius: 20px;
    }

    .TopNews .News img {
        width: 200px;
        height: 200px;
        object-fit: cover;
        margin-bottom: 20px;
        border-radius: 20px;
        border: solid var(--primarycolor) 1px;
    }

    .TopNews .News div {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: flex-start;
        height: 100%;
        width: 100%;
        padding: 0 20px;
    }
</style>
