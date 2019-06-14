<template>
    <div id="title">
        <div id="title-wrap">
            <h1>WELCOME TO OUR B59E PLATFORM!</h1>
            <p>The platform to exchange treasure</p>
            <i class="fas fa-long-arrow-alt-right"><router-link class="router-text" to="/Project">start exchange</router-link></i>
        </div>
    </div>
</template>

<style lang="scss">
#title {
    background-image: url("../../assets/backgroundPic.png");
    background-size: cover;
    background-repeat: no-repeat;
    background-position: center;
    height: 100vh;
    width: 100%;
    object-fit: cover;
    z-index: 0;

    display: flex;
    justify-content: center;
    align-items: center;

    #title-wrap {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: 50%;
        h1 {
            // font-size: 5em;
            text-align: center;
            font-size: 5vw;
            font-family: 'Roboto Condensed', sans-serif;
        }
        p {
            font-size: 1vw;
            font-weight: 500;
        }

        i {
            font-size: 16px;
            color: #ffc815;
        }
        .router-text {
            color: #ffc815;
            font-size: 16px;
            font-family: 'Lato', sans-serif;
            &:link { text-decoration: none}
        }
    }

}

</style>
