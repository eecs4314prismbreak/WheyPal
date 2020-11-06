<template>
  <div>
    <layout
      top-nav
      logo="https://i.ibb.co/5GFxj30/handcuffs-2.png"
      :router="true"
      :side-nav="sideNav"
      toggle="full"
      class="matches"
    >
      <div class="container" slot="content">
        <div class="row">
          <div class="col-md-12">
            <div class="row mb-4">
            <header class="col-md-4 ml-2 mb-3">
              <h1 style="font-size: 4rem;" class="mt-5">Matches</h1>
            </header>
            <img class="col-md-6" src="../assets/matches.png">
            </div>
          </div>
        </div>
      </div>
    </layout>
  </div>
</template>

<script>
import { Layout } from "rbc-wm-framework-vuejs/dist/wm/components";
import sideNav from "../../sidenav.JSON";
export default {
  name: "Matches",
  components: {
    "layout": Layout
  },
  data() {
    return {
      dynamicComponent: "Discover",
      sideNav: sideNav
    };
  },
  methods: {
    hi() {
      return console.log("hi");
    }
  }
};
</script>
<style>
.matches {
  padding-left: 5%;
  padding-top: 0%;
}
</style>
