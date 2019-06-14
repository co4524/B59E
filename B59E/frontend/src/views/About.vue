<template>
  <div id="about">
    <div id="about-word">
      <h2>B59E Platform</h2>
      <h2>ABOUT US</h2>
    </div>
    <div id="about-pic">
      <div id="about-wrap">
        <div class=about-obj>
          <v-icon size="48" color="cyan" class="icon">cloud</v-icon>  
          <p>Decentralized Tech</p>
        </div>
        <div class=about-obj>
          <v-icon size="48" color="orange" class="icon">show_chart</v-icon>  
          <p>B＝ Block</p>
        </div>
        <div class=about-obj>
          <v-icon size="48" color="green accent-2" class="icon">public</v-icon>  
          <p>E＝ Exchange</p>
        </div>
        <div class=about-obj>
          <v-icon size="48" color="blue accent-1" class="icon">cancel</v-icon>  
          <p>No High commission</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style lang="scss">
$main-color: #00b285;

#about {
  height: 100vh;
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1fr;
  z-index: -999;

  background: linear-gradient(
    70deg,
    transparent 53%,
    #37474f 40%,
    #37474f 100%,
    transparent 60%
  );
}

#about-word {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: flex-end;
  h2 {
    font-size: 4.2vw;
    font-family: "Roboto Condensed", sans-serif;
  }
}

#about-pic {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;


  #about-wrap {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    .about-obj {
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 10px;
    }
    margin-left: 100px;
    > div:nth-child(2) {
      margin-left: 30px;
    }
    > div:nth-child(3) {
      margin-left: 60px;
    }
    > div:nth-child(4) {
      margin-left: 90px;
    }

    .icon {
      border: 3px solid;
      border-radius: 50%;
      padding: 6px;
    }


    p {
      font-size: 32px;
      margin: 0 auto;
      margin-left: 10px;
    }
  }
}
</style>
