<template>
  <div>
    <layout
      top-nav
      logo="https://i.ibb.co/5GFxj30/handcuffs-2.png"
      :router="true"
      :side-nav="sideNav"
      toggle="full"
      class="discover"
    >
      <div class="container" slot="content">
        <div class="row">
          <div class="col-md-12">
            <modal
              id="basicModal"
              :active="modalIsOpen"
              title="Error"
              :show-button="false"
              @toggle="modalIsOpen = false"
            >
              <p>{{ this.error }}</p>
            </modal>

            <div class="row mb-4">
            <header class="col-md-4 mb-3 ml-1">
              <h1 style="font-size: 4rem;" class="mt-5">Discover</h1>              
            </header>
            <img class="col-md-6" src="../assets/discover.png">
            </div>

            <div class="row">
              <div
                class="col-md-6"
                v-for="(recommendation, index) in userRecommendations"
                :key="index"
              >
                <card class="mb-3">
                  <div slot="header">
                    <h3>
                      Name:
                      {{
                        recommendation.name ? recommendation.name : "Undefined"
                      }}
                    </h3>
                  </div>
                  <ul class="list-highlight">
                    <li class="my-2">
                      Email:
                      {{
                        recommendation.email
                          ? recommendation.email
                          : "Undefined"
                      }}
                    </li>
                  </ul>
                  <div slot="footer">
                    <custom-button
                      class="mr-2"
                      icon="thumbs-up"
                      @click="swipeRight()"
                      outline
                    >
                    </custom-button>
                    <custom-button
                      class="mr-2"
                      icon="thumbs-down"
                      @click="swipeLeft()"
                      outline
                    >
                    </custom-button>
                  </div>
                </card>
              </div>
            </div>
          </div>
        </div>
      </div>
    </layout>
  </div>
</template>

<script>
import {
  Layout,
  Card,
  Button,
  Modal
} from "rbc-wm-framework-vuejs/dist/wm/components";
import { mapActions, mapState } from "vuex";
import sideNav from "../../sidenav.JSON";
export default {
  name: "Discover",
  components: {
    "layout": Layout,
    "card": Card,
    "custom-button": Button,
    "modal": Modal
  },
  data() {
    return {
      sideNav: sideNav,
      dynamicComponent: "Discover",
      error: "",
      modalIsOpen: false
    };
  },
  computed: {
    ...mapState({
      userRecommendations: state => state.wheypal.userRecommendations,
      userToken: state => state.wheypal.userToken,
      userEmail: state => state.wheypal.userEmail
    })
  },
  methods: {
    ...mapActions(["getRecommendations"])
  },
  async created() {
    this.error = "";
    const body = {
      userEmail: this.userEmail,
      userToken: this.userToken
    };
    try {
      await this.getRecommendations(body);
    } catch (e) {
      this.error = e;
      this.modalIsOpen = !this.modalIsOpen;
    }
  }
};
</script>
<style>
h1 {
  font-size: 3.074rem;
}
.discover {
  padding-left: 5%;
  padding-top: 0%;
}
</style>
