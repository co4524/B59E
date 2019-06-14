<template>
  <div style="display:flex; flex-direction:column; align-item:center; justify-conter:center">
    <v-text-field
        v-model="EAS.id"
        prepend-icon="star"
        placeholder="EASID"
        type="text"
        color="cyan"
        class="input"
        required
    ></v-text-field>
    <v-btn color="info" depressed @click="submit">Submit</v-btn>
  </div>
</template>

<script>
import { retrieveEAS } from "../test.js"
export default {
  name: "RegisterName",
  data() {
    return {
      EAS: {
        id: ""
      }
    };
  },
  methods: {
    async submit() {
      let result = await retrieveEAS( this.EAS.id ) ;
      alert( result ) ;
      console.log( this.EAS.id ) ;
    }
  }
}
</script>


<style lang="scss">
#status {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  h1 {
    color: white;
    size: 28px;
  }

}
#exchange {
  display: grid; //  10% 15% 10% 50% 15%
  grid-template-rows: 60px 90px 60px 200px auto;
  grid-template-columns: 1fr 1fr;

  h2 {
    color: #cfcfcf;
    font-weight: 500;
  }

  .metamask-info {
    p {
      line-height: 10px;
    }
  }
  #currency_header {
    grid-row-start: 1;
    grid-row-end: 2;
    grid-column-start: 1;
    grid-column-end: 3;
    padding: 30px;
  }
  #exchange_header {
    grid-row-start: 3;
    grid-row-end: 4;
    grid-column-start: 1;
    grid-column-end: 3;
    padding: 30px;
  }
  #Lcurrency {
    grid-row-start: 2;
    grid-row-end: 3;
    grid-column-start: 1;
    grid-column-end: 2;
    padding: 30px;
  }
  #Bcurrency {
    grid-row-start: 2;
    grid-row-end: 3;
    grid-column-start: 2;
    grid-column-end: 3;
  }
  #linepoint {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 30px;
  }
  #bcppoint {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 30px;
  }
  #arrowWraper {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  .point {
    display: block;
    width: 84px;
  }
  .arrow {
    display: block;
    width: 40px;
  }
  #btn_submit {
    grid-column-start: 1;
    grid-column-end: 3;
    display: flex;
    justify-content: space-around;
    align-items: flex-start;
  }
  #has-metamask {
    color: green;
  }
  #no-metamask {
    color: red;
  }
}
</style>
