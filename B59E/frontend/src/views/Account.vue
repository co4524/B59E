<template>
  <v-content>
    <v-container fluid fill-height id="account">
      <div id="sidebar">
        <div v-for="(item, idx) in sideOption" :key="idx" class="sidewrap">
          <i :class="item.icon" v-on:click="select($event, item.title)"></i>
          <p>{{ item.title }}</p>
        </div>
      </div>
      <div id="main">
        <div id="header">
          <avatar class="avatar" :username="user.name" :size="64"></avatar>
          <div>
            <h2>{{ user.name }}</h2>
            <span>Member since {{ user.registertime}}</span>
          </div>
        </div>
        <div v-if="selected === 'Overview'">
        <!-- Component -->
          <Overview/>
        <!-------------->
        </div>
        <!-- Component -->
        <div v-else-if="selected === 'Exchange'">
          <Exchange/>
        <!-------------->
        </div>
      </div>
    </v-container>
  </v-content>
</template>

<script>
import Avatar from "vue-avatar";
import { mapState } from "vuex";
import Overview from '@/components/Overview'
import Exchange from '@/components/Exchange'
export default {
  components: {
    Avatar,
    Overview,
    Exchange,
  },
  data() {
    return {
      selected: "Overview",
      lastClick: "",
      sideOption: [
        { title: "Overview", icon: "first-selected icon fas fa-user-alt" },
        { title: "History", icon: "icon fas fa-history" },
        { title: "Detail", icon: "icon fas fa-kiwi-bird" },
        { title: "Exchange", icon: "icon fas fa-search-dollar" }
      ],
    };
  },
  computed: {
    ...mapState('account/',{
      user: state => state.user
    })
  },

  async beforeMount() {
  },


  methods: {
    select(event, name) {
      let lastEvent = this.lastClick;
      if (lastEvent == "") {
        document.getElementsByClassName("first-selected")[0].style.background =
          "#37efba";
        lastEvent = event;
      }
      lastEvent.target.style.background = "#37efba";
      this.selected = name;
      event.target.style.background = "#ffb74d"; // change new click color
      this.lastClick = event;
    },
  },
};
</script>


<style lang="scss">
.input_text {
  width: 80%;
}
.ex_text {
  width: 60%;
}

.divider {
  height: 30px;
}

#account {
  width: 70%;
  display: grid;
  grid-template-columns: 120px auto;
  grid-column-gap: 100px;
}
#sidebar {
  grid-column-start: 1;
  grid-column-end: 2;
  background-color: #424242;
  height: 86%;
  /*flexbox*/
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;

  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
  .sidewrap {
    text-align: center;
    width: 100%;
  }

  .icon {
    cursor: pointer;
    border-radius: 50%;
    font-size: 3em;
    padding: 12px;
    vertical-align: middle;
    color: white;
    background-color: #37efba;
    margin-bottom: 5px;
    -webkit-transition: all 0.3s ease-in-out;
    -moz-transition: all 0.3s ease-in-out;
    -ms-transition: all 0.3s ease-in-out;
    -o-transition: all 0.3s ease-in-out;
    transition: all 0.3s ease-in-out;

    &:hover,
    &:active {
      -webkit-transform: scale(1.1) rotate(360deg);
      -moz-transform: scale(1.1) rotate(360deg);
      -ms-transform: scale(1.1) rotate(360deg);
      -o-transform: scale(1.1) rotate(360deg);
      transform: scale(1.1) rotate(360deg);
    }
  }
  .first-selected {
    background-color: #ffb74d;
  }
}
#main {
  grid-column-start: 2;
  grid-column-end: 3;
  background-color: #424242;
  height: 86%;
  padding: 32px;
  display: grid;
  grid-template-rows: 80px auto;
  box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
  transition: all 0.3s cubic-bezier(.25,.8,.25,1);
  overflow: hidden;

  #header {
    grid-row-start: 1;
    grid-row-end: 2;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    border-bottom: 1px solid #b9b7b7;
    .avatar {
      margin-right: 20px;
    }
  }

  #overview {
    display: grid;
    grid-template-columns: 1fr 1fr;
    h2 {
      color: #cfcfcf;
      font-weight: 500;
    }
    #email-passw {
      padding: 30px;
    }
    #information {
      padding: 30px;
    }
  }

}
</style>
