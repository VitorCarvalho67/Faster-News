<script>
    import { onMount } from "svelte";
    import { GetNews } from "../services/api";
    import Header from "../components/Header.svelte";
</script>

<main>
    <Header />
    <div class="main">
        <div class="box" id="box1">
            
        </div>
        <div class="box" id="box2">

        </div>
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

</style>
